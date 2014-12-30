package pandocfilter

// Duplicator duplicates the input json into output json
//
// Duplicator is the most basic filter possible.
// It makes sense to embed the Duplicator in more advanced filters.
type Duplicator struct {
}

func (d Duplicator) List(key string, json []interface{}) (bool, interface{}) {
	return true, json
}

func (d Duplicator) Map(key string, json map[string]interface{}) (bool, interface{}) {
	return true, json
}

func (d Duplicator) Text(key string, value string) interface{} {
	return value
}

func (d Duplicator) Number(key string, value float64) interface{} {
	return value
}

func (d Duplicator) Bool(key string, value bool) interface{} {
	return value
}
