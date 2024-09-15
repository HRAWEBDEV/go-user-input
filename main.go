package main

import (
	"fmt"

	"example.com/user-input/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func getNoteData() (string, string) {
	title := getUesrInput("note title:")
	content := getUesrInput("note content:")

	return title, content
}

func getUesrInput(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)

	return value
}
