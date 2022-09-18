package usecase

import "github.com/duykb2015/ecom-api/internal/entity"

type UserUsecase struct {
	r UserRepo
}

func New(r UserRepo) *UserUsecase {
	return &UserUsecase{r}
}

func (r *UserUsecase) AuthLogin() (entity.UserResponse, error) {
	return entity.UserResponse{}, nil
}
