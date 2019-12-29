package repository

import (
	"github.com/go-xorm/xorm"
	"xormdemo191226/model"
)

type MovieReporitory struct {
	engine *xorm.Engine
}

func NewMovieRepository(engine *xorm.Engine) *MovieReporitory {
	return &MovieReporitory{
		engine: engine,
	}
}

func (m *MovieReporitory) Get(id int) *model.Movie {
	data := &model.Movie{Id: id}
	ok, err := m.engine.Get(data)
	if ok && err != nil {
		return data
	} else {
		return nil
	}

	return data
}

func (m *MovieReporitory) GetAll() []model.Movie {
	datalist := []model.Movie{}
	err := m.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return nil
	}
	return datalist
}

func (m *MovieReporitory) Delete(id int) error {
	return nil
}
