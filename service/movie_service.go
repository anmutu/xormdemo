package service

import (
	"xormdemo191226/datasource"
	"xormdemo191226/model"
	"xormdemo191226/repository"
)

type MovieService interface {
	GetAll() []model.Movie
	//GetById(id int) model.Movie
	//DeleteById(id int) bool
	//Update(user *model.Movie) error
	//Create(user *model.Movie) error
}

//小写
type movieService struct {
	repository *repository.MovieReporitory
}

func NewMovieService() MovieService {
	return &movieService{
		repository: repository.NewMovieRepository(datasource.InstanceMaster()),
	}
}

func (m *movieService) GetAll() []model.Movie {
	return m.repository.GetAll()
}
