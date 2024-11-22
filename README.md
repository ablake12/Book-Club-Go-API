# Book-Club-Go-API
Hello. This is a REST API that manages a virtual Book Club using SQLite for storing and managing the the data. Users can interact with this API through a User Interface or through postman.
This API has the functionality to:
- Add and delete books
- Add, update, and delete reviews to the books
- Delete all the books in the book club
- Update whether the book club has read a book yet
- Update what current book the book club is on
## Prerequisites
Go 1.23+
## Usage
1. Open terminal and run the program in the current working directory
2. Initialize mod file
```bash
go mod init Book-Club-Go-API
```
3. Install sqlite package
```bash
go get github.com/mattn/go-sqlite3
```
4. Install gin package
```bash
go get github.com/gin-gonic/gin
```
5. The default port is localhost:8080. If you prefer another port, set the preferred port to the environment variable `PORT`
```bash
export PORT = <last four digits of port>
```
6. Run the program
```bash
go run main.go
```

## API Requests
The following routes for this API are:
- ```GET``` <port>/books - Fetches all the books in the book club
- ```GET``` <port>/books/read - Fetches all the books that have been read by the book club
- ```GET``` <port>/books/unread - Fetches all the books that have not been read by the book club yet
- ```GET``` <port>/books/current - Fetches the current book that the book club is reading
- ```GET``` <port>/books/`<book_id>`/reviews- Fetches the reviews for a given book
- ```POST``` <port>/books - Adds a book to the book club
- ```POST``` <port>/books/`<book_id>`/reviews - Adds review for a given book
- ```PUT``` <port>/books/`<book_id>`/status - Updates whether the book club has read a book or not
- ```PUT``` <port>/books/`<book_id>`/current - Updates the current book the book club is on
- ```PUT``` <port>/books/`<book_id>`/reviews/`<review_id>` - Updates a review for a given book
- ```DELETE``` <port>/books/`<book_id>` - Deletes a book from the book club
- ```DELETE``` <port>/books/`<book_id>`/reviews/`<review_id>` - Delete a review for a given book
- ```DELETE``` <port>/books - Deletes all the books in the book club

## API Interfaces
Your options to use the API are:
- Through the interface on the home page at route `<port>/`
- Through Postman

## Postman Instructions
To add a book in Postman, use the following json in the body
```json
{
    "title": "book_title_placeholder",
    "author": "book_author_placeholder",
    "genre": "book_genre_placeholder",
    "description": "book_desc_placeholder",
    "read_status": "book_read_status_placeholder",
    "current_status": "is_current_book_placeholder"
}
```
Note: Only `title`, `author`, `genre` and `description` are required fields. `read_status` and `current_status` are optional fields and will default to N if not included in the json.

### Adding a review in Postman:
To add a review in postman, use the following json in the body:
```json
{
    "user": "user_name_placeholder",
    "review": "book_review_placeholder",
    "rating": "book_rating_placeholder"
}
```

### Updating a review in Postman:
To update a review in Postman, use the following json in the body:
```json
{
    "review": "book_review_placeholder",
    "rating": "book_rating_placeholder"
}
```
