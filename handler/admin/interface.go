package admin

import (
	"context"

	"gorm.io/gorm"
)

type Handler interface {
	InitHandler(ctx context.Context, db *gorm.DB) error
	ResetHandler(ctx context.Context, db *gorm.DB) error
}

type DefaultHandler struct{}

var Default Handler = &DefaultHandler{}
