package handlers

import (
	"crud/cmd/repositories"
)

func fromCreateBookRequest(book CreateBookRequest) repositories.Book {
	return repositories.Book{
		Title:       book.Title,
		Author:      book.Author,
		PublishedAt: book.PublishedAt,
	}
}

func toCreateBookResponse(book repositories.Book) CreateBookResponse {
	return CreateBookResponse{
		Id: book.Id,
	}
}

func fromUpdateBookRequest(book UpdateBookRequest) repositories.Book {
	return repositories.Book{
		Title:       book.Title,
		Author:      book.Author,
		PublishedAt: book.PublishedAt,
	}
}

func toUpdateBookResponse(book repositories.Book) UpdateBookResponse {
	return UpdateBookResponse{
		Id: book.Id,
	}
}

func fromGetBookRequest(book GetBookRequest) repositories.Book {
	return repositories.Book{
		Id: book.Id,
	}
}

func toGetBookResponse(book repositories.Book) GetBookResponse {
	return GetBookResponse{
		Id:          book.Id,
		Title:       book.Title,
		Author:      book.Author,
		PublishedAt: book.PublishedAt,
	}
}

func toGetAllBooksResponse(books []repositories.Book) GetAllBooksResponse {
	var resBooks []BookResponse

	for i := 0; i < len(books); i++ {
		curBook := books[i]

		book := BookResponse{
			Id:          curBook.Id,
			Title:       curBook.Title,
			Author:      curBook.Author,
			PublishedAt: curBook.PublishedAt,
		}

		resBooks = append(resBooks, book)
	}

	return GetAllBooksResponse{
		List: resBooks,
	}

}

func fromDeleteBookRequest(book DeleteBookRequest) repositories.Book {
	return repositories.Book{
		Id: book.Id,
	}
}

func toDeleteBookResponse(book repositories.Book) DeleteBookResponse {
	return DeleteBookResponse{
		Id: book.Id,
	}
}
