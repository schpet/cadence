/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright 2019-2022 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package value_codec

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/big"

	"github.com/onflow/cadence"
	"github.com/onflow/cadence/encoding/custom/common_codec"
	"github.com/onflow/cadence/runtime/common"
)

// A Decoder decodes custom-encoded representations of Cadence values.
type Decoder struct {
	r           common_codec.LocatedReader
	buf         []byte
	memoryGauge common.MemoryGauge
	// TODO abi for cutting down on what needs to be transferred
}

// Decode returns a Cadence value decoded from its custom-encoded representation.
//
// This function returns an error if the bytes represent a custom encoding that
// is malformed, does not conform to the custom Cadence specification, or contains
// an unknown composite type.
func Decode(gauge common.MemoryGauge, b []byte) (cadence.Value, error) {
	r := bytes.NewReader(b)
	dec := NewDecoder(gauge, r)

	v, err := dec.Decode()
	if err != nil {
		return nil, err
	}

	return v, nil
}

// NewDecoder initializes a Decoder that will decode custom-encoded bytes from the
// given io.Reader.
func NewDecoder(memoryGauge common.MemoryGauge, r io.Reader) *Decoder {
	return &Decoder{
		r:           common_codec.NewLocatedReader(r),
		memoryGauge: memoryGauge,
	}
}

// Decode reads custom-encoded bytes from the io.Reader and decodes them to a
// Cadence value.
//
// This function returns an error if the bytes represent a custom encoding that
// is malformed, does not conform to the custom Cadence specification, or contains
// an unknown composite type.
func (d *Decoder) Decode() (value cadence.Value, err error) {
	return d.DecodeValue()
}

// TODO need a way to decode values with known type vs values with unknown type
//      if type is known then no identifier is needed, such as for elements in constant sized array
func (d *Decoder) DecodeValue() (value cadence.Value, err error) {
	identifier, err := d.DecodeIdentifier()
	if err != nil {
		return
	}

	switch identifier {
	case EncodedValueVoid:
		value = cadence.NewMeteredVoid(d.memoryGauge)
	case EncodedValueOptional:
		value, err = d.DecodeOptional()
	case EncodedValueBool:
		value, err = d.DecodeBool()
	case EncodedValueString:
		value, err = d.DecodeString()
	case EncodedValueBytes:
		value, err = d.DecodeBytes()
	case EncodedValueCharacter:
		value, err = d.DecodeCharacter()
	case EncodedValueAddress:
		value, err = d.DecodeAddress()
	case EncodedValueInt:
		value, err = d.DecodeInt()
	case EncodedValueInt8:
		value, err = d.DecodeInt8()
	case EncodedValueInt16:
		value, err = d.DecodeInt16()
	case EncodedValueInt32:
		value, err = d.DecodeInt32()
	case EncodedValueInt64:
		value, err = d.DecodeInt64()
	case EncodedValueInt128:
		value, err = d.DecodeInt128()
	case EncodedValueInt256:
		value, err = d.DecodeInt256()
	case EncodedValueUInt:
		value, err = d.DecodeUInt()
	case EncodedValueUInt8:
		value, err = d.DecodeUInt8()
	case EncodedValueUInt16:
		value, err = d.DecodeUInt16()
	case EncodedValueUInt32:
		value, err = d.DecodeUInt32()
	case EncodedValueUInt64:
		value, err = d.DecodeUInt64()
	case EncodedValueUInt128:
		value, err = d.DecodeUInt128()
	case EncodedValueUInt256:
		value, err = d.DecodeUInt256()
	case EncodedValueWord8:
		value, err = d.DecodeWord8()
	case EncodedValueWord16:
		value, err = d.DecodeWord16()
	case EncodedValueWord32:
		value, err = d.DecodeWord32()
	case EncodedValueWord64:
		value, err = d.DecodeWord64()
	case EncodedValueFix64:
		value, err = d.DecodeFix64()
	case EncodedValueUFix64:
		value, err = d.DecodeUFix64()

	case EncodedValueVariableArray:
		var t cadence.VariableSizedArrayType
		t, err = d.DecodeVariableArrayType()
		if err != nil {
			return
		}
		value, err = d.DecodeVariableArray(t)
	case EncodedValueConstantArray:
		var t cadence.ConstantSizedArrayType
		t, err = d.DecodeConstantArrayType()
		if err != nil {
			return
		}
		value, err = d.DecodeConstantArray(t)
	case EncodedValueDictionary:
		value, err = d.DecodeDictionary()

	default:
		err = fmt.Errorf("unknown cadence.Value: %s", value)
	}

	return
}

func (d *Decoder) DecodeIdentifier() (id EncodedValue, err error) {
	b, err := d.read(1)
	if err != nil {
		return
	}

	id = EncodedValue(b[0])
	return
}

func (d *Decoder) DecodeVoid() (value cadence.Void, err error) {
	_, err = d.read(1)
	value = cadence.NewMeteredVoid(d.memoryGauge)
	return
}

func (d *Decoder) DecodeOptional() (value cadence.Optional, err error) {
	isNil, err := d.DecodeBool()
	if isNil || err != nil {
		return
	}

	innerValue, err := d.DecodeValue()
	value = cadence.NewMeteredOptional(d.memoryGauge, innerValue)
	return
}

func (d *Decoder) DecodeBool() (value cadence.Bool, err error) {
	boolean, err := common_codec.DecodeBool(&d.r)
	if err != nil {
		return
	}

	value = cadence.NewMeteredBool(d.memoryGauge, boolean)
	return
}

func (d *Decoder) DecodeString() (value cadence.String, err error) {
	s, err := common_codec.DecodeString(&d.r)
	if err != nil {
		return
	}

	value, err = cadence.NewMeteredString(
		d.memoryGauge,
		common.NewCadenceStringMemoryUsage(len(s)),
		func() string {
			return s
		},
	)
	return
}

func (d *Decoder) DecodeCharacter() (value cadence.Character, err error) {
	s, err := common_codec.DecodeString(&d.r)
	if err != nil {
		return
	}

	value, err = cadence.NewMeteredCharacter(
		d.memoryGauge,
		common.NewCadenceStringMemoryUsage(len(s)),
		func() string {
			return s
		},
	)
	return
}

func (d *Decoder) DecodeAddress() (value cadence.Address, err error) {
	address, err := common_codec.DecodeAddress(&d.r)
	if err != nil {
		return
	}

	value = cadence.NewMeteredAddress(
		d.memoryGauge,
		address,
	)
	return
}

func (d *Decoder) DecodeInt() (value cadence.Int, err error) {
	i, err := common_codec.DecodeBigInt(&d.r)
	if err != nil {
		return
	}

	value = cadence.NewMeteredIntFromBig(
		d.memoryGauge,
		common.NewBigIntMemoryUsage(common.BigIntByteLength(i)),
		func() *big.Int {
			return i
		},
	)
	return
}

func (d *Decoder) DecodeInt8() (value cadence.Int8, err error) {
	i, err := common_codec.DecodeNumber[int8](&d.r)
	value = cadence.Int8(i)
	return
}

func (d *Decoder) DecodeInt16() (value cadence.Int16, err error) {
	i, err := common_codec.DecodeNumber[int16](&d.r)
	value = cadence.Int16(i)
	return
}

func (d *Decoder) DecodeInt32() (value cadence.Int32, err error) {
	i, err := common_codec.DecodeNumber[int32](&d.r)
	value = cadence.Int32(i)
	return
}

func (d *Decoder) DecodeInt64() (value cadence.Int64, err error) {
	i, err := common_codec.DecodeNumber[int64](&d.r)
	value = cadence.Int64(i)
	return
}

func (d *Decoder) DecodeInt128() (value cadence.Int128, err error) {
	i, err := common_codec.DecodeBigInt(&d.r)
	if err != nil {
		return
	}

	return cadence.NewMeteredInt128FromBig(
		d.memoryGauge,
		func() *big.Int {
			return i
		},
	)
}

func (d *Decoder) DecodeInt256() (value cadence.Int256, err error) {
	i, err := common_codec.DecodeBigInt(&d.r)
	if err != nil {
		return
	}

	return cadence.NewMeteredInt256FromBig(
		d.memoryGauge,
		func() *big.Int {
			return i
		},
	)
}

func (d *Decoder) DecodeUInt() (value cadence.UInt, err error) {
	i, err := common_codec.DecodeBigInt(&d.r)
	if err != nil {
		return
	}

	return cadence.NewMeteredUIntFromBig(
		d.memoryGauge,
		common.NewBigIntMemoryUsage(common.BigIntByteLength(i)),
		func() *big.Int {
			return i
		},
	)
}

func (d *Decoder) DecodeUInt8() (value cadence.UInt8, err error) {
	i, err := common_codec.DecodeNumber[uint8](&d.r)
	value = cadence.UInt8(i)
	return
}

func (d *Decoder) DecodeUInt16() (value cadence.UInt16, err error) {
	i, err := common_codec.DecodeNumber[uint16](&d.r)
	value = cadence.UInt16(i)
	return
}

func (d *Decoder) DecodeUInt32() (value cadence.UInt32, err error) {
	i, err := common_codec.DecodeNumber[uint32](&d.r)
	value = cadence.UInt32(i)
	return
}

func (d *Decoder) DecodeUInt64() (value cadence.UInt64, err error) {
	i, err := common_codec.DecodeNumber[uint64](&d.r)
	value = cadence.UInt64(i)
	return
}

func (d *Decoder) DecodeUInt128() (value cadence.UInt128, err error) {
	i, err := common_codec.DecodeBigInt(&d.r)
	if err != nil {
		return
	}

	return cadence.NewMeteredUInt128FromBig(
		d.memoryGauge,
		func() *big.Int {
			return i
		},
	)
}

func (d *Decoder) DecodeUInt256() (value cadence.UInt256, err error) {
	i, err := common_codec.DecodeBigInt(&d.r)
	if err != nil {
		return
	}

	return cadence.NewMeteredUInt256FromBig(
		d.memoryGauge,
		func() *big.Int {
			return i
		},
	)
}

func (d *Decoder) DecodeWord8() (value cadence.Word8, err error) {
	i, err := common_codec.DecodeNumber[uint8](&d.r)
	if err != nil {
		return
	}
	value = cadence.Word8(i)
	return
}

func (d *Decoder) DecodeWord16() (value cadence.Word16, err error) {
	i, err := common_codec.DecodeNumber[uint16](&d.r)
	if err != nil {
		return
	}
	value = cadence.Word16(i)
	return
}

func (d *Decoder) DecodeWord32() (value cadence.Word32, err error) {
	i, err := common_codec.DecodeNumber[uint32](&d.r)
	if err != nil {
		return
	}
	value = cadence.Word32(i)
	return
}

func (d *Decoder) DecodeWord64() (value cadence.Word64, err error) {
	i, err := common_codec.DecodeNumber[uint64](&d.r)
	if err != nil {
		return
	}
	value = cadence.Word64(i)
	return
}

func (d *Decoder) DecodeFix64() (value cadence.Fix64, err error) {
	i, err := common_codec.DecodeNumber[int64](&d.r)
	if err != nil {
		return
	}
	value = cadence.Fix64(i)
	return
}

func (d *Decoder) DecodeUFix64() (value cadence.UFix64, err error) {
	i, err := common_codec.DecodeNumber[uint64](&d.r)
	if err != nil {
		return
	}
	value = cadence.UFix64(i)
	return
}

func (d *Decoder) DecodeBytes() (value cadence.Bytes, err error) {
	s, err := common_codec.DecodeBytes(&d.r)
	if err != nil {
		return
	}

	value = cadence.NewBytes(s)
	return
}

func (d *Decoder) DecodeVariableArray(arrayType cadence.VariableSizedArrayType) (array cadence.Array, err error) {
	size, err := d.DecodeLength()
	if err != nil {
		return
	}
	array, err = cadence.NewMeteredArray(d.memoryGauge, size, func() (elements []cadence.Value, err error) {
		elements = make([]cadence.Value, 0, size)
		for i := 0; i < size; i++ {
			// TODO if `elementType` is concrete then each element needn't encode its type
			var elementValue cadence.Value
			elementValue, err = d.DecodeValue()
			if err != nil {
				return
			}
			elements = append(elements, elementValue)
		}

		return elements, nil
	})

	array = array.WithType(arrayType)

	return
}

func (d *Decoder) DecodeConstantArray(arrayType cadence.ConstantSizedArrayType) (array cadence.Array, err error) {
	size := int(arrayType.Size)
	array, err = cadence.NewMeteredArray(d.memoryGauge, size, func() (elements []cadence.Value, err error) {
		elements = make([]cadence.Value, 0, size)
		for i := 0; i < size; i++ {
			// TODO if `elementType` is concrete then each element needn't encode its type
			var elementValue cadence.Value
			elementValue, err = d.DecodeValue()
			if err != nil {
				return
			}
			elements = append(elements, elementValue)
		}

		return elements, nil
	})

	array = array.WithType(arrayType)

	return
}

func (d *Decoder) DecodeDictionary() (dict cadence.Dictionary, err error) {
	dictType, err := d.DecodeDictionaryType()
	if err != nil {
		return
	}

	size, err := d.DecodeLength()
	if err != nil {
		return
	}

	dict, err = cadence.NewMeteredDictionary(d.memoryGauge, size, func() (pairs []cadence.KeyValuePair, err error) {
		pairs = make([]cadence.KeyValuePair, 0, size)
		var key, value cadence.Value
		for i := 0; i < size; i++ {
			key, err = d.DecodeValue()
			if err != nil {
				return
			}
			value, err = d.DecodeValue()
			if err != nil {
				return
			}
			pairs = append(pairs, cadence.NewMeteredKeyValuePair(d.memoryGauge, key, value))
		}
		return
	})

	dict = dict.WithType(dictType)

	return
}

//
// Types
//

func (d *Decoder) DecodeType() (t cadence.Type, err error) {
	typeIdentifer, err := d.DecodeTypeIdentifier()

	switch typeIdentifer {
	case EncodedTypeVoid:
		t = cadence.NewMeteredVoidType(d.memoryGauge)
	case EncodedTypeNever:
		t = cadence.NewMeteredNeverType(d.memoryGauge)
	case EncodedTypeOptional:
		t, err = d.DecodeOptionalType()
	case EncodedTypeBool:
		t = cadence.NewMeteredBoolType(d.memoryGauge)
	case EncodedTypeString:
		t = cadence.NewMeteredStringType(d.memoryGauge)
	case EncodedTypeCharacter:
		t = cadence.NewMeteredCharacterType(d.memoryGauge)
	case EncodedTypeBytes:
		t = cadence.NewMeteredBytesType(d.memoryGauge)
	case EncodedTypeAddress:
		t = cadence.NewMeteredAddressType(d.memoryGauge)
	case EncodedTypeNumber:
		t = cadence.NewMeteredNumberType(d.memoryGauge)
	case EncodedTypeSignedNumber:
		t = cadence.NewMeteredSignedNumberType(d.memoryGauge)
	case EncodedTypeInteger:
		t = cadence.NewMeteredIntegerType(d.memoryGauge)
	case EncodedTypeSignedInteger:
		t = cadence.NewMeteredSignedIntegerType(d.memoryGauge)
	case EncodedTypeFixedPoint:
		t = cadence.NewMeteredFixedPointType(d.memoryGauge)
	case EncodedTypeSignedFixedPoint:
		t = cadence.NewMeteredSignedFixedPointType(d.memoryGauge)
	case EncodedTypeInt:
		t = cadence.NewMeteredIntType(d.memoryGauge)
	case EncodedTypeInt8:
		t = cadence.NewMeteredInt8Type(d.memoryGauge)
	case EncodedTypeInt16:
		t = cadence.NewMeteredInt16Type(d.memoryGauge)
	case EncodedTypeInt32:
		t = cadence.NewMeteredInt32Type(d.memoryGauge)
	case EncodedTypeInt64:
		t = cadence.NewMeteredInt64Type(d.memoryGauge)
	case EncodedTypeInt128:
		t = cadence.NewMeteredInt128Type(d.memoryGauge)
	case EncodedTypeInt256:
		t = cadence.NewMeteredInt256Type(d.memoryGauge)
	case EncodedTypeUInt128:
		t = cadence.NewMeteredUInt128Type(d.memoryGauge)
	case EncodedTypeUInt256:
		t = cadence.NewMeteredUInt256Type(d.memoryGauge)
	case EncodedTypeUInt:
		t = cadence.NewMeteredUIntType(d.memoryGauge)
	case EncodedTypeUInt8:
		t = cadence.NewMeteredUInt8Type(d.memoryGauge)
	case EncodedTypeUInt16:
		t = cadence.NewMeteredUInt16Type(d.memoryGauge)
	case EncodedTypeUInt32:
		t = cadence.NewMeteredUInt32Type(d.memoryGauge)
	case EncodedTypeUInt64:
		t = cadence.NewMeteredUInt64Type(d.memoryGauge)
	case EncodedTypeWord8:
		t = cadence.NewMeteredWord8Type(d.memoryGauge)
	case EncodedTypeWord16:
		t = cadence.NewMeteredWord16Type(d.memoryGauge)
	case EncodedTypeWord32:
		t = cadence.NewMeteredWord32Type(d.memoryGauge)
	case EncodedTypeWord64:
		t = cadence.NewMeteredWord64Type(d.memoryGauge)
	case EncodedTypeFix64:
		t = cadence.NewMeteredFix64Type(d.memoryGauge)
	case EncodedTypeUFix64:
		t = cadence.NewMeteredUFix64Type(d.memoryGauge)
	case EncodedTypeVariableSizedArray:
		t, err = d.DecodeVariableArrayType()
	case EncodedTypeConstantSizedArray:
		t, err = d.DecodeConstantArrayType()
	case EncodedTypeDictionary:
		t, err = d.DecodeDictionaryType()

	case EncodedTypeAnyType:
		t = cadence.NewMeteredAnyType(d.memoryGauge)
	case EncodedTypeAnyStructType:
		t = cadence.NewMeteredAnyStructType(d.memoryGauge)
	case EncodedTypeAnyResourceType:
		t = cadence.NewMeteredAnyResourceType(d.memoryGauge)
	default:
		err = fmt.Errorf("unknown type identifier: %d", typeIdentifer)
	}
	return
}

func (d *Decoder) DecodeTypeIdentifier() (t EncodedType, err error) {
	b, err := d.read(1)
	t = EncodedType(b[0])
	return
}

func (d *Decoder) DecodeOptionalType() (t cadence.OptionalType, err error) {
	isNil, err := common_codec.DecodeBool(&d.r)
	if isNil || err != nil {
		return
	}

	elementType, err := d.DecodeType()
	if err != nil {
		return
	}

	t = cadence.NewMeteredOptionalType(d.memoryGauge, elementType)
	return
}

func (d *Decoder) DecodeVariableArrayType() (t cadence.VariableSizedArrayType, err error) {
	elementType, err := d.DecodeType()
	if err != nil {
		return
	}

	t = cadence.NewMeteredVariableSizedArrayType(d.memoryGauge, elementType)
	return
}

func (d *Decoder) DecodeConstantArrayType() (t cadence.ConstantSizedArrayType, err error) {
	elementType, err := d.DecodeType()
	if err != nil {
		return
	}

	size, err := d.DecodeLength()
	if err != nil {
		return
	}
	t = cadence.NewMeteredConstantSizedArrayType(d.memoryGauge, uint(size), elementType)
	return
}

func (d *Decoder) DecodeDictionaryType() (t cadence.DictionaryType, err error) {
	keyType, err := d.DecodeType()
	if err != nil {
		return
	}
	elementType, err := d.DecodeType()
	if err != nil {
		return
	}
	t = cadence.NewMeteredDictionaryType(d.memoryGauge, keyType, elementType)
	return
}

//
// Other
//

func (d *Decoder) DecodeLength() (length int, err error) {
	b, err := d.read(4)
	if err != nil {
		return
	}

	asUint32 := binary.BigEndian.Uint32(b)

	length = int(asUint32)
	return
}

func (d *Decoder) read(howManyBytes int) (b []byte, err error) {
	b = make([]byte, howManyBytes)
	_, err = d.r.Read(b)
	return
}
