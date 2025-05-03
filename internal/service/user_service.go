package service

import (
    "context"
    "errors"
    "time"
    "github.com/google/uuid"
    "golang.org/x/crypto/bcrypt"
    "vibegogo/internal/domain"
    "vibegogo/internal/repository"
    "vibegogo/pkg/logger"
)

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, username, email, password string) (*domain.User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        logger.Error("Failed to hash password")
        return nil, err
    }

    user := &domain.User{
        ID:        uuid.New().String(),
        Username:  username,
        Email:     email,
        Password:  string(hashedPassword),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    if err := s.repo.Create(ctx, user); err != nil {
        logger.Error("Failed to create user")
        return nil, err
    }

    return user, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (*domain.User, error) {
    user, err := s.repo.GetByEmail(ctx, email)
    if err != nil {
        logger.Error("Failed to get user by email")
        return nil, err
    }

    if user == nil {
        return nil, errors.New("user not found")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, errors.New("invalid password")
    }

    return user, nil
}

func (s *UserService) GetProfile(ctx context.Context, userID string) (*domain.User, error) {
    return s.repo.GetByID(ctx, userID)
}

func (s *UserService) UpdateProfile(ctx context.Context, userID string, username string) (*domain.User, error) {
    user, err := s.repo.GetByID(ctx, userID)
    if err != nil {
        return nil, err
    }

    user.Username = username
    user.UpdatedAt = time.Now()

    if err := s.repo.Update(ctx, user); err != nil {
        return nil, err
    }

    return user, nil
}