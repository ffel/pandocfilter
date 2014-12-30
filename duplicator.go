package pandocfilter

// Duplicator duplicates the input json into output json
//
// Duplicator is the most basic filter possible.
// It makes sense to embed the Duplicator in more advanced filters.
type Duplicator struct {
}

func (p Duplicator) List(key string, json []interface{}) (bool, interface{}) {
	return true, json
}

func (p Duplicator) Map(key string, json map[string]interface{}) (bool, interface{}) {
	return true, json
}
