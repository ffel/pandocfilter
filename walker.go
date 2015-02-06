// package to support pandoc filters in Go
package pandocfilter

import (
	"sort"
	"strconv"
)

// Filter prescibes Value method that is called by Walk.
//
// The returned bool tells the walker whether or not to
// decend in the returned interface.
//
// The returned interface{} is either the original value or
// a modified version by Filter.
type Filter interface {
	Value(key string, value interface{}) (bool, interface{})
}

// Walk walks the pandoc json structure, calling filter on every element
func Walk(filter Filter, key string, json interface{}) interface{} {
	list, isList := json.([]interface{})
	set, isSet := json.(map[string]interface{})

	switch {
	case isList:
		decend, result := filter.Value(key, list)

		if !decend {
			return result
		}

		slice := make([]interface{}, 0, len(list))

		for i, v := range list {
			if next := Walk(filter, strconv.Itoa(i), v); next != nil {
				slice = append(slice, next)
			}
		}

		return slice

	case isSet:
		decend, result := filter.Value(key, set)

		if !decend {
			return result
		}

		m := make(map[string]interface{})

		for _, k := range keys(set, true) {
			if next := Walk(filter, k, set[k]); next != nil {
				m[k] = next
			}
		}

		return m

	default:
		_, result := filter.Value(key, json)

		return result
	}
}

// keys return a sorted list of keys in `set` depending on `sorted`
func keys(set map[string]interface{}, sorted bool) []string {
	kk := make([]string, len(set))
	i := 0
	for k, _ := range set {
		kk[i] = k
		i++
	}

	if sorted {
		sort.Strings(kk)
	}

	return kk
}
