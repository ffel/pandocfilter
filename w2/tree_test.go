package w2

import "fmt"

func ExampleTree() {
	json := decode(meta)

	f := NewTree()

	_ = Walk(f, "", json)

	fmt.Println(f.String())

	// Output:
	// + "": list:
	//     + "0": map:
	//         + "unMeta": map:
	//             + "author": map:
	//                 + "c": list:
	//                     + "0": map:
	//                         + "c": list:
	//                             + "0" - Str: "ffel"
	//                         + "t": value: MetaInlines
	//                 + "t": value: MetaList
	//             + "date": map:
	//                 + "c": list:
	//                     + "0" - Str: "december"
	//                     + "1" - Space
	//                     + "2" - Str: "2014"
	//                 + "t": value: MetaInlines
	//             + "title": map:
	//                 + "c": list:
	//                     + "0" - Str: "fix"
	//                     + "1" - Space
	//                     + "2" - Str: "meta"
	//                     + "3" - Space
	//                     + "4" - Str: "header"
	//                 + "t": value: MetaInlines
	//     + "1": list:
	//             + "Header": list:
	//                 + "0": value: 1
	//                 + "1": list:
	//                     + "0": value: hello
	//                     + "1": list:
	//                     + "2": list:
	//                 + "2": list:
	//                     + "0" - Str: "Hello"
	//             + "Para": list:
	//                 + "0" - Str: "Ordinary"
	//                 + "1" - Space
	//                 + "2" - Str: "text"
}
