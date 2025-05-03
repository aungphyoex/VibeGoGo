package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Welcome to VibeGoGo API",
        "spec": map[string]interface{}{
            "register": "POST /api/users/register",
            "login": "POST /api/users/login",
            "get_profile": "GET /api/users/me (auth required)",
            "update_profile": "PUT /api/users/me (auth required)",
            "create_post": "POST /api/posts (auth required)",
            "list_posts": "GET /api/posts (auth required)",
            "get_post_by_id": "GET /api/posts/:id (auth required)",
            "update_post": "PUT /api/posts/:id (auth required)",
            "delete_post": "DELETE /api/posts/:id (auth required)",
        },
        "docs": "Add more details or link to full documentation here.",
    })
}