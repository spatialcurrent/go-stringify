// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package stringify

import (
	"github.com/pkg/errors"
)

// StringifyMapKeys recursively stringifying map keys from interface{} to string.
// This functionality is inspired by work done in https://github.com/gohugoio/hugo, but
// support many more types, including:
//	- []interface{}
//	- [][]interface{}
//	- []map[interface{}]interface{}
//	- map[interface{}]interface{}
//	- map[string]interface{}
//	- map[string][]interface{}
//	- map[string]map[string][]interface{}
//	- map[interface{}]struct{}
// See https://github.com/gohugoio/hugo/pull/4138 for background.
func StringifyMapKeys(in interface{}, stringer Stringer) (interface{}, error) {

	switch in := in.(type) {
	case []interface{}:
		res := make([]interface{}, len(in))
		for i, value := range in {
			newValue, err := StringifyMapKeys(value, stringer)
			if err != nil {
				return in, err
			}
			res[i] = newValue
		}
		return res, nil
	case [][]interface{}:
		res := make([][]interface{}, len(in))
		for i, value := range in {
			newValue, err := StringifyMapKeys(value, stringer)
			if err != nil {
				return in, err
			}
			res[i] = newValue.([]interface{})
		}
		return res, nil
	case []map[interface{}]interface{}:
		res := make([]map[string]interface{}, len(in))
		for i, v := range in {
			newValue, err := StringifyMapKeys(v, stringer)
			if err != nil {
				return in, err
			}
			res[i] = newValue.(map[string]interface{})
		}
		return res, nil
	case map[interface{}]interface{}:
		res := make(map[string]interface{})
		for key, value := range in {
			newKey, err := stringer(key)
			if err != nil {
				return in, errors.Wrapf(err, "error stringifying map key %q", key)
			}
			newValue, err := StringifyMapKeys(value, stringer)
			if err != nil {
				return in, err
			}
			res[newKey] = newValue
		}
		return res, nil
	case map[string]interface{}:
		res := make(map[string]interface{})
		for key, value := range in {
			newKey, err := stringer(key)
			if err != nil {
				return in, errors.Wrapf(err, "error stringifying map key %q", key)
			}
			newValue, err := StringifyMapKeys(value, stringer)
			if err != nil {
				return in, err
			}
			res[newKey] = newValue
		}
		return res, nil
	case map[string][]interface{}:
		res := make(map[string][]interface{})
		for key, value := range in {
			newValue, err := StringifyMapKeys(value, stringer)
			if err != nil {
				return in, err
			}
			res[key] = newValue.([]interface{})
		}
		return res, nil
	case map[string]map[string]interface{}:
		res := make(map[string]map[string]interface{})
		for key, value := range in {
			newValue, err := StringifyMapKeys(value, stringer)
			if err != nil {
				return in, err
			}
			res[key] = newValue.(map[string]interface{})
		}
		return res, nil
	case map[string]map[string][]interface{}:
		res := make(map[string]map[string][]interface{})
		for key, value := range in {
			newValue, err := StringifyMapKeys(value, stringer)
			if err != nil {
				return in, err
			}
			res[key] = newValue.(map[string][]interface{})
		}
		return res, nil
	case map[interface{}]struct{}:
		res := make(map[string]interface{})
		for key, value := range in {
			newKey, err := stringer(key)
			if err != nil {
				return in, errors.Wrapf(err, "error stringifying map key %q", key)
			}
			newValue, err := StringifyMapKeys(value, stringer)
			if err != nil {
				return in, err
			}
			res[newKey] = newValue
		}
		return res, nil
	case map[string]struct{}:
		res := make(map[string]struct{})
		for key, value := range in {
			newKey, err := stringer(key)
			if err != nil {
				return in, errors.Wrapf(err, "error stringifying map key %q", key)
			}
			newValue, err := StringifyMapKeys(value, stringer)
			if err != nil {
				return in, err
			}
			res[newKey] = newValue.(struct{})
		}
		return res, nil
	default:
		return in, nil
	}
}
