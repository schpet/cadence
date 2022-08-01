package value_codec

type EncodedValue byte

const (
	EncodedValueUnknown EncodedValue = iota

	EncodedValueVoid
	EncodedValueOptional
	EncodedValueBool
	EncodedValueString
	EncodedValueBytes
	EncodedValueCharacter
	EncodedValueAddress
	EncodedValueInt
	EncodedValueInt8
	EncodedValueInt16
	EncodedValueInt32
	EncodedValueInt64
	EncodedValueInt128
	EncodedValueInt256
	EncodedValueUInt
	EncodedValueUInt8
	EncodedValueUInt16
	EncodedValueUInt32
	EncodedValueUInt64
	EncodedValueUInt128
	EncodedValueUInt256
	EncodedValueWord8
	EncodedValueWord16
	EncodedValueWord32
	EncodedValueWord64
	EncodedValueFix64
	EncodedValueUFix64
	EncodedValueArray
	EncodedValueDictionary
	EncodedValueStruct
	EncodedValueResource
	EncodedValueEvebt
	EncodedValueContract
	EncodedValueLink
	EncodedValuePath
	EncodedValueCapability
	EncodedValueEnum
)

type EncodedType byte

const (
	EncodedTypeUnknown EncodedType = iota

	// TODO classify these types, probably as simple, complex, or abstract

	// Concrete Types

	EncodedTypeVoid
	EncodedTypeNever
	EncodedTypeBool
	EncodedTypeArray
	EncodedTypeOptional
	EncodedTypeString
	EncodedTypeCharacter
	EncodedTypeBytes
	EncodedTypeAddress
	EncodedTypeNumber
	EncodedTypeSignedNumber
	EncodedTypeInteger
	EncodedTypeSignedInteger
	EncodedTypeFixedPoint
	EncodedTypeSignedFixedPoint
	EncodedTypeInt
	EncodedTypeInt8
	EncodedTypeInt16
	EncodedTypeInt32
	EncodedTypeInt64
	EncodedTypeInt128
	EncodedTypeInt256
	EncodedTypeUInt
	EncodedTypeUInt8
	EncodedTypeUInt16
	EncodedTypeUInt32
	EncodedTypeUInt64
	EncodedTypeUInt128
	EncodedTypeUInt256
	EncodedTypeWord8
	EncodedTypeWord16
	EncodedTypeWord32
	EncodedTypeWord64
	EncodedTypeFix64
	EncodedTypeUFix64
	EncodedTypeVariableSizedArray
	EncodedTypeConstantSizedArray
	EncodedTypeDictionary
	EncodedTypeStruct
	EncodedTypeResource
	EncodedTypeEvent
	EncodedTypeContract
	EncodedTypeStructInterface
	EncodedTypeResourceInterface
	EncodedTypeContractInterface
	EncodedTypeFunction
	EncodedTypeReference
	EncodedTypeRestricted
	EncodedTypeBlock
	EncodedTypeCapabilityPath
	EncodedTypeStoragePath
	EncodedTypePublicPath
	EncodedTypePrivatePath
	EncodedTypeCapability
	EncodedTypeEnum
	EncodedTypeAuthAccount
	EncodedTypePublicAccount
	EncodedTypeDeployedContract
	EncodedTypeAuthAccountContracts
	EncodedTypePublicAccountContracts
	EncodedTypeAuthAccountKeys
	EncodedTypePublicAccountKeys

	// Abstract Types

	EncodedTypeAnyType
	EncodedTypeAnyStructType
	EncodedTypeAnyResourceType

	EncodedTypeArrau     // TODO is this necessary?
	EncodedTypeComposite // TODO is this necessary?
	EncodedTypeInterface // TODO is this necessary?

	// TODO - classify

	EncodedTypeMetaType
)

type EncodedArrayType byte

const (
	EncodedArrayTypeUnknown EncodedArrayType = iota
	EncodedArrayTypeVariable
	EncodedArrayTypeConstant
)
