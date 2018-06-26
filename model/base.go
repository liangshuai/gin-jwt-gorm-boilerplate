package model

import (
	"time"
)

type Model struct {
	ID				uint16		`form:"id" gorm:"primary_key"`
	CreatedAt		time.Time	`binding:"-"`
	UpdatedAt		time.Time	`binding:"-"`
}