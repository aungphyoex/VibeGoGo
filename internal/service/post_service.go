package service

import (
    "context"
    "errors"
    "github.com/google/uuid"
    "time"
    "vibegogo/internal/domain"
    "vibegogo/internal/repository"
    "vibegogo/pkg/logger"
)

type PostService struct {
    repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *PostService {
    return &PostService{repo: repo}
}

func (s *PostService) Create(ctx context.Context, userID, title, content, imageURL string) (*domain.Post, error) {
    post := &domain.Post{
        ID:        uuid.New().String(),
        Title:     title,
        Content:   content,
        ImageURL:  imageURL,
        UserID:    userID,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    if err := s.repo.Create(ctx, post); err != nil {
        logger.Error("Failed to create post")
        return nil, err
    }

    return post, nil
}

func (s *PostService) GetByID(ctx context.Context, id string) (*domain.Post, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *PostService) GetByUserID(ctx context.Context, userID string) ([]*domain.Post, error) {
    return s.repo.GetByUserID(ctx, userID)
}

func (s *PostService) Update(ctx context.Context, id, userID, title, content, imageURL string) (*domain.Post, error) {
    post, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }

    if post.UserID != userID {
        return nil, errors.New("unauthorized")
    }

    post.Title = title
    post.Content = content
    post.ImageURL = imageURL
    post.UpdatedAt = time.Now()

    if err := s.repo.Update(ctx, post); err != nil {
        return nil, err
    }

    return post, nil
}

func (s *PostService) Delete(ctx context.Context, id, userID string) error {
    post, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return err
    }

    if post.UserID != userID {
        return errors.New("unauthorized")
    }

    return s.repo.Delete(ctx, id)
}