package contexts

import (
	"github.com/ElrondNetwork/elrond-go-core/data/vm"
	logger "github.com/ElrondNetwork/elrond-go-logger"
)

func (context *asyncContext) NotifyChildIsComplete(callID []byte, gasToAccumulate uint64) error {
	if logAsync.GetLevel() == logger.LogTrace {
		logAsync.Trace("NofityChildIsComplete")
		logAsync.Trace("", "address", string(context.address))
		logAsync.Trace("", "callID", context.callID) // DebugCallIDAsString
		logAsync.Trace("", "callerAddr", string(context.callerAddr))
		logAsync.Trace("", "callerCallID", context.callerCallID)
		logAsync.Trace("", "callID", callID)
		logAsync.Trace("", "gasToAccumulate", gasToAccumulate)
	}

	context.completeChild(callID, gasToAccumulate)

	if !context.IsComplete() {
		return context.Save()
	}

	return context.complete()
}

func (context *asyncContext) completeChild(callID []byte, gasToAccumulate uint64) error {
	return context.CompleteChildConditional(true, callID, gasToAccumulate)
}

func (context *asyncContext) CompleteChildConditional(isComplete bool, callID []byte, gasToAccumulate uint64) error {
	if !isComplete {
		return nil
	}
	context.DecrementCallsCounter()
	context.accumulateGas(gasToAccumulate)
	if callID != nil {
		err := context.DeleteAsyncCallAndCleanGroup(callID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (context *asyncContext) complete() error {
	// There are no more callbacks to return from other shards. The context can
	// be deleted from storage.
	err := context.DeleteFromAddress(context.address)
	if err != nil {
		return err
	}

	// if we reached first call, stop notification chain
	if context.IsFirstCall() {
		return nil
	}

	currentCallID := context.GetCallID()
	if context.callType == vm.AsynchronousCall {
		vmOutput := context.childResults
		isComplete, _, err := context.callCallback(currentCallID, vmOutput, nil)
		if err != nil {
			return err
		}
		if isComplete {
			return context.NotifyChildIsComplete(currentCallID, 0)
		}
	} else if context.callType == vm.AsynchronousCallBack {
		currentCallID := context.GetCallerCallID()
		context.LoadParentContext()
		return context.NotifyChildIsComplete(currentCallID, context.gasAccumulated)
	} else if context.callType == vm.DirectCall {
		context.LoadParentContext()
		return context.NotifyChildIsComplete(nil, context.gasAccumulated)
	}

	return nil
}
