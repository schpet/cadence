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
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence"
	"github.com/onflow/cadence/encoding/json"
	"github.com/onflow/cadence/runtime/common"
)

func measureValueCodec(value cadence.Value) (
	encodingDuration time.Duration,
	decodingDuration time.Duration,
	size int,
	err error,
) {
	encoder, decoder, buffer := NewTestCodec()

	t0 := time.Now()
	err = encoder.EncodeValue(value)
	t1 := time.Now()
	if err != nil {
		return
	}

	size = buffer.Len()

	t2 := time.Now()
	_, err = decoder.Decode()
	t3 := time.Now()
	if err != nil {
		return
	}

	encodingDuration = t1.Sub(t0)
	decodingDuration = t3.Sub(t2)
	return
}

func measureJsonCodec(value cadence.Value) (
	encodingDuration time.Duration,
	decodingDuration time.Duration,
	size int,
	err error,
) {
	t0 := time.Now()
	blob, err := json.Encode(value)
	t1 := time.Now()
	if err != nil {
		return
	}

	size = len(blob)

	t2 := time.Now()
	_, err = json.Decode(nil, blob)
	t3 := time.Now()
	if err != nil {
		return
	}

	encodingDuration = t1.Sub(t0)
	decodingDuration = t3.Sub(t2)
	return
}

func average(ints []int) int {
	sum := 0

	for _, i := range ints {
		sum += i
	}

	return sum / len(ints)
}

func TestCodecPerformance(t *testing.T) {
	// not parallel

	const iterations = 100
	values := []cadence.Value{
		cadence.String("This sentence is composed of exactly fifty-nine characters."),
		cadence.NewResource([]cadence.Value{
			cadence.Bool(true),
		}).WithType(cadence.NewResourceType(
			common.AddressLocation{
				Address: common.Address{1, 2, 3, 4, 5, 6, 7, 8},
				Name:    "NFT",
			},
			"A.12345678.NFT",
			[]cadence.Field{
				{
					Identifier: "Awesome?",
					Type:       cadence.BoolType{},
				},
			},
			[][]cadence.Parameter{},
		)),
	}

	jsonEncodingDurations := make([][]int, len(values))
	jsonDecodingDurations := make([][]int, len(values))
	jsonEncodedSizes := make([][]int, len(values))

	valueEncodingDurations := make([][]int, len(values))
	valueDecodingDurations := make([][]int, len(values))
	valueEncodedSizes := make([][]int, len(values))

	for i, value := range values {
		jsonEncodingDurations[i] = make([]int, iterations)
		jsonDecodingDurations[i] = make([]int, iterations)
		jsonEncodedSizes[i] = make([]int, iterations)

		valueEncodingDurations[i] = make([]int, iterations)
		valueDecodingDurations[i] = make([]int, iterations)
		valueEncodedSizes[i] = make([]int, iterations)

		for j := 0; j < iterations; j++ {
			encodingDuration, decodingDuration, size, err := measureJsonCodec(value)
			require.NoError(t, err, "json codec error")
			jsonEncodingDurations[i][j] = int(encodingDuration)
			jsonDecodingDurations[i][j] = int(decodingDuration)
			jsonEncodedSizes[i][j] = size

			encodingDuration, decodingDuration, size, err = measureValueCodec(value)
			require.NoError(t, err, "value codec error")
			valueEncodingDurations[i][j] = int(encodingDuration)
			valueDecodingDurations[i][j] = int(decodingDuration)
			valueEncodedSizes[i][j] = size
		}
	}

	averageJsonEncodingDurations := make([]int, len(values))
	averageJsonDecodingDurations := make([]int, len(values))
	averageJsonEncodedSizes := make([]int, len(values))

	averageValueEncodingDurations := make([]int, len(values))
	averageValueDecodingDurations := make([]int, len(values))
	averageValueEncodedSizes := make([]int, len(values))

	diffEncodingDurations := make([]float64, len(values))
	diffDecodingDurations := make([]float64, len(values))
	diffEncodedSizes := make([]float64, len(values))

	minJsonEncodingDurations := make([]int, len(values))
	minJsonDecodingDurations := make([]int, len(values))
	minJsonEncodedSizes := make([]int, len(values))

	minValueEncodingDurations := make([]int, len(values))
	minValueDecodingDurations := make([]int, len(values))
	minValueEncodedSizes := make([]int, len(values))

	maxJsonEncodingDurations := make([]int, len(values))
	maxJsonDecodingDurations := make([]int, len(values))
	maxJsonEncodedSizes := make([]int, len(values))

	maxValueEncodingDurations := make([]int, len(values))
	maxValueDecodingDurations := make([]int, len(values))
	maxValueEncodedSizes := make([]int, len(values))

	for i := range values {
		sort.Ints(jsonEncodingDurations[i])
		sort.Ints(jsonDecodingDurations[i])
		sort.Ints(jsonEncodedSizes[i])

		sort.Ints(valueEncodingDurations[i])
		sort.Ints(valueDecodingDurations[i])
		sort.Ints(valueEncodedSizes[i])

		averageJsonEncodingDurations[i] = average(jsonEncodingDurations[i])
		averageJsonDecodingDurations[i] = average(jsonDecodingDurations[i])
		averageJsonEncodedSizes[i] = average(jsonEncodedSizes[i])

		averageValueEncodingDurations[i] = average(valueEncodingDurations[i])
		averageValueDecodingDurations[i] = average(valueDecodingDurations[i])
		averageValueEncodedSizes[i] = average(valueEncodedSizes[i])

		diffEncodingDurations[i] = float64(averageJsonEncodingDurations[i]) / float64(averageValueEncodingDurations[i])
		diffDecodingDurations[i] = float64(averageJsonDecodingDurations[i]) / float64(averageValueDecodingDurations[i])
		diffEncodedSizes[i] = float64(averageJsonEncodedSizes[i]) / float64(averageValueEncodedSizes[i])

		minJsonEncodingDurations[i] = jsonEncodingDurations[i][0]
		minJsonDecodingDurations[i] = jsonDecodingDurations[i][0]
		minJsonEncodedSizes[i] = jsonEncodedSizes[i][0]

		minValueEncodingDurations[i] = valueEncodingDurations[i][0]
		minValueDecodingDurations[i] = valueDecodingDurations[i][0]
		minValueEncodedSizes[i] = valueEncodedSizes[i][0]

		maxJsonEncodingDurations[i] = jsonEncodingDurations[i][iterations-1]
		maxJsonDecodingDurations[i] = jsonDecodingDurations[i][iterations-1]
		maxJsonEncodedSizes[i] = jsonEncodedSizes[i][iterations-1]

		maxValueEncodingDurations[i] = valueEncodingDurations[i][iterations-1]
		maxValueDecodingDurations[i] = valueDecodingDurations[i][iterations-1]
		maxValueEncodedSizes[i] = valueEncodedSizes[i][iterations-1]
	}

	fmt.Println("Results:")
	fmt.Println("(higher speedup/reduction is better)")

	for i, value := range values {
		fmt.Println("Cadence Value: ", value.String())
		fmt.Println("Encoding Time (ns):")
		fmt.Println("\tJSON Average: ", averageJsonEncodingDurations[i])
		fmt.Println("\tJSON Minimum: ", minJsonEncodingDurations[i])
		fmt.Println("\tJSON Maximum: ", maxJsonEncodingDurations[i])
		fmt.Println("\tValue Average: ", averageValueEncodingDurations[i])
		fmt.Println("\tValue Minimum: ", minValueEncodingDurations[i])
		fmt.Println("\tValue Maximum: ", maxValueEncodingDurations[i])
		fmt.Println("\tSpeedup:", diffEncodingDurations[i])
		fmt.Println("Decoding Time (ns):")
		fmt.Println("\tJSON Average: ", averageJsonDecodingDurations[i])
		fmt.Println("\tJSON Minimum: ", minJsonDecodingDurations[i])
		fmt.Println("\tJSON Maximum: ", maxJsonDecodingDurations[i])
		fmt.Println("\tValue Average: ", averageValueDecodingDurations[i])
		fmt.Println("\tValue Minimum: ", minValueDecodingDurations[i])
		fmt.Println("\tValue Maximum: ", maxValueDecodingDurations[i])
		fmt.Println("\tSpeedup:", diffDecodingDurations[i])
		fmt.Println("Size (bytes):")
		fmt.Println("\tJSON Average: ", averageJsonEncodedSizes[i])
		fmt.Println("\tJSON Minimum: ", minJsonEncodedSizes[i])
		fmt.Println("\tJSON Maximum: ", maxJsonEncodedSizes[i])
		fmt.Println("\tValue Average: ", averageValueEncodedSizes[i])
		fmt.Println("\tValue Minimum: ", minValueEncodedSizes[i])
		fmt.Println("\tValue Maximum: ", maxValueEncodedSizes[i])
		fmt.Println("\tReduction:", diffEncodedSizes[i])
		fmt.Println()
	}

	t.Fail()
}
