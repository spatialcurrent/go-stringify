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

// DecimalValueStringer returns a function that converts an object into a string.
// The returned function converts floats to strings in decimal notation rather than scientific notation.
var DecimalValueStringer = func(noDataValue string) func(object interface{}) (string, error) {
	return func(object interface{}) (string, error) {
		if object == nil {
			return noDataValue, nil
		}
		switch value := object.(type) {
		case float32, float64:
			return fmt.Sprintf("%f", value), nil
		}
		return fmt.Sprint(object), nil
	}
}
