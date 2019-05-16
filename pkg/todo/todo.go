package todo

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	ID      	string    `json:"id" db:"id" gorm:"column:id;primary_key;type:uuid"`
	UserID      string    `json:"user_id" db:"user_id" gorm:"column:user_id;primary_key"`
	Content     string    `json:"content" db:"content" gorm:"varchar(255);not null"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at" gorm:"column:updated_at"`
}

func (Todo) TableName() string {
	return "todos"
}

func (u *Todo) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid.String())
}