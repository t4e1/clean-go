package out

// packages for database implementation
import (
	"github.com/t4e1/clean-go/clean-b/post/internal/core/domain"
	"gorm.io/gorm"
)

type PostgresAdapter struct {
	database *gorm.DB
}

func InitPostgresAdapter(db *gorm.DB) *PostgresAdapter {
	return &PostgresAdapter{
		database: db,
	}
}

func (p *PostgresAdapter) SelectPost() (domain.Post, error) {

	return domain.Post{}, nil
}
