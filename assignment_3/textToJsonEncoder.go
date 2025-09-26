package assignment_3

import "fmt"

type TextToJsonEncoder struct {
	text *Text
}

func (t *TextToJsonEncoder) PrintInJSON() {
	fmt.Println("Encoding Text to JSON")
	t.text.PrintConvertedInJSON()
}
