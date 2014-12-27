package walker

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

/*
wat wil je precies:

1.	het filter op een generieke manier nabouwen, evt zo dat de
	uitkomst minder json maar meer pandoc is.

2.	uiteindelijk wil ik iets waarin ik kan zeggen dat ik string
	x in type y kan vervangen door iets anders, en dan eigenlijk
	zo nodig het complete element.

Hoe kan je met walker nu aanpassingen maken in text dat in block
quotes staat?

1.	Maak een walker die op zoek gaat naar block quoted text en start
	op de json die je daar binnenkrijgt een tweede walker die de
	vervanging doet.
*/

type Walker interface {
	// je kan retouneren of Walk verder door de elementen heen loopt
	// of dat List dat zelf doet: wanneer de structuur bekend verondersteld
	// mag worden, zou je true kunnen teruggeven
	List(key string, json []interface{}, level int) bool
	String(key, value string, level int)
	Number(key string, value float64, level int)
	Bool(key string, value bool, level int)
	Map(key string, json map[string]interface{}, level int) bool // used for meta
}

func Walk(walker Walker, json interface{}, key string, level int) {
	switch elem := json.(type) {
	case []interface{}:
		if walker.List(key, json.([]interface{}), level) {
			for i, v := range json.([]interface{}) {
				Walk(walker, v, strconv.Itoa(i), level+1)
			}
		}
	case map[string]interface{}:
		t, tok := json.(map[string]interface{})["t"]
		c, cok := json.(map[string]interface{})["c"]
		if tok && cok {
			Walk(walker, c, t.(string), level)
		} else {
			if walker.Map(key, json.(map[string]interface{}), level) {
				for k, v := range json.(map[string]interface{}) {
					Walk(walker, v, k, level+1)
				}
			}
		}
	case string:
		walker.String(key, json.(string), level)
	case float64:
		walker.Number(key, json.(float64), level)
	default:
		fmt.Printf("%s* what to do with %T --- %v\n",
			indent(level), elem, elem)
	}
}

func indent(level int) string {
	return strings.Repeat("    ", level)
}

// http://stackoverflow.com/questions/13765797/the-best-way-to-get-a-string-from-a-writer-in-go
type dumper struct {
	w *bytes.Buffer
}

func (d dumper) List(key string, json []interface{}, level int) bool {
	fmt.Fprintf(d.w, "%s+ list %q\n", indent(level), key)

	// let walker continue the traversal
	return true
}

func (d dumper) Map(key string, json map[string]interface{}, level int) bool {
	fmt.Fprintf(d.w, "%s+ map %q\n", indent(level), key)

	// let walker continue the traversal into map
	return true
}

func (d dumper) String(key, value string, level int) {
	fmt.Fprintf(d.w, "%s+ string %q %q\n", indent(level), key, value)
}

func (d dumper) Number(key string, value float64, level int) {
	fmt.Fprintf(d.w, "%s+ number %q %v\n", indent(level), key, value)
}

func (d dumper) Bool(key string, value bool, level int) {
	fmt.Fprintf(d.w, "%s+ bool %q %t\n", indent(level), key, value)
}
