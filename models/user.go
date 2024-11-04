package models

import "time" // Import the time package

// Users represents the user model with necessary fields
type Users struct {
    UserID       uint   `gorm:"primaryKey;column:user_id"` // Primary key for the user
    Name     string `json:"username"`                  // Username of the user
    Password string `json:"password"`                  // Password of the user (consider hashing this)
}

// TableName specifies the exact table name for the Users model
func (Users) TableName() string {
    return "users" // Specify the exact table name
}

// UserLogin represents the user login model with necessary fields
type UserLogin struct {
    ID           uint      `json:"login_id" gorm:"primaryKey;auto_increment"` // Primary key for the user login
    UserID       uint   `gorm:"primaryKey;column:user_id"`                    // User ID; foreign key
    ModifiedTime time.Time `json:"modified_time"`              // Time of last modification
}

// TableName specifies the exact table name for the UserLogin model
func (UserLogin) TableName() string {
    return "user_login" // Specify the exact table name
}
