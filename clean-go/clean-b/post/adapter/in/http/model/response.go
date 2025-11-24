package dto

import (
	"time"

	"github.com/google/uuid"
)

type PostResponse struct {
	Id        uuid.UUID
	Title     string
	Content   string
	AuthorId  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
