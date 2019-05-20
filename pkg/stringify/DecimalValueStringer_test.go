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

func TestDecimalValueStringer(t *testing.T) {

	in := []interface{}{
		"a",
		1,
		1234567890.123,
		true,
		false,
		nil,
	}

	out := []string{
		"a",
		"1",
		"1234567890.123000",
		"true",
		"false",
		"-",
	}

	stringer := DecimalValueStringer("-")

	for i, x := range in {
		str, err := stringer(x)
		assert.Nil(t, err)
		assert.Equal(t, out[i], str)
	}

}
