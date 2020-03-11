package entity

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"
)

const TopicInsertNews  = "insert-news"

type (
	News struct {
		ID      int       `json:"id,omitempty"`
		Author  string    `json:"author,omitempty"`
		Body    string    `json:"body,omitempty"`
		Created time.Time `json:"created,omitempty"`
	}

	NewsSlice []News

	Request struct {
		Page  int `query:"page" json:"page"`
		Limit int `query:"limit" json:"limit"`
	}
)

// Count pagination offset
func (DI *Request) Offset() int {
	return (DI.Page - 1) * DI.Limit
}

// Validate create news
func (DI *News) CommandNewsValidation() error {
	return validation.ValidateStruct(DI,
		validation.Field(&DI.Author, validation.Required),
		validation.Field(&DI.Body, validation.Required),
	)
}

// Validate list news
func (DI *Request) QueryNewsValidation() error {
	return validation.ValidateStruct(DI,
		validation.Field(&DI.Page, validation.Required),
		validation.Field(&DI.Limit, validation.Required),
	)
}
