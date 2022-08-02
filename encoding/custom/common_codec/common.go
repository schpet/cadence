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

package common_codec

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/onflow/cadence"
	"github.com/onflow/cadence/runtime/common"
)

//
// LengthyWriter
//

type LengthyWriter struct {
	w      io.Writer
	length int
}

func NewLengthyWriter(w io.Writer) LengthyWriter {
	return LengthyWriter{w: w}
}

func (l *LengthyWriter) Write(p []byte) (n int, err error) {
	n, err = l.w.Write(p)
	l.length += n
	return
}

func (l *LengthyWriter) Len() int {
	return l.length
}

//
// LocatedReader
//

type LocatedReader struct {
	r        io.Reader
	location int
}

func NewLocatedReader(r io.Reader) LocatedReader {
	return LocatedReader{r: r}
}

func (l *LocatedReader) Read(p []byte) (n int, err error) {
	n, err = l.r.Read(p)
	l.location += n
	return
}

func (l *LocatedReader) Location() int {
	return l.location
}

//
// Bool
//

type EncodedBool byte

const (
	EncodedBoolUnknown EncodedBool = iota
	EncodedBoolFalse
	EncodedBoolTrue
)

func EncodeBool(w io.Writer, boolean bool) (err error) {
	b := EncodedBoolFalse
	if boolean {
		b = EncodedBoolTrue
	}

	_, err = w.Write([]byte{byte(b)})
	return
}

func DecodeBool(r io.Reader) (boolean bool, err error) {
	b := make([]byte, 1)
	_, err = r.Read(b)
	if err != nil {
		return
	}

	switch EncodedBool(b[0]) {
	case EncodedBoolFalse:
		boolean = false
	case EncodedBoolTrue:
		boolean = true
	default:
		err = fmt.Errorf("invalid boolean value: %d", b[0])
	}

	return
}

//
// Length
//

// TODO encode length with variable-sized encoding?
//      e.g. first byte starting with `0` is the last byte in the length
//      will usually save 3 bytes. the question is if it saves or costs encode and/or decode time

// EncodeLength encodes a non-negative length as a uint32.
// It uses 4 bytes.
func EncodeLength(w io.Writer, length int) (err error) {
	if length < 0 { // TODO is this safety check useful?
		return fmt.Errorf("cannot encode length below zero: %d", length)
	}

	l := uint32(length)

	return binary.Write(w, binary.BigEndian, l)
}

func DecodeLength(r io.Reader) (length int, err error) {
	b := make([]byte, 4)
	_, err = r.Read(b)
	if err != nil {
		return
	}

	asUint32 := binary.BigEndian.Uint32(b)
	length = int(asUint32)
	return
}

//
// Bytes
//

func EncodeBytes(w io.Writer, bytes []byte) (err error) {
	err = EncodeLength(w, len(bytes))
	if err != nil {
		return
	}
	_, err = w.Write(bytes)
	return
}

func DecodeBytes(r io.Reader) (bytes []byte, err error) {
	length, err := DecodeLength(r)
	if err != nil {
		return
	}

	bytes = make([]byte, length)

	_, err = r.Read(bytes)
	return
}

//
// String
//

func EncodeString(w io.Writer, s string) (err error) {
	return EncodeBytes(w, []byte(s))
}

func DecodeString(r io.Reader) (s string, err error) {
	b, err := DecodeBytes(r)
	if err != nil {
		return
	}
	s = string(b)
	return
}

//
// Address
//

func EncodeAddress[Address common.Address | cadence.Address](w io.Writer, a Address) (err error) {
	_, err = w.Write(a[:])
	return
}

func DecodeAddress(r io.Reader) (a common.Address, err error) {
	bytes := make([]byte, common.AddressLength)

	_, err = r.Read(bytes)
	if err != nil {
		return
	}

	return common.BytesToAddress(bytes)
}

//
// Int64
//

// TODO use a more efficient encoder than `binary` (they say to in their top source comment)

func EncodeInt64(w io.Writer, i int64) (err error) {
	return binary.Write(w, binary.BigEndian, i)
}

func DecodeInt64(r io.Reader) (i int64, err error) {
	err = binary.Read(r, binary.BigEndian, &i)
	return
}

//
// UInt64
//

func EncodeUInt64(w io.Writer, i uint64) (err error) {
	return binary.Write(w, binary.BigEndian, i)
}

func DecodeUInt64(r io.Reader) (i uint64, err error) {
	err = binary.Read(r, binary.BigEndian, &i)
	return
}

//
// Misc
//

func Concat(deep ...[]byte) []byte {
	length := 0
	for _, b := range deep {
		length += len(b)
	}

	flat := make([]byte, 0, length)
	for _, b := range deep {
		flat = append(flat, b...)
	}

	return flat
}