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

func ExampleStringer() {

	in := []interface{}{
		"a",
		1,
		1234567890.123,
		true,
		false,
		nil,
	}

	nodata := "-"
	decimal := false
	lower := false
	upper := false
	stringer := NewStringer(nodata, decimal, lower, upper)

	for _, x := range in {
		str, err := stringer(x)
		if err != nil {
			panic(err)
		}
		fmt.Println(str)
	}
	// Output: a
	// 1
	// 1.234567890123e+09
	// true
	// false
	// -

}
