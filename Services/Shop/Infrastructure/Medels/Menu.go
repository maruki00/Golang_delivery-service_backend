package shop_infrastructure_medels

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Id        int       `json: "id"`
	LAbel     string    `json: "label"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}
