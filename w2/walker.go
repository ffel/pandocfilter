package w2

import (
	"sort"
	"strconv"
)

// wellicht een idee dat dit nog interne functies zijn, maar
// ik kan me ook helper functies voorstellen waarmee je van
// bijvoorbeeld een Set kan nagaan of er sprake is van een
// string waarde
type Filter interface {
	List(key string, value []interface{}) (bool, interface{})
	Set(key string, value map[string]interface{}) (bool, interface{})
	Value(key string, value interface{}) interface{}
}

// alleen nog maar het doorlopen van de structuur
func Walk(filter Filter, key string, json interface{}) interface{} {
	list, isList := json.([]interface{})
	set, isSet := json.(map[string]interface{})

	switch {
	case isList:
		decend, result := filter.List(key, list)

		if !decend {
			return result
		}

		slice := make([]interface{}, 0, len(list))

		for i, v := range list {
			slice = append(slice, Walk(filter, strconv.Itoa(i), v))
		}

		return slice

	case isSet:
		decend, result := filter.Set(key, set)

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
		return filter.Value(key, json)
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
