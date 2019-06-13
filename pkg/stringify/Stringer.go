// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package stringify

import (
	"fmt"
	"strings"
)

// Stringer is a type alias for a function that converts an object into a string and returns an error if any.
type Stringer func(object interface{}) (string, error)

// NewStringer creates a new Stringer.
// If decimal is true, the stringer converts floats to strings in decimal notation rather than scientific notation.
// If lower, then returned strings are converted to lower case.
// If upper, then returned strings are converted to upper case.
func NewStringer(noDataValue string, decimal bool, lower bool, upper bool) Stringer {
	if decimal {

		if lower {
			return func(object interface{}) (string, error) {
				if object == nil {
					return noDataValue, nil
				}
				switch value := object.(type) {
				case float32, float64:
					return fmt.Sprintf("%f", value), nil
				}
				return strings.ToLower(fmt.Sprint(object)), nil
			}
		}

		if upper {
			return func(object interface{}) (string, error) {
				if object == nil {
					return noDataValue, nil
				}
				switch value := object.(type) {
				case float32, float64:
					return fmt.Sprintf("%f", value), nil
				}
				return strings.ToUpper(fmt.Sprint(object)), nil
			}
		}

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

	if lower {
		return func(object interface{}) (string, error) {
			if object == nil {
				return noDataValue, nil
			}
			return strings.ToLower(fmt.Sprint(object)), nil
		}
	}

	if upper {
		return func(object interface{}) (string, error) {
			if object == nil {
				return noDataValue, nil
			}
			return strings.ToUpper(fmt.Sprint(object)), nil
		}
	}

	return func(object interface{}) (string, error) {
		if object == nil {
			return noDataValue, nil
		}
		return fmt.Sprint(object), nil
	}

}

// NewDefaultStringer returns a new default stringer.
func NewDefaultStringer() Stringer {
	return NewStringer("", false, false, false)
}

// NewDecimalStringer returns a new stringer that uses decimal notation.
func NewDecimalStringer() Stringer {
	return NewStringer("", false, false, false)
}
