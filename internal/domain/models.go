package domain

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
    Username  string         `gorm:"type:varchar(255);not null" json:"username"`
    Email     string         `gorm:"type:varchar(255);unique;not null" json:"email"`
    Password  string         `gorm:"type:varchar(255);not null" json:"-"`
    CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Post struct {
    ID        string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
    Title     string         `gorm:"type:varchar(255);not null" json:"title"`
    Content   string         `gorm:"type:text;not null" json:"content"`
    ImageURL  string         `gorm:"type:varchar(255)" json:"image_url"`
    UserID    string         `gorm:"type:uuid;not null" json:"user_id"`
    User      User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
    CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}