package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// using struct tags

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}
	return Note{
		title,
		content,
		time.Now(),
	}, nil
}

// struct methods
func (n Note) Display() {
	fmt.Printf("your note titled %v has the following content:\n\n%v", n.Title, n.Content)
}

func (n Note) Save() error {
	filename := strings.ReplaceAll(n.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}
