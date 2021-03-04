package comment

import "github.com/jinzhu/gorm"

// Service - the struct for our ocmment service
type Service struct {
	DB *gorm.DB
}

// Comment - defines our comment structure
type Comment struct {
	gorm.Model
	// Slug - path to where the comment is posted
	Slug   string
	Body   string
	Author string
}

// CommentService - the interface for our comment service
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

// NewService - returns a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetComment - retrieves comments by their ID from the database
func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// GetCommentsBySlug - retrieves all comments by slug (path - /article/name/)
func (s *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

// PostComment - Saves the
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}
