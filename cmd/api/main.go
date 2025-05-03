package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
    "go.uber.org/zap"
    "log"
    "os"
    "vibegogo/internal/config"
    "vibegogo/internal/domain"
    "vibegogo/internal/handler"
    "vibegogo/internal/middleware"
    "vibegogo/internal/repository/postgres"
    "vibegogo/internal/service"
    "vibegogo/pkg/logger"
)

func main() {
    // Hot reload is enabled - any changes will trigger automatic rebuild
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    logger.Init()

    router := gin.Default()
    router.Use(middleware.LoggerMiddleware())

    // Initialize database connection
    db, err := config.NewDBConnection()
    if err != nil {
        logger.Error("Failed to connect to database", zap.Error(err))
        log.Fatal(err)
    }

    // Auto migrate the schema
    if err := db.AutoMigrate(&domain.User{}, &domain.Post{}); err != nil {
        logger.Error("Failed to migrate database schema", zap.Error(err))
        log.Fatal(err)
    }

    // Repositories
    userRepo := postgres.NewUserRepo(db)
    postRepo := postgres.NewPostRepo(db)

    // Services
    userService := service.NewUserService(userRepo)
    postService := service.NewPostService(postRepo)

    // Handlers
    userHandler := handler.NewUserHandler(userService)
    postHandler := handler.NewPostHandler(postService)

    // Routes
    api := router.Group("/api")
    {
        // Public routes
        api.POST("/auth/register", userHandler.Register)
        api.POST("/auth/login", userHandler.Login)

        // Protected routes
        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware(os.Getenv("JWT_SECRET")))
        {
            // User routes
            protected.GET("/users/me", userHandler.GetProfile)
            protected.PUT("/users/me", userHandler.UpdateProfile)

            // Post routes
            protected.POST("/posts", postHandler.Create)
            protected.GET("/posts", postHandler.List)
            protected.GET("/posts/:id", postHandler.GetByID)
            protected.PUT("/posts/:id", postHandler.Update)
            protected.DELETE("/posts/:id", postHandler.Delete)
        }
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    logger.Info("Server starting on port " + port)
    if err := router.Run(":" + port); err != nil {
        logger.Error("Server failed to start")
        log.Fatal(err)
    }
}