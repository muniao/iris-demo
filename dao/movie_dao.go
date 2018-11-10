package dao

import (
	"github.com/go-xorm/xorm"
	"org.mm/iris-demo/models"
)

type MovieDao struct {
	engine *xorm.Engine
}

func NewMovieDao(engine *xorm.Engine) *MovieDao {
	return &MovieDao{
		engine: engine,
	}
}

/**
 *
 * 通过id获取电影信息
 **/
func (d *MovieDao) Get(id int) *models.Movie {
	data := &models.Movie{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (d *MovieDao) GetAll() []models.Movie{
	datalist := make([]models.Movie, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MovieDao) Search(name string) []models.Movie {
	datalist := make([]models.Movie, 0)
	err := d.engine.Where("name like ?", name).
		Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

func (d *MovieDao) Delete(id int) error {
	data := &models.Movie{Id: id}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *MovieDao) Update(data *models.Movie, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *MovieDao) Create(data *models.Movie) error {
	_, err := d.engine.Insert(data)
	return err
}
