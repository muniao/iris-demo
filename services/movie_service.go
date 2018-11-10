package services

import (
	"org.mm/iris-demo/dao"
	"org.mm/iris-demo/datasource"
	"org.mm/iris-demo/models"
)

type MovieService interface {
	GetAll() []models.Movie
	Search(name string) []models.Movie
	Get(id int) *models.Movie
	Delete(id int) error
	Update(user *models.Movie, columns []string) error
	Create(user *models.Movie) error
}

type movieService struct {
	dao *dao.MovieDao
}

func NewMovieService() MovieService {
	return &movieService{
		dao: dao.NewMovieDao(datasource.Instance()),
	}
}

func (s *movieService) GetAll() []models.Movie {
	return s.dao.GetAll()
}

func (s *movieService) Search(name string) []models.Movie {
	return s.dao.Search(name)
}

func (s *movieService) Get(id int) *models.Movie {
	return s.dao.Get(id)
}

func (s *movieService) Update(m *models.Movie, columns []string) error{
	return s.dao.Update(m, columns)
}

func (s *movieService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *movieService) Create(m *models.Movie) error {
	return s.dao.Create(m)
}
