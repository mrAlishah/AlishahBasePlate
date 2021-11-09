package services

import "context"

//01- implimet service method or impliment into other file
func (s service) Create(ctx context.Context, video Video) (Video, error) {
	//fmt.Printf("Hi %s", name)
	//msg := fmt.Sprintf("Hi %s", name)
	return video, nil

}
