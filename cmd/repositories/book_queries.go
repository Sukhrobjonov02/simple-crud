package repositories

const (
	createBook = `
	INSERT INTO book(title, author, published_at)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	updateBook = `
	UPDATE book
	SET title = $2, author = $3, published_at = $4
	WHERE id = $1
	RETURNING id
	`

	getBook = `
	SELECT *
	FROM book
	WHERE id = $1
	`

	getAllBooks = `
	SELECT *
	FROM book
	`

	deleteBook = `
	DELETE FROM book
	WHERE id=$1
	`
)
