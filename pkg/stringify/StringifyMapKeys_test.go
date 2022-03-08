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

func TestStringifyMapKeys(t *testing.T) {

	in := map[interface{}]interface{}{"a": "x", "b": "y", "c": "z"}
	expected := map[string]interface{}{"a": "x", "b": "y", "c": "z"}

	stringer := NewStringer("", false, false, false)
	got, err := StringifyMapKeys(in, stringer)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)

}

func TestStringifyMapKeysUpper(t *testing.T) {

	in := map[interface{}]interface{}{"a": "x", "b": "y", "c": "z"}
	expected := map[string]interface{}{"A": "x", "B": "y", "C": "z"}

	stringer := NewStringer("", false, false, true)
	got, err := StringifyMapKeys(in, stringer)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)

}
