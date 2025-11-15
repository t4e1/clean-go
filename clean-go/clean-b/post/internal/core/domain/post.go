package domain

import (
	"errors"
	"time"
)

// Post represents a blog post entity
type Post struct {
	ID        string
	Title     string
	Content   string
	AuthorID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewPost creates a new post with validation
func NewPost(id, title, content, authorID string) (*Post, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if authorID == "" {
		return nil, errors.New("authorID cannot be empty")
	}

	now := time.Now()
	return &Post{
		ID:        id,
		Title:     title,
		Content:   content,
		AuthorID:  authorID,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// Update modifies the post content and updates timestamp
func (p *Post) Update(title, content string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}

	p.Title = title
	p.Content = content
	p.UpdatedAt = time.Now()
	return nil
}
