package wasmer

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!!!!!!!!!!!!!!!!!!!!! AUTO-GENERATED FILE !!!!!!!!!!!!!!!!!!!!!!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

// // Declare the function signatures (see [cgo](https://golang.org/cmd/cgo/)).
//
// #include <stdlib.h>
// typedef int int32_t;
//
// extern int32_t   wasmer1_bigFloatNewFromParts(void* context, int32_t integralPart, int32_t fractionalPart, int32_t exponent);
// extern int32_t   wasmer1_bigFloatNewFromFrac(void* context, long long numerator, long long denominator);
// extern int32_t   wasmer1_bigFloatNewFromSci(void* context, long long significand, long long exponent);
// extern void      wasmer1_bigFloatAdd(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigFloatSub(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigFloatMul(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigFloatDiv(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigFloatNeg(void* context, int32_t destinationHandle, int32_t opHandle);
// extern void      wasmer1_bigFloatClone(void* context, int32_t destinationHandle, int32_t opHandle);
// extern int32_t   wasmer1_bigFloatCmp(void* context, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigFloatAbs(void* context, int32_t destinationHandle, int32_t opHandle);
// extern int32_t   wasmer1_bigFloatSign(void* context, int32_t opHandle);
// extern void      wasmer1_bigFloatSqrt(void* context, int32_t destinationHandle, int32_t opHandle);
// extern void      wasmer1_bigFloatPow(void* context, int32_t destinationHandle, int32_t opHandle, int32_t exponent);
// extern void      wasmer1_bigFloatFloor(void* context, int32_t destBigIntHandle, int32_t opHandle);
// extern void      wasmer1_bigFloatCeil(void* context, int32_t destBigIntHandle, int32_t opHandle);
// extern void      wasmer1_bigFloatTruncate(void* context, int32_t destBigIntHandle, int32_t opHandle);
// extern void      wasmer1_bigFloatSetInt64(void* context, int32_t destinationHandle, long long value);
// extern int32_t   wasmer1_bigFloatIsInt(void* context, int32_t opHandle);
// extern void      wasmer1_bigFloatSetBigInt(void* context, int32_t destinationHandle, int32_t bigIntHandle);
// extern void      wasmer1_bigFloatGetConstPi(void* context, int32_t destinationHandle);
// extern void      wasmer1_bigFloatGetConstE(void* context, int32_t destinationHandle);
// extern void      wasmer1_bigIntGetUnsignedArgument(void* context, int32_t id, int32_t destinationHandle);
// extern void      wasmer1_bigIntGetSignedArgument(void* context, int32_t id, int32_t destinationHandle);
// extern int32_t   wasmer1_bigIntStorageStoreUnsigned(void* context, int32_t keyOffset, int32_t keyLength, int32_t sourceHandle);
// extern int32_t   wasmer1_bigIntStorageLoadUnsigned(void* context, int32_t keyOffset, int32_t keyLength, int32_t destinationHandle);
// extern void      wasmer1_bigIntGetCallValue(void* context, int32_t destinationHandle);
// extern void      wasmer1_bigIntGetESDTCallValue(void* context, int32_t destination);
// extern void      wasmer1_bigIntGetESDTCallValueByIndex(void* context, int32_t destinationHandle, int32_t index);
// extern void      wasmer1_bigIntGetExternalBalance(void* context, int32_t addressOffset, int32_t result);
// extern void      wasmer1_bigIntGetESDTExternalBalance(void* context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce, int32_t resultHandle);
// extern int32_t   wasmer1_bigIntNew(void* context, long long smallValue);
// extern int32_t   wasmer1_bigIntUnsignedByteLength(void* context, int32_t referenceHandle);
// extern int32_t   wasmer1_bigIntSignedByteLength(void* context, int32_t referenceHandle);
// extern int32_t   wasmer1_bigIntGetUnsignedBytes(void* context, int32_t referenceHandle, int32_t byteOffset);
// extern int32_t   wasmer1_bigIntGetSignedBytes(void* context, int32_t referenceHandle, int32_t byteOffset);
// extern void      wasmer1_bigIntSetUnsignedBytes(void* context, int32_t destinationHandle, int32_t byteOffset, int32_t byteLength);
// extern void      wasmer1_bigIntSetSignedBytes(void* context, int32_t destinationHandle, int32_t byteOffset, int32_t byteLength);
// extern int32_t   wasmer1_bigIntIsInt64(void* context, int32_t destinationHandle);
// extern long long wasmer1_bigIntGetInt64(void* context, int32_t destinationHandle);
// extern void      wasmer1_bigIntSetInt64(void* context, int32_t destinationHandle, long long value);
// extern void      wasmer1_bigIntAdd(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigIntSub(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigIntMul(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigIntTDiv(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigIntTMod(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigIntEDiv(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigIntEMod(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigIntSqrt(void* context, int32_t destinationHandle, int32_t opHandle);
// extern void      wasmer1_bigIntPow(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern int32_t   wasmer1_bigIntLog2(void* context, int32_t op1Handle);
// extern void      wasmer1_bigIntAbs(void* context, int32_t destinationHandle, int32_t opHandle);
// extern void      wasmer1_bigIntNeg(void* context, int32_t destinationHandle, int32_t opHandle);
// extern int32_t   wasmer1_bigIntSign(void* context, int32_t opHandle);
// extern int32_t   wasmer1_bigIntCmp(void* context, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigIntNot(void* context, int32_t destinationHandle, int32_t opHandle);
// extern void      wasmer1_bigIntAnd(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigIntOr(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigIntXor(void* context, int32_t destinationHandle, int32_t op1Handle, int32_t op2Handle);
// extern void      wasmer1_bigIntShr(void* context, int32_t destinationHandle, int32_t opHandle, int32_t bits);
// extern void      wasmer1_bigIntShl(void* context, int32_t destinationHandle, int32_t opHandle, int32_t bits);
// extern void      wasmer1_bigIntFinishUnsigned(void* context, int32_t referenceHandle);
// extern void      wasmer1_bigIntFinishSigned(void* context, int32_t referenceHandle);
// extern void      wasmer1_bigIntToString(void* context, int32_t bigIntHandle, int32_t destinationHandle);
// extern long long wasmer1_getGasLeft(void* context);
// extern void      wasmer1_getSCAddress(void* context, int32_t resultOffset);
// extern void      wasmer1_getOwnerAddress(void* context, int32_t resultOffset);
// extern int32_t   wasmer1_getShardOfAddress(void* context, int32_t addressOffset);
// extern int32_t   wasmer1_isSmartContract(void* context, int32_t addressOffset);
// extern void      wasmer1_signalError(void* context, int32_t messageOffset, int32_t messageLength);
// extern void      wasmer1_getExternalBalance(void* context, int32_t addressOffset, int32_t resultOffset);
// extern int32_t   wasmer1_blockHash(void* context, long long nonce, int32_t resultOffset);
// extern int32_t   wasmer1_getESDTBalance(void* context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce, int32_t resultOffset);
// extern int32_t   wasmer1_getESDTNFTNameLength(void* context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce);
// extern int32_t   wasmer1_getESDTNFTAttributeLength(void* context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce);
// extern int32_t   wasmer1_getESDTNFTURILength(void* context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce);
// extern int32_t   wasmer1_getESDTTokenData(void* context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen, long long nonce, int32_t valueHandle, int32_t propertiesOffset, int32_t hashOffset, int32_t nameOffset, int32_t attributesOffset, int32_t creatorOffset, int32_t royaltiesHandle, int32_t urisOffset);
// extern long long wasmer1_getESDTLocalRoles(void* context, int32_t tokenIdHandle);
// extern int32_t   wasmer1_validateTokenIdentifier(void* context, int32_t tokenIdHandle);
// extern int32_t   wasmer1_transferValue(void* context, int32_t destOffset, int32_t valueOffset, int32_t dataOffset, int32_t length);
// extern int32_t   wasmer1_transferValueExecute(void* context, int32_t destOffset, int32_t valueOffset, long long gasLimit, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t   wasmer1_transferESDTExecute(void* context, int32_t destOffset, int32_t tokenIDOffset, int32_t tokenIDLen, int32_t valueOffset, long long gasLimit, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t   wasmer1_transferESDTNFTExecute(void* context, int32_t destOffset, int32_t tokenIDOffset, int32_t tokenIDLen, int32_t valueOffset, long long nonce, long long gasLimit, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t   wasmer1_multiTransferESDTNFTExecute(void* context, int32_t destOffset, int32_t numTokenTransfers, int32_t tokenTransfersArgsLengthOffset, int32_t tokenTransferDataOffset, long long gasLimit, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t   wasmer1_createAsyncCall(void* context, int32_t destOffset, int32_t valueOffset, int32_t dataOffset, int32_t dataLength, int32_t successOffset, int32_t successLength, int32_t errorOffset, int32_t errorLength, long long gas, long long extraGasForCallback);
// extern int32_t   wasmer1_setAsyncContextCallback(void* context, int32_t callback, int32_t callbackLength, int32_t data, int32_t dataLength, long long gas);
// extern void      wasmer1_upgradeContract(void* context, int32_t destOffset, long long gasLimit, int32_t valueOffset, int32_t codeOffset, int32_t codeMetadataOffset, int32_t length, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern void      wasmer1_upgradeFromSourceContract(void* context, int32_t destOffset, long long gasLimit, int32_t valueOffset, int32_t sourceContractAddressOffset, int32_t codeMetadataOffset, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern void      wasmer1_deleteContract(void* context, int32_t destOffset, long long gasLimit, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern void      wasmer1_asyncCall(void* context, int32_t destOffset, int32_t valueOffset, int32_t dataOffset, int32_t length);
// extern int32_t   wasmer1_getArgumentLength(void* context, int32_t id);
// extern int32_t   wasmer1_getArgument(void* context, int32_t id, int32_t argOffset);
// extern int32_t   wasmer1_getFunction(void* context, int32_t functionOffset);
// extern int32_t   wasmer1_getNumArguments(void* context);
// extern int32_t   wasmer1_storageStore(void* context, int32_t keyOffset, int32_t keyLength, int32_t dataOffset, int32_t dataLength);
// extern int32_t   wasmer1_storageLoadLength(void* context, int32_t keyOffset, int32_t keyLength);
// extern int32_t   wasmer1_storageLoadFromAddress(void* context, int32_t addressOffset, int32_t keyOffset, int32_t keyLength, int32_t dataOffset);
// extern int32_t   wasmer1_storageLoad(void* context, int32_t keyOffset, int32_t keyLength, int32_t dataOffset);
// extern int32_t   wasmer1_setStorageLock(void* context, int32_t keyOffset, int32_t keyLength, long long lockTimestamp);
// extern long long wasmer1_getStorageLock(void* context, int32_t keyOffset, int32_t keyLength);
// extern int32_t   wasmer1_isStorageLocked(void* context, int32_t keyOffset, int32_t keyLength);
// extern int32_t   wasmer1_clearStorageLock(void* context, int32_t keyOffset, int32_t keyLength);
// extern void      wasmer1_getCaller(void* context, int32_t resultOffset);
// extern void      wasmer1_checkNoPayment(void* context);
// extern int32_t   wasmer1_callValue(void* context, int32_t resultOffset);
// extern int32_t   wasmer1_getESDTValue(void* context, int32_t resultOffset);
// extern int32_t   wasmer1_getESDTValueByIndex(void* context, int32_t resultOffset, int32_t index);
// extern int32_t   wasmer1_getESDTTokenName(void* context, int32_t resultOffset);
// extern int32_t   wasmer1_getESDTTokenNameByIndex(void* context, int32_t resultOffset, int32_t index);
// extern long long wasmer1_getESDTTokenNonce(void* context);
// extern long long wasmer1_getESDTTokenNonceByIndex(void* context, int32_t index);
// extern long long wasmer1_getCurrentESDTNFTNonce(void* context, int32_t addressOffset, int32_t tokenIDOffset, int32_t tokenIDLen);
// extern int32_t   wasmer1_getESDTTokenType(void* context);
// extern int32_t   wasmer1_getESDTTokenTypeByIndex(void* context, int32_t index);
// extern int32_t   wasmer1_getNumESDTTransfers(void* context);
// extern int32_t   wasmer1_getCallValueTokenName(void* context, int32_t callValueOffset, int32_t tokenNameOffset);
// extern int32_t   wasmer1_getCallValueTokenNameByIndex(void* context, int32_t callValueOffset, int32_t tokenNameOffset, int32_t index);
// extern void      wasmer1_writeLog(void* context, int32_t dataPointer, int32_t dataLength, int32_t topicPtr, int32_t numTopics);
// extern void      wasmer1_writeEventLog(void* context, int32_t numTopics, int32_t topicLengthsOffset, int32_t topicOffset, int32_t dataOffset, int32_t dataLength);
// extern long long wasmer1_getBlockTimestamp(void* context);
// extern long long wasmer1_getBlockNonce(void* context);
// extern long long wasmer1_getBlockRound(void* context);
// extern long long wasmer1_getBlockEpoch(void* context);
// extern void      wasmer1_getBlockRandomSeed(void* context, int32_t pointer);
// extern void      wasmer1_getStateRootHash(void* context, int32_t pointer);
// extern long long wasmer1_getPrevBlockTimestamp(void* context);
// extern long long wasmer1_getPrevBlockNonce(void* context);
// extern long long wasmer1_getPrevBlockRound(void* context);
// extern long long wasmer1_getPrevBlockEpoch(void* context);
// extern void      wasmer1_getPrevBlockRandomSeed(void* context, int32_t pointer);
// extern void      wasmer1_returnData(void* context, int32_t pointer, int32_t length);
// extern int32_t   wasmer1_executeOnSameContext(void* context, long long gasLimit, int32_t addressOffset, int32_t valueOffset, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t   wasmer1_executeOnDestContext(void* context, long long gasLimit, int32_t addressOffset, int32_t valueOffset, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t   wasmer1_executeReadOnly(void* context, long long gasLimit, int32_t addressOffset, int32_t functionOffset, int32_t functionLength, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t   wasmer1_createContract(void* context, long long gasLimit, int32_t valueOffset, int32_t codeOffset, int32_t codeMetadataOffset, int32_t length, int32_t resultOffset, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t   wasmer1_deployFromSourceContract(void* context, long long gasLimit, int32_t valueOffset, int32_t sourceContractAddressOffset, int32_t codeMetadataOffset, int32_t resultAddressOffset, int32_t numArguments, int32_t argumentsLengthOffset, int32_t dataOffset);
// extern int32_t   wasmer1_getNumReturnData(void* context);
// extern int32_t   wasmer1_getReturnDataSize(void* context, int32_t resultID);
// extern int32_t   wasmer1_getReturnData(void* context, int32_t resultID, int32_t dataOffset);
// extern void      wasmer1_cleanReturnData(void* context);
// extern void      wasmer1_deleteFromReturnData(void* context, int32_t resultID);
// extern void      wasmer1_getOriginalTxHash(void* context, int32_t dataOffset);
// extern void      wasmer1_getCurrentTxHash(void* context, int32_t dataOffset);
// extern void      wasmer1_getPrevTxHash(void* context, int32_t dataOffset);
// extern void      wasmer1_managedSCAddress(void* context, int32_t destinationHandle);
// extern void      wasmer1_managedOwnerAddress(void* context, int32_t destinationHandle);
// extern void      wasmer1_managedCaller(void* context, int32_t destinationHandle);
// extern void      wasmer1_managedSignalError(void* context, int32_t errHandle);
// extern void      wasmer1_managedWriteLog(void* context, int32_t topicsHandle, int32_t dataHandle);
// extern void      wasmer1_managedGetOriginalTxHash(void* context, int32_t resultHandle);
// extern void      wasmer1_managedGetStateRootHash(void* context, int32_t resultHandle);
// extern void      wasmer1_managedGetBlockRandomSeed(void* context, int32_t resultHandle);
// extern void      wasmer1_managedGetPrevBlockRandomSeed(void* context, int32_t resultHandle);
// extern void      wasmer1_managedGetReturnData(void* context, int32_t resultID, int32_t resultHandle);
// extern void      wasmer1_managedGetMultiESDTCallValue(void* context, int32_t multiCallValueHandle);
// extern void      wasmer1_managedGetESDTBalance(void* context, int32_t addressHandle, int32_t tokenIDHandle, long long nonce, int32_t valueHandle);
// extern void      wasmer1_managedGetESDTTokenData(void* context, int32_t addressHandle, int32_t tokenIDHandle, long long nonce, int32_t valueHandle, int32_t propertiesHandle, int32_t hashHandle, int32_t nameHandle, int32_t attributesHandle, int32_t creatorHandle, int32_t royaltiesHandle, int32_t urisHandle);
// extern void      wasmer1_managedAsyncCall(void* context, int32_t destHandle, int32_t valueHandle, int32_t functionHandle, int32_t argumentsHandle);
// extern int32_t   wasmer1_managedCreateAsyncCall(void* context, int32_t destHandle, int32_t valueHandle, int32_t functionHandle, int32_t argumentsHandle, int32_t successOffset, int32_t successLength, int32_t errorOffset, int32_t errorLength, long long gas, long long extraGasForCallback, int32_t callbackClosureHandle);
// extern void      wasmer1_managedGetCallbackClosure(void* context, int32_t callbackClosureHandle);
// extern void      wasmer1_managedUpgradeFromSourceContract(void* context, int32_t destHandle, long long gas, int32_t valueHandle, int32_t addressHandle, int32_t codeMetadataHandle, int32_t argumentsHandle, int32_t resultHandle);
// extern void      wasmer1_managedUpgradeContract(void* context, int32_t destHandle, long long gas, int32_t valueHandle, int32_t codeHandle, int32_t codeMetadataHandle, int32_t argumentsHandle, int32_t resultHandle);
// extern void      wasmer1_managedDeleteContract(void* context, int32_t destHandle, long long gasLimit, int32_t argumentsHandle);
// extern int32_t   wasmer1_managedDeployFromSourceContract(void* context, long long gas, int32_t valueHandle, int32_t addressHandle, int32_t codeMetadataHandle, int32_t argumentsHandle, int32_t resultAddressHandle, int32_t resultHandle);
// extern int32_t   wasmer1_managedCreateContract(void* context, long long gas, int32_t valueHandle, int32_t codeHandle, int32_t codeMetadataHandle, int32_t argumentsHandle, int32_t resultAddressHandle, int32_t resultHandle);
// extern int32_t   wasmer1_managedExecuteReadOnly(void* context, long long gas, int32_t addressHandle, int32_t functionHandle, int32_t argumentsHandle, int32_t resultHandle);
// extern int32_t   wasmer1_managedExecuteOnSameContext(void* context, long long gas, int32_t addressHandle, int32_t valueHandle, int32_t functionHandle, int32_t argumentsHandle, int32_t resultHandle);
// extern int32_t   wasmer1_managedExecuteOnDestContext(void* context, long long gas, int32_t addressHandle, int32_t valueHandle, int32_t functionHandle, int32_t argumentsHandle, int32_t resultHandle);
// extern int32_t   wasmer1_managedMultiTransferESDTNFTExecute(void* context, int32_t dstHandle, int32_t tokenTransfersHandle, long long gasLimit, int32_t functionHandle, int32_t argumentsHandle);
// extern int32_t   wasmer1_managedTransferValueExecute(void* context, int32_t dstHandle, int32_t valueHandle, long long gasLimit, int32_t functionHandle, int32_t argumentsHandle);
// extern int32_t   wasmer1_managedIsESDTFrozen(void* context, int32_t addressHandle, int32_t tokenIDHandle, long long nonce);
// extern int32_t   wasmer1_managedIsESDTLimitedTransfer(void* context, int32_t tokenIDHandle);
// extern int32_t   wasmer1_managedIsESDTPaused(void* context, int32_t tokenIDHandle);
// extern void      wasmer1_managedBufferToHex(void* context, int32_t sourceHandle, int32_t destHandle);
// extern int32_t   wasmer1_mBufferNew(void* context);
// extern int32_t   wasmer1_mBufferNewFromBytes(void* context, int32_t dataOffset, int32_t dataLength);
// extern int32_t   wasmer1_mBufferGetLength(void* context, int32_t mBufferHandle);
// extern int32_t   wasmer1_mBufferGetBytes(void* context, int32_t mBufferHandle, int32_t resultOffset);
// extern int32_t   wasmer1_mBufferGetByteSlice(void* context, int32_t sourceHandle, int32_t startingPosition, int32_t sliceLength, int32_t resultOffset);
// extern int32_t   wasmer1_mBufferCopyByteSlice(void* context, int32_t sourceHandle, int32_t startingPosition, int32_t sliceLength, int32_t destinationHandle);
// extern int32_t   wasmer1_mBufferEq(void* context, int32_t mBufferHandle1, int32_t mBufferHandle2);
// extern int32_t   wasmer1_mBufferSetBytes(void* context, int32_t mBufferHandle, int32_t dataOffset, int32_t dataLength);
// extern int32_t   wasmer1_mBufferSetByteSlice(void* context, int32_t mBufferHandle, int32_t startingPosition, int32_t dataLength, int32_t dataOffset);
// extern int32_t   wasmer1_mBufferAppend(void* context, int32_t accumulatorHandle, int32_t dataHandle);
// extern int32_t   wasmer1_mBufferAppendBytes(void* context, int32_t accumulatorHandle, int32_t dataOffset, int32_t dataLength);
// extern int32_t   wasmer1_mBufferToBigIntUnsigned(void* context, int32_t mBufferHandle, int32_t bigIntHandle);
// extern int32_t   wasmer1_mBufferToBigIntSigned(void* context, int32_t mBufferHandle, int32_t bigIntHandle);
// extern int32_t   wasmer1_mBufferFromBigIntUnsigned(void* context, int32_t mBufferHandle, int32_t bigIntHandle);
// extern int32_t   wasmer1_mBufferFromBigIntSigned(void* context, int32_t mBufferHandle, int32_t bigIntHandle);
// extern int32_t   wasmer1_mBufferToBigFloat(void* context, int32_t mBufferHandle, int32_t bigFloatHandle);
// extern int32_t   wasmer1_mBufferFromBigFloat(void* context, int32_t mBufferHandle, int32_t bigFloatHandle);
// extern int32_t   wasmer1_mBufferStorageStore(void* context, int32_t keyHandle, int32_t sourceHandle);
// extern int32_t   wasmer1_mBufferStorageLoad(void* context, int32_t keyHandle, int32_t destinationHandle);
// extern void      wasmer1_mBufferStorageLoadFromAddress(void* context, int32_t addressHandle, int32_t keyHandle, int32_t destinationHandle);
// extern int32_t   wasmer1_mBufferGetArgument(void* context, int32_t id, int32_t destinationHandle);
// extern int32_t   wasmer1_mBufferFinish(void* context, int32_t sourceHandle);
// extern int32_t   wasmer1_mBufferSetRandom(void* context, int32_t destinationHandle, int32_t length);
// extern long long wasmer1_smallIntGetUnsignedArgument(void* context, int32_t id);
// extern long long wasmer1_smallIntGetSignedArgument(void* context, int32_t id);
// extern void      wasmer1_smallIntFinishUnsigned(void* context, long long value);
// extern void      wasmer1_smallIntFinishSigned(void* context, long long value);
// extern int32_t   wasmer1_smallIntStorageStoreUnsigned(void* context, int32_t keyOffset, int32_t keyLength, long long value);
// extern int32_t   wasmer1_smallIntStorageStoreSigned(void* context, int32_t keyOffset, int32_t keyLength, long long value);
// extern long long wasmer1_smallIntStorageLoadUnsigned(void* context, int32_t keyOffset, int32_t keyLength);
// extern long long wasmer1_smallIntStorageLoadSigned(void* context, int32_t keyOffset, int32_t keyLength);
// extern long long wasmer1_int64getArgument(void* context, int32_t id);
// extern void      wasmer1_int64finish(void* context, long long value);
// extern int32_t   wasmer1_int64storageStore(void* context, int32_t keyOffset, int32_t keyLength, long long value);
// extern long long wasmer1_int64storageLoad(void* context, int32_t keyOffset, int32_t keyLength);
// extern int32_t   wasmer1_sha256(void* context, int32_t dataOffset, int32_t length, int32_t resultOffset);
// extern int32_t   wasmer1_managedSha256(void* context, int32_t inputHandle, int32_t outputHandle);
// extern int32_t   wasmer1_keccak256(void* context, int32_t dataOffset, int32_t length, int32_t resultOffset);
// extern int32_t   wasmer1_managedKeccak256(void* context, int32_t inputHandle, int32_t outputHandle);
// extern int32_t   wasmer1_ripemd160(void* context, int32_t dataOffset, int32_t length, int32_t resultOffset);
// extern int32_t   wasmer1_managedRipemd160(void* context, int32_t inputHandle, int32_t outputHandle);
// extern int32_t   wasmer1_verifyBLS(void* context, int32_t keyOffset, int32_t messageOffset, int32_t messageLength, int32_t sigOffset);
// extern int32_t   wasmer1_managedVerifyBLS(void* context, int32_t keyHandle, int32_t messageHandle, int32_t sigHandle);
// extern int32_t   wasmer1_verifyEd25519(void* context, int32_t keyOffset, int32_t messageOffset, int32_t messageLength, int32_t sigOffset);
// extern int32_t   wasmer1_managedVerifyEd25519(void* context, int32_t keyHandle, int32_t messageHandle, int32_t sigHandle);
// extern int32_t   wasmer1_verifyCustomSecp256k1(void* context, int32_t keyOffset, int32_t keyLength, int32_t messageOffset, int32_t messageLength, int32_t sigOffset, int32_t hashType);
// extern int32_t   wasmer1_managedVerifyCustomSecp256k1(void* context, int32_t keyHandle, int32_t messageHandle, int32_t sigHandle, int32_t hashType);
// extern int32_t   wasmer1_verifySecp256k1(void* context, int32_t keyOffset, int32_t keyLength, int32_t messageOffset, int32_t messageLength, int32_t sigOffset);
// extern int32_t   wasmer1_managedVerifySecp256k1(void* context, int32_t keyHandle, int32_t messageHandle, int32_t sigHandle);
// extern int32_t   wasmer1_encodeSecp256k1DerSignature(void* context, int32_t rOffset, int32_t rLength, int32_t sOffset, int32_t sLength, int32_t sigOffset);
// extern int32_t   wasmer1_managedEncodeSecp256k1DerSignature(void* context, int32_t rHandle, int32_t sHandle, int32_t sigHandle);
// extern void      wasmer1_addEC(void* context, int32_t xResultHandle, int32_t yResultHandle, int32_t ecHandle, int32_t fstPointXHandle, int32_t fstPointYHandle, int32_t sndPointXHandle, int32_t sndPointYHandle);
// extern void      wasmer1_doubleEC(void* context, int32_t xResultHandle, int32_t yResultHandle, int32_t ecHandle, int32_t pointXHandle, int32_t pointYHandle);
// extern int32_t   wasmer1_isOnCurveEC(void* context, int32_t ecHandle, int32_t pointXHandle, int32_t pointYHandle);
// extern int32_t   wasmer1_scalarBaseMultEC(void* context, int32_t xResultHandle, int32_t yResultHandle, int32_t ecHandle, int32_t dataOffset, int32_t length);
// extern int32_t   wasmer1_managedScalarBaseMultEC(void* context, int32_t xResultHandle, int32_t yResultHandle, int32_t ecHandle, int32_t dataHandle);
// extern int32_t   wasmer1_scalarMultEC(void* context, int32_t xResultHandle, int32_t yResultHandle, int32_t ecHandle, int32_t pointXHandle, int32_t pointYHandle, int32_t dataOffset, int32_t length);
// extern int32_t   wasmer1_managedScalarMultEC(void* context, int32_t xResultHandle, int32_t yResultHandle, int32_t ecHandle, int32_t pointXHandle, int32_t pointYHandle, int32_t dataHandle);
// extern int32_t   wasmer1_marshalEC(void* context, int32_t xPairHandle, int32_t yPairHandle, int32_t ecHandle, int32_t resultOffset);
// extern int32_t   wasmer1_managedMarshalEC(void* context, int32_t xPairHandle, int32_t yPairHandle, int32_t ecHandle, int32_t resultHandle);
// extern int32_t   wasmer1_marshalCompressedEC(void* context, int32_t xPairHandle, int32_t yPairHandle, int32_t ecHandle, int32_t resultOffset);
// extern int32_t   wasmer1_managedMarshalCompressedEC(void* context, int32_t xPairHandle, int32_t yPairHandle, int32_t ecHandle, int32_t resultHandle);
// extern int32_t   wasmer1_unmarshalEC(void* context, int32_t xResultHandle, int32_t yResultHandle, int32_t ecHandle, int32_t dataOffset, int32_t length);
// extern int32_t   wasmer1_managedUnmarshalEC(void* context, int32_t xResultHandle, int32_t yResultHandle, int32_t ecHandle, int32_t dataHandle);
// extern int32_t   wasmer1_unmarshalCompressedEC(void* context, int32_t xResultHandle, int32_t yResultHandle, int32_t ecHandle, int32_t dataOffset, int32_t length);
// extern int32_t   wasmer1_managedUnmarshalCompressedEC(void* context, int32_t xResultHandle, int32_t yResultHandle, int32_t ecHandle, int32_t dataHandle);
// extern int32_t   wasmer1_generateKeyEC(void* context, int32_t xPubKeyHandle, int32_t yPubKeyHandle, int32_t ecHandle, int32_t resultOffset);
// extern int32_t   wasmer1_managedGenerateKeyEC(void* context, int32_t xPubKeyHandle, int32_t yPubKeyHandle, int32_t ecHandle, int32_t resultHandle);
// extern int32_t   wasmer1_createEC(void* context, int32_t dataOffset, int32_t dataLength);
// extern int32_t   wasmer1_managedCreateEC(void* context, int32_t dataHandle);
// extern int32_t   wasmer1_getCurveLengthEC(void* context, int32_t ecHandle);
// extern int32_t   wasmer1_getPrivKeyByteLengthEC(void* context, int32_t ecHandle);
// extern int32_t   wasmer1_ellipticCurveGetValues(void* context, int32_t ecHandle, int32_t fieldOrderHandle, int32_t basePointOrderHandle, int32_t eqConstantHandle, int32_t xBasePointHandle, int32_t yBasePointHandle);
import "C"

import (
	"unsafe"
)

// export wasmer1_bigFloatNewFromParts
func wasmer1_bigFloatNewFromParts(context unsafe.Pointer, integralPart int32, fractionalPart int32, exponent int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigFloatNewFromParts(integralPart, fractionalPart, exponent)
}

// export wasmer1_bigFloatNewFromFrac
func wasmer1_bigFloatNewFromFrac(context unsafe.Pointer, numerator int64, denominator int64) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigFloatNewFromFrac(numerator, denominator)
}

// export wasmer1_bigFloatNewFromSci
func wasmer1_bigFloatNewFromSci(context unsafe.Pointer, significand int64, exponent int64) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigFloatNewFromSci(significand, exponent)
}

// export wasmer1_bigFloatAdd
func wasmer1_bigFloatAdd(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatAdd(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigFloatSub
func wasmer1_bigFloatSub(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatSub(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigFloatMul
func wasmer1_bigFloatMul(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatMul(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigFloatDiv
func wasmer1_bigFloatDiv(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatDiv(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigFloatNeg
func wasmer1_bigFloatNeg(context unsafe.Pointer, destinationHandle int32, opHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatNeg(destinationHandle, opHandle)
}

// export wasmer1_bigFloatClone
func wasmer1_bigFloatClone(context unsafe.Pointer, destinationHandle int32, opHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatClone(destinationHandle, opHandle)
}

// export wasmer1_bigFloatCmp
func wasmer1_bigFloatCmp(context unsafe.Pointer, op1Handle int32, op2Handle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigFloatCmp(op1Handle, op2Handle)
}

// export wasmer1_bigFloatAbs
func wasmer1_bigFloatAbs(context unsafe.Pointer, destinationHandle int32, opHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatAbs(destinationHandle, opHandle)
}

// export wasmer1_bigFloatSign
func wasmer1_bigFloatSign(context unsafe.Pointer, opHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigFloatSign(opHandle)
}

// export wasmer1_bigFloatSqrt
func wasmer1_bigFloatSqrt(context unsafe.Pointer, destinationHandle int32, opHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatSqrt(destinationHandle, opHandle)
}

// export wasmer1_bigFloatPow
func wasmer1_bigFloatPow(context unsafe.Pointer, destinationHandle int32, opHandle int32, exponent int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatPow(destinationHandle, opHandle, exponent)
}

// export wasmer1_bigFloatFloor
func wasmer1_bigFloatFloor(context unsafe.Pointer, destBigIntHandle int32, opHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatFloor(destBigIntHandle, opHandle)
}

// export wasmer1_bigFloatCeil
func wasmer1_bigFloatCeil(context unsafe.Pointer, destBigIntHandle int32, opHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatCeil(destBigIntHandle, opHandle)
}

// export wasmer1_bigFloatTruncate
func wasmer1_bigFloatTruncate(context unsafe.Pointer, destBigIntHandle int32, opHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatTruncate(destBigIntHandle, opHandle)
}

// export wasmer1_bigFloatSetInt64
func wasmer1_bigFloatSetInt64(context unsafe.Pointer, destinationHandle int32, value int64) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatSetInt64(destinationHandle, value)
}

// export wasmer1_bigFloatIsInt
func wasmer1_bigFloatIsInt(context unsafe.Pointer, opHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigFloatIsInt(opHandle)
}

// export wasmer1_bigFloatSetBigInt
func wasmer1_bigFloatSetBigInt(context unsafe.Pointer, destinationHandle int32, bigIntHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatSetBigInt(destinationHandle, bigIntHandle)
}

// export wasmer1_bigFloatGetConstPi
func wasmer1_bigFloatGetConstPi(context unsafe.Pointer, destinationHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatGetConstPi(destinationHandle)
}

// export wasmer1_bigFloatGetConstE
func wasmer1_bigFloatGetConstE(context unsafe.Pointer, destinationHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigFloatGetConstE(destinationHandle)
}

// export wasmer1_bigIntGetUnsignedArgument
func wasmer1_bigIntGetUnsignedArgument(context unsafe.Pointer, id int32, destinationHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntGetUnsignedArgument(id, destinationHandle)
}

// export wasmer1_bigIntGetSignedArgument
func wasmer1_bigIntGetSignedArgument(context unsafe.Pointer, id int32, destinationHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntGetSignedArgument(id, destinationHandle)
}

// export wasmer1_bigIntStorageStoreUnsigned
func wasmer1_bigIntStorageStoreUnsigned(context unsafe.Pointer, keyOffset int32, keyLength int32, sourceHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigIntStorageStoreUnsigned(keyOffset, keyLength, sourceHandle)
}

// export wasmer1_bigIntStorageLoadUnsigned
func wasmer1_bigIntStorageLoadUnsigned(context unsafe.Pointer, keyOffset int32, keyLength int32, destinationHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigIntStorageLoadUnsigned(keyOffset, keyLength, destinationHandle)
}

// export wasmer1_bigIntGetCallValue
func wasmer1_bigIntGetCallValue(context unsafe.Pointer, destinationHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntGetCallValue(destinationHandle)
}

// export wasmer1_bigIntGetESDTCallValue
func wasmer1_bigIntGetESDTCallValue(context unsafe.Pointer, destination int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntGetESDTCallValue(destination)
}

// export wasmer1_bigIntGetESDTCallValueByIndex
func wasmer1_bigIntGetESDTCallValueByIndex(context unsafe.Pointer, destinationHandle int32, index int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntGetESDTCallValueByIndex(destinationHandle, index)
}

// export wasmer1_bigIntGetExternalBalance
func wasmer1_bigIntGetExternalBalance(context unsafe.Pointer, addressOffset int32, result int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntGetExternalBalance(addressOffset, result)
}

// export wasmer1_bigIntGetESDTExternalBalance
func wasmer1_bigIntGetESDTExternalBalance(context unsafe.Pointer, addressOffset int32, tokenIDOffset int32, tokenIDLen int32, nonce int64, resultHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntGetESDTExternalBalance(addressOffset, tokenIDOffset, tokenIDLen, nonce, resultHandle)
}

// export wasmer1_bigIntNew
func wasmer1_bigIntNew(context unsafe.Pointer, smallValue int64) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigIntNew(smallValue)
}

// export wasmer1_bigIntUnsignedByteLength
func wasmer1_bigIntUnsignedByteLength(context unsafe.Pointer, referenceHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigIntUnsignedByteLength(referenceHandle)
}

// export wasmer1_bigIntSignedByteLength
func wasmer1_bigIntSignedByteLength(context unsafe.Pointer, referenceHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigIntSignedByteLength(referenceHandle)
}

// export wasmer1_bigIntGetUnsignedBytes
func wasmer1_bigIntGetUnsignedBytes(context unsafe.Pointer, referenceHandle int32, byteOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigIntGetUnsignedBytes(referenceHandle, byteOffset)
}

// export wasmer1_bigIntGetSignedBytes
func wasmer1_bigIntGetSignedBytes(context unsafe.Pointer, referenceHandle int32, byteOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigIntGetSignedBytes(referenceHandle, byteOffset)
}

// export wasmer1_bigIntSetUnsignedBytes
func wasmer1_bigIntSetUnsignedBytes(context unsafe.Pointer, destinationHandle int32, byteOffset int32, byteLength int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntSetUnsignedBytes(destinationHandle, byteOffset, byteLength)
}

// export wasmer1_bigIntSetSignedBytes
func wasmer1_bigIntSetSignedBytes(context unsafe.Pointer, destinationHandle int32, byteOffset int32, byteLength int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntSetSignedBytes(destinationHandle, byteOffset, byteLength)
}

// export wasmer1_bigIntIsInt64
func wasmer1_bigIntIsInt64(context unsafe.Pointer, destinationHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigIntIsInt64(destinationHandle)
}

// export wasmer1_bigIntGetInt64
func wasmer1_bigIntGetInt64(context unsafe.Pointer, destinationHandle int32) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigIntGetInt64(destinationHandle)
}

// export wasmer1_bigIntSetInt64
func wasmer1_bigIntSetInt64(context unsafe.Pointer, destinationHandle int32, value int64) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntSetInt64(destinationHandle, value)
}

// export wasmer1_bigIntAdd
func wasmer1_bigIntAdd(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntAdd(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigIntSub
func wasmer1_bigIntSub(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntSub(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigIntMul
func wasmer1_bigIntMul(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntMul(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigIntTDiv
func wasmer1_bigIntTDiv(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntTDiv(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigIntTMod
func wasmer1_bigIntTMod(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntTMod(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigIntEDiv
func wasmer1_bigIntEDiv(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntEDiv(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigIntEMod
func wasmer1_bigIntEMod(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntEMod(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigIntSqrt
func wasmer1_bigIntSqrt(context unsafe.Pointer, destinationHandle int32, opHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntSqrt(destinationHandle, opHandle)
}

// export wasmer1_bigIntPow
func wasmer1_bigIntPow(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntPow(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigIntLog2
func wasmer1_bigIntLog2(context unsafe.Pointer, op1Handle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigIntLog2(op1Handle)
}

// export wasmer1_bigIntAbs
func wasmer1_bigIntAbs(context unsafe.Pointer, destinationHandle int32, opHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntAbs(destinationHandle, opHandle)
}

// export wasmer1_bigIntNeg
func wasmer1_bigIntNeg(context unsafe.Pointer, destinationHandle int32, opHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntNeg(destinationHandle, opHandle)
}

// export wasmer1_bigIntSign
func wasmer1_bigIntSign(context unsafe.Pointer, opHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigIntSign(opHandle)
}

// export wasmer1_bigIntCmp
func wasmer1_bigIntCmp(context unsafe.Pointer, op1Handle int32, op2Handle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BigIntCmp(op1Handle, op2Handle)
}

// export wasmer1_bigIntNot
func wasmer1_bigIntNot(context unsafe.Pointer, destinationHandle int32, opHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntNot(destinationHandle, opHandle)
}

// export wasmer1_bigIntAnd
func wasmer1_bigIntAnd(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntAnd(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigIntOr
func wasmer1_bigIntOr(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntOr(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigIntXor
func wasmer1_bigIntXor(context unsafe.Pointer, destinationHandle int32, op1Handle int32, op2Handle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntXor(destinationHandle, op1Handle, op2Handle)
}

// export wasmer1_bigIntShr
func wasmer1_bigIntShr(context unsafe.Pointer, destinationHandle int32, opHandle int32, bits int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntShr(destinationHandle, opHandle, bits)
}

// export wasmer1_bigIntShl
func wasmer1_bigIntShl(context unsafe.Pointer, destinationHandle int32, opHandle int32, bits int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntShl(destinationHandle, opHandle, bits)
}

// export wasmer1_bigIntFinishUnsigned
func wasmer1_bigIntFinishUnsigned(context unsafe.Pointer, referenceHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntFinishUnsigned(referenceHandle)
}

// export wasmer1_bigIntFinishSigned
func wasmer1_bigIntFinishSigned(context unsafe.Pointer, referenceHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntFinishSigned(referenceHandle)
}

// export wasmer1_bigIntToString
func wasmer1_bigIntToString(context unsafe.Pointer, bigIntHandle int32, destinationHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.BigIntToString(bigIntHandle, destinationHandle)
}

// export wasmer1_getGasLeft
func wasmer1_getGasLeft(context unsafe.Pointer) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetGasLeft()
}

// export wasmer1_getSCAddress
func wasmer1_getSCAddress(context unsafe.Pointer, resultOffset int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.GetSCAddress(resultOffset)
}

// export wasmer1_getOwnerAddress
func wasmer1_getOwnerAddress(context unsafe.Pointer, resultOffset int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.GetOwnerAddress(resultOffset)
}

// export wasmer1_getShardOfAddress
func wasmer1_getShardOfAddress(context unsafe.Pointer, addressOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetShardOfAddress(addressOffset)
}

// export wasmer1_isSmartContract
func wasmer1_isSmartContract(context unsafe.Pointer, addressOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.IsSmartContract(addressOffset)
}

// export wasmer1_signalError
func wasmer1_signalError(context unsafe.Pointer, messageOffset int32, messageLength int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.SignalError(messageOffset, messageLength)
}

// export wasmer1_getExternalBalance
func wasmer1_getExternalBalance(context unsafe.Pointer, addressOffset int32, resultOffset int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.GetExternalBalance(addressOffset, resultOffset)
}

// export wasmer1_blockHash
func wasmer1_blockHash(context unsafe.Pointer, nonce int64, resultOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.BlockHash(nonce, resultOffset)
}

// export wasmer1_getESDTBalance
func wasmer1_getESDTBalance(context unsafe.Pointer, addressOffset int32, tokenIDOffset int32, tokenIDLen int32, nonce int64, resultOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTBalance(addressOffset, tokenIDOffset, tokenIDLen, nonce, resultOffset)
}

// export wasmer1_getESDTNFTNameLength
func wasmer1_getESDTNFTNameLength(context unsafe.Pointer, addressOffset int32, tokenIDOffset int32, tokenIDLen int32, nonce int64) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTNFTNameLength(addressOffset, tokenIDOffset, tokenIDLen, nonce)
}

// export wasmer1_getESDTNFTAttributeLength
func wasmer1_getESDTNFTAttributeLength(context unsafe.Pointer, addressOffset int32, tokenIDOffset int32, tokenIDLen int32, nonce int64) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTNFTAttributeLength(addressOffset, tokenIDOffset, tokenIDLen, nonce)
}

// export wasmer1_getESDTNFTURILength
func wasmer1_getESDTNFTURILength(context unsafe.Pointer, addressOffset int32, tokenIDOffset int32, tokenIDLen int32, nonce int64) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTNFTURILength(addressOffset, tokenIDOffset, tokenIDLen, nonce)
}

// export wasmer1_getESDTTokenData
func wasmer1_getESDTTokenData(context unsafe.Pointer, addressOffset int32, tokenIDOffset int32, tokenIDLen int32, nonce int64, valueHandle int32, propertiesOffset int32, hashOffset int32, nameOffset int32, attributesOffset int32, creatorOffset int32, royaltiesHandle int32, urisOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTTokenData(addressOffset, tokenIDOffset, tokenIDLen, nonce, valueHandle, propertiesOffset, hashOffset, nameOffset, attributesOffset, creatorOffset, royaltiesHandle, urisOffset)
}

// export wasmer1_getESDTLocalRoles
func wasmer1_getESDTLocalRoles(context unsafe.Pointer, tokenIdHandle int32) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTLocalRoles(tokenIdHandle)
}

// export wasmer1_validateTokenIdentifier
func wasmer1_validateTokenIdentifier(context unsafe.Pointer, tokenIdHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ValidateTokenIdentifier(tokenIdHandle)
}

// export wasmer1_transferValue
func wasmer1_transferValue(context unsafe.Pointer, destOffset int32, valueOffset int32, dataOffset int32, length int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.TransferValue(destOffset, valueOffset, dataOffset, length)
}

// export wasmer1_transferValueExecute
func wasmer1_transferValueExecute(context unsafe.Pointer, destOffset int32, valueOffset int32, gasLimit int64, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.TransferValueExecute(destOffset, valueOffset, gasLimit, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
}

// export wasmer1_transferESDTExecute
func wasmer1_transferESDTExecute(context unsafe.Pointer, destOffset int32, tokenIDOffset int32, tokenIDLen int32, valueOffset int32, gasLimit int64, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.TransferESDTExecute(destOffset, tokenIDOffset, tokenIDLen, valueOffset, gasLimit, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
}

// export wasmer1_transferESDTNFTExecute
func wasmer1_transferESDTNFTExecute(context unsafe.Pointer, destOffset int32, tokenIDOffset int32, tokenIDLen int32, valueOffset int32, nonce int64, gasLimit int64, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.TransferESDTNFTExecute(destOffset, tokenIDOffset, tokenIDLen, valueOffset, nonce, gasLimit, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
}

// export wasmer1_multiTransferESDTNFTExecute
func wasmer1_multiTransferESDTNFTExecute(context unsafe.Pointer, destOffset int32, numTokenTransfers int32, tokenTransfersArgsLengthOffset int32, tokenTransferDataOffset int32, gasLimit int64, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MultiTransferESDTNFTExecute(destOffset, numTokenTransfers, tokenTransfersArgsLengthOffset, tokenTransferDataOffset, gasLimit, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
}

// export wasmer1_createAsyncCall
func wasmer1_createAsyncCall(context unsafe.Pointer, destOffset int32, valueOffset int32, dataOffset int32, dataLength int32, successOffset int32, successLength int32, errorOffset int32, errorLength int32, gas int64, extraGasForCallback int64) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.CreateAsyncCall(destOffset, valueOffset, dataOffset, dataLength, successOffset, successLength, errorOffset, errorLength, gas, extraGasForCallback)
}

// export wasmer1_setAsyncContextCallback
func wasmer1_setAsyncContextCallback(context unsafe.Pointer, callback int32, callbackLength int32, data int32, dataLength int32, gas int64) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.SetAsyncContextCallback(callback, callbackLength, data, dataLength, gas)
}

// export wasmer1_upgradeContract
func wasmer1_upgradeContract(context unsafe.Pointer, destOffset int32, gasLimit int64, valueOffset int32, codeOffset int32, codeMetadataOffset int32, length int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.UpgradeContract(destOffset, gasLimit, valueOffset, codeOffset, codeMetadataOffset, length, numArguments, argumentsLengthOffset, dataOffset)
}

// export wasmer1_upgradeFromSourceContract
func wasmer1_upgradeFromSourceContract(context unsafe.Pointer, destOffset int32, gasLimit int64, valueOffset int32, sourceContractAddressOffset int32, codeMetadataOffset int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.UpgradeFromSourceContract(destOffset, gasLimit, valueOffset, sourceContractAddressOffset, codeMetadataOffset, numArguments, argumentsLengthOffset, dataOffset)
}

// export wasmer1_deleteContract
func wasmer1_deleteContract(context unsafe.Pointer, destOffset int32, gasLimit int64, numArguments int32, argumentsLengthOffset int32, dataOffset int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.DeleteContract(destOffset, gasLimit, numArguments, argumentsLengthOffset, dataOffset)
}

// export wasmer1_asyncCall
func wasmer1_asyncCall(context unsafe.Pointer, destOffset int32, valueOffset int32, dataOffset int32, length int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.AsyncCall(destOffset, valueOffset, dataOffset, length)
}

// export wasmer1_getArgumentLength
func wasmer1_getArgumentLength(context unsafe.Pointer, id int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetArgumentLength(id)
}

// export wasmer1_getArgument
func wasmer1_getArgument(context unsafe.Pointer, id int32, argOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetArgument(id, argOffset)
}

// export wasmer1_getFunction
func wasmer1_getFunction(context unsafe.Pointer, functionOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetFunction(functionOffset)
}

// export wasmer1_getNumArguments
func wasmer1_getNumArguments(context unsafe.Pointer) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetNumArguments()
}

// export wasmer1_storageStore
func wasmer1_storageStore(context unsafe.Pointer, keyOffset int32, keyLength int32, dataOffset int32, dataLength int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.StorageStore(keyOffset, keyLength, dataOffset, dataLength)
}

// export wasmer1_storageLoadLength
func wasmer1_storageLoadLength(context unsafe.Pointer, keyOffset int32, keyLength int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.StorageLoadLength(keyOffset, keyLength)
}

// export wasmer1_storageLoadFromAddress
func wasmer1_storageLoadFromAddress(context unsafe.Pointer, addressOffset int32, keyOffset int32, keyLength int32, dataOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.StorageLoadFromAddress(addressOffset, keyOffset, keyLength, dataOffset)
}

// export wasmer1_storageLoad
func wasmer1_storageLoad(context unsafe.Pointer, keyOffset int32, keyLength int32, dataOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.StorageLoad(keyOffset, keyLength, dataOffset)
}

// export wasmer1_setStorageLock
func wasmer1_setStorageLock(context unsafe.Pointer, keyOffset int32, keyLength int32, lockTimestamp int64) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.SetStorageLock(keyOffset, keyLength, lockTimestamp)
}

// export wasmer1_getStorageLock
func wasmer1_getStorageLock(context unsafe.Pointer, keyOffset int32, keyLength int32) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetStorageLock(keyOffset, keyLength)
}

// export wasmer1_isStorageLocked
func wasmer1_isStorageLocked(context unsafe.Pointer, keyOffset int32, keyLength int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.IsStorageLocked(keyOffset, keyLength)
}

// export wasmer1_clearStorageLock
func wasmer1_clearStorageLock(context unsafe.Pointer, keyOffset int32, keyLength int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ClearStorageLock(keyOffset, keyLength)
}

// export wasmer1_getCaller
func wasmer1_getCaller(context unsafe.Pointer, resultOffset int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.GetCaller(resultOffset)
}

// export wasmer1_checkNoPayment
func wasmer1_checkNoPayment(context unsafe.Pointer) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.CheckNoPayment()
}

// export wasmer1_callValue
func wasmer1_callValue(context unsafe.Pointer, resultOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.CallValue(resultOffset)
}

// export wasmer1_getESDTValue
func wasmer1_getESDTValue(context unsafe.Pointer, resultOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTValue(resultOffset)
}

// export wasmer1_getESDTValueByIndex
func wasmer1_getESDTValueByIndex(context unsafe.Pointer, resultOffset int32, index int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTValueByIndex(resultOffset, index)
}

// export wasmer1_getESDTTokenName
func wasmer1_getESDTTokenName(context unsafe.Pointer, resultOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTTokenName(resultOffset)
}

// export wasmer1_getESDTTokenNameByIndex
func wasmer1_getESDTTokenNameByIndex(context unsafe.Pointer, resultOffset int32, index int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTTokenNameByIndex(resultOffset, index)
}

// export wasmer1_getESDTTokenNonce
func wasmer1_getESDTTokenNonce(context unsafe.Pointer) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTTokenNonce()
}

// export wasmer1_getESDTTokenNonceByIndex
func wasmer1_getESDTTokenNonceByIndex(context unsafe.Pointer, index int32) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTTokenNonceByIndex(index)
}

// export wasmer1_getCurrentESDTNFTNonce
func wasmer1_getCurrentESDTNFTNonce(context unsafe.Pointer, addressOffset int32, tokenIDOffset int32, tokenIDLen int32) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetCurrentESDTNFTNonce(addressOffset, tokenIDOffset, tokenIDLen)
}

// export wasmer1_getESDTTokenType
func wasmer1_getESDTTokenType(context unsafe.Pointer) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTTokenType()
}

// export wasmer1_getESDTTokenTypeByIndex
func wasmer1_getESDTTokenTypeByIndex(context unsafe.Pointer, index int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetESDTTokenTypeByIndex(index)
}

// export wasmer1_getNumESDTTransfers
func wasmer1_getNumESDTTransfers(context unsafe.Pointer) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetNumESDTTransfers()
}

// export wasmer1_getCallValueTokenName
func wasmer1_getCallValueTokenName(context unsafe.Pointer, callValueOffset int32, tokenNameOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetCallValueTokenName(callValueOffset, tokenNameOffset)
}

// export wasmer1_getCallValueTokenNameByIndex
func wasmer1_getCallValueTokenNameByIndex(context unsafe.Pointer, callValueOffset int32, tokenNameOffset int32, index int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetCallValueTokenNameByIndex(callValueOffset, tokenNameOffset, index)
}

// export wasmer1_writeLog
func wasmer1_writeLog(context unsafe.Pointer, dataPointer int32, dataLength int32, topicPtr int32, numTopics int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.WriteLog(dataPointer, dataLength, topicPtr, numTopics)
}

// export wasmer1_writeEventLog
func wasmer1_writeEventLog(context unsafe.Pointer, numTopics int32, topicLengthsOffset int32, topicOffset int32, dataOffset int32, dataLength int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.WriteEventLog(numTopics, topicLengthsOffset, topicOffset, dataOffset, dataLength)
}

// export wasmer1_getBlockTimestamp
func wasmer1_getBlockTimestamp(context unsafe.Pointer) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetBlockTimestamp()
}

// export wasmer1_getBlockNonce
func wasmer1_getBlockNonce(context unsafe.Pointer) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetBlockNonce()
}

// export wasmer1_getBlockRound
func wasmer1_getBlockRound(context unsafe.Pointer) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetBlockRound()
}

// export wasmer1_getBlockEpoch
func wasmer1_getBlockEpoch(context unsafe.Pointer) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetBlockEpoch()
}

// export wasmer1_getBlockRandomSeed
func wasmer1_getBlockRandomSeed(context unsafe.Pointer, pointer int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.GetBlockRandomSeed(pointer)
}

// export wasmer1_getStateRootHash
func wasmer1_getStateRootHash(context unsafe.Pointer, pointer int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.GetStateRootHash(pointer)
}

// export wasmer1_getPrevBlockTimestamp
func wasmer1_getPrevBlockTimestamp(context unsafe.Pointer) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetPrevBlockTimestamp()
}

// export wasmer1_getPrevBlockNonce
func wasmer1_getPrevBlockNonce(context unsafe.Pointer) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetPrevBlockNonce()
}

// export wasmer1_getPrevBlockRound
func wasmer1_getPrevBlockRound(context unsafe.Pointer) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetPrevBlockRound()
}

// export wasmer1_getPrevBlockEpoch
func wasmer1_getPrevBlockEpoch(context unsafe.Pointer) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetPrevBlockEpoch()
}

// export wasmer1_getPrevBlockRandomSeed
func wasmer1_getPrevBlockRandomSeed(context unsafe.Pointer, pointer int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.GetPrevBlockRandomSeed(pointer)
}

// export wasmer1_returnData
func wasmer1_returnData(context unsafe.Pointer, pointer int32, length int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ReturnData(pointer, length)
}

// export wasmer1_executeOnSameContext
func wasmer1_executeOnSameContext(context unsafe.Pointer, gasLimit int64, addressOffset int32, valueOffset int32, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ExecuteOnSameContext(gasLimit, addressOffset, valueOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
}

// export wasmer1_executeOnDestContext
func wasmer1_executeOnDestContext(context unsafe.Pointer, gasLimit int64, addressOffset int32, valueOffset int32, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ExecuteOnDestContext(gasLimit, addressOffset, valueOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
}

// export wasmer1_executeReadOnly
func wasmer1_executeReadOnly(context unsafe.Pointer, gasLimit int64, addressOffset int32, functionOffset int32, functionLength int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ExecuteReadOnly(gasLimit, addressOffset, functionOffset, functionLength, numArguments, argumentsLengthOffset, dataOffset)
}

// export wasmer1_createContract
func wasmer1_createContract(context unsafe.Pointer, gasLimit int64, valueOffset int32, codeOffset int32, codeMetadataOffset int32, length int32, resultOffset int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.CreateContract(gasLimit, valueOffset, codeOffset, codeMetadataOffset, length, resultOffset, numArguments, argumentsLengthOffset, dataOffset)
}

// export wasmer1_deployFromSourceContract
func wasmer1_deployFromSourceContract(context unsafe.Pointer, gasLimit int64, valueOffset int32, sourceContractAddressOffset int32, codeMetadataOffset int32, resultAddressOffset int32, numArguments int32, argumentsLengthOffset int32, dataOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.DeployFromSourceContract(gasLimit, valueOffset, sourceContractAddressOffset, codeMetadataOffset, resultAddressOffset, numArguments, argumentsLengthOffset, dataOffset)
}

// export wasmer1_getNumReturnData
func wasmer1_getNumReturnData(context unsafe.Pointer) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetNumReturnData()
}

// export wasmer1_getReturnDataSize
func wasmer1_getReturnDataSize(context unsafe.Pointer, resultID int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetReturnDataSize(resultID)
}

// export wasmer1_getReturnData
func wasmer1_getReturnData(context unsafe.Pointer, resultID int32, dataOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetReturnData(resultID, dataOffset)
}

// export wasmer1_cleanReturnData
func wasmer1_cleanReturnData(context unsafe.Pointer) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.CleanReturnData()
}

// export wasmer1_deleteFromReturnData
func wasmer1_deleteFromReturnData(context unsafe.Pointer, resultID int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.DeleteFromReturnData(resultID)
}

// export wasmer1_getOriginalTxHash
func wasmer1_getOriginalTxHash(context unsafe.Pointer, dataOffset int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.GetOriginalTxHash(dataOffset)
}

// export wasmer1_getCurrentTxHash
func wasmer1_getCurrentTxHash(context unsafe.Pointer, dataOffset int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.GetCurrentTxHash(dataOffset)
}

// export wasmer1_getPrevTxHash
func wasmer1_getPrevTxHash(context unsafe.Pointer, dataOffset int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.GetPrevTxHash(dataOffset)
}

// export wasmer1_managedSCAddress
func wasmer1_managedSCAddress(context unsafe.Pointer, destinationHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedSCAddress(destinationHandle)
}

// export wasmer1_managedOwnerAddress
func wasmer1_managedOwnerAddress(context unsafe.Pointer, destinationHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedOwnerAddress(destinationHandle)
}

// export wasmer1_managedCaller
func wasmer1_managedCaller(context unsafe.Pointer, destinationHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedCaller(destinationHandle)
}

// export wasmer1_managedSignalError
func wasmer1_managedSignalError(context unsafe.Pointer, errHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedSignalError(errHandle)
}

// export wasmer1_managedWriteLog
func wasmer1_managedWriteLog(context unsafe.Pointer, topicsHandle int32, dataHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedWriteLog(topicsHandle, dataHandle)
}

// export wasmer1_managedGetOriginalTxHash
func wasmer1_managedGetOriginalTxHash(context unsafe.Pointer, resultHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedGetOriginalTxHash(resultHandle)
}

// export wasmer1_managedGetStateRootHash
func wasmer1_managedGetStateRootHash(context unsafe.Pointer, resultHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedGetStateRootHash(resultHandle)
}

// export wasmer1_managedGetBlockRandomSeed
func wasmer1_managedGetBlockRandomSeed(context unsafe.Pointer, resultHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedGetBlockRandomSeed(resultHandle)
}

// export wasmer1_managedGetPrevBlockRandomSeed
func wasmer1_managedGetPrevBlockRandomSeed(context unsafe.Pointer, resultHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedGetPrevBlockRandomSeed(resultHandle)
}

// export wasmer1_managedGetReturnData
func wasmer1_managedGetReturnData(context unsafe.Pointer, resultID int32, resultHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedGetReturnData(resultID, resultHandle)
}

// export wasmer1_managedGetMultiESDTCallValue
func wasmer1_managedGetMultiESDTCallValue(context unsafe.Pointer, multiCallValueHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedGetMultiESDTCallValue(multiCallValueHandle)
}

// export wasmer1_managedGetESDTBalance
func wasmer1_managedGetESDTBalance(context unsafe.Pointer, addressHandle int32, tokenIDHandle int32, nonce int64, valueHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedGetESDTBalance(addressHandle, tokenIDHandle, nonce, valueHandle)
}

// export wasmer1_managedGetESDTTokenData
func wasmer1_managedGetESDTTokenData(context unsafe.Pointer, addressHandle int32, tokenIDHandle int32, nonce int64, valueHandle int32, propertiesHandle int32, hashHandle int32, nameHandle int32, attributesHandle int32, creatorHandle int32, royaltiesHandle int32, urisHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedGetESDTTokenData(addressHandle, tokenIDHandle, nonce, valueHandle, propertiesHandle, hashHandle, nameHandle, attributesHandle, creatorHandle, royaltiesHandle, urisHandle)
}

// export wasmer1_managedAsyncCall
func wasmer1_managedAsyncCall(context unsafe.Pointer, destHandle int32, valueHandle int32, functionHandle int32, argumentsHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedAsyncCall(destHandle, valueHandle, functionHandle, argumentsHandle)
}

// export wasmer1_managedCreateAsyncCall
func wasmer1_managedCreateAsyncCall(context unsafe.Pointer, destHandle int32, valueHandle int32, functionHandle int32, argumentsHandle int32, successOffset int32, successLength int32, errorOffset int32, errorLength int32, gas int64, extraGasForCallback int64, callbackClosureHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedCreateAsyncCall(destHandle, valueHandle, functionHandle, argumentsHandle, successOffset, successLength, errorOffset, errorLength, gas, extraGasForCallback, callbackClosureHandle)
}

// export wasmer1_managedGetCallbackClosure
func wasmer1_managedGetCallbackClosure(context unsafe.Pointer, callbackClosureHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedGetCallbackClosure(callbackClosureHandle)
}

// export wasmer1_managedUpgradeFromSourceContract
func wasmer1_managedUpgradeFromSourceContract(context unsafe.Pointer, destHandle int32, gas int64, valueHandle int32, addressHandle int32, codeMetadataHandle int32, argumentsHandle int32, resultHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedUpgradeFromSourceContract(destHandle, gas, valueHandle, addressHandle, codeMetadataHandle, argumentsHandle, resultHandle)
}

// export wasmer1_managedUpgradeContract
func wasmer1_managedUpgradeContract(context unsafe.Pointer, destHandle int32, gas int64, valueHandle int32, codeHandle int32, codeMetadataHandle int32, argumentsHandle int32, resultHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedUpgradeContract(destHandle, gas, valueHandle, codeHandle, codeMetadataHandle, argumentsHandle, resultHandle)
}

// export wasmer1_managedDeleteContract
func wasmer1_managedDeleteContract(context unsafe.Pointer, destHandle int32, gasLimit int64, argumentsHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedDeleteContract(destHandle, gasLimit, argumentsHandle)
}

// export wasmer1_managedDeployFromSourceContract
func wasmer1_managedDeployFromSourceContract(context unsafe.Pointer, gas int64, valueHandle int32, addressHandle int32, codeMetadataHandle int32, argumentsHandle int32, resultAddressHandle int32, resultHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedDeployFromSourceContract(gas, valueHandle, addressHandle, codeMetadataHandle, argumentsHandle, resultAddressHandle, resultHandle)
}

// export wasmer1_managedCreateContract
func wasmer1_managedCreateContract(context unsafe.Pointer, gas int64, valueHandle int32, codeHandle int32, codeMetadataHandle int32, argumentsHandle int32, resultAddressHandle int32, resultHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedCreateContract(gas, valueHandle, codeHandle, codeMetadataHandle, argumentsHandle, resultAddressHandle, resultHandle)
}

// export wasmer1_managedExecuteReadOnly
func wasmer1_managedExecuteReadOnly(context unsafe.Pointer, gas int64, addressHandle int32, functionHandle int32, argumentsHandle int32, resultHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedExecuteReadOnly(gas, addressHandle, functionHandle, argumentsHandle, resultHandle)
}

// export wasmer1_managedExecuteOnSameContext
func wasmer1_managedExecuteOnSameContext(context unsafe.Pointer, gas int64, addressHandle int32, valueHandle int32, functionHandle int32, argumentsHandle int32, resultHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedExecuteOnSameContext(gas, addressHandle, valueHandle, functionHandle, argumentsHandle, resultHandle)
}

// export wasmer1_managedExecuteOnDestContext
func wasmer1_managedExecuteOnDestContext(context unsafe.Pointer, gas int64, addressHandle int32, valueHandle int32, functionHandle int32, argumentsHandle int32, resultHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedExecuteOnDestContext(gas, addressHandle, valueHandle, functionHandle, argumentsHandle, resultHandle)
}

// export wasmer1_managedMultiTransferESDTNFTExecute
func wasmer1_managedMultiTransferESDTNFTExecute(context unsafe.Pointer, dstHandle int32, tokenTransfersHandle int32, gasLimit int64, functionHandle int32, argumentsHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedMultiTransferESDTNFTExecute(dstHandle, tokenTransfersHandle, gasLimit, functionHandle, argumentsHandle)
}

// export wasmer1_managedTransferValueExecute
func wasmer1_managedTransferValueExecute(context unsafe.Pointer, dstHandle int32, valueHandle int32, gasLimit int64, functionHandle int32, argumentsHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedTransferValueExecute(dstHandle, valueHandle, gasLimit, functionHandle, argumentsHandle)
}

// export wasmer1_managedIsESDTFrozen
func wasmer1_managedIsESDTFrozen(context unsafe.Pointer, addressHandle int32, tokenIDHandle int32, nonce int64) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedIsESDTFrozen(addressHandle, tokenIDHandle, nonce)
}

// export wasmer1_managedIsESDTLimitedTransfer
func wasmer1_managedIsESDTLimitedTransfer(context unsafe.Pointer, tokenIDHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedIsESDTLimitedTransfer(tokenIDHandle)
}

// export wasmer1_managedIsESDTPaused
func wasmer1_managedIsESDTPaused(context unsafe.Pointer, tokenIDHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedIsESDTPaused(tokenIDHandle)
}

// export wasmer1_managedBufferToHex
func wasmer1_managedBufferToHex(context unsafe.Pointer, sourceHandle int32, destHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.ManagedBufferToHex(sourceHandle, destHandle)
}

// export wasmer1_mBufferNew
func wasmer1_mBufferNew(context unsafe.Pointer) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferNew()
}

// export wasmer1_mBufferNewFromBytes
func wasmer1_mBufferNewFromBytes(context unsafe.Pointer, dataOffset int32, dataLength int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferNewFromBytes(dataOffset, dataLength)
}

// export wasmer1_mBufferGetLength
func wasmer1_mBufferGetLength(context unsafe.Pointer, mBufferHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferGetLength(mBufferHandle)
}

// export wasmer1_mBufferGetBytes
func wasmer1_mBufferGetBytes(context unsafe.Pointer, mBufferHandle int32, resultOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferGetBytes(mBufferHandle, resultOffset)
}

// export wasmer1_mBufferGetByteSlice
func wasmer1_mBufferGetByteSlice(context unsafe.Pointer, sourceHandle int32, startingPosition int32, sliceLength int32, resultOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferGetByteSlice(sourceHandle, startingPosition, sliceLength, resultOffset)
}

// export wasmer1_mBufferCopyByteSlice
func wasmer1_mBufferCopyByteSlice(context unsafe.Pointer, sourceHandle int32, startingPosition int32, sliceLength int32, destinationHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferCopyByteSlice(sourceHandle, startingPosition, sliceLength, destinationHandle)
}

// export wasmer1_mBufferEq
func wasmer1_mBufferEq(context unsafe.Pointer, mBufferHandle1 int32, mBufferHandle2 int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferEq(mBufferHandle1, mBufferHandle2)
}

// export wasmer1_mBufferSetBytes
func wasmer1_mBufferSetBytes(context unsafe.Pointer, mBufferHandle int32, dataOffset int32, dataLength int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferSetBytes(mBufferHandle, dataOffset, dataLength)
}

// export wasmer1_mBufferSetByteSlice
func wasmer1_mBufferSetByteSlice(context unsafe.Pointer, mBufferHandle int32, startingPosition int32, dataLength int32, dataOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferSetByteSlice(mBufferHandle, startingPosition, dataLength, dataOffset)
}

// export wasmer1_mBufferAppend
func wasmer1_mBufferAppend(context unsafe.Pointer, accumulatorHandle int32, dataHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferAppend(accumulatorHandle, dataHandle)
}

// export wasmer1_mBufferAppendBytes
func wasmer1_mBufferAppendBytes(context unsafe.Pointer, accumulatorHandle int32, dataOffset int32, dataLength int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferAppendBytes(accumulatorHandle, dataOffset, dataLength)
}

// export wasmer1_mBufferToBigIntUnsigned
func wasmer1_mBufferToBigIntUnsigned(context unsafe.Pointer, mBufferHandle int32, bigIntHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferToBigIntUnsigned(mBufferHandle, bigIntHandle)
}

// export wasmer1_mBufferToBigIntSigned
func wasmer1_mBufferToBigIntSigned(context unsafe.Pointer, mBufferHandle int32, bigIntHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferToBigIntSigned(mBufferHandle, bigIntHandle)
}

// export wasmer1_mBufferFromBigIntUnsigned
func wasmer1_mBufferFromBigIntUnsigned(context unsafe.Pointer, mBufferHandle int32, bigIntHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferFromBigIntUnsigned(mBufferHandle, bigIntHandle)
}

// export wasmer1_mBufferFromBigIntSigned
func wasmer1_mBufferFromBigIntSigned(context unsafe.Pointer, mBufferHandle int32, bigIntHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferFromBigIntSigned(mBufferHandle, bigIntHandle)
}

// export wasmer1_mBufferToBigFloat
func wasmer1_mBufferToBigFloat(context unsafe.Pointer, mBufferHandle int32, bigFloatHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferToBigFloat(mBufferHandle, bigFloatHandle)
}

// export wasmer1_mBufferFromBigFloat
func wasmer1_mBufferFromBigFloat(context unsafe.Pointer, mBufferHandle int32, bigFloatHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferFromBigFloat(mBufferHandle, bigFloatHandle)
}

// export wasmer1_mBufferStorageStore
func wasmer1_mBufferStorageStore(context unsafe.Pointer, keyHandle int32, sourceHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferStorageStore(keyHandle, sourceHandle)
}

// export wasmer1_mBufferStorageLoad
func wasmer1_mBufferStorageLoad(context unsafe.Pointer, keyHandle int32, destinationHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferStorageLoad(keyHandle, destinationHandle)
}

// export wasmer1_mBufferStorageLoadFromAddress
func wasmer1_mBufferStorageLoadFromAddress(context unsafe.Pointer, addressHandle int32, keyHandle int32, destinationHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.MBufferStorageLoadFromAddress(addressHandle, keyHandle, destinationHandle)
}

// export wasmer1_mBufferGetArgument
func wasmer1_mBufferGetArgument(context unsafe.Pointer, id int32, destinationHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferGetArgument(id, destinationHandle)
}

// export wasmer1_mBufferFinish
func wasmer1_mBufferFinish(context unsafe.Pointer, sourceHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferFinish(sourceHandle)
}

// export wasmer1_mBufferSetRandom
func wasmer1_mBufferSetRandom(context unsafe.Pointer, destinationHandle int32, length int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MBufferSetRandom(destinationHandle, length)
}

// export wasmer1_smallIntGetUnsignedArgument
func wasmer1_smallIntGetUnsignedArgument(context unsafe.Pointer, id int32) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.SmallIntGetUnsignedArgument(id)
}

// export wasmer1_smallIntGetSignedArgument
func wasmer1_smallIntGetSignedArgument(context unsafe.Pointer, id int32) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.SmallIntGetSignedArgument(id)
}

// export wasmer1_smallIntFinishUnsigned
func wasmer1_smallIntFinishUnsigned(context unsafe.Pointer, value int64) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.SmallIntFinishUnsigned(value)
}

// export wasmer1_smallIntFinishSigned
func wasmer1_smallIntFinishSigned(context unsafe.Pointer, value int64) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.SmallIntFinishSigned(value)
}

// export wasmer1_smallIntStorageStoreUnsigned
func wasmer1_smallIntStorageStoreUnsigned(context unsafe.Pointer, keyOffset int32, keyLength int32, value int64) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.SmallIntStorageStoreUnsigned(keyOffset, keyLength, value)
}

// export wasmer1_smallIntStorageStoreSigned
func wasmer1_smallIntStorageStoreSigned(context unsafe.Pointer, keyOffset int32, keyLength int32, value int64) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.SmallIntStorageStoreSigned(keyOffset, keyLength, value)
}

// export wasmer1_smallIntStorageLoadUnsigned
func wasmer1_smallIntStorageLoadUnsigned(context unsafe.Pointer, keyOffset int32, keyLength int32) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.SmallIntStorageLoadUnsigned(keyOffset, keyLength)
}

// export wasmer1_smallIntStorageLoadSigned
func wasmer1_smallIntStorageLoadSigned(context unsafe.Pointer, keyOffset int32, keyLength int32) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.SmallIntStorageLoadSigned(keyOffset, keyLength)
}

// export wasmer1_int64getArgument
func wasmer1_int64getArgument(context unsafe.Pointer, id int32) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.Int64getArgument(id)
}

// export wasmer1_int64finish
func wasmer1_int64finish(context unsafe.Pointer, value int64) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.Int64finish(value)
}

// export wasmer1_int64storageStore
func wasmer1_int64storageStore(context unsafe.Pointer, keyOffset int32, keyLength int32, value int64) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.Int64storageStore(keyOffset, keyLength, value)
}

// export wasmer1_int64storageLoad
func wasmer1_int64storageLoad(context unsafe.Pointer, keyOffset int32, keyLength int32) int64 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.Int64storageLoad(keyOffset, keyLength)
}

// export wasmer1_sha256
func wasmer1_sha256(context unsafe.Pointer, dataOffset int32, length int32, resultOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.Sha256(dataOffset, length, resultOffset)
}

// export wasmer1_managedSha256
func wasmer1_managedSha256(context unsafe.Pointer, inputHandle int32, outputHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedSha256(inputHandle, outputHandle)
}

// export wasmer1_keccak256
func wasmer1_keccak256(context unsafe.Pointer, dataOffset int32, length int32, resultOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.Keccak256(dataOffset, length, resultOffset)
}

// export wasmer1_managedKeccak256
func wasmer1_managedKeccak256(context unsafe.Pointer, inputHandle int32, outputHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedKeccak256(inputHandle, outputHandle)
}

// export wasmer1_ripemd160
func wasmer1_ripemd160(context unsafe.Pointer, dataOffset int32, length int32, resultOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.Ripemd160(dataOffset, length, resultOffset)
}

// export wasmer1_managedRipemd160
func wasmer1_managedRipemd160(context unsafe.Pointer, inputHandle int32, outputHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedRipemd160(inputHandle, outputHandle)
}

// export wasmer1_verifyBLS
func wasmer1_verifyBLS(context unsafe.Pointer, keyOffset int32, messageOffset int32, messageLength int32, sigOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.VerifyBLS(keyOffset, messageOffset, messageLength, sigOffset)
}

// export wasmer1_managedVerifyBLS
func wasmer1_managedVerifyBLS(context unsafe.Pointer, keyHandle int32, messageHandle int32, sigHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedVerifyBLS(keyHandle, messageHandle, sigHandle)
}

// export wasmer1_verifyEd25519
func wasmer1_verifyEd25519(context unsafe.Pointer, keyOffset int32, messageOffset int32, messageLength int32, sigOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.VerifyEd25519(keyOffset, messageOffset, messageLength, sigOffset)
}

// export wasmer1_managedVerifyEd25519
func wasmer1_managedVerifyEd25519(context unsafe.Pointer, keyHandle int32, messageHandle int32, sigHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedVerifyEd25519(keyHandle, messageHandle, sigHandle)
}

// export wasmer1_verifyCustomSecp256k1
func wasmer1_verifyCustomSecp256k1(context unsafe.Pointer, keyOffset int32, keyLength int32, messageOffset int32, messageLength int32, sigOffset int32, hashType int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.VerifyCustomSecp256k1(keyOffset, keyLength, messageOffset, messageLength, sigOffset, hashType)
}

// export wasmer1_managedVerifyCustomSecp256k1
func wasmer1_managedVerifyCustomSecp256k1(context unsafe.Pointer, keyHandle int32, messageHandle int32, sigHandle int32, hashType int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedVerifyCustomSecp256k1(keyHandle, messageHandle, sigHandle, hashType)
}

// export wasmer1_verifySecp256k1
func wasmer1_verifySecp256k1(context unsafe.Pointer, keyOffset int32, keyLength int32, messageOffset int32, messageLength int32, sigOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.VerifySecp256k1(keyOffset, keyLength, messageOffset, messageLength, sigOffset)
}

// export wasmer1_managedVerifySecp256k1
func wasmer1_managedVerifySecp256k1(context unsafe.Pointer, keyHandle int32, messageHandle int32, sigHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedVerifySecp256k1(keyHandle, messageHandle, sigHandle)
}

// export wasmer1_encodeSecp256k1DerSignature
func wasmer1_encodeSecp256k1DerSignature(context unsafe.Pointer, rOffset int32, rLength int32, sOffset int32, sLength int32, sigOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.EncodeSecp256k1DerSignature(rOffset, rLength, sOffset, sLength, sigOffset)
}

// export wasmer1_managedEncodeSecp256k1DerSignature
func wasmer1_managedEncodeSecp256k1DerSignature(context unsafe.Pointer, rHandle int32, sHandle int32, sigHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedEncodeSecp256k1DerSignature(rHandle, sHandle, sigHandle)
}

// export wasmer1_addEC
func wasmer1_addEC(context unsafe.Pointer, xResultHandle int32, yResultHandle int32, ecHandle int32, fstPointXHandle int32, fstPointYHandle int32, sndPointXHandle int32, sndPointYHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.AddEC(xResultHandle, yResultHandle, ecHandle, fstPointXHandle, fstPointYHandle, sndPointXHandle, sndPointYHandle)
}

// export wasmer1_doubleEC
func wasmer1_doubleEC(context unsafe.Pointer, xResultHandle int32, yResultHandle int32, ecHandle int32, pointXHandle int32, pointYHandle int32) {
	callbacks := importsInterfaceFromRaw(context)
	callbacks.DoubleEC(xResultHandle, yResultHandle, ecHandle, pointXHandle, pointYHandle)
}

// export wasmer1_isOnCurveEC
func wasmer1_isOnCurveEC(context unsafe.Pointer, ecHandle int32, pointXHandle int32, pointYHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.IsOnCurveEC(ecHandle, pointXHandle, pointYHandle)
}

// export wasmer1_scalarBaseMultEC
func wasmer1_scalarBaseMultEC(context unsafe.Pointer, xResultHandle int32, yResultHandle int32, ecHandle int32, dataOffset int32, length int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ScalarBaseMultEC(xResultHandle, yResultHandle, ecHandle, dataOffset, length)
}

// export wasmer1_managedScalarBaseMultEC
func wasmer1_managedScalarBaseMultEC(context unsafe.Pointer, xResultHandle int32, yResultHandle int32, ecHandle int32, dataHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedScalarBaseMultEC(xResultHandle, yResultHandle, ecHandle, dataHandle)
}

// export wasmer1_scalarMultEC
func wasmer1_scalarMultEC(context unsafe.Pointer, xResultHandle int32, yResultHandle int32, ecHandle int32, pointXHandle int32, pointYHandle int32, dataOffset int32, length int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ScalarMultEC(xResultHandle, yResultHandle, ecHandle, pointXHandle, pointYHandle, dataOffset, length)
}

// export wasmer1_managedScalarMultEC
func wasmer1_managedScalarMultEC(context unsafe.Pointer, xResultHandle int32, yResultHandle int32, ecHandle int32, pointXHandle int32, pointYHandle int32, dataHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedScalarMultEC(xResultHandle, yResultHandle, ecHandle, pointXHandle, pointYHandle, dataHandle)
}

// export wasmer1_marshalEC
func wasmer1_marshalEC(context unsafe.Pointer, xPairHandle int32, yPairHandle int32, ecHandle int32, resultOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MarshalEC(xPairHandle, yPairHandle, ecHandle, resultOffset)
}

// export wasmer1_managedMarshalEC
func wasmer1_managedMarshalEC(context unsafe.Pointer, xPairHandle int32, yPairHandle int32, ecHandle int32, resultHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedMarshalEC(xPairHandle, yPairHandle, ecHandle, resultHandle)
}

// export wasmer1_marshalCompressedEC
func wasmer1_marshalCompressedEC(context unsafe.Pointer, xPairHandle int32, yPairHandle int32, ecHandle int32, resultOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.MarshalCompressedEC(xPairHandle, yPairHandle, ecHandle, resultOffset)
}

// export wasmer1_managedMarshalCompressedEC
func wasmer1_managedMarshalCompressedEC(context unsafe.Pointer, xPairHandle int32, yPairHandle int32, ecHandle int32, resultHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedMarshalCompressedEC(xPairHandle, yPairHandle, ecHandle, resultHandle)
}

// export wasmer1_unmarshalEC
func wasmer1_unmarshalEC(context unsafe.Pointer, xResultHandle int32, yResultHandle int32, ecHandle int32, dataOffset int32, length int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.UnmarshalEC(xResultHandle, yResultHandle, ecHandle, dataOffset, length)
}

// export wasmer1_managedUnmarshalEC
func wasmer1_managedUnmarshalEC(context unsafe.Pointer, xResultHandle int32, yResultHandle int32, ecHandle int32, dataHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedUnmarshalEC(xResultHandle, yResultHandle, ecHandle, dataHandle)
}

// export wasmer1_unmarshalCompressedEC
func wasmer1_unmarshalCompressedEC(context unsafe.Pointer, xResultHandle int32, yResultHandle int32, ecHandle int32, dataOffset int32, length int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.UnmarshalCompressedEC(xResultHandle, yResultHandle, ecHandle, dataOffset, length)
}

// export wasmer1_managedUnmarshalCompressedEC
func wasmer1_managedUnmarshalCompressedEC(context unsafe.Pointer, xResultHandle int32, yResultHandle int32, ecHandle int32, dataHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedUnmarshalCompressedEC(xResultHandle, yResultHandle, ecHandle, dataHandle)
}

// export wasmer1_generateKeyEC
func wasmer1_generateKeyEC(context unsafe.Pointer, xPubKeyHandle int32, yPubKeyHandle int32, ecHandle int32, resultOffset int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GenerateKeyEC(xPubKeyHandle, yPubKeyHandle, ecHandle, resultOffset)
}

// export wasmer1_managedGenerateKeyEC
func wasmer1_managedGenerateKeyEC(context unsafe.Pointer, xPubKeyHandle int32, yPubKeyHandle int32, ecHandle int32, resultHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedGenerateKeyEC(xPubKeyHandle, yPubKeyHandle, ecHandle, resultHandle)
}

// export wasmer1_createEC
func wasmer1_createEC(context unsafe.Pointer, dataOffset int32, dataLength int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.CreateEC(dataOffset, dataLength)
}

// export wasmer1_managedCreateEC
func wasmer1_managedCreateEC(context unsafe.Pointer, dataHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.ManagedCreateEC(dataHandle)
}

// export wasmer1_getCurveLengthEC
func wasmer1_getCurveLengthEC(context unsafe.Pointer, ecHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetCurveLengthEC(ecHandle)
}

// export wasmer1_getPrivKeyByteLengthEC
func wasmer1_getPrivKeyByteLengthEC(context unsafe.Pointer, ecHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.GetPrivKeyByteLengthEC(ecHandle)
}

// export wasmer1_ellipticCurveGetValues
func wasmer1_ellipticCurveGetValues(context unsafe.Pointer, ecHandle int32, fieldOrderHandle int32, basePointOrderHandle int32, eqConstantHandle int32, xBasePointHandle int32, yBasePointHandle int32) int32 {
	callbacks := importsInterfaceFromRaw(context)
	return callbacks.EllipticCurveGetValues(ecHandle, fieldOrderHandle, basePointOrderHandle, eqConstantHandle, xBasePointHandle, yBasePointHandle)
}
