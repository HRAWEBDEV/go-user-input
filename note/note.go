package note

import (
	"errors"
	"fmt"
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
func (n *Note) Display() {
	fmt.Printf("your note titled %v has the following content:\n\n%v", n.title, n.content)
}
