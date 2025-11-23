package dto

import (
	"time"

	"github.com/google/uuid"
)

type PostResponse struct {
	id        uuid.UUID
	title     string
	content   string
	authorId  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
