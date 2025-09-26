package assignment_3

func Print() {
	client := &Client{}
	json := &JSON{}

	client.PrintText(json)

	text := &Text{}
	textToJsonEncoder := &TextToJsonEncoder{
		text: text,
	}

	client.PrintText(textToJsonEncoder)
}
