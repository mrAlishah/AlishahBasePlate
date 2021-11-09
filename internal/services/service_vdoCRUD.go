package services

import (
	"context"
)

func (s service) Create(ctx context.Context, video Video) (Video, error) {
	//vdo, err := s.postgres.CreateVideo(Video{
	//	Title: "test@test.com",
	//	Desc:  "test",
	//	Url:   "test pour",
	//	//CreatedAt: time.Time{},
	//	//UpdatedAt: time.Time{},
	//	//DeletedAt: gorm.DeletedAt{},
	//})
	vdo, err := s.postgres.CreateVideo(video)
	if err != nil {
		return vdo, err
	}
	return vdo, nil
}
