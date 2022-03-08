// =================================================================
//
// Copyright (C) 2022 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package stringify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceSliceToStringSlice(t *testing.T) {
	is := []interface{}{"a", "b", "c"}
	ss := []string{"a", "b", "c"}
	assert.Equal(t, ss, InterfaceSliceToStringSlice(is))
}
