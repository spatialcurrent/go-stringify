// =================================================================
//
// Copyright (C) 2022 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package stringify

import (
	"fmt"
)

func ExampleStringifyMapKeys_default() {

	in := map[interface{}]interface{}{"a": "x", "b": "y", "c": "z"}

	nodata := ""
	decimal := false
	lower := false
	upper := false
	stringer := NewStringer(nodata, decimal, lower, upper)
	out, err := StringifyMapKeys(in, stringer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", out)
	// Output: map[string]interface {}{"a":"x", "b":"y", "c":"z"}
}

func ExampleStringifyMapKeys_upper() {

	in := map[interface{}]interface{}{"a": "x", "b": "y", "c": "z"}

	nodata := ""
	decimal := false
	lower := false
	upper := true
	stringer := NewStringer(nodata, decimal, lower, upper)
	out, err := StringifyMapKeys(in, stringer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", out)
	// Output: map[string]interface {}{"A":"x", "B":"y", "C":"z"}
}
