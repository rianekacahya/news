package repository

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/rianekacahya/news/internal/v1/news/entity"
	"github.com/rianekacahya/news/pkg/elasticsearch"
)

func (DI *News) List(limit int, offset int) (entity.NewsSlice, error) {
	var data entity.NewsSlice

	search := elasticsearch.GetConnection().Search().
		Index("news").
		SortBy(elastic.NewFieldSort("created").Desc().SortMode("max")).
		From(offset).Size(limit)

	ctx := context.Background()
	result, err := search.Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, hit := range result.Hits.Hits {
		var news entity.News
		json.Unmarshal(hit.Source, &news)
		data = append(data, news)
	}

	return data, nil
}

func (DI *News) InsertElasticsearch(news *entity.News) error {
	payload, _ := json.Marshal(news)
	ctx := context.Background()
	if _, err := elasticsearch.GetConnection().Index().Index("news").BodyJson(string(payload)).Do(ctx); err != nil {
		return err
	}

	return nil
}
