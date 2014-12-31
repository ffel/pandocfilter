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
	//                             + "0": map:
	//                                 + "c": value: ffel
	//                                 + "t": value: Str
	//                         + "t": value: MetaInlines
	//                 + "t": value: MetaList
	//             + "date": map:
	//                 + "c": list:
	//                     + "0": map:
	//                         + "c": value: december
	//                         + "t": value: Str
	//                     + "1": map:
	//                         + "c": list:
	//                         + "t": value: Space
	//                     + "2": map:
	//                         + "c": value: 2014
	//                         + "t": value: Str
	//                 + "t": value: MetaInlines
	//             + "title": map:
	//                 + "c": list:
	//                     + "0": map:
	//                         + "c": value: fix
	//                         + "t": value: Str
	//                     + "1": map:
	//                         + "c": list:
	//                         + "t": value: Space
	//                     + "2": map:
	//                         + "c": value: meta
	//                         + "t": value: Str
	//                     + "3": map:
	//                         + "c": list:
	//                         + "t": value: Space
	//                     + "4": map:
	//                         + "c": value: header
	//                         + "t": value: Str
	//                 + "t": value: MetaInlines
	//     + "1": list:
	//         + "0": map:
	//             + "c": list:
	//                 + "0": value: 1
	//                 + "1": list:
	//                     + "0": value: hello
	//                     + "1": list:
	//                     + "2": list:
	//                 + "2": list:
	//                     + "0": map:
	//                         + "c": value: Hello
	//                         + "t": value: Str
	//             + "t": value: Header
	//         + "1": map:
	//             + "c": list:
	//                 + "0": map:
	//                     + "c": value: Ordinary
	//                     + "t": value: Str
	//                 + "1": map:
	//                     + "c": list:
	//                     + "t": value: Space
	//                 + "2": map:
	//                     + "c": value: text
	//                     + "t": value: Str
	//             + "t": value: Para
}
