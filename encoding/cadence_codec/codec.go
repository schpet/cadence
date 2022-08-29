package cadence_codec

import (
	"github.com/onflow/flow-go/fvm/errors"

	"github.com/onflow/cadence"
	"github.com/onflow/cadence/encoding"
	customCodec "github.com/onflow/cadence/encoding/cbf/cbf_codec"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/cadence/runtime/common"
)

type CadenceCodec struct {
	Encoder encoding.Codec
}

func (c CadenceCodec) Encode(value cadence.Value) ([]byte, error) {
	return c.Encoder.Encode(value)
}

func (c CadenceCodec) MustEncode(value cadence.Value) []byte {
	return c.Encoder.MustEncode(value)
}

func (c CadenceCodec) Decode(gauge common.MemoryGauge, bytes []byte) (cadence.Value, error) {
	codec, err := c.chooseCodec(bytes)
	if err != nil {
		return nil, err
	}
	return codec.Decode(gauge, bytes)
}

func (c CadenceCodec) MustDecode(gauge common.MemoryGauge, bytes []byte) cadence.Value {
	codec, err := c.chooseCodec(bytes)
	if err != nil {
		panic(err)
	}
	return codec.MustDecode(gauge, bytes)
}

func (c CadenceCodec) chooseCodec(bytes []byte) (codec encoding.Codec, err error) {
	if len(bytes) == 0 {
		err = errors.NewInvalidArgumentErrorf("cannot decode empty argument")
		return
	}

	if bytes[0] == '{' {
		codec = jsoncdc.JsonCodec{}
	} else {
		codec = customCodec.CadenceBinaryFormatCodec{}
	}
	return
}

var _ encoding.Codec = CadenceCodec{}
