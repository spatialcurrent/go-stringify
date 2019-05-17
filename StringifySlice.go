// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package stringify

import (
	"fmt"
	"reflect"
)

// StringifySlice stringifies all the objects in a slice or array.
// Returns a []string, and error if any.
func StringifySlice(in interface{}) ([]string, error) {
	v := reflect.ValueOf(in)
	k := v.Type().Kind()
	if k != reflect.Array && k != reflect.Slice {
		return make([]string, 0), fmt.Errorf("StringifySlice cannot stringify %T", in)
	}
	out := make([]string, 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		out = append(out, fmt.Sprint(v.Index(i).Interface()))
	}
	return out, nil
}
