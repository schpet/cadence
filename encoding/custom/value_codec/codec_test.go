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

package value_codec_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence"
	"github.com/onflow/cadence/encoding/custom/common_codec"

	"github.com/onflow/cadence/encoding/custom/value_codec"
)

func TestValueCodecVoid(t *testing.T) {
	t.Parallel()

	t.Run("value", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		value := cadence.NewVoid()

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedValueVoid)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewVoidType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeVoid)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestValueCodecBool(t *testing.T) {
	t.Parallel()

	t.Run("false", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		value := cadence.NewBool(false)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedValueBool),
				byte(common_codec.EncodedBoolFalse),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("true", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		value := cadence.NewBool(true)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedValueBool),
				byte(common_codec.EncodedBoolTrue),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewBoolType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeBool)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestValueCodecOptional(t *testing.T) {
	t.Parallel()

	t.Run("Optional(Void)", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		innerValue := cadence.NewVoid()
		value := cadence.NewOptional(innerValue)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedValueOptional),
				byte(common_codec.EncodedBoolFalse),
				byte(value_codec.EncodedValueVoid),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("Optional(bool)", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		innerValue := cadence.NewBool(true)
		value := cadence.NewOptional(innerValue)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedValueOptional),
				byte(common_codec.EncodedBoolFalse),
				byte(value_codec.EncodedValueBool),
				byte(common_codec.EncodedBoolTrue),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("Optional(nil)", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		value := cadence.NewOptional(nil)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedValueOptional),
				byte(common_codec.EncodedBoolTrue),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		innerType := cadence.NewBoolType()
		typ := cadence.NewOptionalType(innerType)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedTypeOptional),
				byte(common_codec.EncodedBoolFalse),
				byte(value_codec.EncodedTypeBool),
			},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestValueCodecString(t *testing.T) {
	t.Parallel()

	t.Run("len=0", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		s := ""
		value, _ := cadence.NewString(s)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedValueString),
				0, 0, 0, 0,
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("len>0", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		s := "wot\x00 now"
		value, _ := cadence.NewString(s)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueString)},
				[]byte{0, 0, 0, byte(len(s))},
				[]byte(s),
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewStringType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeString)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestValueCodecBytes(t *testing.T) {
	t.Parallel()

	t.Run("len=0", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		s := []byte("")
		value := cadence.NewBytes(s)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedValueBytes),
				0, 0, 0, 0,
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("len>0", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		s := []byte("wot\x00 now")
		value := cadence.NewBytes(s)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueBytes)},
				[]byte{0, 0, 0, byte(len(s))},
				s,
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewBytesType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeBytes)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestValueCodecCharacter(t *testing.T) {
	t.Parallel()

	t.Run("len=1", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		s := "W"
		value, err := cadence.NewCharacter(s)

		require.NoError(t, err)

		err = encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueCharacter)},
				[]byte{0, 0, 0, byte(len(s))},
				[]byte(s),
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("len>1", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		s := "ᄀᄀᄀ각ᆨᆨ"
		value, err := cadence.NewCharacter(s)

		require.NoError(t, err)

		err = encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueCharacter)},
				[]byte{0, 0, 0, byte(len(s))},
				[]byte(s),
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewCharacterType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeCharacter)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestValueCodecAddress(t *testing.T) {
	t.Parallel()

	t.Run("null address", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		value := cadence.NewAddress([8]byte{})

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueAddress)},
				value.Bytes(),
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("some address", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		value := cadence.NewAddress([8]byte{255, 127, 62, 28, 8, 4, 2, 1})

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueAddress)},
				value.Bytes(),
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewAddressType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeAddress)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestValueCodecInt(t *testing.T) {
	t.Parallel()

	t.Run("small positive", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := 255
		value := cadence.NewInt(i0)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueInt)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 1},
				[]byte{byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("large positive", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := 0
		i1 := 1
		value := cadence.NewInt(256)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueInt)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 2},
				[]byte{byte(i1), byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("negative", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := -4
		value := cadence.NewInt(i0)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueInt)},
				[]byte{byte(common_codec.EncodedBoolTrue)}, // negative
				[]byte{0, 0, 0, 1},
				[]byte{byte(-i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewIntType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeInt)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestValueCodecUInt(t *testing.T) {
	t.Parallel()

	t.Run("positive", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := uint(255)
		value := cadence.NewUInt(i0)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueUInt)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 1},
				[]byte{byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("large", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := 0
		i1 := 1
		value := cadence.NewUInt(256)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueUInt)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 2},
				[]byte{byte(i1), byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewUIntType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeUInt)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestValueCodecNumber(t *testing.T) {
	t.Parallel()

	t.Run("value int8", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i := int8(99)
		value := cadence.Int8(i)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueInt8)},
				[]byte{byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type int8", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.Int8Type{}

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeInt8)},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("value int16", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i := int16(99)
		value := cadence.Int16(i)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueInt16)},
				[]byte{0, byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type int16", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.Int16Type{}

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeInt16)},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("value int32", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i := int32(99)
		value := cadence.Int32(i)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueInt32)},
				[]byte{0, 0, 0, byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type int32", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.Int32Type{}

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeInt32)},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("value int64", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i := int64(99)
		value := cadence.Int64(i)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueInt64)},
				[]byte{0, 0, 0, 0, 0, 0, 0, byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type int64", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.Int64Type{}

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeInt64)},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("value uint8", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i := uint8(99)
		value := cadence.UInt8(i)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueUInt8)},
				[]byte{byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type uint8", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.UInt8Type{}

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeUInt8)},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("value uint16", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i := uint16(99)
		value := cadence.UInt16(i)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueUInt16)},
				[]byte{0, byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type uint16", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.UInt16Type{}

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeUInt16)},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("value uint32", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i := uint32(99)
		value := cadence.UInt32(i)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueUInt32)},
				[]byte{0, 0, 0, byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type uint32", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.UInt32Type{}

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeUInt32)},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("value uint64", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i := uint64(99)
		value := cadence.UInt64(i)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(value_codec.EncodedValueUInt64)},
				[]byte{0, 0, 0, 0, 0, 0, 0, byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type uint64", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.UInt64Type{}

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(value_codec.EncodedTypeUInt64)},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

}

func TestValueCodecArray(t *testing.T) {
	t.Parallel()

	t.Run("Variable Array, len=0", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		elements := make([]cadence.Value, 0)

		value := cadence.NewArray(elements).
			WithType(cadence.NewVariableSizedArrayType(cadence.NewAnyType()))

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedValueVariableArray),
				byte(value_codec.EncodedTypeAnyType),
				0, 0, 0, byte(len(elements)),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("Variable Array, len=2", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		elements := []cadence.Value{
			cadence.NewVoid(),
			cadence.NewBool(true),
		}

		value := cadence.NewArray(elements).
			WithType(cadence.NewVariableSizedArrayType(cadence.NewAnyType()))

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedValueVariableArray),
				byte(value_codec.EncodedTypeAnyType),
				0, 0, 0, byte(len(elements)),

				byte(value_codec.EncodedValueVoid),

				byte(value_codec.EncodedValueBool),
				byte(common_codec.EncodedBoolTrue),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("variable type", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		innerType := cadence.NewAnyType()
		typ := cadence.NewVariableSizedArrayType(innerType)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedTypeVariableSizedArray),
				byte(value_codec.EncodedTypeAnyType),
			},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("Constant Array, len=0", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		elements := make([]cadence.Value, 0)

		value := cadence.NewArray(elements).
			WithType(cadence.NewConstantSizedArrayType(uint(len(elements)), cadence.NewAnyStructType()))

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedValueConstantArray),
				byte(value_codec.EncodedTypeAnyStructType),
				0, 0, 0, byte(len(elements)),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("Constant Array, len=2", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		elements := []cadence.Value{
			cadence.NewVoid(),
			cadence.NewBool(true),
		}

		value := cadence.NewArray(elements).
			WithType(cadence.NewConstantSizedArrayType(uint(len(elements)), cadence.NewAnyStructType()))

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedValueConstantArray),
				byte(value_codec.EncodedTypeAnyStructType),
				0, 0, 0, byte(len(elements)),

				byte(value_codec.EncodedValueVoid),

				byte(value_codec.EncodedValueBool),
				byte(common_codec.EncodedBoolTrue),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.Decode()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("constant type", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		size := uint(12)
		innerType := cadence.NewAnyType()
		typ := cadence.NewConstantSizedArrayType(size, innerType)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(value_codec.EncodedTypeConstantSizedArray),
				byte(value_codec.EncodedTypeAnyType),
				0, 0, 0, byte(size),
			},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func NewTestCodec() (encoder *value_codec.Encoder, decoder *value_codec.Decoder, buffer *bytes.Buffer) {
	var w bytes.Buffer
	buffer = &w
	encoder = value_codec.NewEncoder(buffer)
	decoder = value_codec.NewDecoder(nil, buffer)
	return
}
