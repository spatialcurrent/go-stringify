// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package stringify

import (
	"reflect"
)

// Concat recursively flattens, stringifies, and concats an array or array of arrays.
func Concat(in interface{}, stringer Stringer) (string, error) {
	v := reflect.ValueOf(in)
	if k := v.Type().Kind(); k == reflect.Array || k == reflect.Slice {
		out := ""
		for i := 0; i < v.Len(); i++ {
			vi := v.Index(i).Interface()
			if ki := reflect.ValueOf(vi).Type().Kind(); ki == reflect.Array || ki == reflect.Slice {
				str, err := Concat(vi, stringer)
				if err != nil {
					return out, err
				}
				out += str
				continue
			}
			str, err := stringer(vi)
			if err != nil {
				return out, err
			}
			out += str
		}
		return out, nil
	}
	return stringer(in)
}
