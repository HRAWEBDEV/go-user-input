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
	title     string
	content   string
	createdAt time.Time
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
	fmt.Printf("your note titled %v has the following content:\n\n%v", n.title, n.content)
}

func (n Note) Save() error {
	filename := strings.ReplaceAll(n.title, " ", "_")
	filename = strings.ToLower(n.title)

	json, err := json.Marshal(n)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}
