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

// InterfaceSliceToStringSlice converts a slice of interface{} to a slice of strings using fmt.Sprint.
// Use StringifySlice for options for stringifying.
func InterfaceSliceToStringSlice(values []interface{}) []string {
	stringSlice := make([]string, 0, len(values))
	for _, v := range values {
		stringSlice = append(stringSlice, fmt.Sprint(v))
	}
	return stringSlice
}
