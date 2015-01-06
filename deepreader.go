package pandocfilter

import (
	"errors"
	"strconv"
)

// GetString retrieves a sting from a deeply nested json
func GetString(json interface{}, indices ...string) (string, error) {
	// walk all bu
	for _, index := range indices /*[:len(slice)-1]*/ {
		i, err := strconv.Atoi(index)
		if err == nil {
			// json is a slice?
			s, sok := json.(Jslice)

			if !sok {
				return "", errors.New("GetString error - no slice for index " + index)
			}

			if i < 0 || i > len(s) {
				return "", errors.New("GetString error - slice out-of-range index " + index)
			}

			json = s[i]

		} else {
			// json is a map?
			m, mok := json.(Jmap)

			if !mok {
				return "", errors.New("GetString error - no map for index " + index)
			}

			_, ok := m[index]

			if !ok {
				return "", errors.New("GetString error reading map - no such key " + index)
			}

			json = m[index]
		}
	}

	str, strok := json.(string)

	if !strok {
		return "", errors.New("GetString error - no string")
	}

	return str, nil
}
