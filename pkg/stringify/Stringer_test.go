// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package stringify

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

var (
	testInput = []interface{}{
		"a",
		1,
		1234567890.123,
		true,
		false,
		nil,
	}
)

func TestStringer(t *testing.T) {

	out := []string{
		"a",
		"1",
		"1.234567890123e+09",
		"true",
		"false",
		"-",
	}

	stringer := NewStringer("-", false, false, false)

	for i, x := range testInput {
		str, err := stringer(x)
		assert.Nil(t, err)
		assert.Equal(t, out[i], str)
	}

}

func TestStringerUpper(t *testing.T) {

	out := []string{
		"A",
		"1",
		"1.234567890123E+09",
		"TRUE",
		"FALSE",
		"-",
	}

	stringer := NewStringer("-", false, false, true)

	for i, x := range testInput {
		str, err := stringer(x)
		assert.Nil(t, err)
		assert.Equal(t, out[i], str)
	}

}

func TestStringerDecimal(t *testing.T) {

	out := []string{
		"a",
		"1",
		"1234567890.123000",
		"true",
		"false",
		"-",
	}

	stringer := NewStringer("-", true, false, false)

	for i, x := range testInput {
		str, err := stringer(x)
		assert.Nil(t, err)
		assert.Equal(t, out[i], str)
	}

}
