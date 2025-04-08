package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string `json:"title"`
	Content   string  `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNote(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) OutputNoteData() {

	fmt.Println("Title:", note.Title)
	fmt.Println("Content:", note.Content)
}

func (note Note) SaveToFile() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// os.WriteFile(fileName, []byte(note.content), 0644)  writing to a file

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644) 

}
