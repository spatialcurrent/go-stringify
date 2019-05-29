// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package stringify

// StringSliceToInterfaceSlice converts a slice of strings to a slice of interface{}.
func StringSliceToInterfaceSlice(values []string) []interface{} {
	interfaceSlice := make([]interface{}, 0, len(values))
	for _, v := range values {
		interfaceSlice = append(interfaceSlice, v)
	}
	return interfaceSlice
}
