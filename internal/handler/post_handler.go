package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "vibegogo/internal/service"
    "vibegogo/pkg/logger"
)

type PostHandler struct {
    service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
    return &PostHandler{service: service}
}

type createPostRequest struct {
    Title    string `json:"title" binding:"required"`
    Content  string `json:"content" binding:"required"`
    ImageURL string `json:"image_url"`
}

func (h *PostHandler) Create(c *gin.Context) {
    var req createPostRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetString("user_id")
    post, err := h.service.Create(c.Request.Context(), userID, req.Title, req.Content, req.ImageURL)
    if err != nil {
        logger.Error("Failed to create post")
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, post)
}

func (h *PostHandler) List(c *gin.Context) {
    userID := c.GetString("user_id")
    posts, err := h.service.GetByUserID(c.Request.Context(), userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) GetByID(c *gin.Context) {
    postID := c.Param("id")
    post, err := h.service.GetByID(c.Request.Context(), postID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, post)
}

type updatePostRequest struct {
    Title    string `json:"title" binding:"required"`
    Content  string `json:"content" binding:"required"`
    ImageURL string `json:"image_url"`
}

func (h *PostHandler) Update(c *gin.Context) {
    var req updatePostRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    postID := c.Param("id")
    userID := c.GetString("user_id")
    post, err := h.service.Update(c.Request.Context(), postID, userID, req.Title, req.Content, req.ImageURL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, post)
}

func (h *PostHandler) Delete(c *gin.Context) {
    postID := c.Param("id")
    userID := c.GetString("user_id")
    
    if err := h.service.Delete(c.Request.Context(), postID, userID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}