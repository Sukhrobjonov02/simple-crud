package repositories

import (
	"crud/cmd/storage"
)

func CreateBook(book Book) (Book, error) {
	db := storage.GetDB()

	err := db.QueryRow(createBook, book.Title, book.Author, book.PublishedAt).
		Scan(&book.Id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func UpdateBook(book Book, id int) (Book, error) {
	db := storage.GetDB()

	err := db.QueryRow(updateBook, id, book.Title, book.Author, book.PublishedAt).
		Scan(&book.Id)
	if err != nil {
		return Book{}, err
	}
	book.Id = id
	return book, nil
}

func GetBook(book Book) (Book, error) {
	db := storage.GetDB()

	err := db.QueryRow(getBook, book.Id).
		Scan(&book.Id, &book.Title, &book.Author, &book.PublishedAt)
	if err != nil {
		return book, err
	}

	return book, nil
}

func GetAllBooks() ([]Book, error) {
	db := storage.GetDB()

	var books []Book

	rows, err := db.Query(getAllBooks)

	if err != nil {
		return books, err
	}

	defer rows.Close()

	for rows.Next() {
		var book Book

		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.PublishedAt)

		if err != nil {
			return books, err
		}

		books = append(books, book)
	}

	return books, nil
}

func DeleteBook(book Book) (Book, error) {

	db := storage.GetDB()

	res, err := db.Exec(deleteBook, book.Id)

	if err != nil {
		return book, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil || rowsAffected == 0 {
		book.Id = 0
		return book, err
	}

	return book, nil
}
