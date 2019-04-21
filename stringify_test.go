// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package stringify

import (
	"reflect"
	"testing"
)

func TestStringify(t *testing.T) {

	testCases := []struct {
		Input  interface{}
		Output interface{}
	}{
		struct {
			Input  interface{}
			Output interface{}
		}{
			Input:  map[interface{}]interface{}{"foo": "bar"},
			Output: map[string]interface{}{"foo": "bar"},
		},
	}

	for _, testCase := range testCases {

		got := StringifyMapKeys(testCase.Input)
		if !reflect.DeepEqual(got, testCase.Output) {
			t.Errorf("StringifyMapKeys(%v) == %v (%v), want %v (%s)", testCase.Input, got, reflect.TypeOf(got), testCase.Output, reflect.TypeOf(testCase.Output))
		}
	}

}
