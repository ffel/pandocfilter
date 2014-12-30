package pandocfilter

import "fmt"

func ExampleTree() {
	json := decode(data)

	f := NewTree()

	_ = Walk(f, "", json)

	fmt.Println(f.String())

	// Output:
	// + list ""
	//     + map "0"
	//         + map "unMeta"
	//     + list "1"
	//         + list "Header"
	//             + number "0": 1
	//             + list "1"
	//                 + text "0": "hallo"
	//                 + list "1"
	//                 + list "2"
	//             + list "2"
	//                 + text "Str": "Hallo"
	//         + list "Para"
	//             + text "Str": "Hallo"
	//             + list "Space"
	//             + text "Str": "Wereld!"
}
