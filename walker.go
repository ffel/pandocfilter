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
//
// Typically, you don't want the walker to decend in case
// the original data is modified.
type Filter interface {
	Value(level int, key string, value interface{}) (bool, interface{})
}

func Walk(filter Filter, json interface{}) interface{} {
	return walk(filter, 0, "", json)
}

// Walk walks the pandoc json structure, calling filter on every element
func walk(filter Filter, level int, key string, json interface{}) interface{} {
	list, isList := json.([]interface{})
	set, isSet := json.(map[string]interface{})

	switch {
	case isList:
		decend, result := filter.Value(level, key, list)

		if !decend {
			return result
		}

		slice := make([]interface{}, 0, len(list))

		for i, v := range list {
			if next := walk(filter, level+1, strconv.Itoa(i), v); next != nil {
				slice = append(slice, next)
			}
		}

		return slice

	case isSet:
		decend, result := filter.Value(level, key, set)

		if !decend {
			return result
		}

		m := make(map[string]interface{})

		for _, k := range keys(set, true) {
			if next := walk(filter, level+1, k, set[k]); next != nil {
				m[k] = next
			}
		}

		return m

	default:
		_, result := filter.Value(level, key, json)

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
