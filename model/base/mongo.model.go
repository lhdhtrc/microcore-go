package base

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoTableModel struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at" bson:"deleted_at"`
}

func (m *MongoTableModel) DefaultId() {
	if m.ID.IsZero() {
		m.ID = primitive.NewObjectID()
	}
}

func (m *MongoTableModel) DefaultCreatedAt() {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now().Local()
	}
}

func (m *MongoTableModel) DefaultUpdatedAt() {
	m.UpdatedAt = time.Now().Local()
}
