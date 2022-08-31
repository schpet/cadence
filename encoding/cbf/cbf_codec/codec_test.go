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

package cbf_codec_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence"
	"github.com/onflow/cadence/encoding/cbf/cbf_codec"
	"github.com/onflow/cadence/encoding/cbf/common_codec"
	"github.com/onflow/cadence/runtime/common"
)

func TestCadenceBinaryFormatCodecVoid(t *testing.T) {
	t.Parallel()

	t.Run("value", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		value := cadence.NewVoid()

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedValueVoid)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewVoidType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeVoid)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecBool(t *testing.T) {
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
				byte(cbf_codec.EncodedValueBool),
				byte(common_codec.EncodedBoolFalse),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
				byte(cbf_codec.EncodedValueBool),
				byte(common_codec.EncodedBoolTrue),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewBoolType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeBool)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecOptional(t *testing.T) {
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
				byte(cbf_codec.EncodedValueOptional),
				byte(common_codec.EncodedBoolFalse),
				byte(cbf_codec.EncodedValueVoid),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
				byte(cbf_codec.EncodedValueOptional),
				byte(common_codec.EncodedBoolFalse),
				byte(cbf_codec.EncodedValueBool),
				byte(common_codec.EncodedBoolTrue),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
				byte(cbf_codec.EncodedValueOptional),
				byte(common_codec.EncodedBoolTrue),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		innerType := cadence.NewBoolType()
		typ := cadence.NewOptionalType(innerType)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(cbf_codec.EncodedTypeOptional),
				byte(common_codec.EncodedBoolFalse),
				byte(cbf_codec.EncodedTypeBool),
			},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecString(t *testing.T) {
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
				byte(cbf_codec.EncodedValueString),
				0, 0, 0, 0,
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
				[]byte{byte(cbf_codec.EncodedValueString)},
				[]byte{0, 0, 0, byte(len(s))},
				[]byte(s),
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewStringType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeString)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecBytes(t *testing.T) {
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
				byte(cbf_codec.EncodedValueBytes),
				0, 0, 0, 0,
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
				[]byte{byte(cbf_codec.EncodedValueBytes)},
				[]byte{0, 0, 0, byte(len(s))},
				s,
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewBytesType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeBytes)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecCharacter(t *testing.T) {
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
				[]byte{byte(cbf_codec.EncodedValueCharacter)},
				[]byte{0, 0, 0, byte(len(s))},
				[]byte(s),
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
				[]byte{byte(cbf_codec.EncodedValueCharacter)},
				[]byte{0, 0, 0, byte(len(s))},
				[]byte(s),
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewCharacterType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeCharacter)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecAddress(t *testing.T) {
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
				[]byte{byte(cbf_codec.EncodedValueAddress)},
				value.Bytes(),
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
				[]byte{byte(cbf_codec.EncodedValueAddress)},
				value.Bytes(),
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewAddressType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeAddress)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecInt(t *testing.T) {
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
				[]byte{byte(cbf_codec.EncodedValueInt)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 1},
				[]byte{byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
				[]byte{byte(cbf_codec.EncodedValueInt)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 2},
				[]byte{byte(i1), byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
				[]byte{byte(cbf_codec.EncodedValueInt)},
				[]byte{byte(common_codec.EncodedBoolTrue)}, // negative
				[]byte{0, 0, 0, 1},
				[]byte{byte(-i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewIntType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeInt)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecInt128(t *testing.T) {
	t.Parallel()

	t.Run("small positive", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := 255
		value := cadence.NewInt128(i0)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueInt128)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 1},
				[]byte{byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("large positive", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := 0
		i1 := 1
		value := cadence.NewInt128(256)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueInt128)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 2},
				[]byte{byte(i1), byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("negative", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := -4
		value := cadence.NewInt128(i0)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueInt128)},
				[]byte{byte(common_codec.EncodedBoolTrue)}, // negative
				[]byte{0, 0, 0, 1},
				[]byte{byte(-i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewInt128Type()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeInt128)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecInt256(t *testing.T) {
	t.Parallel()

	t.Run("small positive", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := 255
		value := cadence.NewInt256(i0)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueInt256)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 1},
				[]byte{byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("large positive", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := 0
		i1 := 1
		value := cadence.NewInt256(256)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueInt256)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 2},
				[]byte{byte(i1), byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("negative", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := -4
		value := cadence.NewInt256(i0)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueInt256)},
				[]byte{byte(common_codec.EncodedBoolTrue)}, // negative
				[]byte{0, 0, 0, 1},
				[]byte{byte(-i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewInt256Type()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeInt256)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecUInt128(t *testing.T) {
	t.Parallel()

	t.Run("small positive", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := uint(255)
		value := cadence.NewUInt128(i0)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueUInt128)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 1},
				[]byte{byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("large positive", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := 0
		i1 := 1
		value := cadence.NewUInt128(256)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueUInt128)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 2},
				[]byte{byte(i1), byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewUInt128Type()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeUInt128)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecUInt256(t *testing.T) {
	t.Parallel()

	t.Run("small positive", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := uint(255)
		value := cadence.NewUInt256(i0)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueUInt256)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 1},
				[]byte{byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("large positive", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		i0 := 0
		i1 := 1
		value := cadence.NewUInt256(256)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueUInt256)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 2},
				[]byte{byte(i1), byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewUInt256Type()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeUInt256)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecUInt(t *testing.T) {
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
				[]byte{byte(cbf_codec.EncodedValueUInt)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 1},
				[]byte{byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
				[]byte{byte(cbf_codec.EncodedValueUInt)},
				[]byte{byte(common_codec.EncodedBoolFalse)}, // not negative
				[]byte{0, 0, 0, 2},
				[]byte{byte(i1), byte(i0)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewUIntType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeUInt)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecNumber(t *testing.T) {
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
				[]byte{byte(cbf_codec.EncodedValueInt8)},
				[]byte{byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
			[]byte{byte(cbf_codec.EncodedTypeInt8)},
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
				[]byte{byte(cbf_codec.EncodedValueInt16)},
				[]byte{0, byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
			[]byte{byte(cbf_codec.EncodedTypeInt16)},
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
				[]byte{byte(cbf_codec.EncodedValueInt32)},
				[]byte{0, 0, 0, byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
			[]byte{byte(cbf_codec.EncodedTypeInt32)},
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
				[]byte{byte(cbf_codec.EncodedValueInt64)},
				[]byte{0, 0, 0, 0, 0, 0, 0, byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
			[]byte{byte(cbf_codec.EncodedTypeInt64)},
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
				[]byte{byte(cbf_codec.EncodedValueUInt8)},
				[]byte{byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
			[]byte{byte(cbf_codec.EncodedTypeUInt8)},
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
				[]byte{byte(cbf_codec.EncodedValueUInt16)},
				[]byte{0, byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
			[]byte{byte(cbf_codec.EncodedTypeUInt16)},
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
				[]byte{byte(cbf_codec.EncodedValueUInt32)},
				[]byte{0, 0, 0, byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
			[]byte{byte(cbf_codec.EncodedTypeUInt32)},
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
				[]byte{byte(cbf_codec.EncodedValueUInt64)},
				[]byte{0, 0, 0, 0, 0, 0, 0, byte(i)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
			[]byte{byte(cbf_codec.EncodedTypeUInt64)},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

}

func TestCadenceBinaryFormatCodecArray(t *testing.T) {
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
				byte(cbf_codec.EncodedValueVariableArray),
				byte(cbf_codec.EncodedTypeAnyType),
				0, 0, 0, byte(len(elements)),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
				byte(cbf_codec.EncodedValueVariableArray),
				byte(cbf_codec.EncodedTypeAnyType),
				0, 0, 0, byte(len(elements)),

				byte(cbf_codec.EncodedValueVoid),

				byte(cbf_codec.EncodedValueBool),
				byte(common_codec.EncodedBoolTrue),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("variable type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		innerType := cadence.NewAnyType()
		typ := cadence.NewVariableSizedArrayType(innerType)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(cbf_codec.EncodedTypeVariableSizedArray),
				byte(cbf_codec.EncodedTypeAnyType),
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
				byte(cbf_codec.EncodedValueConstantArray),
				byte(cbf_codec.EncodedTypeAnyStructType),
				0, 0, 0, byte(len(elements)),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
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
				byte(cbf_codec.EncodedValueConstantArray),
				byte(cbf_codec.EncodedTypeAnyStructType),
				0, 0, 0, byte(len(elements)),

				byte(cbf_codec.EncodedValueVoid),

				byte(cbf_codec.EncodedValueBool),
				byte(common_codec.EncodedBoolTrue),
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("constant type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		size := uint(12)
		innerType := cadence.NewAnyType()
		typ := cadence.NewConstantSizedArrayType(size, innerType)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(cbf_codec.EncodedTypeConstantSizedArray),
				byte(cbf_codec.EncodedTypeAnyType),
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

func TestCadenceBinaryFormatCodecDictionary(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		dictionaryType := cadence.NewDictionaryType(cadence.Fix64Type{}, cadence.FixedPointType{})

		value := cadence.NewDictionary([]cadence.KeyValuePair{}).
			WithType(dictionaryType)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(cbf_codec.EncodedValueDictionary),
				byte(cbf_codec.EncodedTypeFix64),
				byte(cbf_codec.EncodedTypeFixedPoint),
				0, 0, 0, 0,
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("two elements", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		dictionaryType := cadence.NewDictionaryType(cadence.Fix64Type{}, cadence.FixedPointType{})

		pairs := []cadence.KeyValuePair{
			{
				Key:   cadence.Fix64(8),
				Value: cadence.UFix64(3),
			},
			{
				Key:   cadence.Fix64(7),
				Value: cadence.Fix64(18),
			},
		}
		value := cadence.NewDictionary(pairs).
			WithType(dictionaryType)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(cbf_codec.EncodedValueDictionary),
				byte(cbf_codec.EncodedTypeFix64),
				byte(cbf_codec.EncodedTypeFixedPoint),
				0, 0, 0, byte(len(pairs)),
				byte(cbf_codec.EncodedValueFix64),
				0, 0, 0, 0, 0, 0, 0, 8,
				byte(cbf_codec.EncodedValueUFix64),
				0, 0, 0, 0, 0, 0, 0, 3,
				byte(cbf_codec.EncodedValueFix64),
				0, 0, 0, 0, 0, 0, 0, 7,
				byte(cbf_codec.EncodedValueFix64),
				0, 0, 0, 0, 0, 0, 0, 18,
			},
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewDictionaryType(cadence.AnyResourceType{}, cadence.SignedNumberType{})

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{
				byte(cbf_codec.EncodedTypeDictionary),
				byte(cbf_codec.EncodedTypeAnyResourceType),
				byte(cbf_codec.EncodedTypeSignedNumber),
			},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecStruct(t *testing.T) {
	t.Parallel()

	t.Run("value", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		location := common.REPLLocation{}
		qualifiedIdentifier := "neon"
		fieldsTypes := []cadence.Field{
			{
				Identifier: "argon",
				Type:       cadence.UInt16Type{},
			},
		}
		initializers := [][]cadence.Parameter{
			{
				{
					Label:      "lebal",
					Identifier: "home",
					Type:       cadence.Word8Type{},
				},
			},
		}
		structType := cadence.NewStructType(
			location,
			qualifiedIdentifier,
			fieldsTypes,
			initializers,
		)

		fieldValue := uint16(12)
		fields := []cadence.Value{
			cadence.NewUInt16(fieldValue),
		}
		value := cadence.NewStruct(fields).
			WithType(structType)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueStruct)},
				[]byte{common.REPLLocationPrefix[0]},
				[]byte{0, 0, 0, byte(len(qualifiedIdentifier))},
				[]byte(qualifiedIdentifier),
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(fieldsTypes[0].Identifier))},
				[]byte(fieldsTypes[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeUInt16)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(initializers[0][0].Label))},
				[]byte(initializers[0][0].Label),
				[]byte{0, 0, 0, byte(len(initializers[0][0].Identifier))},
				[]byte(initializers[0][0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeWord8)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{byte(cbf_codec.EncodedValueUInt16)},
				[]byte{0, byte(fieldValue)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		location := common.REPLLocation{}
		qualifiedIdentifier := "neon"
		fieldsTypes := []cadence.Field{
			{
				Identifier: "argon",
				Type:       cadence.UInt16Type{},
			},
		}
		initializers := [][]cadence.Parameter{
			{
				{
					Label:      "lebal",
					Identifier: "home",
					Type:       cadence.Word8Type{},
				},
			},
		}
		typ := cadence.NewStructType(
			location,
			qualifiedIdentifier,
			fieldsTypes,
			initializers,
		)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypeStruct)},
				[]byte{common.REPLLocationPrefix[0]},
				[]byte{0, 0, 0, byte(len(qualifiedIdentifier))},
				[]byte(qualifiedIdentifier),
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(fieldsTypes[0].Identifier))},
				[]byte(fieldsTypes[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeUInt16)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(initializers[0][0].Label))},
				[]byte(initializers[0][0].Label),
				[]byte{0, 0, 0, byte(len(initializers[0][0].Identifier))},
				[]byte(initializers[0][0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeWord8)},
			),
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecResource(t *testing.T) {
	t.Parallel()

	t.Run("value", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		location := common.REPLLocation{}
		qualifiedIdentifier := "neon"
		fieldsTypes := []cadence.Field{
			{
				Identifier: "argon",
				Type:       cadence.UInt16Type{},
			},
		}
		initializers := [][]cadence.Parameter{
			{
				{
					Label:      "lebal",
					Identifier: "home",
					Type:       cadence.Word8Type{},
				},
			},
		}
		resourceType := cadence.NewResourceType(
			location,
			qualifiedIdentifier,
			fieldsTypes,
			initializers,
		)

		fieldValue := uint16(12)
		fields := []cadence.Value{
			cadence.NewUInt16(fieldValue),
		}
		value := cadence.NewResource(fields).
			WithType(resourceType)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueResource)},
				[]byte{common.REPLLocationPrefix[0]},
				[]byte{0, 0, 0, byte(len(qualifiedIdentifier))},
				[]byte(qualifiedIdentifier),
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(fieldsTypes[0].Identifier))},
				[]byte(fieldsTypes[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeUInt16)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(initializers[0][0].Label))},
				[]byte(initializers[0][0].Label),
				[]byte{0, 0, 0, byte(len(initializers[0][0].Identifier))},
				[]byte(initializers[0][0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeWord8)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{byte(cbf_codec.EncodedValueUInt16)},
				[]byte{0, byte(fieldValue)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		location := common.REPLLocation{}
		qualifiedIdentifier := "neon"
		fieldsTypes := []cadence.Field{
			{
				Identifier: "argon",
				Type:       cadence.UInt16Type{},
			},
		}
		initializers := [][]cadence.Parameter{
			{
				{
					Label:      "lebal",
					Identifier: "home",
					Type:       cadence.Word8Type{},
				},
			},
		}
		typ := cadence.NewResourceType(
			location,
			qualifiedIdentifier,
			fieldsTypes,
			initializers,
		)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypeResource)},
				[]byte{common.REPLLocationPrefix[0]},
				[]byte{0, 0, 0, byte(len(qualifiedIdentifier))},
				[]byte(qualifiedIdentifier),
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(fieldsTypes[0].Identifier))},
				[]byte(fieldsTypes[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeUInt16)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(initializers[0][0].Label))},
				[]byte(initializers[0][0].Label),
				[]byte{0, 0, 0, byte(len(initializers[0][0].Identifier))},
				[]byte(initializers[0][0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeWord8)},
			),
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecEvent(t *testing.T) {
	t.Parallel()

	t.Run("value", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		location := common.REPLLocation{}
		qualifiedIdentifier := "neon"
		fieldsTypes := []cadence.Field{
			{
				Identifier: "argon",
				Type:       cadence.UInt16Type{},
			},
		}
		initializer := []cadence.Parameter{
			{
				Label:      "lebal",
				Identifier: "home",
				Type:       cadence.Word8Type{},
			},
		}
		eventType := cadence.NewEventType(
			location,
			qualifiedIdentifier,
			fieldsTypes,
			initializer,
		)

		fieldValue := uint16(12)
		fields := []cadence.Value{
			cadence.NewUInt16(fieldValue),
		}
		value := cadence.NewEvent(fields).
			WithType(eventType)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueEvent)},
				[]byte{common.REPLLocationPrefix[0]},
				[]byte{0, 0, 0, byte(len(qualifiedIdentifier))},
				[]byte(qualifiedIdentifier),
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(fieldsTypes[0].Identifier))},
				[]byte(fieldsTypes[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeUInt16)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(initializer[0].Label))},
				[]byte(initializer[0].Label),
				[]byte{0, 0, 0, byte(len(initializer[0].Identifier))},
				[]byte(initializer[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeWord8)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{byte(cbf_codec.EncodedValueUInt16)},
				[]byte{0, byte(fieldValue)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		location := common.REPLLocation{}
		qualifiedIdentifier := "neon"
		fieldsTypes := []cadence.Field{
			{
				Identifier: "argon",
				Type:       cadence.UInt16Type{},
			},
		}
		initializer := []cadence.Parameter{
			{
				Label:      "lebal",
				Identifier: "home",
				Type:       cadence.Word8Type{},
			},
		}
		typ := cadence.NewEventType(
			location,
			qualifiedIdentifier,
			fieldsTypes,
			initializer,
		)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypeEvent)},
				[]byte{common.REPLLocationPrefix[0]},
				[]byte{0, 0, 0, byte(len(qualifiedIdentifier))},
				[]byte(qualifiedIdentifier),
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(fieldsTypes[0].Identifier))},
				[]byte(fieldsTypes[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeUInt16)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(initializer[0].Label))},
				[]byte(initializer[0].Label),
				[]byte{0, 0, 0, byte(len(initializer[0].Identifier))},
				[]byte(initializer[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeWord8)},
			),
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecContract(t *testing.T) {
	t.Parallel()

	t.Run("value", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		location := common.REPLLocation{}
		qualifiedIdentifier := "neon"
		fieldsTypes := []cadence.Field{
			{
				Identifier: "argon",
				Type:       cadence.UInt16Type{},
			},
		}
		initializers := [][]cadence.Parameter{
			{
				{
					Label:      "lebal",
					Identifier: "home",
					Type:       cadence.Word8Type{},
				},
			},
		}
		contractType := cadence.NewContractType(
			location,
			qualifiedIdentifier,
			fieldsTypes,
			initializers,
		)

		fieldValue := uint16(12)
		fields := []cadence.Value{
			cadence.NewUInt16(fieldValue),
		}
		value := cadence.NewContract(fields).
			WithType(contractType)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueContract)},
				[]byte{common.REPLLocationPrefix[0]},
				[]byte{0, 0, 0, byte(len(qualifiedIdentifier))},
				[]byte(qualifiedIdentifier),
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(fieldsTypes[0].Identifier))},
				[]byte(fieldsTypes[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeUInt16)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(initializers[0][0].Label))},
				[]byte(initializers[0][0].Label),
				[]byte{0, 0, 0, byte(len(initializers[0][0].Identifier))},
				[]byte(initializers[0][0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeWord8)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{byte(cbf_codec.EncodedValueUInt16)},
				[]byte{0, byte(fieldValue)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		location := common.REPLLocation{}
		qualifiedIdentifier := "neon"
		fieldsTypes := []cadence.Field{
			{
				Identifier: "argon",
				Type:       cadence.UInt16Type{},
			},
		}
		initializers := [][]cadence.Parameter{
			{
				{
					Label:      "lebal",
					Identifier: "home",
					Type:       cadence.Word8Type{},
				},
			},
		}
		typ := cadence.NewContractType(
			location,
			qualifiedIdentifier,
			fieldsTypes,
			initializers,
		)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypeContract)},
				[]byte{common.REPLLocationPrefix[0]},
				[]byte{0, 0, 0, byte(len(qualifiedIdentifier))},
				[]byte(qualifiedIdentifier),
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(fieldsTypes[0].Identifier))},
				[]byte(fieldsTypes[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeUInt16)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(initializers[0][0].Label))},
				[]byte(initializers[0][0].Label),
				[]byte{0, 0, 0, byte(len(initializers[0][0].Identifier))},
				[]byte(initializers[0][0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeWord8)},
			),
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecLink(t *testing.T) {
	t.Parallel()

	t.Run("value", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		targetPath := cadence.NewPath("domi", "le nom")
		borrowType := "borrow'd"
		value := cadence.NewLink(targetPath, borrowType)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueLink)},
				[]byte{0, 0, 0, byte(len(targetPath.Domain))},
				[]byte(targetPath.Domain),
				[]byte{0, 0, 0, byte(len(targetPath.Identifier))},
				[]byte(targetPath.Identifier),
				[]byte{0, 0, 0, byte(len(borrowType))},
				[]byte(borrowType),
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})
}

func TestCadenceBinaryFormatCodecPath(t *testing.T) {
	t.Parallel()

	t.Run("value", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		value := cadence.NewPath("domi", "le nom")

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValuePath)},
				[]byte{0, 0, 0, byte(len(value.Domain))},
				[]byte(value.Domain),
				[]byte{0, 0, 0, byte(len(value.Identifier))},
				[]byte(value.Identifier),
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type(capability)", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewCapabilityPathType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypeCapabilityPath)},
			),
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("type(storage)", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewStoragePathType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypeStoragePath)},
			),
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("type(public)", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewPublicPathType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypePublicPath)},
			),
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("type(private)", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewPrivatePathType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypePrivatePath)},
			),
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecCapability(t *testing.T) {
	t.Parallel()

	t.Run("value", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		path := cadence.NewPath("demesne", "pointer")
		address := cadence.NewAddress([8]byte{1, 2, 3, 4, 5, 6, 7, 8})
		borrowType := cadence.NewIntType()

		value := cadence.NewCapability(path, address, borrowType)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueCapability)},
				[]byte{0, 0, 0, byte(len(path.Domain))},
				[]byte(path.Domain),
				[]byte{0, 0, 0, byte(len(path.Identifier))},
				[]byte(path.Identifier),
				address.Bytes(),
				[]byte{byte(cbf_codec.EncodedTypeInt)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewCapabilityType(cadence.NewAddressType())

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypeCapability)},
				[]byte{byte(cbf_codec.EncodedTypeAddress)},
			),
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecEnum(t *testing.T) {
	t.Parallel()

	t.Run("value", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		location := common.REPLLocation{}
		qualifiedIdentifier := "neon"
		rawType := cadence.NewNeverType()
		fieldsTypes := []cadence.Field{
			{
				Identifier: "argon",
				Type:       cadence.UInt16Type{},
			},
		}
		initializers := [][]cadence.Parameter{
			{
				{
					Label:      "lebal",
					Identifier: "home",
					Type:       cadence.Word8Type{},
				},
			},
		}
		enumType := cadence.NewEnumType(
			location,
			qualifiedIdentifier,
			rawType,
			fieldsTypes,
			initializers,
		)

		fieldValue := uint16(12)
		fields := []cadence.Value{
			cadence.NewUInt16(fieldValue),
		}
		value := cadence.NewEnum(fields).
			WithType(enumType)

		err := encoder.Encode(value)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedValueEnum)},
				[]byte{common.REPLLocationPrefix[0]},
				[]byte{0, 0, 0, byte(len(qualifiedIdentifier))},
				[]byte(qualifiedIdentifier),
				[]byte{byte(cbf_codec.EncodedTypeNever)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(fieldsTypes[0].Identifier))},
				[]byte(fieldsTypes[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeUInt16)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(initializers[0][0].Label))},
				[]byte(initializers[0][0].Label),
				[]byte{0, 0, 0, byte(len(initializers[0][0].Identifier))},
				[]byte(initializers[0][0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeWord8)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{byte(cbf_codec.EncodedValueUInt16)},
				[]byte{0, byte(fieldValue)},
			),
			buffer.Bytes(), "encoded bytes differ")

		output, err := decoder.DecodeValue()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, value, output, "decoded value differs")
	})

	t.Run("type", func(t *testing.T) {
		t.Parallel()

		encoder, decoder, buffer := NewTestCodec()

		location := common.REPLLocation{}
		qualifiedIdentifier := "neon"
		rawType := cadence.NewNeverType()
		fieldsTypes := []cadence.Field{
			{
				Identifier: "argon",
				Type:       cadence.UInt16Type{},
			},
		}
		initializers := [][]cadence.Parameter{
			{
				{
					Label:      "lebal",
					Identifier: "home",
					Type:       cadence.Word8Type{},
				},
			},
		}
		typ := cadence.NewEnumType(
			location,
			qualifiedIdentifier,
			rawType,
			fieldsTypes,
			initializers,
		)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypeEnum)},
				[]byte{common.REPLLocationPrefix[0]},
				[]byte{0, 0, 0, byte(len(qualifiedIdentifier))},
				[]byte(qualifiedIdentifier),
				[]byte{byte(cbf_codec.EncodedTypeNever)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(fieldsTypes[0].Identifier))},
				[]byte(fieldsTypes[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeUInt16)},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{byte(common_codec.EncodedBoolFalse), 0, 0, 0, 1},
				[]byte{0, 0, 0, byte(len(initializers[0][0].Label))},
				[]byte(initializers[0][0].Label),
				[]byte{0, 0, 0, byte(len(initializers[0][0].Identifier))},
				[]byte(initializers[0][0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeWord8)},
			),
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecInterfaceType(t *testing.T) {
	t.Parallel()

	t.Run("StructInterfaceType", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewMeteredStructInterfaceType(
			nil,
			common.REPLLocation{},
			"qid",
			[]cadence.Field{
				{
					Identifier: "i",
					Type:       cadence.StringType{},
				},
			},
			[][]cadence.Parameter{
				{{
					Label:      "l",
					Identifier: "i2",
					Type:       cadence.BoolType{},
				}},
			},
		)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypeStructInterface)},
				[]byte{common.REPLLocationPrefix[0]},
				[]byte{0, 0, 0, byte(len(typ.QualifiedIdentifier))},
				[]byte(typ.QualifiedIdentifier),

				[]byte{byte(common_codec.EncodedBoolFalse)},
				[]byte{0, 0, 0, byte(len(typ.Fields))},
				[]byte{0, 0, 0, byte(len(typ.Fields[0].Identifier))},
				[]byte(typ.Fields[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeString)},

				[]byte{byte(common_codec.EncodedBoolFalse)},
				[]byte{0, 0, 0, byte(len(typ.Initializers))},
				[]byte{byte(common_codec.EncodedBoolFalse)},
				[]byte{0, 0, 0, byte(len(typ.Initializers[0]))},
				[]byte{0, 0, 0, byte(len(typ.Initializers[0][0].Label))},
				[]byte(typ.Initializers[0][0].Label),
				[]byte{0, 0, 0, byte(len(typ.Initializers[0][0].Identifier))},
				[]byte(typ.Initializers[0][0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeBool)},
			),
			buffer.Bytes(),
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("ResourceInterfaceType", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewMeteredResourceInterfaceType(
			nil,
			common.REPLLocation{},
			"qid",
			[]cadence.Field{
				{
					Identifier: "i",
					Type:       cadence.StringType{},
				},
			},
			[][]cadence.Parameter{
				{{
					Label:      "l",
					Identifier: "i2",
					Type:       cadence.BoolType{},
				}},
			},
		)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypeResourceInterface)},
				[]byte{common.REPLLocationPrefix[0]},
				[]byte{0, 0, 0, byte(len(typ.QualifiedIdentifier))},
				[]byte(typ.QualifiedIdentifier),

				[]byte{byte(common_codec.EncodedBoolFalse)},
				[]byte{0, 0, 0, byte(len(typ.Fields))},
				[]byte{0, 0, 0, byte(len(typ.Fields[0].Identifier))},
				[]byte(typ.Fields[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeString)},

				[]byte{byte(common_codec.EncodedBoolFalse)},
				[]byte{0, 0, 0, byte(len(typ.Initializers))},
				[]byte{byte(common_codec.EncodedBoolFalse)},
				[]byte{0, 0, 0, byte(len(typ.Initializers[0]))},
				[]byte{0, 0, 0, byte(len(typ.Initializers[0][0].Label))},
				[]byte(typ.Initializers[0][0].Label),
				[]byte{0, 0, 0, byte(len(typ.Initializers[0][0].Identifier))},
				[]byte(typ.Initializers[0][0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeBool)},
			),
			buffer.Bytes(),
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("ContractInterfaceType", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewMeteredContractInterfaceType(
			nil,
			common.REPLLocation{},
			"qid",
			[]cadence.Field{
				{
					Identifier: "i",
					Type:       cadence.StringType{},
				},
			},
			[][]cadence.Parameter{
				{{
					Label:      "l",
					Identifier: "i2",
					Type:       cadence.BoolType{},
				}},
			},
		)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypeContractInterface)},
				[]byte{common.REPLLocationPrefix[0]},
				[]byte{0, 0, 0, byte(len(typ.QualifiedIdentifier))},
				[]byte(typ.QualifiedIdentifier),

				[]byte{byte(common_codec.EncodedBoolFalse)},
				[]byte{0, 0, 0, byte(len(typ.Fields))},
				[]byte{0, 0, 0, byte(len(typ.Fields[0].Identifier))},
				[]byte(typ.Fields[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeString)},

				[]byte{byte(common_codec.EncodedBoolFalse)},
				[]byte{0, 0, 0, byte(len(typ.Initializers))},
				[]byte{byte(common_codec.EncodedBoolFalse)},
				[]byte{0, 0, 0, byte(len(typ.Initializers[0]))},
				[]byte{0, 0, 0, byte(len(typ.Initializers[0][0].Label))},
				[]byte(typ.Initializers[0][0].Label),
				[]byte{0, 0, 0, byte(len(typ.Initializers[0][0].Identifier))},
				[]byte(typ.Initializers[0][0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeBool)},
			),
			buffer.Bytes(),
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecFunctionType(t *testing.T) {
	t.Parallel()

	t.Run("full", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewFunctionType(
			"tid",
			[]cadence.Parameter{
				{
					Label:      "el",
					Identifier: "ihden",
					Type:       cadence.UIntType{},
				},
			},
			cadence.IntType{},
		)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypeFunction)},

				[]byte{0, 0, 0, byte(len(typ.ID()))},
				[]byte(typ.ID()),

				[]byte{byte(common_codec.EncodedBoolFalse)},
				[]byte{0, 0, 0, byte(len(typ.Parameters))},
				[]byte{0, 0, 0, byte(len(typ.Parameters[0].Label))},
				[]byte(typ.Parameters[0].Label),
				[]byte{0, 0, 0, byte(len(typ.Parameters[0].Identifier))},
				[]byte(typ.Parameters[0].Identifier),
				[]byte{byte(cbf_codec.EncodedTypeUInt)},
				[]byte{byte(cbf_codec.EncodedTypeInt)},
			),
			buffer.Bytes(),
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecReferenceType(t *testing.T) {
	t.Parallel()

	t.Run("authorized", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewReferenceType(
			true,
			cadence.Int64Type{},
		)

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			common_codec.Concat(
				[]byte{byte(cbf_codec.EncodedTypeReference)},
				[]byte{byte(common_codec.EncodedBoolTrue)},
				[]byte{byte(cbf_codec.EncodedTypeInt64)},
			),
			buffer.Bytes(),
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func TestCadenceBinaryFormatCodecAbstractTypes(t *testing.T) {
	t.Parallel()

	t.Run("Never", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewNeverType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeNever)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("Number", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewNumberType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeNumber)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("SignedNumber", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewSignedNumberType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeSignedNumber)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("Integer", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewIntegerType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeInteger)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("SignedInteger", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewSignedIntegerType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeSignedInteger)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("FixedPoint", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewFixedPointType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeFixedPoint)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("SignedFixedPoint", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewSignedFixedPointType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypeSignedFixedPoint)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})

	t.Run("Path", func(t *testing.T) {
		encoder, decoder, buffer := NewTestCodec()

		typ := cadence.NewPathType()

		err := encoder.EncodeType(typ)
		require.NoError(t, err, "encoding error")

		assert.Equal(
			t,
			[]byte{byte(cbf_codec.EncodedTypePath)},
			buffer.Bytes(),
			"encoded bytes differ",
		)

		output, err := decoder.DecodeType()
		require.NoError(t, err, "decoding error")

		assert.Equal(t, typ, output, "decoded type differs")
	})
}

func NewTestCodec() (encoder *cbf_codec.Encoder, decoder *cbf_codec.Decoder, buffer *bytes.Buffer) {
	var w bytes.Buffer
	buffer = &w
	encoder = cbf_codec.NewEncoder(buffer)
	decoder = cbf_codec.NewDecoder(nil, buffer)
	return
}
