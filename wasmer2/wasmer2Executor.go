package wasmer2

import (
	"unsafe"

	vmcommon "github.com/multiversx/mx-chain-vm-common-go"
	"github.com/multiversx/mx-chain-vm-go/executor"
)

var _ executor.Executor = (*Wasmer2Executor)(nil)

// WasmerExecutor oversees the creation of Wasmer instances and execution.
type Wasmer2Executor struct {
	cgoExecutor *cWasmerExecutorT

	vmHookPointers *cWasmerVmHookPointers
	vmHooks        executor.VMHooks
	vmHooksPtr     uintptr
	vmHooksPtrPtr  unsafe.Pointer

	opcodeCost *OpcodeCost
}

// CreateExecutor creates a new wasmer executor.
func CreateExecutor() (*Wasmer2Executor, error) {
	vmHookPointers := populateCgoFunctionPointers()
	localPtr := uintptr(unsafe.Pointer(vmHookPointers))
	localPtrPtr := unsafe.Pointer(&localPtr)

	var c_executor *cWasmerExecutorT

	var result = cWasmerNewExecutor(
		&c_executor,
		localPtrPtr,
	)

	if result != cWasmerOk {
		return nil, newWrappedError(ErrFailedInstantiation)
	}

	executor := &Wasmer2Executor{
		cgoExecutor:    c_executor,
		vmHookPointers: vmHookPointers,
	}

	return executor, nil
}

// SetOpcodeCosts sets gas costs globally inside the Wasmer executor.
func (wasmerExecutor *Wasmer2Executor) SetOpcodeCosts(wasmOps *executor.WASMOpcodeCost) {
	// extract only wasmer2 opcodes
	wasmerExecutor.opcodeCost = wasmerExecutor.extractOpcodeCost(wasmOps)
	cWasmerExecutorSetOpcodeCost(
		wasmerExecutor.cgoExecutor,
		(*cWasmerOpcodeCostT)(unsafe.Pointer(wasmerExecutor.opcodeCost)),
	)
}

// SetRkyvSerializationEnabled controls a Wasmer flag.
func (wasmerExecutor *Wasmer2Executor) SetRkyvSerializationEnabled(enabled bool) {
}

// SetSIGSEGVPassthrough controls a Wasmer flag.
func (wasmerExecutor *Wasmer2Executor) SetSIGSEGVPassthrough() {
}

func (wasmerExecutor *Wasmer2Executor) FunctionNames() vmcommon.FunctionNames {
	return functionNames
}

// NewInstanceWithOptions creates a new Wasmer instance from WASM bytecode,
// respecting the provided options
func (wasmerExecutor *Wasmer2Executor) NewInstanceWithOptions(
	contractCode []byte,
	options executor.CompilationOptions,
) (executor.Instance, error) {
	var c_instance *cWasmerInstanceT

	if len(contractCode) == 0 {
		return nil, newWrappedError(ErrInvalidBytecode)
	}

	cOptions := unsafe.Pointer(&options)
	var compileResult = cWasmerInstantiateWithOptions(
		wasmerExecutor.cgoExecutor,
		&c_instance,
		(*cUchar)(unsafe.Pointer(&contractCode[0])),
		cUint(len(contractCode)),
		(*cWasmerCompilationOptions)(cOptions),
	)

	if compileResult != cWasmerOk {
		return nil, newWrappedError(ErrFailedInstantiation)
	}

	return newInstance(c_instance)
}

// NewInstanceFromCompiledCodeWithOptions creates a new Wasmer instance from
// precompiled machine code, respecting the provided options
func (wasmerExecutor *Wasmer2Executor) NewInstanceFromCompiledCodeWithOptions(
	compiledCode []byte,
	options executor.CompilationOptions,
) (executor.Instance, error) {
	var c_instance *cWasmerInstanceT

	if len(compiledCode) == 0 {
		return nil, newWrappedError(ErrInvalidBytecode)
	}

	cOptions := unsafe.Pointer(&options)
	var compileResult = cWasmerInstanceFromCache(
		wasmerExecutor.cgoExecutor,
		&c_instance,
		(*cUchar)(unsafe.Pointer(&compiledCode[0])),
		cUint32T(len(compiledCode)),
		(*cWasmerCompilationOptions)(cOptions),
	)

	if compileResult != cWasmerOk {
		return nil, newWrappedError(ErrFailedInstantiation)
	}

	return newInstance(c_instance)
}

// IsInterfaceNil returns true if underlying object is nil
func (wasmerExecutor *Wasmer2Executor) IsInterfaceNil() bool {
	return wasmerExecutor == nil
}

// InitVMHooks inits the VM hooks
func (wasmerExecutor *Wasmer2Executor) initVMHooks(vmHooks executor.VMHooks) {
	wasmerExecutor.vmHooks = vmHooks
	localPtr := uintptr(unsafe.Pointer(&wasmerExecutor.vmHooks))
	wasmerExecutor.vmHooksPtr = localPtr
	wasmerExecutor.vmHooksPtrPtr = unsafe.Pointer(&localPtr)
	cWasmerExecutorContextDataSet(wasmerExecutor.cgoExecutor, wasmerExecutor.vmHooksPtrPtr)
}

func (wasmerExecutor *Wasmer2Executor) extractOpcodeCost(wasmOps *executor.WASMOpcodeCost) *OpcodeCost {
	return &OpcodeCost{
		AtomicFence:               wasmOps.AtomicFence,
		Block:                     wasmOps.Block,
		Br:                        wasmOps.Br,
		BrIf:                      wasmOps.BrIf,
		BrTable:                   wasmOps.BrTable,
		Call:                      wasmOps.Call,
		CallIndirect:              wasmOps.CallIndirect,
		Catch:                     wasmOps.Catch,
		CatchAll:                  wasmOps.CatchAll,
		DataDrop:                  wasmOps.DataDrop,
		Delegate:                  wasmOps.Delegate,
		Drop:                      wasmOps.Drop,
		ElemDrop:                  wasmOps.ElemDrop,
		Else:                      wasmOps.Else,
		End:                       wasmOps.End,
		F32Abs:                    wasmOps.F32Abs,
		F32Add:                    wasmOps.F32Add,
		F32Ceil:                   wasmOps.F32Ceil,
		F32Const:                  wasmOps.F32Const,
		F32ConvertI32S:            wasmOps.F32ConvertI32S,
		F32ConvertI32U:            wasmOps.F32ConvertI32U,
		F32ConvertI64S:            wasmOps.F32ConvertI64S,
		F32ConvertI64U:            wasmOps.F32ConvertI64U,
		F32Copysign:               wasmOps.F32Copysign,
		F32DemoteF64:              wasmOps.F32DemoteF64,
		F32Div:                    wasmOps.F32Div,
		F32Eq:                     wasmOps.F32Eq,
		F32Floor:                  wasmOps.F32Floor,
		F32Ge:                     wasmOps.F32Ge,
		F32Gt:                     wasmOps.F32Gt,
		F32Le:                     wasmOps.F32Le,
		F32Load:                   wasmOps.F32Load,
		F32Lt:                     wasmOps.F32Lt,
		F32Max:                    wasmOps.F32Max,
		F32Min:                    wasmOps.F32Min,
		F32Mul:                    wasmOps.F32Mul,
		F32Ne:                     wasmOps.F32Ne,
		F32Nearest:                wasmOps.F32Nearest,
		F32Neg:                    wasmOps.F32Neg,
		F32ReinterpretI32:         wasmOps.F32ReinterpretI32,
		F32Sqrt:                   wasmOps.F32Sqrt,
		F32Store:                  wasmOps.F32Store,
		F32Sub:                    wasmOps.F32Sub,
		F32Trunc:                  wasmOps.F32Trunc,
		F32x4Abs:                  wasmOps.F32x4Abs,
		F32x4Add:                  wasmOps.F32x4Add,
		F32x4Ceil:                 wasmOps.F32x4Ceil,
		F32x4ConvertI32x4S:        wasmOps.F32x4ConvertI32x4S,
		F32x4ConvertI32x4U:        wasmOps.F32x4ConvertI32x4U,
		F32x4DemoteF64x2Zero:      wasmOps.F32x4DemoteF64x2Zero,
		F32x4Div:                  wasmOps.F32x4Div,
		F32x4Eq:                   wasmOps.F32x4Eq,
		F32x4ExtractLane:          wasmOps.F32x4ExtractLane,
		F32x4Floor:                wasmOps.F32x4Floor,
		F32x4Ge:                   wasmOps.F32x4Ge,
		F32x4Gt:                   wasmOps.F32x4Gt,
		F32x4Le:                   wasmOps.F32x4Le,
		F32x4Lt:                   wasmOps.F32x4Lt,
		F32x4Max:                  wasmOps.F32x4Max,
		F32x4Min:                  wasmOps.F32x4Min,
		F32x4Mul:                  wasmOps.F32x4Mul,
		F32x4Ne:                   wasmOps.F32x4Ne,
		F32x4Nearest:              wasmOps.F32x4Nearest,
		F32x4Neg:                  wasmOps.F32x4Neg,
		F32x4PMax:                 wasmOps.F32x4PMax,
		F32x4PMin:                 wasmOps.F32x4PMin,
		F32x4ReplaceLane:          wasmOps.F32x4ReplaceLane,
		F32x4Splat:                wasmOps.F32x4Splat,
		F32x4Sqrt:                 wasmOps.F32x4Sqrt,
		F32x4Sub:                  wasmOps.F32x4Sub,
		F32x4Trunc:                wasmOps.F32x4Trunc,
		F64Abs:                    wasmOps.F64Abs,
		F64Add:                    wasmOps.F64Add,
		F64Ceil:                   wasmOps.F64Ceil,
		F64Const:                  wasmOps.F64Const,
		F64ConvertI32S:            wasmOps.F64ConvertI32S,
		F64ConvertI32U:            wasmOps.F64ConvertI32U,
		F64ConvertI64S:            wasmOps.F64ConvertI64S,
		F64ConvertI64U:            wasmOps.F64ConvertI64U,
		F64Copysign:               wasmOps.F64Copysign,
		F64Div:                    wasmOps.F64Div,
		F64Eq:                     wasmOps.F64Eq,
		F64Floor:                  wasmOps.F64Floor,
		F64Ge:                     wasmOps.F64Ge,
		F64Gt:                     wasmOps.F64Gt,
		F64Le:                     wasmOps.F64Le,
		F64Load:                   wasmOps.F64Load,
		F64Lt:                     wasmOps.F64Lt,
		F64Max:                    wasmOps.F64Max,
		F64Min:                    wasmOps.F64Min,
		F64Mul:                    wasmOps.F64Mul,
		F64Ne:                     wasmOps.F64Ne,
		F64Nearest:                wasmOps.F64Nearest,
		F64Neg:                    wasmOps.F64Neg,
		F64PromoteF32:             wasmOps.F64PromoteF32,
		F64ReinterpretI64:         wasmOps.F64ReinterpretI64,
		F64Sqrt:                   wasmOps.F64Sqrt,
		F64Store:                  wasmOps.F64Store,
		F64Sub:                    wasmOps.F64Sub,
		F64Trunc:                  wasmOps.F64Trunc,
		F64x2Abs:                  wasmOps.F64x2Abs,
		F64x2Add:                  wasmOps.F64x2Add,
		F64x2Ceil:                 wasmOps.F64x2Ceil,
		F64x2ConvertLowI32x4S:     wasmOps.F64x2ConvertLowI32x4S,
		F64x2ConvertLowI32x4U:     wasmOps.F64x2ConvertLowI32x4U,
		F64x2Div:                  wasmOps.F64x2Div,
		F64x2Eq:                   wasmOps.F64x2Eq,
		F64x2ExtractLane:          wasmOps.F64x2ExtractLane,
		F64x2Floor:                wasmOps.F64x2Floor,
		F64x2Ge:                   wasmOps.F64x2Ge,
		F64x2Gt:                   wasmOps.F64x2Gt,
		F64x2Le:                   wasmOps.F64x2Le,
		F64x2Lt:                   wasmOps.F64x2Lt,
		F64x2Max:                  wasmOps.F64x2Max,
		F64x2Min:                  wasmOps.F64x2Min,
		F64x2Mul:                  wasmOps.F64x2Mul,
		F64x2Ne:                   wasmOps.F64x2Ne,
		F64x2Nearest:              wasmOps.F64x2Nearest,
		F64x2Neg:                  wasmOps.F64x2Neg,
		F64x2PMax:                 wasmOps.F64x2PMax,
		F64x2PMin:                 wasmOps.F64x2PMin,
		F64x2PromoteLowF32x4:      wasmOps.F64x2PromoteLowF32x4,
		F64x2ReplaceLane:          wasmOps.F64x2ReplaceLane,
		F64x2Splat:                wasmOps.F64x2Splat,
		F64x2Sqrt:                 wasmOps.F64x2Sqrt,
		F64x2Sub:                  wasmOps.F64x2Sub,
		F64x2Trunc:                wasmOps.F64x2Trunc,
		GlobalGet:                 wasmOps.GlobalGet,
		GlobalSet:                 wasmOps.GlobalSet,
		I16x8Abs:                  wasmOps.I16x8Abs,
		I16x8Add:                  wasmOps.I16x8Add,
		I16x8AddSatS:              wasmOps.I16x8AddSatS,
		I16x8AddSatU:              wasmOps.I16x8AddSatU,
		I16x8AllTrue:              wasmOps.I16x8AllTrue,
		I16x8Bitmask:              wasmOps.I16x8Bitmask,
		I16x8Eq:                   wasmOps.I16x8Eq,
		I16x8ExtAddPairwiseI8x16S: wasmOps.I16x8ExtAddPairwiseI8x16S,
		I16x8ExtAddPairwiseI8x16U: wasmOps.I16x8ExtAddPairwiseI8x16U,
		I16x8ExtMulHighI8x16S:     wasmOps.I16x8ExtMulHighI8x16S,
		I16x8ExtMulHighI8x16U:     wasmOps.I16x8ExtMulHighI8x16U,
		I16x8ExtMulLowI8x16S:      wasmOps.I16x8ExtMulLowI8x16S,
		I16x8ExtMulLowI8x16U:      wasmOps.I16x8ExtMulLowI8x16U,
		I16x8ExtendHighI8x16S:     wasmOps.I16x8ExtendHighI8x16S,
		I16x8ExtendHighI8x16U:     wasmOps.I16x8ExtendHighI8x16U,
		I16x8ExtendLowI8x16S:      wasmOps.I16x8ExtendLowI8x16S,
		I16x8ExtendLowI8x16U:      wasmOps.I16x8ExtendLowI8x16U,
		I16x8ExtractLaneS:         wasmOps.I16x8ExtractLaneS,
		I16x8ExtractLaneU:         wasmOps.I16x8ExtractLaneU,
		I16x8GeS:                  wasmOps.I16x8GeS,
		I16x8GeU:                  wasmOps.I16x8GeU,
		I16x8GtS:                  wasmOps.I16x8GtS,
		I16x8GtU:                  wasmOps.I16x8GtU,
		I16x8LeS:                  wasmOps.I16x8LeS,
		I16x8LeU:                  wasmOps.I16x8LeU,
		I16x8LtS:                  wasmOps.I16x8LtS,
		I16x8LtU:                  wasmOps.I16x8LtU,
		I16x8MaxS:                 wasmOps.I16x8MaxS,
		I16x8MaxU:                 wasmOps.I16x8MaxU,
		I16x8MinS:                 wasmOps.I16x8MinS,
		I16x8MinU:                 wasmOps.I16x8MinU,
		I16x8Mul:                  wasmOps.I16x8Mul,
		I16x8NarrowI32x4S:         wasmOps.I16x8NarrowI32x4S,
		I16x8NarrowI32x4U:         wasmOps.I16x8NarrowI32x4U,
		I16x8Ne:                   wasmOps.I16x8Ne,
		I16x8Neg:                  wasmOps.I16x8Neg,
		I16x8Q15MulrSatS:          wasmOps.I16x8Q15MulrSatS,
		I16x8ReplaceLane:          wasmOps.I16x8ReplaceLane,
		I16x8RoundingAverageU:     wasmOps.I16x8RoundingAverageU,
		I16x8Shl:                  wasmOps.I16x8Shl,
		I16x8ShrS:                 wasmOps.I16x8ShrS,
		I16x8ShrU:                 wasmOps.I16x8ShrU,
		I16x8Splat:                wasmOps.I16x8Splat,
		I16x8Sub:                  wasmOps.I16x8Sub,
		I16x8SubSatS:              wasmOps.I16x8SubSatS,
		I16x8SubSatU:              wasmOps.I16x8SubSatU,
		I32Add:                    wasmOps.I32Add,
		I32And:                    wasmOps.I32And,
		I32AtomicLoad:             wasmOps.I32AtomicLoad,
		I32AtomicLoad16U:          wasmOps.I32AtomicLoad16U,
		I32AtomicLoad8U:           wasmOps.I32AtomicLoad8U,
		I32AtomicRmw16AddU:        wasmOps.I32AtomicRmw16AddU,
		I32AtomicRmw16AndU:        wasmOps.I32AtomicRmw16AndU,
		I32AtomicRmw16CmpxchgU:    wasmOps.I32AtomicRmw16CmpxchgU,
		I32AtomicRmw16OrU:         wasmOps.I32AtomicRmw16OrU,
		I32AtomicRmw16SubU:        wasmOps.I32AtomicRmw16SubU,
		I32AtomicRmw16XchgU:       wasmOps.I32AtomicRmw16XchgU,
		I32AtomicRmw16XorU:        wasmOps.I32AtomicRmw16XorU,
		I32AtomicRmw8AddU:         wasmOps.I32AtomicRmw8AddU,
		I32AtomicRmw8AndU:         wasmOps.I32AtomicRmw8AndU,
		I32AtomicRmw8CmpxchgU:     wasmOps.I32AtomicRmw8CmpxchgU,
		I32AtomicRmw8OrU:          wasmOps.I32AtomicRmw8OrU,
		I32AtomicRmw8SubU:         wasmOps.I32AtomicRmw8SubU,
		I32AtomicRmw8XchgU:        wasmOps.I32AtomicRmw8XchgU,
		I32AtomicRmw8XorU:         wasmOps.I32AtomicRmw8XorU,
		I32AtomicRmwAdd:           wasmOps.I32AtomicRmwAdd,
		I32AtomicRmwAnd:           wasmOps.I32AtomicRmwAnd,
		I32AtomicRmwCmpxchg:       wasmOps.I32AtomicRmwCmpxchg,
		I32AtomicRmwOr:            wasmOps.I32AtomicRmwOr,
		I32AtomicRmwSub:           wasmOps.I32AtomicRmwSub,
		I32AtomicRmwXchg:          wasmOps.I32AtomicRmwXchg,
		I32AtomicRmwXor:           wasmOps.I32AtomicRmwXor,
		I32AtomicStore:            wasmOps.I32AtomicStore,
		I32AtomicStore16:          wasmOps.I32AtomicStore16,
		I32AtomicStore8:           wasmOps.I32AtomicStore8,
		I32Clz:                    wasmOps.I32Clz,
		I32Const:                  wasmOps.I32Const,
		I32Ctz:                    wasmOps.I32Ctz,
		I32DivS:                   wasmOps.I32DivS,
		I32DivU:                   wasmOps.I32DivU,
		I32Eq:                     wasmOps.I32Eq,
		I32Eqz:                    wasmOps.I32Eqz,
		I32Extend16S:              wasmOps.I32Extend16S,
		I32Extend8S:               wasmOps.I32Extend8S,
		I32GeS:                    wasmOps.I32GeS,
		I32GeU:                    wasmOps.I32GeU,
		I32GtS:                    wasmOps.I32GtS,
		I32GtU:                    wasmOps.I32GtU,
		I32LeS:                    wasmOps.I32LeS,
		I32LeU:                    wasmOps.I32LeU,
		I32Load:                   wasmOps.I32Load,
		I32Load16S:                wasmOps.I32Load16S,
		I32Load16U:                wasmOps.I32Load16U,
		I32Load8S:                 wasmOps.I32Load8S,
		I32Load8U:                 wasmOps.I32Load8U,
		I32LtS:                    wasmOps.I32LtS,
		I32LtU:                    wasmOps.I32LtU,
		I32Mul:                    wasmOps.I32Mul,
		I32Ne:                     wasmOps.I32Ne,
		I32Or:                     wasmOps.I32Or,
		I32Popcnt:                 wasmOps.I32Popcnt,
		I32ReinterpretF32:         wasmOps.I32ReinterpretF32,
		I32RemS:                   wasmOps.I32RemS,
		I32RemU:                   wasmOps.I32RemU,
		I32Rotl:                   wasmOps.I32Rotl,
		I32Rotr:                   wasmOps.I32Rotr,
		I32Shl:                    wasmOps.I32Shl,
		I32ShrS:                   wasmOps.I32ShrS,
		I32ShrU:                   wasmOps.I32ShrU,
		I32Store:                  wasmOps.I32Store,
		I32Store16:                wasmOps.I32Store16,
		I32Store8:                 wasmOps.I32Store8,
		I32Sub:                    wasmOps.I32Sub,
		I32TruncF32S:              wasmOps.I32TruncF32S,
		I32TruncF32U:              wasmOps.I32TruncF32U,
		I32TruncF64S:              wasmOps.I32TruncF64S,
		I32TruncF64U:              wasmOps.I32TruncF64U,
		I32TruncSatF32S:           wasmOps.I32TruncSatF32S,
		I32TruncSatF32U:           wasmOps.I32TruncSatF32U,
		I32TruncSatF64S:           wasmOps.I32TruncSatF64S,
		I32TruncSatF64U:           wasmOps.I32TruncSatF64U,
		I32WrapI64:                wasmOps.I32WrapI64,
		I32Xor:                    wasmOps.I32Xor,
		I32x4Abs:                  wasmOps.I32x4Abs,
		I32x4Add:                  wasmOps.I32x4Add,
		I32x4AllTrue:              wasmOps.I32x4AllTrue,
		I32x4Bitmask:              wasmOps.I32x4Bitmask,
		I32x4DotI16x8S:            wasmOps.I32x4DotI16x8S,
		I32x4Eq:                   wasmOps.I32x4Eq,
		I32x4ExtAddPairwiseI16x8S: wasmOps.I32x4ExtAddPairwiseI16x8S,
		I32x4ExtAddPairwiseI16x8U: wasmOps.I32x4ExtAddPairwiseI16x8U,
		I32x4ExtMulHighI16x8S:     wasmOps.I32x4ExtMulHighI16x8S,
		I32x4ExtMulHighI16x8U:     wasmOps.I32x4ExtMulHighI16x8U,
		I32x4ExtMulLowI16x8S:      wasmOps.I32x4ExtMulLowI16x8S,
		I32x4ExtMulLowI16x8U:      wasmOps.I32x4ExtMulLowI16x8U,
		I32x4ExtendHighI16x8S:     wasmOps.I32x4ExtendHighI16x8S,
		I32x4ExtendHighI16x8U:     wasmOps.I32x4ExtendHighI16x8U,
		I32x4ExtendLowI16x8S:      wasmOps.I32x4ExtendLowI16x8S,
		I32x4ExtendLowI16x8U:      wasmOps.I32x4ExtendLowI16x8U,
		I32x4ExtractLane:          wasmOps.I32x4ExtractLane,
		I32x4GeS:                  wasmOps.I32x4GeS,
		I32x4GeU:                  wasmOps.I32x4GeU,
		I32x4GtS:                  wasmOps.I32x4GtS,
		I32x4GtU:                  wasmOps.I32x4GtU,
		I32x4LeS:                  wasmOps.I32x4LeS,
		I32x4LeU:                  wasmOps.I32x4LeU,
		I32x4LtS:                  wasmOps.I32x4LtS,
		I32x4LtU:                  wasmOps.I32x4LtU,
		I32x4MaxS:                 wasmOps.I32x4MaxS,
		I32x4MaxU:                 wasmOps.I32x4MaxU,
		I32x4MinS:                 wasmOps.I32x4MinS,
		I32x4MinU:                 wasmOps.I32x4MinU,
		I32x4Mul:                  wasmOps.I32x4Mul,
		I32x4Ne:                   wasmOps.I32x4Ne,
		I32x4Neg:                  wasmOps.I32x4Neg,
		I32x4ReplaceLane:          wasmOps.I32x4ReplaceLane,
		I32x4Shl:                  wasmOps.I32x4Shl,
		I32x4ShrS:                 wasmOps.I32x4ShrS,
		I32x4ShrU:                 wasmOps.I32x4ShrU,
		I32x4Splat:                wasmOps.I32x4Splat,
		I32x4Sub:                  wasmOps.I32x4Sub,
		I32x4TruncSatF32x4S:       wasmOps.I32x4TruncSatF32x4S,
		I32x4TruncSatF32x4U:       wasmOps.I32x4TruncSatF32x4U,
		I32x4TruncSatF64x2SZero:   wasmOps.I32x4TruncSatF64x2SZero,
		I32x4TruncSatF64x2UZero:   wasmOps.I32x4TruncSatF64x2UZero,
		I64Add:                    wasmOps.I64Add,
		I64And:                    wasmOps.I64And,
		I64AtomicLoad:             wasmOps.I64AtomicLoad,
		I64AtomicLoad16U:          wasmOps.I64AtomicLoad16U,
		I64AtomicLoad32U:          wasmOps.I64AtomicLoad32U,
		I64AtomicLoad8U:           wasmOps.I64AtomicLoad8U,
		I64AtomicRmw16AddU:        wasmOps.I64AtomicRmw16AddU,
		I64AtomicRmw16AndU:        wasmOps.I64AtomicRmw16AndU,
		I64AtomicRmw16CmpxchgU:    wasmOps.I64AtomicRmw16CmpxchgU,
		I64AtomicRmw16OrU:         wasmOps.I64AtomicRmw16OrU,
		I64AtomicRmw16SubU:        wasmOps.I64AtomicRmw16SubU,
		I64AtomicRmw16XchgU:       wasmOps.I64AtomicRmw16XchgU,
		I64AtomicRmw16XorU:        wasmOps.I64AtomicRmw16XorU,
		I64AtomicRmw32AddU:        wasmOps.I64AtomicRmw32AddU,
		I64AtomicRmw32AndU:        wasmOps.I64AtomicRmw32AndU,
		I64AtomicRmw32CmpxchgU:    wasmOps.I64AtomicRmw32CmpxchgU,
		I64AtomicRmw32OrU:         wasmOps.I64AtomicRmw32OrU,
		I64AtomicRmw32SubU:        wasmOps.I64AtomicRmw32SubU,
		I64AtomicRmw32XchgU:       wasmOps.I64AtomicRmw32XchgU,
		I64AtomicRmw32XorU:        wasmOps.I64AtomicRmw32XorU,
		I64AtomicRmw8AddU:         wasmOps.I64AtomicRmw8AddU,
		I64AtomicRmw8AndU:         wasmOps.I64AtomicRmw8AndU,
		I64AtomicRmw8CmpxchgU:     wasmOps.I64AtomicRmw8CmpxchgU,
		I64AtomicRmw8OrU:          wasmOps.I64AtomicRmw8OrU,
		I64AtomicRmw8SubU:         wasmOps.I64AtomicRmw8SubU,
		I64AtomicRmw8XchgU:        wasmOps.I64AtomicRmw8XchgU,
		I64AtomicRmw8XorU:         wasmOps.I64AtomicRmw8XorU,
		I64AtomicRmwAdd:           wasmOps.I64AtomicRmwAdd,
		I64AtomicRmwAnd:           wasmOps.I64AtomicRmwAnd,
		I64AtomicRmwCmpxchg:       wasmOps.I64AtomicRmwCmpxchg,
		I64AtomicRmwOr:            wasmOps.I64AtomicRmwOr,
		I64AtomicRmwSub:           wasmOps.I64AtomicRmwSub,
		I64AtomicRmwXchg:          wasmOps.I64AtomicRmwXchg,
		I64AtomicRmwXor:           wasmOps.I64AtomicRmwXor,
		I64AtomicStore:            wasmOps.I64AtomicStore,
		I64AtomicStore16:          wasmOps.I64AtomicStore16,
		I64AtomicStore32:          wasmOps.I64AtomicStore32,
		I64AtomicStore8:           wasmOps.I64AtomicStore8,
		I64Clz:                    wasmOps.I64Clz,
		I64Const:                  wasmOps.I64Const,
		I64Ctz:                    wasmOps.I64Ctz,
		I64DivS:                   wasmOps.I64DivS,
		I64DivU:                   wasmOps.I64DivU,
		I64Eq:                     wasmOps.I64Eq,
		I64Eqz:                    wasmOps.I64Eqz,
		I64Extend16S:              wasmOps.I64Extend16S,
		I64Extend32S:              wasmOps.I64Extend32S,
		I64Extend8S:               wasmOps.I64Extend8S,
		I64ExtendI32S:             wasmOps.I64ExtendI32S,
		I64ExtendI32U:             wasmOps.I64ExtendI32U,
		I64GeS:                    wasmOps.I64GeS,
		I64GeU:                    wasmOps.I64GeU,
		I64GtS:                    wasmOps.I64GtS,
		I64GtU:                    wasmOps.I64GtU,
		I64LeS:                    wasmOps.I64LeS,
		I64LeU:                    wasmOps.I64LeU,
		I64Load:                   wasmOps.I64Load,
		I64Load16S:                wasmOps.I64Load16S,
		I64Load16U:                wasmOps.I64Load16U,
		I64Load32S:                wasmOps.I64Load32S,
		I64Load32U:                wasmOps.I64Load32U,
		I64Load8S:                 wasmOps.I64Load8S,
		I64Load8U:                 wasmOps.I64Load8U,
		I64LtS:                    wasmOps.I64LtS,
		I64LtU:                    wasmOps.I64LtU,
		I64Mul:                    wasmOps.I64Mul,
		I64Ne:                     wasmOps.I64Ne,
		I64Or:                     wasmOps.I64Or,
		I64Popcnt:                 wasmOps.I64Popcnt,
		I64ReinterpretF64:         wasmOps.I64ReinterpretF64,
		I64RemS:                   wasmOps.I64RemS,
		I64RemU:                   wasmOps.I64RemU,
		I64Rotl:                   wasmOps.I64Rotl,
		I64Rotr:                   wasmOps.I64Rotr,
		I64Shl:                    wasmOps.I64Shl,
		I64ShrS:                   wasmOps.I64ShrS,
		I64ShrU:                   wasmOps.I64ShrU,
		I64Store:                  wasmOps.I64Store,
		I64Store16:                wasmOps.I64Store16,
		I64Store32:                wasmOps.I64Store32,
		I64Store8:                 wasmOps.I64Store8,
		I64Sub:                    wasmOps.I64Sub,
		I64TruncF32S:              wasmOps.I64TruncF32S,
		I64TruncF32U:              wasmOps.I64TruncF32U,
		I64TruncF64S:              wasmOps.I64TruncF64S,
		I64TruncF64U:              wasmOps.I64TruncF64U,
		I64TruncSatF32S:           wasmOps.I64TruncSatF32S,
		I64TruncSatF32U:           wasmOps.I64TruncSatF32U,
		I64TruncSatF64S:           wasmOps.I64TruncSatF64S,
		I64TruncSatF64U:           wasmOps.I64TruncSatF64U,
		I64Xor:                    wasmOps.I64Xor,
		I64x2Abs:                  wasmOps.I64x2Abs,
		I64x2Add:                  wasmOps.I64x2Add,
		I64x2AllTrue:              wasmOps.I64x2AllTrue,
		I64x2Bitmask:              wasmOps.I64x2Bitmask,
		I64x2Eq:                   wasmOps.I64x2Eq,
		I64x2ExtMulHighI32x4S:     wasmOps.I64x2ExtMulHighI32x4S,
		I64x2ExtMulHighI32x4U:     wasmOps.I64x2ExtMulHighI32x4U,
		I64x2ExtMulLowI32x4S:      wasmOps.I64x2ExtMulLowI32x4S,
		I64x2ExtMulLowI32x4U:      wasmOps.I64x2ExtMulLowI32x4U,
		I64x2ExtendHighI32x4S:     wasmOps.I64x2ExtendHighI32x4S,
		I64x2ExtendHighI32x4U:     wasmOps.I64x2ExtendHighI32x4U,
		I64x2ExtendLowI32x4S:      wasmOps.I64x2ExtendLowI32x4S,
		I64x2ExtendLowI32x4U:      wasmOps.I64x2ExtendLowI32x4U,
		I64x2ExtractLane:          wasmOps.I64x2ExtractLane,
		I64x2GeS:                  wasmOps.I64x2GeS,
		I64x2GtS:                  wasmOps.I64x2GtS,
		I64x2LeS:                  wasmOps.I64x2LeS,
		I64x2LtS:                  wasmOps.I64x2LtS,
		I64x2Mul:                  wasmOps.I64x2Mul,
		I64x2Ne:                   wasmOps.I64x2Ne,
		I64x2Neg:                  wasmOps.I64x2Neg,
		I64x2ReplaceLane:          wasmOps.I64x2ReplaceLane,
		I64x2Shl:                  wasmOps.I64x2Shl,
		I64x2ShrS:                 wasmOps.I64x2ShrS,
		I64x2ShrU:                 wasmOps.I64x2ShrU,
		I64x2Splat:                wasmOps.I64x2Splat,
		I64x2Sub:                  wasmOps.I64x2Sub,
		I8x16Abs:                  wasmOps.I8x16Abs,
		I8x16Add:                  wasmOps.I8x16Add,
		I8x16AddSatS:              wasmOps.I8x16AddSatS,
		I8x16AddSatU:              wasmOps.I8x16AddSatU,
		I8x16AllTrue:              wasmOps.I8x16AllTrue,
		I8x16Bitmask:              wasmOps.I8x16Bitmask,
		I8x16Eq:                   wasmOps.I8x16Eq,
		I8x16ExtractLaneS:         wasmOps.I8x16ExtractLaneS,
		I8x16ExtractLaneU:         wasmOps.I8x16ExtractLaneU,
		I8x16GeS:                  wasmOps.I8x16GeS,
		I8x16GeU:                  wasmOps.I8x16GeU,
		I8x16GtS:                  wasmOps.I8x16GtS,
		I8x16GtU:                  wasmOps.I8x16GtU,
		I8x16LeS:                  wasmOps.I8x16LeS,
		I8x16LeU:                  wasmOps.I8x16LeU,
		I8x16LtS:                  wasmOps.I8x16LtS,
		I8x16LtU:                  wasmOps.I8x16LtU,
		I8x16MaxS:                 wasmOps.I8x16MaxS,
		I8x16MaxU:                 wasmOps.I8x16MaxU,
		I8x16MinS:                 wasmOps.I8x16MinS,
		I8x16MinU:                 wasmOps.I8x16MinU,
		I8x16NarrowI16x8S:         wasmOps.I8x16NarrowI16x8S,
		I8x16NarrowI16x8U:         wasmOps.I8x16NarrowI16x8U,
		I8x16Ne:                   wasmOps.I8x16Ne,
		I8x16Neg:                  wasmOps.I8x16Neg,
		I8x16Popcnt:               wasmOps.I8x16Popcnt,
		I8x16ReplaceLane:          wasmOps.I8x16ReplaceLane,
		I8x16RoundingAverageU:     wasmOps.I8x16RoundingAverageU,
		I8x16Shl:                  wasmOps.I8x16Shl,
		I8x16ShrS:                 wasmOps.I8x16ShrS,
		I8x16ShrU:                 wasmOps.I8x16ShrU,
		I8x16Shuffle:              wasmOps.I8x16Shuffle,
		I8x16Splat:                wasmOps.I8x16Splat,
		I8x16Sub:                  wasmOps.I8x16Sub,
		I8x16SubSatS:              wasmOps.I8x16SubSatS,
		I8x16SubSatU:              wasmOps.I8x16SubSatU,
		I8x16Swizzle:              wasmOps.I8x16Swizzle,
		If:                        wasmOps.If,
		LocalGet:                  wasmOps.LocalGet,
		LocalSet:                  wasmOps.LocalSet,
		LocalTee:                  wasmOps.LocalTee,
		LocalAllocate:             wasmOps.LocalAllocate,
		Loop:                      wasmOps.Loop,
		MemoryAtomicNotify:        wasmOps.MemoryAtomicNotify,
		MemoryAtomicWait32:        wasmOps.MemoryAtomicWait32,
		MemoryAtomicWait64:        wasmOps.MemoryAtomicWait64,
		MemoryCopy:                wasmOps.MemoryCopy,
		MemoryFill:                wasmOps.MemoryFill,
		MemoryGrow:                wasmOps.MemoryGrow,
		MemoryInit:                wasmOps.MemoryInit,
		MemorySize:                wasmOps.MemorySize,
		Nop:                       wasmOps.Nop,
		RefFunc:                   wasmOps.RefFunc,
		RefIsNull:                 wasmOps.RefIsNull,
		RefNull:                   wasmOps.RefNull,
		Rethrow:                   wasmOps.Rethrow,
		Return:                    wasmOps.Return,
		ReturnCall:                wasmOps.ReturnCall,
		ReturnCallIndirect:        wasmOps.ReturnCallIndirect,
		Select:                    wasmOps.Select,
		TableCopy:                 wasmOps.TableCopy,
		TableFill:                 wasmOps.TableFill,
		TableGet:                  wasmOps.TableGet,
		TableGrow:                 wasmOps.TableGrow,
		TableInit:                 wasmOps.TableInit,
		TableSet:                  wasmOps.TableSet,
		TableSize:                 wasmOps.TableSize,
		Throw:                     wasmOps.Throw,
		Try:                       wasmOps.Try,
		TypedSelect:               wasmOps.TypedSelect,
		Unreachable:               wasmOps.Unreachable,
		Unwind:                    wasmOps.Unwind,
		V128And:                   wasmOps.V128And,
		V128AndNot:                wasmOps.V128AndNot,
		V128AnyTrue:               wasmOps.V128AnyTrue,
		V128Bitselect:             wasmOps.V128Bitselect,
		V128Const:                 wasmOps.V128Const,
		V128Load:                  wasmOps.V128Load,
		V128Load16Lane:            wasmOps.V128Load16Lane,
		V128Load16Splat:           wasmOps.V128Load16Splat,
		V128Load16x4S:             wasmOps.V128Load16x4S,
		V128Load16x4U:             wasmOps.V128Load16x4U,
		V128Load32Lane:            wasmOps.V128Load32Lane,
		V128Load32Splat:           wasmOps.V128Load32Splat,
		V128Load32Zero:            wasmOps.V128Load32Zero,
		V128Load32x2S:             wasmOps.V128Load32x2S,
		V128Load32x2U:             wasmOps.V128Load32x2U,
		V128Load64Lane:            wasmOps.V128Load64Lane,
		V128Load64Splat:           wasmOps.V128Load64Splat,
		V128Load64Zero:            wasmOps.V128Load64Zero,
		V128Load8Lane:             wasmOps.V128Load8Lane,
		V128Load8Splat:            wasmOps.V128Load8Splat,
		V128Load8x8S:              wasmOps.V128Load8x8S,
		V128Load8x8U:              wasmOps.V128Load8x8U,
		V128Not:                   wasmOps.V128Not,
		V128Or:                    wasmOps.V128Or,
		V128Store:                 wasmOps.V128Store,
		V128Store16Lane:           wasmOps.V128Store16Lane,
		V128Store32Lane:           wasmOps.V128Store32Lane,
		V128Store64Lane:           wasmOps.V128Store64Lane,
		V128Store8Lane:            wasmOps.V128Store8Lane,
		V128Xor:                   wasmOps.V128Xor,
	}
}

// IsInterfaceNil returns true if there is no value under the interface
func (wasmerExecutor *Wasmer2Executor) IsInterfaceNil() bool {
	return wasmerExecutor == nil
}