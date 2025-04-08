package main

import (
	"fmt"
    "example.com/note/note"
	"bufio"
	"os"
	"strings"
)

func main() {

	title, content := getNoteDate()

	userNote, err := note.NewNote(title, content) // calling the constructor

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.OutputNoteData() // calling the method of the struct

	err = userNote.SaveToFile() // calling the method of the struct

	if err != nil {
		fmt.Println("saving to file error:", err)
		return
	}

}

func getNoteDate() (string, string) {

	title := getUserInput("Enter the title of the note")

	content := getUserInput("Enter the content of the note")

	return title, content

}

func getUserInput(prompt string) string {
	
	fmt.Println(prompt)
	
	reader := bufio.NewReader(os.Stdin) // Create a new reader
	text, err := reader.ReadString('\n') // Read user input until newline

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	text = strings.TrimSuffix(text,"\n") // Remove any leading/trailing whitespace
	text = strings.TrimSuffix(text,"\r") // Remove any leading/trailing whitespace


	return text
}
