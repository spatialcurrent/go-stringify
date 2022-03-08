// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package stringify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	in := []interface{}{
		[]interface{}{
			"a",
			"b",
			"c",
			[]interface{}{
				"d",
				"e",
				"f",
			},
			[]interface{}{
				"g",
				"h",
				"i",
			},
			"j",
			"k",
		},
		"l",
		"m",
		"n",
	}
	stringer := NewStringer("-", false, false, false)
	str, err := Concat(in, stringer)
	assert.NoError(t, err)
	assert.Equal(t, str, "abcdefghijklmn")
}
