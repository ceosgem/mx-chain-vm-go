package testcommon

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/data/vm"
	logger "github.com/ElrondNetwork/elrond-go-logger"
	vmcommon "github.com/ElrondNetwork/elrond-vm-common"
	"github.com/ElrondNetwork/elrond-vm-common/builtInFunctions"
	"github.com/ElrondNetwork/elrond-vm-common/parsers"
	"github.com/ElrondNetwork/wasm-vm/arwen"
	arwenHost "github.com/ElrondNetwork/wasm-vm/arwen/host"
	"github.com/ElrondNetwork/wasm-vm/arwen/mock"
	"github.com/ElrondNetwork/wasm-vm/config"
	"github.com/ElrondNetwork/wasm-vm/crypto/hashing"
	contextmock "github.com/ElrondNetwork/wasm-vm/mock/context"
	worldmock "github.com/ElrondNetwork/wasm-vm/mock/world"
	"github.com/ElrondNetwork/wasm-vm/wasmer"
	"github.com/stretchr/testify/require"
)

var log = logger.GetOrCreate("arwen/host")

// DefaultVMType is an exposed value to use in tests
var DefaultVMType = []byte{0xF, 0xF}

// ErrAccountNotFound is an exposed value to use in tests
var ErrAccountNotFound = errors.New("account not found")

// UserAddress is an exposed value to use in tests
var UserAddress = []byte("userAccount.....................")
var UserAddress2 = []byte("userAccount2....................")

// AddressSize is the size of an account address, in bytes.
const AddressSize = 32

// SCAddressPrefix is the prefix of any smart contract address used for testing.
var SCAddressPrefix = []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x0f\x0f")

// ParentAddress is an exposed value to use in tests
var ParentAddress = MakeTestSCAddress("parentSC")

// ChildAddress is an exposed value to use in tests
var ChildAddress = MakeTestSCAddress("childSC")

// NephewAddress is an exposed value to use in tests
var NephewAddress = MakeTestSCAddress("NephewAddress")

var customGasSchedule = config.GasScheduleMap(nil)

// ESDTTransferGasCost is an exposed value to use in tests
var ESDTTransferGasCost = uint64(1)

// ESDTTestTokenName is an exposed value to use in tests
var ESDTTestTokenName = []byte("TTT-010101")

// DefaultCodeMetadata is an exposed value to use in tests
var DefaultCodeMetadata = []byte{3, 0}

// MakeTestSCAddress generates a new smart contract address to be used for
// testing based on the given identifier.
func MakeTestSCAddress(identifier string) []byte {
	numberOfTrailingDots := AddressSize - len(SCAddressPrefix) - len(identifier)
	leftBytes := SCAddressPrefix
	rightBytes := []byte(identifier + strings.Repeat(".", numberOfTrailingDots))
	return append(leftBytes, rightBytes...)
}

// GetSCCode retrieves the bytecode of a WASM module from a file
func GetSCCode(fileName string) []byte {
	code, err := ioutil.ReadFile(filepath.Clean(fileName))
	if err != nil {
		panic(fmt.Sprintf("GetSCCode(): %s", fileName))
	}

	return code
}

// GetTestSCCode retrieves the bytecode of a WASM testing contract
func GetTestSCCode(scName string, prefixToTestSCs ...string) []byte {
	var searchedPaths []string
	for _, prefixToTestSC := range prefixToTestSCs {
		pathToSC := prefixToTestSC + "test/contracts/" + scName + "/output/" + scName + ".wasm"
		searchedPaths = append(searchedPaths, pathToSC)
		code, err := ioutil.ReadFile(filepath.Clean(pathToSC))
		if err == nil {
			return code
		}
	}
	panic(fmt.Sprintf("GetSCCode(): %s", searchedPaths))
}

// GetTestSCCodeModule retrieves the bytecode of a WASM testing contract, given
// a specific name of the WASM module
func GetTestSCCodeModule(scName string, moduleName string, prefixToTestSCs string) []byte {
	pathToSC := prefixToTestSCs + "test/contracts/" + scName + "/output/" + moduleName + ".wasm"
	return GetSCCode(pathToSC)
}

// BuildSCModule invokes erdpy to build the contract into a WASM module
func BuildSCModule(scName string, prefixToTestSCs string) {
	pathToSCDir := prefixToTestSCs + "test/contracts/" + scName
	out, err := exec.Command("erdpy", "contract", "build", "--no-optimization", pathToSCDir).Output()
	if err != nil {
		log.Error("error building contract", "err", err, "contract", pathToSCDir)
		return
	}

	log.Info("contract built", "output", fmt.Sprintf("\n%s", out))
}

// DefaultTestArwenForDeployment creates an Arwen vmHost configured for testing deployments
func DefaultTestArwenForDeployment(t *testing.T, _ uint64, newAddress []byte) (arwen.VMHost, *contextmock.BlockchainHookStub) {
	stubBlockchainHook := &contextmock.BlockchainHookStub{}
	stubBlockchainHook.GetUserAccountCalled = func(address []byte) (vmcommon.UserAccountHandler, error) {
		return &contextmock.StubAccount{
			Nonce: 24,
		}, nil
	}
	stubBlockchainHook.NewAddressCalled = func(creatorAddress []byte, nonce uint64, vmType []byte) ([]byte, error) {
		return newAddress, nil
	}

	host := DefaultTestArwen(t, stubBlockchainHook)
	return host, stubBlockchainHook
}

// DefaultTestArwenForCall creates a BlockchainHookStub
func DefaultTestArwenForCall(tb testing.TB, code []byte, balance *big.Int) (arwen.VMHost, *contextmock.BlockchainHookStub) {
	stubBlockchainHook := &contextmock.BlockchainHookStub{}
	stubBlockchainHook.GetUserAccountCalled = func(scAddress []byte) (vmcommon.UserAccountHandler, error) {
		if bytes.Equal(scAddress, ParentAddress) {
			return &contextmock.StubAccount{
				Balance: balance,
			}, nil
		}
		return nil, ErrAccountNotFound
	}
	stubBlockchainHook.GetCodeCalled = func(account vmcommon.UserAccountHandler) []byte {
		return code
	}

	host := DefaultTestArwen(tb, stubBlockchainHook)
	return host, stubBlockchainHook
}

// DefaultTestArwenForCall creates a BlockchainHookStub
func DefaultTestArwenForCallSigSegv(tb testing.TB, code []byte, balance *big.Int) (arwen.VMHost, *contextmock.BlockchainHookStub) {
	stubBlockchainHook := &contextmock.BlockchainHookStub{}
	stubBlockchainHook.GetUserAccountCalled = func(scAddress []byte) (vmcommon.UserAccountHandler, error) {
		if bytes.Equal(scAddress, ParentAddress) {
			return &contextmock.StubAccount{
				Balance: balance,
			}, nil
		}
		return nil, ErrAccountNotFound
	}
	stubBlockchainHook.GetCodeCalled = func(account vmcommon.UserAccountHandler) []byte {
		return code
	}

	customGasSchedule := config.GasScheduleMap(nil)
	host := DefaultTestArwenWithGasSchedule(tb, stubBlockchainHook, customGasSchedule, true)
	return host, stubBlockchainHook
}

// DefaultTestArwenForCallWithInstanceRecorderMock creates an InstanceBuilderRecorderMock
func DefaultTestArwenForCallWithInstanceRecorderMock(tb testing.TB, code []byte, balance *big.Int) (arwen.VMHost, *contextmock.InstanceBuilderRecorderMock) {
	// this uses a Blockchain Hook Stub that does not cache the compiled code
	host, _ := DefaultTestArwenForCall(tb, code, balance)

	instanceBuilderRecorderMock := contextmock.NewInstanceBuilderRecorderMock()
	host.Runtime().ReplaceInstanceBuilder(instanceBuilderRecorderMock)

	return host, instanceBuilderRecorderMock
}

// DefaultTestArwenForCallWithInstanceMocks creates an InstanceBuilderMock
func DefaultTestArwenForCallWithInstanceMocks(tb testing.TB) (arwen.VMHost, *contextmock.InstanceBuilderMock) {
	world := worldmock.NewMockWorld()
	return DefaultTestArwenForCallWithInstanceMocksAndWorld(tb, world)
}

// DefaultTestArwenForCallWithInstanceMocksAndWorld creates an InstanceBuilderMock
func DefaultTestArwenForCallWithInstanceMocksAndWorld(tb testing.TB, world *worldmock.MockWorld) (arwen.VMHost, *contextmock.InstanceBuilderMock) {
	if world == nil {
		world = worldmock.NewMockWorld()
	}
	host := DefaultTestArwen(tb, world)

	instanceBuilderMock := contextmock.NewInstanceBuilderMock(world)
	host.Runtime().ReplaceInstanceBuilder(instanceBuilderMock)

	return host, instanceBuilderMock
}

// DefaultTestArwenForCallWithWorldMock creates a MockWorld
func DefaultTestArwenForCallWithWorldMock(tb testing.TB, code []byte, balance *big.Int) (arwen.VMHost, *worldmock.MockWorld) {
	world := worldmock.NewMockWorld()
	host := DefaultTestArwen(tb, world)

	err := world.InitBuiltinFunctions(host.GetGasScheduleMap())
	require.Nil(tb, err)

	host.SetBuiltInFunctionsContainer(world.BuiltinFuncs.Container)

	parentAccount := world.AcctMap.CreateSmartContractAccount(UserAddress, ParentAddress, code, world)
	parentAccount.Balance = balance

	return host, world
}

// DefaultTestArwenForTwoSCs creates an Arwen vmHost configured for testing calls between 2 SmartContracts
func DefaultTestArwenForTwoSCs(
	t *testing.T,
	parentCode []byte,
	childCode []byte,
	parentSCBalance *big.Int,
	childSCBalance *big.Int,
) (arwen.VMHost, *contextmock.BlockchainHookStub) {
	stubBlockchainHook := &contextmock.BlockchainHookStub{}

	if parentSCBalance == nil {
		parentSCBalance = big.NewInt(1000)
	}

	if childSCBalance == nil {
		childSCBalance = big.NewInt(1000)
	}

	stubBlockchainHook.GetUserAccountCalled = func(scAddress []byte) (vmcommon.UserAccountHandler, error) {
		if bytes.Equal(scAddress, ParentAddress) {
			return &contextmock.StubAccount{
				Address: ParentAddress,
				Balance: parentSCBalance,
			}, nil
		}
		if bytes.Equal(scAddress, ChildAddress) {
			return &contextmock.StubAccount{
				Address: ChildAddress,
				Balance: childSCBalance,
			}, nil
		}

		return nil, ErrAccountNotFound
	}
	stubBlockchainHook.GetCodeCalled = func(account vmcommon.UserAccountHandler) []byte {
		if bytes.Equal(account.AddressBytes(), ParentAddress) {
			return parentCode
		}
		if bytes.Equal(account.AddressBytes(), ChildAddress) {
			return childCode
		}
		return nil
	}

	host := DefaultTestArwen(t, stubBlockchainHook)
	return host, stubBlockchainHook
}

func defaultTestArwenForContracts(
	tb testing.TB,
	contracts []*InstanceTestSmartContract,
	gasSchedule config.GasScheduleMap,
	wasmerSIGSEGVPassthrough bool,
) (arwen.VMHost, *contextmock.BlockchainHookStub) {

	stubBlockchainHook := &contextmock.BlockchainHookStub{}

	contractsMap := make(map[string]*contextmock.StubAccount)
	codeMap := make(map[string]*[]byte)

	for _, contract := range contracts {
		codeHash, _ := hashing.NewHasher().Sha256(contract.code)
		contractsMap[string(contract.address)] = &contextmock.StubAccount{
			Address:      contract.address,
			Balance:      big.NewInt(contract.balance),
			CodeHash:     codeHash,
			CodeMetadata: DefaultCodeMetadata,
			OwnerAddress: ParentAddress,
		}
		codeMap[string(contract.address)] = &contract.code
	}

	stubBlockchainHook.GetUserAccountCalled = func(scAddress []byte) (vmcommon.UserAccountHandler, error) {
		contract, found := contractsMap[string(scAddress)]
		if found {
			return contract, nil
		}
		return nil, ErrAccountNotFound
	}
	stubBlockchainHook.GetCodeCalled = func(account vmcommon.UserAccountHandler) []byte {
		code, found := codeMap[string(account.AddressBytes())]
		if found {
			return *code
		}
		return nil
	}

	host := DefaultTestArwenWithGasSchedule(tb, stubBlockchainHook, gasSchedule, wasmerSIGSEGVPassthrough)
	return host, stubBlockchainHook
}

// DefaultTestArwenWithWorldMock creates a host configured with a mock world
func DefaultTestArwenWithWorldMock(tb testing.TB) (arwen.VMHost, *worldmock.MockWorld) {
	customGasSchedule := config.GasScheduleMap(nil)
	return DefaultTestArwenWithWorldMockWithGasSchedule(tb, customGasSchedule)
}

// DefaultTestArwenWithWorldMockWithGasSchedule creates a host configured with a mock world
func DefaultTestArwenWithWorldMockWithGasSchedule(tb testing.TB, customGasSchedule config.GasScheduleMap) (arwen.VMHost, *worldmock.MockWorld) {
	world := worldmock.NewMockWorld()
	gasSchedule := customGasSchedule
	if gasSchedule == nil {
		gasSchedule = config.MakeGasMapForTests()
	}
	err := world.InitBuiltinFunctions(gasSchedule)
	require.Nil(tb, err)

	esdtTransferParser, _ := parsers.NewESDTTransferParser(worldmock.WorldMarshalizer)
	host, err := arwenHost.NewArwenVM(
		world,
		wasmer.NewExecutor(),
		&arwen.VMHostParameters{
			VMType:                   DefaultVMType,
			BlockGasLimit:            uint64(1000),
			GasSchedule:              gasSchedule,
			BuiltInFuncContainer:     world.BuiltinFuncs.Container,
			ElrondProtectedKeyPrefix: []byte("ELROND"),
			ESDTTransferParser:       esdtTransferParser,
			EpochNotifier:            &mock.EpochNotifierStub{},
			EnableEpochsHandler: &worldmock.EnableEpochsHandlerStub{
				IsStorageAPICostOptimizationFlagEnabledField:     true,
				IsMultiESDTTransferFixOnCallBackFlagEnabledField: true,
				IsFixOOGReturnCodeFlagEnabledField:               true,
				IsRemoveNonUpdatedStorageFlagEnabledField:        true,
				IsCreateNFTThroughExecByCallerFlagEnabledField:   true,
				IsManagedCryptoAPIsFlagEnabledField:              true,
				IsFailExecutionOnEveryAPIErrorFlagEnabledField:   true,
				IsESDTTransferRoleFlagEnabledField:               true,
				IsSendAlwaysFlagEnabledField:                     true,
				IsGlobalMintBurnFlagEnabledField:                 true,
				IsCheckFunctionArgumentFlagEnabledField:          true,
				IsCheckExecuteOnReadOnlyFlagEnabledField:         true,
			},
			WasmerSIGSEGVPassthrough: false,
		})
	require.Nil(tb, err)
	require.NotNil(tb, host)

	return host, world
}

// DefaultTestArwen creates a host configured with a configured blockchain hook
func DefaultTestArwen(tb testing.TB, blockchain vmcommon.BlockchainHook) arwen.VMHost {
	customGasSchedule := config.GasScheduleMap(nil)
	return DefaultTestArwenWithGasSchedule(tb, blockchain, customGasSchedule, false)
}

func DefaultTestArwenWithGasSchedule(
	tb testing.TB,
	blockchain vmcommon.BlockchainHook,
	customGasSchedule config.GasScheduleMap,
	wasmerSIGSEGVPassthrough bool,
) arwen.VMHost {
	gasSchedule := customGasSchedule
	if gasSchedule == nil {
		gasSchedule = config.MakeGasMapForTests()
	}

	esdtTransferParser, _ := parsers.NewESDTTransferParser(worldmock.WorldMarshalizer)
	host, err := arwenHost.NewArwenVM(
		blockchain,
		wasmer.NewExecutor(),
		&arwen.VMHostParameters{
			VMType:                   DefaultVMType,
			BlockGasLimit:            uint64(1000),
			GasSchedule:              gasSchedule,
			BuiltInFuncContainer:     builtInFunctions.NewBuiltInFunctionContainer(),
			ElrondProtectedKeyPrefix: []byte("ELROND"),
			ESDTTransferParser:       esdtTransferParser,
			EpochNotifier:            &mock.EpochNotifierStub{},
			EnableEpochsHandler: &worldmock.EnableEpochsHandlerStub{
				IsStorageAPICostOptimizationFlagEnabledField:         true,
				IsMultiESDTTransferFixOnCallBackFlagEnabledField:     true,
				IsFixOOGReturnCodeFlagEnabledField:                   true,
				IsRemoveNonUpdatedStorageFlagEnabledField:            true,
				IsCreateNFTThroughExecByCallerFlagEnabledField:       true,
				IsManagedCryptoAPIsFlagEnabledField:                  true,
				IsFailExecutionOnEveryAPIErrorFlagEnabledField:       true,
				IsRefactorContextFlagEnabledField:                    true,
				IsCheckCorrectTokenIDForTransferRoleFlagEnabledField: true,
				IsDisableExecByCallerFlagEnabledField:                true,
				IsESDTTransferRoleFlagEnabledField:                   true,
				IsSendAlwaysFlagEnabledField:                         true,
				IsGlobalMintBurnFlagEnabledField:                     true,
				IsCheckFunctionArgumentFlagEnabledField:              true,
				IsCheckExecuteOnReadOnlyFlagEnabledField:             true,
			},
			WasmerSIGSEGVPassthrough: wasmerSIGSEGVPassthrough,
		})
	require.Nil(tb, err)
	require.NotNil(tb, host)

	return host
}

// AddTestSmartContractToWorld directly deploys the provided code into the
// given MockWorld under a SC address built with the given identifier.
func AddTestSmartContractToWorld(world *worldmock.MockWorld, identifier string, code []byte) *worldmock.Account {
	address := MakeTestSCAddress(identifier)
	return world.AcctMap.CreateSmartContractAccount(UserAddress, address, code, world)
}

// MakeEmptyContractCallInput instantiates an empty ContractCallInput
func MakeEmptyContractCallInput() *vmcommon.ContractCallInput {
	return &vmcommon.ContractCallInput{
		VMInput: vmcommon.VMInput{
			CallerAddr:           nil,
			Arguments:            make([][]byte, 0),
			CallValue:            big.NewInt(0),
			CallType:             vm.DirectCall,
			GasPrice:             1,
			GasProvided:          0,
			ReturnCallAfterError: false,
		},
		RecipientAddr: nil,
		Function:      "",
	}
}

// MakeContractCallInput creates a ContractCallInput and sets the provided arguments
func MakeContractCallInput(
	caller []byte,
	recipient []byte,
	function string,
	value int,
) *vmcommon.ContractCallInput {
	input := MakeEmptyContractCallInput()
	SetCallParties(input, caller, recipient)
	input.Function = function
	input.CallValue = big.NewInt(int64(value))
	return input
}

// SetCallParties sets the caller and recipient of the given ContractCallInput
func SetCallParties(input *vmcommon.ContractCallInput, caller []byte, recipient []byte) {
	input.CallerAddr = caller
	input.RecipientAddr = recipient
}

// AddArgument adds the provided argument to the ContractCallInput
func AddArgument(input *vmcommon.ContractCallInput, argument []byte) {
	if input.Arguments == nil {
		input.Arguments = make([][]byte, 0)
	}
	input.Arguments = append(input.Arguments, argument)
}

// CopyTxHashes copies the tx hashes from a source ContractCallInput into another
func CopyTxHashes(input *vmcommon.ContractCallInput, sourceInput *vmcommon.ContractCallInput) {
	input.CurrentTxHash = sourceInput.CurrentTxHash
	input.PrevTxHash = sourceInput.PrevTxHash
	input.OriginalTxHash = sourceInput.OriginalTxHash
}

// DefaultTestContractCreateInput creates a vmcommon.ContractCreateInput struct
// with default values.
func DefaultTestContractCreateInput() *vmcommon.ContractCreateInput {
	return &vmcommon.ContractCreateInput{
		VMInput: vmcommon.VMInput{
			CallerAddr: []byte("caller"),
			Arguments: [][]byte{
				[]byte("argument 1"),
				[]byte("argument 2"),
			},
			CallValue:   big.NewInt(0),
			CallType:    vm.DirectCall,
			GasPrice:    0,
			GasProvided: 0,
		},
		ContractCode: []byte("contract"),
	}
}

// DefaultTestContractCallInput creates a vmcommon.ContractCallInput struct
// with default values.
func DefaultTestContractCallInput() *vmcommon.ContractCallInput {
	return &vmcommon.ContractCallInput{
		VMInput: vmcommon.VMInput{
			CallerAddr:  UserAddress,
			Arguments:   make([][]byte, 0),
			CallValue:   big.NewInt(0),
			CallType:    vm.DirectCall,
			GasPrice:    0,
			GasProvided: 0,
		},
		RecipientAddr: ParentAddress,
		Function:      "function",
	}
}

// ContractCallInputBuilder extends a ContractCallInput for extra building functionality during testing
type ContractCallInputBuilder struct {
	vmcommon.ContractCallInput
}

// CreateTestContractCallInputBuilder is a builder for ContractCallInputBuilder
func CreateTestContractCallInputBuilder() *ContractCallInputBuilder {
	return &ContractCallInputBuilder{
		ContractCallInput: *DefaultTestContractCallInput(),
	}
}

// WithRecipientAddr provides the recepient address of ContractCallInputBuilder
func (contractInput *ContractCallInputBuilder) WithRecipientAddr(address []byte) *ContractCallInputBuilder {
	contractInput.ContractCallInput.RecipientAddr = address
	return contractInput
}

// WithCallerAddr provides the caller address of ContractCallInputBuilder
func (contractInput *ContractCallInputBuilder) WithCallerAddr(address []byte) *ContractCallInputBuilder {
	contractInput.ContractCallInput.CallerAddr = address
	return contractInput
}

// WithCallValue provides the value transferred to the called contract
func (contractInput *ContractCallInputBuilder) WithCallValue(value int64) *ContractCallInputBuilder {
	contractInput.ContractCallInput.CallValue = big.NewInt(value)
	return contractInput
}

// WithGasProvided provides the gas of ContractCallInputBuilder
func (contractInput *ContractCallInputBuilder) WithGasProvided(gas uint64) *ContractCallInputBuilder {
	contractInput.ContractCallInput.VMInput.GasProvided = gas
	return contractInput
}

// WithGasLocked provides the locked gas of ContractCallInputBuilder
func (contractInput *ContractCallInputBuilder) WithGasLocked(gas uint64) *ContractCallInputBuilder {
	contractInput.ContractCallInput.VMInput.GasLocked = gas
	return contractInput
}

// WithFunction provides the function to be called for ContractCallInputBuilder
func (contractInput *ContractCallInputBuilder) WithFunction(function string) *ContractCallInputBuilder {
	contractInput.ContractCallInput.Function = function
	return contractInput
}

// WithArguments provides the arguments to be called for ContractCallInputBuilder
func (contractInput *ContractCallInputBuilder) WithArguments(arguments ...[]byte) *ContractCallInputBuilder {
	contractInput.ContractCallInput.VMInput.Arguments = arguments
	return contractInput
}

// WithAsyncArguments provides the async arguments to be called for ContractCallInputBuilder
func (contractInput *ContractCallInputBuilder) WithAsyncArguments(arguments *vmcommon.AsyncArguments) *ContractCallInputBuilder {
	contractInput.ContractCallInput.VMInput.AsyncArguments = arguments
	return contractInput
}

// WithCallType provides the arguments to be called for ContractCallInputBuilder
func (contractInput *ContractCallInputBuilder) WithCallType(callType vm.CallType) *ContractCallInputBuilder {
	contractInput.ContractCallInput.VMInput.CallType = callType
	return contractInput
}

// WithCurrentTxHash provides the CurrentTxHash for ContractCallInputBuilder
func (contractInput *ContractCallInputBuilder) WithCurrentTxHash(txHash []byte) *ContractCallInputBuilder {
	contractInput.ContractCallInput.CurrentTxHash = txHash
	return contractInput
}

// WithPrevTxHash provides the PrevTxHash for ContractCallInputBuilder
func (contractInput *ContractCallInputBuilder) WithPrevTxHash(txHash []byte) *ContractCallInputBuilder {
	contractInput.ContractCallInput.PrevTxHash = txHash
	return contractInput
}

func (contractInput *ContractCallInputBuilder) initESDTTransferIfNeeded() {
	if len(contractInput.ESDTTransfers) == 0 {
		contractInput.ESDTTransfers = make([]*vmcommon.ESDTTransfer, 1)
		contractInput.ESDTTransfers[0] = &vmcommon.ESDTTransfer{}
	}
}

// WithESDTValue provides the ESDTValue for ContractCallInputBuilder
func (contractInput *ContractCallInputBuilder) WithESDTValue(esdtValue *big.Int) *ContractCallInputBuilder {
	contractInput.initESDTTransferIfNeeded()
	contractInput.ContractCallInput.ESDTTransfers[0].ESDTValue = esdtValue
	return contractInput
}

// WithESDTTokenName provides the ESDTTokenName for ContractCallInputBuilder
func (contractInput *ContractCallInputBuilder) WithESDTTokenName(esdtTokenName []byte) *ContractCallInputBuilder {
	contractInput.initESDTTransferIfNeeded()
	contractInput.ContractCallInput.ESDTTransfers[0].ESDTTokenName = esdtTokenName
	return contractInput
}

// Build completes the build of a ContractCallInput
func (contractInput *ContractCallInputBuilder) Build() *vmcommon.ContractCallInput {
	return &contractInput.ContractCallInput
}

// ContractCreateInputBuilder extends a ContractCreateInput for extra building functionality during testing
type ContractCreateInputBuilder struct {
	vmcommon.ContractCreateInput
}

// CreateTestContractCreateInputBuilder is a builder for ContractCreateInputBuilder
func CreateTestContractCreateInputBuilder() *ContractCreateInputBuilder {
	return &ContractCreateInputBuilder{
		ContractCreateInput: *DefaultTestContractCreateInput(),
	}
}

// WithGasProvided provides the GasProvided for a ContractCreateInputBuilder
func (contractInput *ContractCreateInputBuilder) WithGasProvided(gas uint64) *ContractCreateInputBuilder {
	contractInput.ContractCreateInput.GasProvided = gas
	return contractInput
}

// WithContractCode provides the ContractCode for a ContractCreateInputBuilder
func (contractInput *ContractCreateInputBuilder) WithContractCode(code []byte) *ContractCreateInputBuilder {
	contractInput.ContractCreateInput.ContractCode = code
	return contractInput
}

// WithCallerAddr provides the CallerAddr for a ContractCreateInputBuilder
func (contractInput *ContractCreateInputBuilder) WithCallerAddr(address []byte) *ContractCreateInputBuilder {
	contractInput.ContractCreateInput.CallerAddr = address
	return contractInput
}

// WithCallValue provides the CallValue for a ContractCreateInputBuilder
func (contractInput *ContractCreateInputBuilder) WithCallValue(callValue int64) *ContractCreateInputBuilder {
	contractInput.ContractCreateInput.CallValue = big.NewInt(callValue)
	return contractInput
}

// WithArguments provides the Arguments for a ContractCreateInputBuilder
func (contractInput *ContractCreateInputBuilder) WithArguments(arguments ...[]byte) *ContractCreateInputBuilder {
	contractInput.ContractCreateInput.Arguments = arguments
	return contractInput
}

// Build completes the build of a ContractCreateInput
func (contractInput *ContractCreateInputBuilder) Build() *vmcommon.ContractCreateInput {
	return &contractInput.ContractCreateInput
}

// OpenFile method opens the file from given path - does not close the file
func OpenFile(relativePath string) (*os.File, error) {
	path, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Printf("cannot create absolute path for the provided file: %s", err.Error())
		return nil, err
	}
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}

	return f, nil
}