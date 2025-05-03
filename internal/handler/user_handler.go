package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
    "time"
    "github.com/golang-jwt/jwt/v5"
    "vibegogo/internal/service"
    "vibegogo/pkg/logger"
)

type UserHandler struct {
    service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
    return &UserHandler{service: service}
}

type registerRequest struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

func (h *UserHandler) Register(c *gin.Context) {
    var req registerRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.service.Register(c.Request.Context(), req.Username, req.Email, req.Password)
    if err != nil {
        logger.Error("Failed to register user")
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, user)
}

type loginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

func (h *UserHandler) Login(c *gin.Context) {
    var req loginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.service.Login(c.Request.Context(), req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    // Generate JWT token here
    token := ""
    jwtSecret := []byte(os.Getenv("JWT_SECRET"))
    claims := jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }
    jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err = jwtToken.SignedString(jwtSecret)
    if err != nil {
        logger.Error("Failed to generate JWT token")
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": token,
        "user":  user,
    })
}

func (h *UserHandler) GetProfile(c *gin.Context) {
    userID := c.GetString("user_id")
    user, err := h.service.GetProfile(c.Request.Context(), userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

type updateProfileRequest struct {
    Username string `json:"username" binding:"required"`
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
    var req updateProfileRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetString("user_id")
    user, err := h.service.UpdateProfile(c.Request.Context(), userID, req.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}