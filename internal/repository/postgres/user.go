package postgres

import (
    "context"
    "vibegogo/internal/domain"
    "vibegogo/pkg/logger"
    "gorm.io/gorm"
    "go.uber.org/zap"
)

type UserRepo struct {
    db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
    return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *domain.User) error {
    logger.Info("Creating new user", zap.String("email", user.Email))
    
    result := r.db.WithContext(ctx).Create(user)
    if result.Error != nil {
        logger.Error("Failed to create user", zap.Error(result.Error))
        return result.Error
    }
    
    return nil
}

func (r *UserRepo) GetByID(ctx context.Context, id string) (*domain.User, error) {
    var user domain.User
    
    result := r.db.WithContext(ctx).First(&user, "id = ?", id)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, nil
        }
        logger.Error("Failed to get user by ID", zap.Error(result.Error))
        return nil, result.Error
    }
    
    return &user, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
    var user domain.User
    
    result := r.db.WithContext(ctx).First(&user, "email = ?", email)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, nil
        }
        logger.Error("Failed to get user by email", zap.Error(result.Error))
        return nil, result.Error
    }
    
    return &user, nil
}

func (r *UserRepo) Update(ctx context.Context, user *domain.User) error {
    result := r.db.WithContext(ctx).Save(user)
    if result.Error != nil {
        logger.Error("Failed to update user", zap.Error(result.Error))
        return result.Error
    }
    
    return nil
}

func (r *UserRepo) Delete(ctx context.Context, id string) error {
    result := r.db.WithContext(ctx).Delete(&domain.User{}, "id = ?", id)
    if result.Error != nil {
        logger.Error("Failed to delete user", zap.Error(result.Error))
        return result.Error
    }
    
    return nil
}