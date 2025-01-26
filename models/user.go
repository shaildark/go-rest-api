package models

import (
	"time"

	"example.com/go-api/db"
	"gorm.io/gorm"
)

type User struct {
	ID            int            `gorm:"primaryKey;autoIncrement" json:"id"` // Primary key with auto-increment
	IRoleId       int            `gorm:"column:iRoleId;default:1" json:"iRoleId"`
	VUserName     string         `gorm:"column:vUserName;size:150" json:"vUserName"`
	VFirstName    string         `gorm:"column:vFirstName;size:150" json:"vFirstName"`
	VLastName     string         `gorm:"column:vLastName;size:150" json:"vLastName"`
	Email         string         `gorm:"column:email;size:255" json:"email"`
	Password      string         `gorm:"column:password;size:255" json:"password"`
	VProfileImage string         `gorm:"column:vProfileImage;size:255" json:"vProfileImage"`
	CreatedAt     time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"` // Automatically set on creation
	UpdatedAt     time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // Automatically updated on update
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`          // Soft delete field
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}

func ExcludeDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("deleted_at IS NULL")
}

func (User) DefaultScopes() []func(*gorm.DB) *gorm.DB {
	return []func(*gorm.DB) *gorm.DB{
		ExcludeDeleted,
	}
}

func (u *User) Create() error {

	db, dbError := db.Connection()

	if dbError != nil {
		return dbError
	}

	result := db.Create(&u)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
