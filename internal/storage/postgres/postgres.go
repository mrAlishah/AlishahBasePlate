package postgres

import (
	"AlishahBasePlate/internal/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	database *gorm.DB
}

var models = []interface{}{
	&services.Video{},
}

func CreateRepository(db *gorm.DB) (*Repository, error) {
	repo := &Repository{
		database: db,
	}
	logrus.Infof("current db name: %s", db.Migrator().CurrentDatabase())
	err := db.AutoMigrate(models...)
	if err != nil {
		return repo, errors.Wrap(err, "failed to auto migrate models")
	}
	return repo, nil
}

//02-
func (r *Repository) CreateVideo(video services.Video) (services.Video, error) {
	vdo := services.Video{
		Title: video.Title,
		Desc:  video.Desc,
		Path:  video.Path,
	}
	err := r.database.Create(&vdo).Error
	if err != nil {
		return vdo, errors.Wrap(err, "failed to create a video")
	}
	return vdo, nil
}
