package repository

import (
	"github.com/rianekacahya/news/internal/v1/news/entity"
	"github.com/rianekacahya/news/pkg/mysql"
	"strings"
)

func (DI *News) Detail(newsID int) (*entity.News, error) {
	var (
		err    error
		query  strings.Builder
		bind   []interface{}
		result = new(entity.News)
	)

	query.WriteString(`SELECT author, body FROM news WHERE id = ?`)
	bind = append(bind, newsID)

	if err = mysql.GetConnection().QueryRow(query.String(), bind...).Scan(
		&result.Author, &result.Body,
	); err != nil {
		mysql.Debug(query.String(), bind, err)
		return result, mysql.Error(err)
	}

	// Print query
	mysql.Debug(query.String(), bind, nil)

	return result, nil
}

func (DI *News) InsertSQL(news *entity.News) (int, error) {
	var (
		newsID int
		err    error
		query  strings.Builder
		bind   []interface{}
	)

	query.WriteString(`INSERT INTO news (author, body, created) VALUES (?, ?, ?)`)
	bind = append(bind,
		news.Author,
		news.Body,
		news.Created,
	)

	// insert data news
	stmtNews, err := mysql.GetConnection().Prepare(query.String())
	if err != nil {
		mysql.Debug(query.String(), bind, err)
		return newsID, mysql.Error(err)
	}
	defer stmtNews.Close()

	// execute statement
	newsResult, err := stmtNews.Exec(bind...)
	if err != nil {
		mysql.Debug(query.String(), bind, err)
		return newsID, mysql.Error(err)
	}

	// get last insert ID
	insertID, err := newsResult.LastInsertId()
	if err != nil {
		mysql.Debug(query.String(), bind, err)
		return newsID, mysql.Error(err)
	}

	// Print query insert module
	mysql.Debug(query.String(), bind, nil)

	newsID = int(insertID)

	return newsID, nil
}
