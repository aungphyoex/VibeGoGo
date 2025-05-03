package repository

import (
    "context"
    "vibegogo/internal/domain"
)

type UserRepository interface {
    Create(ctx context.Context, user *domain.User) error
    GetByID(ctx context.Context, id string) (*domain.User, error)
    GetByEmail(ctx context.Context, email string) (*domain.User, error)
    Update(ctx context.Context, user *domain.User) error
    Delete(ctx context.Context, id string) error
}

type PostRepository interface {
    Create(ctx context.Context, post *domain.Post) error
    GetByID(ctx context.Context, id string) (*domain.Post, error)
    GetByUserID(ctx context.Context, userID string) ([]*domain.Post, error)
    Update(ctx context.Context, post *domain.Post) error
    Delete(ctx context.Context, id string) error
}