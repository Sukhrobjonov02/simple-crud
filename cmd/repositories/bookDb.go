package repositories

import (
	"crud/cmd/models"
	"crud/cmd/storage"
	"fmt"
)

func CreateBook(book models.Book) (models.Book, error) {
	db := storage.GetDB()

	sqlStatement := `INSERT INTO book (title, author, published_at) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(sqlStatement, book.Title, book.Author, book.PublishedAt).
		Scan(&book.Id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func UpdateBook(book models.Book, id int) (models.Book, error) {
	db := storage.GetDB()

	sqlStatement := `
	  UPDATE book
	  SET title = $2, author = $3, published_at = $4
	  WHERE id = $1
	  RETURNING id`
	err := db.QueryRow(sqlStatement, id, book.Title, book.Author, book.PublishedAt).
		Scan(&id)
	if err != nil {
		return models.Book{}, err
	}
	book.Id = id
	return book, nil
}

func GetBook(book models.Book, id int) (models.Book, error) {
	db := storage.GetDB()

	sqlStatement := `
	  SELECT *
	  FROM book
	  WHERE id = $1`

	err := db.QueryRow(sqlStatement, id).
		Scan(&book.Id, &book.Title, &book.Author, &book.PublishedAt)
	if err != nil {
		return book, err
	}

	return book, nil
}

func GetBooks() ([]models.Book, error) {
	db := storage.GetDB()

	var books []models.Book

	sqlStatement := `
	  SELECT *
	  FROM book`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return books, err
	}

	defer rows.Close()

	for rows.Next() {
		var book models.Book

		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.PublishedAt)

		if err != nil {
			return books, err
		}

		books = append(books, book)
	}

	return books, nil
}

func DeleteBook(id int) (int, error) {

	db := storage.GetDB()

	sqlStatement := `DELETE FROM book WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return 0, err
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return id, nil
}
