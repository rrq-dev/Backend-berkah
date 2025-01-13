package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Location struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Address     string `gorm:"type:text;not null" json:"address"`
	Description string `gorm:"type:text" json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ActiveToken struct {
	ID        uint      `gorm:"primaryKey"`               // ID unik untuk setiap token
	UserID    uint      `gorm:"not null"`                // ID user yang memiliki token ini (relasi ke tabel User)
	Token     string    `gorm:"unique;not null"`         // Token JWT
	ExpiresAt time.Time `gorm:"not null"`                // Waktu kedaluwarsa token
	CreatedAt time.Time `gorm:"autoCreateTime"`          // Waktu token dibuat
	UpdatedAt time.Time `gorm:"autoUpdateTime"`          // Waktu token terakhir diperbarui
}

type BlacklistToken struct {
    ID        uint      `gorm:"primaryKey"`
    Token     string    `gorm:"unique;not null"`
    ExpiresAt time.Time `gorm:"not null"`
    CreatedAt time.Time
}
type Token struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"` // Relasi ke User
	User      User      `gorm:"foreignKey:UserID"`
	Token     string    `gorm:"unique;not null"`
	Role      string    `gorm:"not null"`
	CreatedAt time.Time
	ExpiresAt time.Time // Token Expiry Time
}
type UpdateInput struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

type Input struct {
	ID 			uint   `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"`
}
type LoginInput struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type RequestData struct {
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type ContextKey string

const (
	UserIDKey ContextKey = "userID"
	RoleKey   ContextKey = "role"
)
type Feedback struct {  
    ID       uint   `json:"id" gorm:"primaryKey"`  
    LocationID uint   `json:"location_id"` // ID of the mosque being rated  
    UserID   uint   `json:"user_id"`   // ID of the user providing the feedback  
    Rating   int    `json:"rating"`     // Rating value (e.g., 1 to 5)  
    Comment  string `json:"comment"`    // Optional comment
	CreatedAt  time.Time `json:"created_at"`  
}  