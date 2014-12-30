package pandocfilter

import (
	"log"
	"strconv"
)

// Filter defines the methods to process pandoc json.  Methods List
// and Map ask whether Walk should walk its members or not.  In
// case not, it returns its own members
//
// Methods for basic types, String, Bool, Number return an interface
// to allow these to return a larger object (say, a link)
type Filter interface {
	List(key string, json []interface{}) (bool, interface{})         // announce start of list
	Map(key string, json map[string]interface{}) (bool, interface{}) // announce non ct maps
	String(key string, value string) interface{}
	Number(key string, value float64) interface{}
	Bool(key string, value bool) interface{}
}

func Walk(filter Filter, key string, json interface{}) interface{} {
	switch elem := json.(type) {
	case []interface{}:
		decend, result := filter.List(key, json.([]interface{}))
		if !decend {
			return result
		}

		// okay, we do the walk into the slice elements ourselves
		slice := make([]interface{}, 0, len(json.([]interface{})))
		for i, v := range json.([]interface{}) {
			slice = append(slice, Walk(filter, strconv.Itoa(i), v))
		}
		return slice
	case map[string]interface{}:
		// check for common c(ontent) t(ype) json object
		typekey, tok := json.(map[string]interface{})["t"]
		contents, cok := json.(map[string]interface{})["c"]

		if tok && cok {
			result := make(map[string]interface{})
			result["t"] = typekey.(string)
			result["c"] = Walk(filter, typekey.(string), contents)

			return result
		}

		// check if filter.Map returns its own result (decend == false)
		decend, result := filter.Map(key, json.(map[string]interface{}))
		if !decend {
			return result
		}

		// okay, we do the walk into object
		m := make(map[string]interface{})
		for k, v := range json.(map[string]interface{}) {
			m[k] = Walk(filter, k, v)
		}
		return m

	case string:
		return filter.String(key, json.(string))

	case float64:
		return filter.Number(key, json.(float64))

	case bool:
		return filter.Bool(key, json.(bool))

	default:
		log.Printf("no support for %T %#v\n", elem, elem)
		return json
	}

}
