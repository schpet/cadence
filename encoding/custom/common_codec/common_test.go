package common_codec_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence/encoding/custom/common_codec"
	"github.com/onflow/cadence/runtime/common"
)

func TestSemaCodecMiscValues(t *testing.T) {
	t.Parallel()

	t.Run("length (1 byte)", func(t *testing.T) {
		t.Parallel()

		length := 10

		var w bytes.Buffer
		err := common_codec.EncodeLength(&w, 10)
		require.NoError(t, err, "encoding error")

		assert.Equal(t, w.Bytes(), []byte{0, 0, 0, byte(length)}, "encoded bytes differ")

		output, err := common_codec.DecodeLength(&w)
		require.NoError(t, err, "decoding error")

		assert.Equal(t, length, output)
	})

	t.Run("length (2 bytes)", func(t *testing.T) {
		t.Parallel()

		length0 := 5
		length1 := 10
		length := length0 + (length1 << 8)

		var w bytes.Buffer

		err := common_codec.EncodeLength(&w, length)
		require.NoError(t, err, "encoding error")

		assert.Equal(t, w.Bytes(), []byte{0, 0, byte(length1), byte(length0)}, "encoded bytes differ")

		output, err := common_codec.DecodeLength(&w)
		require.NoError(t, err, "decoding error")

		assert.Equal(t, length, output)
	})

	t.Run("length error: negative", func(t *testing.T) {
		t.Parallel()

		var w bytes.Buffer

		err := common_codec.EncodeLength(&w, -1)
		assert.ErrorContains(t, err, "cannot encode length below zero: -1")
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()

		var w bytes.Buffer

		s := "some string \x00 foo \t \n\r\n $ 5"

		err := common_codec.EncodeString(&w, s)
		require.NoError(t, err, "encoding error")

		assert.Equal(t, w.Bytes(), common_codec.Concat(
			[]byte{0, 0, 0, byte(len(s))},
			[]byte(s),
		), "encoded bytes differ")

		output, err := common_codec.DecodeString(&w)
		require.NoError(t, err, "decoding error")

		assert.Equal(t, s, output)
	})

	t.Run("string len=0", func(t *testing.T) {
		t.Parallel()

		var w bytes.Buffer

		s := ""

		err := common_codec.EncodeString(&w, s)
		require.NoError(t, err, "encoding error")

		assert.Equal(t, w.Bytes(), common_codec.Concat(
			[]byte{0, 0, 0, byte(len(s))},
			[]byte(s),
		), "encoded bytes differ")

		output, err := common_codec.DecodeString(&w)
		require.NoError(t, err, "decoding error")

		assert.Equal(t, s, output)
	})

	t.Run("bytes", func(t *testing.T) {
		t.Parallel()

		var w bytes.Buffer

		s := []byte("some string \x00 foo \t \n\r\n $ 5")

		err := common_codec.EncodeBytes(&w, s)
		require.NoError(t, err, "encoding error")

		assert.Equal(t, w.Bytes(), common_codec.Concat(
			[]byte{0, 0, 0, byte(len(s))},
			s,
		), "encoded bytes differ")

		output, err := common_codec.DecodeBytes(&w)
		require.NoError(t, err, "decoding error")

		assert.Equal(t, s, output)
	})

	t.Run("bool true", func(t *testing.T) {
		t.Parallel()

		var w bytes.Buffer

		var b = true

		err := common_codec.EncodeBool(&w, b)
		require.NoError(t, err, "encoding error")

		assert.Equal(t, w.Bytes(), []byte{byte(common_codec.EncodedBoolTrue)}, "encoded bytes differ")

		output, err := common_codec.DecodeBool(&w)
		require.NoError(t, err, "decoding error")

		assert.Equal(t, b, output)
	})

	t.Run("bool false", func(t *testing.T) {
		t.Parallel()

		var w bytes.Buffer

		var b = false

		err := common_codec.EncodeBool(&w, b)
		require.NoError(t, err, "encoding error")

		assert.Equal(t, w.Bytes(), []byte{byte(common_codec.EncodedBoolFalse)}, "encoded bytes differ")

		output, err := common_codec.DecodeBool(&w)
		require.NoError(t, err, "decoding error")

		assert.Equal(t, b, output)
	})

	t.Run("address", func(t *testing.T) {
		t.Parallel()

		var w bytes.Buffer

		addr := common.MustBytesToAddress([]byte{0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00})

		err := common_codec.EncodeAddress(&w, addr)
		require.NoError(t, err, "encoding error")

		assert.Equal(t, w.Bytes(), addr.Bytes(), "encoded bytes differ")

		output, err := common_codec.DecodeAddress(&w)
		require.NoError(t, err, "decoding error")

		assert.Equal(t, addr, output)
	})

	t.Run("uint64", func(t *testing.T) {
		t.Parallel()

		var w bytes.Buffer

		i := uint64(1<<63) + 17

		err := common_codec.EncodeUInt64(&w, i)
		require.NoError(t, err, "encoding error")

		assert.Equal(t, w.Bytes(), []byte{128, 0, 0, 0, 0, 0, 0, 17}, "encoded bytes differ")

		output, err := common_codec.DecodeUInt64(&w)
		require.NoError(t, err, "decoding error")

		assert.Equal(t, i, output)
	})

	t.Run("int64 positive", func(t *testing.T) {
		t.Parallel()

		var w bytes.Buffer

		i := int64(1<<62) + 17

		err := common_codec.EncodeInt64(&w, i)
		require.NoError(t, err, "encoding error")

		assert.Equal(t, w.Bytes(), []byte{64, 0, 0, 0, 0, 0, 0, 17}, "encoded bytes differ")

		output, err := common_codec.DecodeInt64(&w)
		require.NoError(t, err, "decoding error")

		assert.Equal(t, i, output)
	})

	t.Run("int64 negative", func(t *testing.T) {
		t.Parallel()

		var w bytes.Buffer

		i := -(int64(1<<62) + 17)

		err := common_codec.EncodeInt64(&w, i)
		require.NoError(t, err, "encoding error")

		assert.Equal(t, w.Bytes(), []byte{0xff - 64, 0xff - 0, 0xff - 0, 0xff - 0, 0xff - 0, 0xff - 0, 0xff - 0, 0xff - 17 + 1}, "encoded bytes differ")

		output, err := common_codec.DecodeInt64(&w)
		require.NoError(t, err, "decoding error")

		assert.Equal(t, i, output)
	})
}