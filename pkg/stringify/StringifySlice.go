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

import (
	"github.com/pkg/errors"
)

// StringifySlice converts all the objects in a slice or array to their default string representatoin.
// Returns a []string, and error if any.
func StringifySlice(in interface{}, valueStringer func(object interface{}) (string, error)) ([]string, error) {
	v := reflect.ValueOf(in)
	k := v.Type().Kind()
	if k != reflect.Array && k != reflect.Slice {
		return make([]string, 0), fmt.Errorf("StringifySlice cannot stringify %T", in)
	}
	out := make([]string, 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		str, err := valueStringer(v.Index(i).Interface())
		if err != nil {
			return out, errors.Wrap(err, fmt.Sprintf("error stringifying element at index %d", i))
		}
		out = append(out, str)
	}
	return out, nil
}
