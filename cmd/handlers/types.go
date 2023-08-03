package handlers

import (
	"time"
)

type CreateBookRequest struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
}

type CreateBookResponse struct {
	Id int `json:"id"`
}

type UpdateBookRequest struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
}

type UpdateBookResponse struct {
	Id int `json:"id"`
}

type GetBookRequest struct {
	Id int `json:"id"`
}

type GetBookResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
}

type GetAllBooksResponse struct {
	List []BookResponse `json:"list"`
}

type BookResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
}

type DeleteBookRequest struct {
	Id int `json:"id"`
}

type DeleteBookResponse struct {
	Id int `json:"id"`
}
