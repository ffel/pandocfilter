// package to support pandoc filters in Go
package pandocfilter

import (
	"sort"
	"strconv"
)

// Filter prescibes Value method that is called by Walk
// Value should return true or meaningfull 2nd return value
// (for leaves 1st return argument is ignored)
type Filter interface {
	Value(key string, value interface{}) (bool, interface{})
}

// alleen nog maar het doorlopen van de structuur
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
			slice = append(slice, Walk(filter, strconv.Itoa(i), v))
		}

		return slice

	case isSet:
		decend, result := filter.Value(key, set)

		if !decend {
			return result
		}

		m := make(map[string]interface{})

		for _, k := range keys(set, true) {
			m[k] = Walk(filter, k, set[k])
		}

		return m

	default:
		// log.Printf("unexpected value %T %#v\n", json, json)
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
