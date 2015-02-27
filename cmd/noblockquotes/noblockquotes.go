// pandoc filter that removes block quotes from the output

package main

import "github.com/ffel/pandocfilter"

func main() {
	pandocfilter.Run(remover{})
}

type remover struct{}

func (r remover) Value(level int, key string, value interface{}) (bool, interface{}) {
	ok, t, _ := pandocfilter.IsTypeContents(value)

	if ok && t == "BlockQuote" {
		return false, nil
	}

	return true, value
}
