// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package stringify

import (
	"fmt"
)

// DefaultValueStringer returns a function that converts an object into a string.
// The returned function converts certain floats to strings in scientific notation, rather than just decimal notation.
var DefaultValueStringer = func(noDataValue string) func(object interface{}) (string, error) {
	return func(object interface{}) (string, error) {
		if object == nil {
			return noDataValue, nil
		}
		return fmt.Sprint(object), nil
	}
}
