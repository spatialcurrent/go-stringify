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

func ExampleStringifySlice_default() {
	in := []interface{}{
		"a",
		"b",
		"c",
		1,
		2,
		3,
		1234567890.123,
		true,
		false,
		nil,
	}
	nodata := "-"
	decimal := false
	lower := false
	upper := true
	stringer := NewStringer(nodata, decimal, lower, upper)
	out, err := StringifySlice(in, stringer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", out)
	// Output: []string{"A", "B", "C", "1", "2", "3", "1.234567890123E+09", "TRUE", "FALSE", "-"}
}

func ExampleStringifySlice_decimal() {
	in := []interface{}{
		"a",
		"b",
		"c",
		1,
		2,
		3,
		1234567890.123,
		true,
		false,
		nil,
	}
	nodata := "-"
	decimal := true
	lower := false
	upper := true
	stringer := NewStringer(nodata, decimal, lower, upper)
	out, err := StringifySlice(in, stringer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", out)
	// Output: []string{"A", "B", "C", "1", "2", "3", "1234567890.123000", "TRUE", "FALSE", "-"}
}
