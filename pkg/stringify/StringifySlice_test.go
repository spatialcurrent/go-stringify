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

func TestStringifySlice(t *testing.T) {
	in := []interface{}{
		"a",
		"b",
		"c",
		1,
		2,
		3,
		true,
		false,
	}
	out, err := StringifySlice(in, DefaultValueStringer(""))
	assert.Nil(t, err)
	assert.Equal(t, []string{"a", "b", "c", "1", "2", "3", "true", "false"}, out)
}

func TestStringifySliceInvalid(t *testing.T) {
	in := map[string]interface{}{
		"a": "b",
	}
	out, err := StringifySlice(in, DefaultValueStringer(""))
	assert.Equal(t, "StringifySlice cannot stringify map[string]interface {}", err.Error())
	assert.Equal(t, []string{}, out)
}
