package w2

import "testing"

/*
hello is formatted json produced by pandoc of the following

	Hello
	=====

	world!
*/
const hello = `[ { "unMeta" : {  } },
  [ { "c" : [ 1,
          [ "hello",
            [  ],
            [  ]
          ],
          [ { "c" : "Hello",
              "t" : "Str"
            } ]
        ],
      "t" : "Header"
    },
    { "c" : [ { "c" : "world!",
            "t" : "Str"
          } ],
      "t" : "Para"
    }
  ]
]`

// a very simple mod test
func Test_modify(t *testing.T) {
	json := decode(hello)

	out := Walk(mod{}, "", json)

	expected := `[{"unMeta":{}},[{"c":[1,["hello",[],[]],[{"c":"Hello","t":"Str"}]],"t":"Header"},{"c":[{"c":"Universe!!","t":"Str"}],"t":"Para"}]]`

	if got := encode(out); expected+"\n" != got {
		t.Errorf("unexpected mod value\n%s\n%s\n", got, expected)
	}
}

type mod struct{}

func (m mod) Value(key string, value interface{}) (bool, interface{}) {
	// CString is a little bit easier than IsTypeContents(value)
	ok, t, c := CString(value)

	if ok && t == Str && c == "world!" {
		return false, WrapCString(Str, "Universe!!")
	}

	return true, value
}
