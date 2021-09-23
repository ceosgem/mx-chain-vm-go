package hosttest

import (
	"bytes"
	"fmt"
	"math/big"
	"strconv"
	"testing"

	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/arwen"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/arwen/contexts"
	worldmock "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/mock/world"
	"github.com/ElrondNetwork/arwen-wasm-vm/v1_4/testcommon"
	test "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/testcommon"
	"github.com/ElrondNetwork/elrond-go-core/data/vm"
	vmcommon "github.com/ElrondNetwork/elrond-vm-common"
	"github.com/ElrondNetwork/elrond-vm-common/parsers"
	"github.com/ElrondNetwork/elrond-vm-common/txDataBuilder"
	"github.com/stretchr/testify/require"
)

// TODO matei-p add error test cases

func TestGasUsed_SyncCalls_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestSyncCalls()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_SyncCalls2_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestSyncCalls2()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_OneAsyncCall_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestOneAsyncCall()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_OneAsyncCallCrossShard_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestOneAsyncCallCrossShard()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_TwoAsyncCalls_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestTwoAsyncCalls()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_TwoAsyncCalls_LocalCross_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestTwoAsyncCallsLocalCross()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_TwoAsyncCalls_CrossLocal_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestTwoAsyncCallsCrossLocal()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_TwoAsyncCallsCrossShard_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestTwoAsyncCallsCrossShard()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_AsyncCallsAsync_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestAsyncCallsAsync()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_AsyncCallsAsync_CrossLocal_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestAsyncCallsAsyncCrossLocal()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_AsyncCallsAsync_LocalCross_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestAsyncCallsAsyncLocalCross()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_AsyncCallsAsyncCrossShard_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestAsyncCallsAsyncCrossShard()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_DifferentTypeOfCallsToSameFunction_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestDifferentTypeOfCallsToSameFunction()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_CallbackCallsSync_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestCallbackCallsSync()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_CallbackCallsAsync_LocalLocal_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestCallbackCallsAsyncLocalLocal()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_CallbackCallsAsync_LocalCross_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestCallbackCallsAsyncLocalCross()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_CallbackCallsAsync_CrossLocal_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestCallbackCallsAsyncCrossLocal()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_CallbackCallsAsync_CrossCross_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestCallbackCallsAsyncCrossCross()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_SyncAndAsync1_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestSyncAndAsync1()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_SyncAndAsync2_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestSyncAndAsync2()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_SyncAndAsync3_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestSyncAndAsync3()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_SyncAndAsync4_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestSyncAndAsync4()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_SyncAndAsync5_CallGraph(t *testing.T) {
	t.Skip()
	callGraph := test.CreateGraphTestSyncAndAsync5()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_SyncAndAsync6_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestSyncAndAsync6()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_SyncAndAsync7_CallGraph(t *testing.T) {
	callGraph := test.CreateGraphTestSyncAndAsync7()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_GraphTest1_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTest1()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_GraphTest2_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTest2()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_AsyncCall2_CrossShard_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestAsyncCallsCrossShard2()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_AsyncCall3_CrossShard_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestAsyncCallsCrossShard3()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_AsyncCall4_CrossShard_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestAsyncCallsCrossShard4()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_AsyncCall5_CrossShard_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestAsyncCallsCrossShard5()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_AsyncCall6_CrossShard_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestAsyncCallsCrossShard6()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_AsyncCall7_CrossShard_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestAsyncCallsCrossShard7()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_AsyncCall8_CrossShard_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestAsyncCallsCrossShard8()
	runGraphCallTestTemplate(t, callGraph)
}

func TestGasUsed_AsyncCall9_CrossShard_CallGraph(t *testing.T) {
	// t.Skip()
	callGraph := test.CreateGraphTestAsyncCallsCrossShard9()
	runGraphCallTestTemplate(t, callGraph)
}

func runGraphCallTestTemplate(t *testing.T, callGraph *test.TestCallGraph) {
	testConfig := makeTestConfig()
	testConfig.GasProvided = callGraph.StartNode.GasLimit
	testConfig.GasLockCost = test.DefaultCallGraphLockedGas

	executionGraph := callGraph.CreateExecutionGraphFromCallGraph()

	gasGraph := executionGraph.ComputeGasGraphFromExecutionGraph()
	gasGraph.AssignExecutionRounds()
	gasGraph.ComputeRemainingGasBeforeCallbacks()
	gasGraph.ComputeRemainingGasAfterCallbacks()

	startNode := gasGraph.GetStartNode()
	crossShardCallsQueue := test.NewCrossShardCallQueue()
	crossShardCallsQueue.Enqueue(test.UserAddress, startNode, vm.DirectCall, []byte{})

	computeCallIDs(gasGraph)

	// compute execution order (return data) assertions and compute gas assertions
	totalGasUsed, totalGasRemaining, expectedReturnData := computeExpectedValues(gasGraph)

	// expected gas sanity check
	require.Equal(t, int(gasGraph.StartNode.GasLimit), int(totalGasUsed+totalGasRemaining), "Expected Gas Sanity Check")

	// account -> (key -> value)
	storage := make(map[string]map[string][]byte)

	globalReturnData := make([][]byte, 0)
	crtTxNumber := 0

	var currentVMOutput *vmcommon.VMOutput
	var lastErr error

	var crossShardCall *test.CrossShardCall
	for !crossShardCallsQueue.IsEmpty() {
		crossShardCall = crossShardCallsQueue.Dequeue()
		startNode = crossShardCall.StartNode

		crtTxNumber++
		crtTxHash := big.NewInt(int64(crtTxNumber)).Bytes()
		crossShardCall.StartNode.CrtTxHash = crtTxHash

		crossShardEdges := getCrossShardEdgesFromSubtree(gasGraph, startNode, crossShardCallsQueue)

		arguments := [][]byte{}
		if len(crossShardCall.Data) != 0 {
			_, parsedArguments, err := parsers.NewCallArgsParser().ParseData(string(crossShardCall.Data))
			if err != nil {
				panic(err)
			}
			arguments = parsedArguments
		}

		currentVMOutput, lastErr = test.BuildMockInstanceCallTest(t).
			WithContracts(
				test.CreateMockContractsFromAsyncTestCallGraph(callGraph, testConfig)...,
			).
			WithInput(test.CreateTestContractCallInputBuilder().
				WithCallerAddr(crossShardCall.CallerAddress).
				WithRecipientAddr([]byte(startNode.Call.ContractAddress)).
				WithFunction(startNode.Call.FunctionName).
				WithGasProvided(startNode.GasLimit).
				WithGasLocked(startNode.GasLocked).
				WithCallType(crossShardCall.CallType).
				WithArguments(arguments...).
				WithPrevTxHash(big.NewInt(int64(crtTxNumber - 1)).Bytes()).
				WithCurrentTxHash(crtTxHash).
				Build()).
			WithSetup(func(host arwen.VMHost, world *worldmock.MockWorld) {
				world.SelfShardID = world.GetShardOfAddress(startNode.Call.ContractAddress)
				persistStorageUpdatesToWorld(storage, world)
				setZeroCodeCosts(host)
				setAsyncCosts(host, testConfig.GasLockCost)
			}).
			AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
				verify.Ok()
			})

		extractStores(currentVMOutput, storage)
		globalReturnData = append(globalReturnData, currentVMOutput.ReturnData...)

		extractOuptutTransferCalls(currentVMOutput, crossShardEdges, crossShardCallsQueue)
	}

	checkReturnDataWithGasValuesForGraphTesting(t, expectedReturnData, globalReturnData)
	test.NewVMOutputVerifier(t, currentVMOutput, lastErr).
		Ok() //.
	// GasRemaining(callGraph.StartNode.GasLimit - totalGasUsed)
}

func persistStorageUpdatesToWorld(storage map[string]map[string][]byte, world *worldmock.MockWorld) {
	for address, store := range storage {
		account := world.AcctMap.GetAccount([]byte(address))
		if account == nil {
			continue
		}
		for key, value := range store {
			account.SaveKeyValue([]byte(key), value)
		}
	}
}

func getCrossShardEdgesFromSubtree(gasGraph *test.TestCallGraph, startNode *test.TestCallNode, crossShardCallsQueue *test.CrossShardCallsQueue) []*test.TestCallEdge {
	crossShardEdges := make([]*test.TestCallEdge, 0)
	visits := make(map[uint]bool)
	gasGraph.DfsGraphFromNode(startNode, func(path []*test.TestCallNode, parent *test.TestCallNode, node *test.TestCallNode, incomingEdge *test.TestCallEdge) *test.TestCallNode {
		for _, edge := range node.AdjacentEdges {
			if edge.Type == test.AsyncCrossShard || edge.Type == test.CallbackCrossShard {
				crossShardEdges = append(crossShardEdges, edge)
			}
		}
		return node
	}, visits, false /* don't followCrossShardEdges */)

	// if a parent context async exists, add it's callback edge also
	if startNode.IncomingEdgeType == test.CallbackCrossShard &&
		startNode.Parent != nil && startNode.Parent.Parent != nil {
		prevPrevNode := startNode.Parent.Parent
		if prevPrevNode.IncomingEdgeType == test.Async ||
			prevPrevNode.IncomingEdgeType == test.AsyncCrossShard {
			for _, edge := range prevPrevNode.AdjacentEdges {
				if edge.Type == test.Callback || edge.Type == test.CallbackCrossShard {
					crossShardEdges = append(crossShardEdges, edge)
				}
			}
		}
	}

	return crossShardEdges
}

func executionOrderTraversal(gasGraph *test.TestCallGraph, nodeProcessing func(parent *test.TestCallNode, node *test.TestCallNode) *test.TestCallNode) {
	crossShardCallsQueue := test.NewCrossShardCallQueue()
	crossShardCallsQueue.Enqueue(test.UserAddress, gasGraph.StartNode, vm.DirectCall, []byte{})
	var crossShardCall *test.CrossShardCall
	visits := make(map[uint]bool)
	for !crossShardCallsQueue.IsEmpty() {
		crossShardCall = crossShardCallsQueue.Dequeue()
		startNode := crossShardCall.StartNode

		if (startNode.IncomingEdgeType == test.Callback || startNode.IncomingEdgeType == test.CallbackCrossShard) &&
			!crossShardCallsQueue.CanExecuteLocalCallback(startNode) {
			crossShardCallsQueue.Requeue(crossShardCall)
			continue
		}

		gasGraph.DfsGraphFromNode(startNode, func(path []*test.TestCallNode, parent *test.TestCallNode, node *test.TestCallNode, incomingEdge *test.TestCallEdge) *test.TestCallNode {
			for _, edge := range node.AdjacentEdges {
				if edge.Type == test.AsyncCrossShard || edge.Type == test.CallbackCrossShard {
					destinationNode := edge.To
					var callType vm.CallType
					switch edge.Type {
					case test.AsyncCrossShard:
						callType = vm.AsynchronousCall
					case test.CallbackCrossShard:
						callType = vm.AsynchronousCallBack
					}
					crossShardCallsQueue.Enqueue(node.Call.ContractAddress, destinationNode, callType, nil)
				}
			}

			if incomingEdge != nil && incomingEdge.Type == test.Callback {
				if !crossShardCallsQueue.CanExecuteLocalCallback(node) {
					crossShardCallsQueue.Enqueue(node.Call.ContractAddress, incomingEdge.To, vm.AsynchronousCallBack, nil)
					// stop DFS for this branch
					return nil
				}
			}

			return nodeProcessing(parent, node)
		}, visits, false)
	}
}

func computeCallIDs(gasGraph *test.TestCallGraph) {
	executionOrderTraversal(gasGraph, func(parent *test.TestCallNode, node *test.TestCallNode) *test.TestCallNode {
		if node.IsLeaf() || node.Parent == nil {
			return node
		}

		if node.IncomingEdgeType == test.Callback {
			parent = node.Parent.Parent
		} else {
			parent = node.Parent
		}
		// fmt.Println("compute callID for", string(node.Call.ContractAddress)+"_"+node.Call.FunctionName, "gasLimit", parent.GasLimit)

		if parent != nil {
			parent.NonGasEdgeCounter++
			newCallID := append(parent.Call.CallID, big.NewInt(parent.NonGasEdgeCounter).Bytes()...)
			newCallID, _ = gasGraph.Crypto.Sha256(newCallID)
			node.Call.CallID = newCallID
			// fmt.Println("label", node.Label, "CallID", newCallID)
		}

		return node
	})
}

func computeExpectedValues(gasGraph *test.TestCallGraph) (uint64, uint64, [][]byte) {
	totalGasUsed := uint64(0)
	totalGasRemaining := uint64(0)
	expectedReturnData := make([][]byte, 0)

	executionOrderTraversal(gasGraph, func(parent *test.TestCallNode, node *test.TestCallNode) *test.TestCallNode {
		if !node.IsLeaf() {
			if parent == nil {
				totalGasRemaining += node.GasRemaining + node.GasAccumulated
			} else if node.IncomingEdgeType == testcommon.Callback || node.IncomingEdgeType == testcommon.CallbackCrossShard {
				totalGasRemaining += node.GasAccumulated
			}
			return node
		}

		totalGasUsed += node.GasUsed

		expectedNodeRetData := txDataBuilder.NewBuilder()
		expectedNodeRetData.Func(parent.Call.FunctionName)
		expectedNodeRetData.Str(string(parent.Call.ContractAddress) + "_" + parent.Call.FunctionName + test.TestReturnDataSuffix)
		expectedNodeRetData.Int64(int64(parent.GasLimit))
		expectedNodeRetData.Int64(int64(parent.GasRemaining))
		expectedReturnData = append(expectedReturnData, expectedNodeRetData.ToBytes())

		return node
	})

	return totalGasUsed, totalGasRemaining, expectedReturnData
}

func computeGasRemainingForNode(node *test.TestCallNode) {
	var gasUsed uint64
	for _, edge := range node.AdjacentEdges {
		destNode := edge.To
		if edge.To.IsGasLeaf() {
			gasUsed += destNode.GasUsed
		} else if edge.Type != testcommon.Callback && edge.Type != testcommon.CallbackCrossShard {
			gasUsed += destNode.GasLimit + destNode.GasLocked - destNode.GasRemaining
		}
	}
	node.GasRemaining = node.GasLimit - gasUsed
}

func extractOuptutTransferCalls(vmOutput *vmcommon.VMOutput, crossShardEdges []*test.TestCallEdge, crossShardCallsQueue *test.CrossShardCallsQueue) {
	for _, crossShardEdge := range crossShardEdges {
		edgeFromAddress := string(crossShardEdge.To.Parent.Call.ContractAddress)
		edgeToAddress := string(crossShardEdge.To.Call.ContractAddress)
		for _, outputAccount := range vmOutput.OutputAccounts {
			transferDestinationAddress := string(outputAccount.Address)
			if edgeToAddress != transferDestinationAddress {
				continue
			}
			for _, outputTransfer := range outputAccount.OutputTransfers {
				callType := outputTransfer.CallType
				transferSenderAddress := string(outputTransfer.SenderAddress)

				argParser := parsers.NewCallArgsParser()
				function, parsedArgs, _ := argParser.ParseData(string(outputTransfer.Data))

				callID := parsedArgs[0]

				var encodedArgs []byte
				if edgeFromAddress == transferSenderAddress &&
					bytes.Equal(callID, crossShardEdge.To.Call.CallID) &&
					(callType == vm.AsynchronousCall || callType == vm.AsynchronousCallBack) {
					fmt.Println(
						"Found transfer from sender", string(outputTransfer.SenderAddress),
						"to", string(outputAccount.Address),
						"gas limit", outputTransfer.GasLimit,
						"callType", callType,
						"data", contexts.DebugCallIDAsString(outputTransfer.Data))
					if callType == vm.AsynchronousCall {
						encodedArgs = outputTransfer.Data
					} else if callType == vm.AsynchronousCallBack {
						callData := txDataBuilder.NewBuilder()
						callData.Func(function)
						// will be read and removed by arwen
						callData.Bytes(parsedArgs[0])
						callData.Bytes(parsedArgs[1])
						callData.Bytes(parsedArgs[2])
						callData.Bytes(parsedArgs[3])
						// return code
						callData.Bytes(parsedArgs[4])
						// testing framework info
						callData.Bytes(big.NewInt(int64(crossShardEdge.Type)).Bytes())
						callData.Int64(int64(crossShardEdge.GasUsedByCallback))

						encodedArgs = callData.ToBytes()
					}
					crossShardCallsQueue.Enqueue(outputTransfer.SenderAddress, crossShardEdge.To, callType, encodedArgs)
					// "delete" info so we use the output transfer only once
					outputTransfer.SenderAddress = nil
				}
			}
		}
	}
}

func extractStores(vmOutput *vmcommon.VMOutput, storage map[string]map[string][]byte) {
	for _, outputAccount := range vmOutput.OutputAccounts {
		for _, storageUpdate := range outputAccount.StorageUpdates {
			accountStorage := storage[string(outputAccount.Address)]
			if accountStorage == nil {
				accountStorage = make(map[string][]byte)
				storage[string(outputAccount.Address)] = accountStorage
			}
			storage[string(outputAccount.Address)][string(storageUpdate.Offset)] = storageUpdate.Data
		}
	}
}

func extractGasUsedPerContract(vmOutput *vmcommon.VMOutput, gasUsedPerContract map[string]uint64) {
	for _, outputAccount := range vmOutput.OutputAccounts {
		if _, ok := gasUsedPerContract[string(outputAccount.Address)]; !ok {
			gasUsedPerContract[string(outputAccount.Address)] = 0
		}
		gasUsedPerContract[string(outputAccount.Address)] += outputAccount.GasUsed
	}
}

type ReturnDataItem struct {
	contractAndFunction          string
	gasProvided                  uint64
	gasRemaining                 uint64
	callID                       []byte
	callbackAsyncInitiatorCallID []byte
	isCrossShard                 bool
}

func checkReturnDataWithGasValuesForGraphTesting(t testing.TB, expectedReturnData [][]byte, returnData [][]byte) {
	processedReturnData := make([]*ReturnDataItem, 0)
	argParser := parsers.NewCallArgsParser()

	// eliminte from the final return data the gas used for callback arguments
	// in order to be able to compare them with the provided return data
	for i := 0; i < len(returnData); i++ {
		retDataItem := returnData[i]
		if len(retDataItem) == 1 && retDataItem[0] == test.Callback {
			i++ // jump over next item
			continue
		}

		_, parsedRetData, _ := argParser.ParseData(string(retDataItem))
		isCrossShard, _ := strconv.ParseBool(string(parsedRetData[5]))
		processedReturnData = append(processedReturnData, &ReturnDataItem{
			contractAndFunction:          string(parsedRetData[0]),
			gasProvided:                  big.NewInt(0).SetBytes(parsedRetData[1]).Uint64(),
			gasRemaining:                 big.NewInt(0).SetBytes(parsedRetData[2]).Uint64(),
			callID:                       parsedRetData[3],
			callbackAsyncInitiatorCallID: parsedRetData[4],
			isCrossShard:                 isCrossShard,
		})
	}

	require.Equal(t, len(expectedReturnData), len(processedReturnData), "ReturnData length")
	for idx := range expectedReturnData {
		_, expRetData, _ := argParser.ParseData(string(expectedReturnData[idx]))
		actualReturnData := processedReturnData[idx]
		expectedContractAndFunction := string(expRetData[0])
		require.Equal(t, expectedContractAndFunction, actualReturnData.contractAndFunction, "ReturnData - Call")
		expectedGasLimitForCall := big.NewInt(0).SetBytes(expRetData[1]).Uint64()
		require.Equal(t, int(expectedGasLimitForCall), int(actualReturnData.gasProvided), fmt.Sprintf("ReturnData - Gas Limit for '%s'", expRetData[0]))
		expectedGasRemainingForCall := big.NewInt(0).SetBytes(expRetData[2]).Uint64()
		require.Equal(t, int(expectedGasRemainingForCall), int(actualReturnData.gasRemaining), fmt.Sprintf("ReturnData - Gas Remaining for '%s'", expRetData[0]))
	}
}

func CheckUsedGasPerContract(t testing.TB, expectedGasUsagePerContract map[string]uint64, gasUsedPerContract map[string]uint64) {
	for expectedContract, expectedGas := range expectedGasUsagePerContract {
		require.Equal(t, int(expectedGas), int(gasUsedPerContract[expectedContract]), fmt.Sprintf("Used gas for contract %s", expectedContract))
	}
}
