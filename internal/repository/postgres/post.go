package postgres

import (
    "context"
    "vibegogo/internal/domain"
    "vibegogo/pkg/logger"
    "gorm.io/gorm"
    "go.uber.org/zap"
)

type PostRepo struct {
    db *gorm.DB
}

func NewPostRepo(db *gorm.DB) *PostRepo {
    return &PostRepo{db: db}
}

func (r *PostRepo) Create(ctx context.Context, post *domain.Post) error {
    logger.Info("Creating new post", zap.String("user_id", post.UserID))
    
    result := r.db.WithContext(ctx).Create(post)
    if result.Error != nil {
        logger.Error("Failed to create post", zap.Error(result.Error))
        return result.Error
    }
    
    return nil
}

func (r *PostRepo) GetByID(ctx context.Context, id string) (*domain.Post, error) {
    var post domain.Post
    
    result := r.db.WithContext(ctx).First(&post, "id = ?", id)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, nil
        }
        logger.Error("Failed to get post by ID", zap.Error(result.Error))
        return nil, result.Error
    }
    
    return &post, nil
}

func (r *PostRepo) GetByUserID(ctx context.Context, userID string) ([]*domain.Post, error) {
    var posts []*domain.Post
    
    result := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&posts)
    if result.Error != nil {
        logger.Error("Failed to get posts by user ID", zap.Error(result.Error))
        return nil, result.Error
    }
    
    return posts, nil
}

func (r *PostRepo) Update(ctx context.Context, post *domain.Post) error {
    result := r.db.WithContext(ctx).Save(post)
    if result.Error != nil {
        logger.Error("Failed to update post", zap.Error(result.Error))
        return result.Error
    }
    
    return nil
}

func (r *PostRepo) Delete(ctx context.Context, id string) error {
    result := r.db.WithContext(ctx).Delete(&domain.Post{}, "id = ?", id)
    if result.Error != nil {
        logger.Error("Failed to delete post", zap.Error(result.Error))
        return result.Error
    }
    
    return nil
}