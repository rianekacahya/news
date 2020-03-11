package news

import (
	"github.com/rianekacahya/news/internal/v1/news/entity"
	"github.com/rianekacahya/news/internal/v1/news/repository"
)

type news struct {
	repository Repository
}

func Initialize() *news {
	return &news{
		repository: repository.NewNews(),
	}
}

type Usecase interface {
	List(*entity.Request) (entity.NewsSlice, error)
	Create(*entity.News, string) (*entity.News, error)
	Insert(*entity.News) error
}

type Repository interface {
	List(int, int) (entity.NewsSlice, error)
	Detail(int) (*entity.News, error)
	InsertSQL(*entity.News) (int, error)
	InsertElasticsearch(*entity.News) error
}