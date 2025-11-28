package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Post represents a blog post entity
type Post struct {
	id        uuid.UUID
	title     string
	content   string
	authorId  string
	createdAt time.Time
	updatedAt time.Time
}

const idErrMsg string = "id cannot be empty"

func (p *Post) GetPost(id uuid.UUID) (*Post, error) {
	if len(id) == 0 {
		return nil, errors.New(idErrMsg)
	}

	return nil, nil
}

// NewPost creates a new post with validation
func (p *Post) NewPost(id uuid.UUID, title, content, authorID string) (*Post, error) {
	if len(id) == 0 {
		return nil, errors.New(idErrMsg)
	}
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if authorID == "" {
		return nil, errors.New("authorID cannot be empty")
	}

	now := time.Now()
	return &Post{
		id:        id,
		title:     title,
		content:   content,
		authorId:  authorID,
		createdAt: now,
		updatedAt: now,
	}, nil
}

// Update modifies the post content and updates timestamp
func (p *Post) Update(title, content string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}

	p.title = title
	p.content = content
	p.updatedAt = time.Now()
	return nil
}

func (p *Post) Delete(id uuid.UUID) (bool, error) {
	if len(id) == 0 {
		return false, errors.New(idErrMsg)
	}
	return false, nil
}
