package base

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type DBTableEntity struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type DBTableUUIDEntity struct {
	ID        string         `json:"id" gorm:"primarykey;size:36;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (s *DBTableUUIDEntity) BeforeCreate(_ *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	return
}
