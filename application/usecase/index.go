package usecase

import (
	"context"
	"html/template"

	"gorm.io/gorm"
)

//go:generate mockgen -source=./index.go -destination=./mock/index.go -package=mock
type IndexUsecase interface {
	Index(ctx context.Context) (*template.Template, error)
	Ready(ctx context.Context) error
}

type indexUsecase struct {
	db *gorm.DB
}

func NewIndexUsecase(db *gorm.DB) IndexUsecase {
	return &indexUsecase{db: db}
}

func (u indexUsecase) Index(ctx context.Context) (*template.Template, error) {
	tmpl, err := template.ParseFiles("web/index.html")
	if err != nil {
		return nil, err
	}

	return tmpl, err
}

func (u indexUsecase) Ready(ctx context.Context) error {
	var i int
	if err := u.db.Raw("SELECT 1").Scan(&i).Error; err != nil {
		return err
	}

	return nil
}
