package service

import "xormdemo191226/model"

type MovieService interface {
	GetAll() []model.Movie
	GetById(id int) model.Movie
	DeleteById(id int) bool
	Update(user *model.Movie) error
	Create(user *model.Movie) error
}
