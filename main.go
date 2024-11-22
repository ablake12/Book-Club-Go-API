package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func initialize_db() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	book_table := `
    CREATE TABLE IF NOT EXISTS books (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        author TEXT NOT NULL,
        genre TEXT,
        description TEXT,
		is_read TEXT NOT NULL DEFAULT 'N',
        current_book TEXT NOT NULL DEFAULT 'N',
        rating TEXT
    );`
	_, err = db.Exec(book_table)
	if err != nil {
		log.Fatal(err)
	}

	review_table := `
    CREATE TABLE IF NOT EXISTS reviews (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        book_id INTEGER,
        user TEXT,
        review TEXT,
        rating TEXT,
        FOREIGN KEY (book_id) REFERENCES books(id)
    );`
	_, err = db.Exec(review_table)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database, book table, and review table created successfully!")
}

type book struct {
	ID          int
	Title       string `json:"title"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	IsRead      string `json:"read_status"`
	CurrentBook string `json:"current_status"`
	Rating      string `json:"rating"`
	Review      []review
}

type review struct {
	ID     int
	BookID int
	User   string `json:"user"`
	Review string `json:"review"`
	Rating string `json:"rating"`
}

func getHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Book Club API",
	})
}

func getBooks(c *gin.Context) {
	// Open sqlite db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Get books
	book_rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer book_rows.Close()

	var books []book = []book{}
	// Iterate through books to add to the books slice
	for book_rows.Next() {
		var book book
		err := book_rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Description, &book.IsRead, &book.CurrentBook, &book.Rating)
		if err != nil {
			log.Fatal(err)
		}
		// Get reviews for current book
		review_rows, err := db.Query("SELECT * FROM reviews WHERE book_id = ?", book.ID)
		if err != nil {
			log.Fatal(err)
		}
		defer review_rows.Close()
		var reviews []review = []review{}
		// Add any reviews for the book to its review attribute in book struct
		for review_rows.Next() {
			var review review
			err := review_rows.Scan(&review.ID, &review.BookID, &review.User, &review.Review, &review.Rating)
			if err != nil {
				log.Fatal(err)
			}
			reviews = append(reviews, review)
		}
		book.Review = reviews

		books = append(books, book)
	}

	// Check for any errors from any of the rows
	if err = book_rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Return 200 message
	c.IndentedJSON(http.StatusOK, books)
}

func getReadBooks(c *gin.Context) {
	// Open sqlite db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Get read books
	rows, err := db.Query("SELECT id, title, author, genre, description, is_read, current_book, rating FROM books WHERE is_read = 'Y'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []book = []book{}
	// Iterate through books to add to the books slice
	for rows.Next() {
		var book book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Description, &book.IsRead, &book.CurrentBook, &book.Rating)
		if err != nil {
			log.Fatal(err)
		}
		// Get reviews for current book
		review_rows, err := db.Query("SELECT * FROM reviews WHERE book_id = ?", book.ID)
		if err != nil {
			log.Fatal(err)
		}
		defer review_rows.Close()
		var reviews []review = []review{}
		// Add any reviews for the book to its review attribute in book struct
		for review_rows.Next() {
			var review review
			err := review_rows.Scan(&review.ID, &review.BookID, &review.User, &review.Review, &review.Rating)
			if err != nil {
				log.Fatal(err)
			}
			reviews = append(reviews, review)
		}
		book.Review = reviews

		books = append(books, book)
	}

	// Check for errors encountered during iteration
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	// Return 200 message
	c.IndentedJSON(http.StatusOK, books)
}

func getUnreadBooks(c *gin.Context) {
	// Open sqlite db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Get unread book
	rows, err := db.Query("SELECT id, title, author, genre, description, is_read, current_book, rating FROM books WHERE is_read = 'N'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []book = []book{}
	// Iterate through books to add to the books slice
	for rows.Next() {
		var book book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Description, &book.IsRead, &book.CurrentBook, &book.Rating)
		if err != nil {
			log.Fatal(err)
		}
		// Get reviews for current book
		review_rows, err := db.Query("SELECT * FROM reviews WHERE book_id = ?", book.ID)
		if err != nil {
			log.Fatal(err)
		}
		defer review_rows.Close()

		var reviews []review = []review{}
		// Add any reviews for the book to its review attribute in book struct
		for review_rows.Next() {
			var review review
			err := review_rows.Scan(&review.ID, &review.BookID, &review.User, &review.Review, &review.Rating)
			if err != nil {
				log.Fatal(err)
			}
			reviews = append(reviews, review)
		}
		book.Review = reviews

		books = append(books, book)
	}

	// Check for errors encountered during iteration
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	// Return 200 message
	c.IndentedJSON(http.StatusOK, books)
}

func getCurrentBook(c *gin.Context) {
	// Open sqlite db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Get current book
	rows, err := db.Query("SELECT id, title, author, genre, description, is_read, current_book, rating FROM books WHERE current_book = 'Y'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []book = []book{}
	// Adding current book and its attributes to slice
	for rows.Next() {
		var book book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Description, &book.IsRead, &book.CurrentBook, &book.Rating)
		if err != nil {
			log.Fatal(err)
		}
		// Get reviews for current book
		review_rows, err := db.Query("SELECT * FROM reviews WHERE book_id = ?", book.ID)
		if err != nil {
			log.Fatal(err)
		}
		defer review_rows.Close()
		var reviews []review = []review{}
		// Adding any reviews for the book to a slice
		for review_rows.Next() {
			var review review
			err := review_rows.Scan(&review.ID, &review.BookID, &review.User, &review.Review, &review.Rating)
			if err != nil {
				log.Fatal(err)
			}
			reviews = append(reviews, review)
		}
		// Add reviews slice to review in book struct
		book.Review = reviews

		books = append(books, book)
	}

	// Check for errors encountered during iteration
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	// Return 200 message
	c.IndentedJSON(http.StatusOK, books)
}

func getReviews(c *gin.Context) {
	// Open sqlite db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bookID := c.Param("book_id") // get book_id parameter from url

	// Call query to make sure book exist
	var bookTitle, bookAuthor string
	err = db.QueryRow("SELECT title, author FROM books WHERE id = ?", bookID).Scan(&bookTitle, &bookAuthor)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Get reviews for the book
	review_query := "SELECT * FROM reviews WHERE book_id = ?"
	review_rows, err := db.Query(review_query, bookID)
	if err != nil {
		log.Fatal(err)
	}
	defer review_rows.Close()
	var reviews []review = []review{}
	// Adding reviews from database to slice
	for review_rows.Next() {
		var review review
		err := review_rows.Scan(&review.ID, &review.BookID, &review.User, &review.Review, &review.Rating)
		if err != nil {
			log.Fatal(err)
		}
		reviews = append(reviews, review)
	}

	// Check for any errors from any of the rows
	if err = review_rows.Err(); err != nil {
		log.Fatal(err)
	}
	if len(reviews) > 0 {
		c.IndentedJSON(http.StatusOK, reviews)
	} else {
		c.IndentedJSON(http.StatusOK, []review{})
	}
}

func postBook(c *gin.Context) {
	var newBook book
	// Retrieve body request for review
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Open sqlite db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check whether the user added current and read statuses
	if newBook.IsRead == "" {
		if newBook.CurrentBook == "" {
			// If either values are not in the request then set to default of No
			book_query := `INSERT INTO books (title, author, genre, description, is_read, current_book, rating) VALUES (?, ?, ?, ?, 'N', 'N', 'N/A')`
			_, err = db.Exec(book_query, newBook.Title, newBook.Author, newBook.Genre, newBook.Description)
		} else {
			newBook.CurrentBook = strings.ToUpper(newBook.CurrentBook)
			if !(newBook.CurrentBook == "Y" || newBook.CurrentBook == "N") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "current_status is not 'Y' or 'N'"})
				return
			}
			// Set the current book to no if the added book is the new current book that the club is reading
			if newBook.CurrentBook == "Y" {
				_, err = db.Exec("UPDATE books SET current_book = 'N' WHERE current_book = 'Y'")
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update current book"})
					return
				}
			}
			// Read status is defaulted to no
			book_query := `INSERT INTO books (title, author, genre, description, is_read, current_book, rating) VALUES (?, ?, ?, ?, 'N', ?, 'N/A')`
			_, err = db.Exec(book_query, newBook.Title, newBook.Author, newBook.Genre, newBook.Description, newBook.CurrentBook)
		}
	} else {
		newBook.IsRead = strings.ToUpper(newBook.IsRead)
		if !(newBook.IsRead == "Y" || newBook.IsRead == "N") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "read_status is not 'Y' or 'N'"})
			return
		}
		if newBook.CurrentBook == "" {
			// Current book is defaulted to no
			book_query := `INSERT INTO books (title, author, genre, description, is_read, current_book, rating) VALUES (?, ?, ?, ?, ?, 'N', 'N/A')`
			_, err = db.Exec(book_query, newBook.Title, newBook.Author, newBook.Genre, newBook.Description, newBook.IsRead)
		} else {
			newBook.CurrentBook = strings.ToUpper(newBook.CurrentBook)
			if !(newBook.CurrentBook == "Y" || newBook.CurrentBook == "N") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "current_status is not 'Y' or 'N'"})
				return
			}
			// Set the current book to no if the added book is the new current book that the club is reading
			if newBook.CurrentBook == "Y" {
				_, err = db.Exec("UPDATE books SET current_book = 'N' WHERE current_book = 'Y'")
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update current book"})
					return
				}
			}
			// Neither values are defaulted
			book_query := `INSERT INTO books (title, author, genre, description, is_read, current_book, rating) VALUES (?, ?, ?, ?, ?, ?, 'N/A')`
			_, err = db.Exec(book_query, newBook.Title, newBook.Author, newBook.Genre, newBook.Description, newBook.IsRead, newBook.CurrentBook)
		}
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book to the Book Club"})
		return
	}

	// Return 200 message
	c.IndentedJSON(http.StatusCreated, gin.H{"Message": newBook.Title + " by " + newBook.Author + " added to the Book Club"})
}

func postReview(c *gin.Context) {
	var newReview review

	// Retrieve body request for review
	if err := c.BindJSON(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Open sqlite db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// get book_id parameter from url
	bookID := c.Param("book_id")

	// Check if book exist and get book information
	var bookTitle, bookAuthor string
	err = db.QueryRow("SELECT title, author FROM books WHERE id = ?", bookID).Scan(&bookTitle, &bookAuthor)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Add new review
	review_query := `INSERT INTO reviews (book_id, user, review, rating) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(review_query, bookID, newReview.User, newReview.Review, newReview.Rating)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add review to the Book Club"})
		return
	}

	// Update book's average rating after adding a new book
	ratings, err := db.Query("SELECT rating FROM reviews WHERE book_id = ?", bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}
	defer ratings.Close()

	// Recalculate book's average rating once review is added
	var totalRating float64
	var count int
	for ratings.Next() {
		var rating float64
		if err := ratings.Scan(&rating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
			return
		}
		totalRating += rating
		count++
	}

	avgRating := totalRating / float64(count)
	avgRatingStr := fmt.Sprintf("%.1f", avgRating)

	// Update book's average rating after updating a book's rating
	_, err = db.Exec("UPDATE books SET rating = ? WHERE id = ?", avgRatingStr, bookID)
	if err != nil {
		log.Println("Error updating book rating:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	// Return 200 message
	c.IndentedJSON(http.StatusCreated, gin.H{"Message": "Review for " + bookTitle + " by " + bookAuthor + " added to the Book Club"})
}

func putReadStatus(c *gin.Context) {
	// Open sqlite db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bookID := c.Param("book_id") // get book_id parameter from url

	// Check is book exist and get book info
	var bookTitle, bookAuthor, oldReadStatus string
	err = db.QueryRow("SELECT title, author, is_read FROM books WHERE id = ?", bookID).Scan(&bookTitle, &bookAuthor, &oldReadStatus)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Update to the opposite of the existing read status
	if oldReadStatus == "Y" {
		_, err = db.Exec("UPDATE books SET is_read = 'N' WHERE id = ?", bookID)
	} else {
		_, err = db.Exec("UPDATE books SET is_read = 'Y' WHERE id = ?", bookID)
	}

	if err != nil {
		log.Println("Error updating book status:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	// Return 200 message
	if oldReadStatus == "Y" {
		c.IndentedJSON(http.StatusOK, gin.H{"Message": "Status for " + bookTitle + " by " + bookAuthor + " updated to No"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"Message": "Status for " + bookTitle + " by " + bookAuthor + " updated to Yes"})
	}
}

func putCurrStatus(c *gin.Context) {
	// Open sqlite db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// get book_id parameter from url
	bookID := c.Param("book_id")

	// Check if book exist and get book information
	var bookTitle, bookAuthor string
	err = db.QueryRow("SELECT title, author FROM books WHERE id = ?", bookID).Scan(&bookTitle, &bookAuthor)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Update the current book
	_, err = db.Exec("UPDATE books SET current_book = 'N' WHERE current_book = 'Y'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update old current book"})
		return
	}

	// Set book to current book status
	_, err = db.Exec("UPDATE books SET current_book = 'Y' WHERE id = ?", bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update new current book"})
		return
	}

	// Return 200 message
	c.IndentedJSON(http.StatusOK, gin.H{"Message": bookTitle + " by " + bookAuthor + " updated to current book"})
}

func putReview(c *gin.Context) {

	var updatedReview review

	// Retrieve body request for review
	if err := c.BindJSON(&updatedReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Open sqlite db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bookID := c.Param("book_id")     // get book_id parameter from url
	reviewID := c.Param("review_id") // get review_id parameter from url

	// Check if book exist and get book information
	var bookTitle, bookAuthor string
	err = db.QueryRow("SELECT title, author FROM books WHERE id = ?", bookID).Scan(&bookTitle, &bookAuthor)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Check if review exist
	var id int
	err = db.QueryRow("SELECT id FROM reviews WHERE id = ?", reviewID).Scan(&id)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	// Update review
	review_query := `UPDATE reviews SET review = ?, rating = ? WHERE id = ?`
	_, err = db.Exec(review_query, updatedReview.Review, updatedReview.Rating, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update review to the Book Club"})
		return
	}

	// Update book's average rating after updating a book's rating
	ratings, err := db.Query("SELECT rating FROM reviews WHERE book_id = ?", bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}
	defer ratings.Close()

	// Recalculate book's average rating once review is updated
	var totalRating float64
	var count int
	for ratings.Next() {
		var rating float64
		if err := ratings.Scan(&rating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
			return
		}
		totalRating += rating
		count++
	}

	avgRating := totalRating / float64(count)
	avgRatingStr := fmt.Sprintf("%.1f", avgRating)

	// Update book average rating
	_, err = db.Exec("UPDATE books SET rating = ? WHERE id = ?", avgRatingStr, bookID)
	if err != nil {
		log.Println("Error updating book rating:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	// Return 200 message
	c.IndentedJSON(http.StatusOK, gin.H{"Message": "Review " + reviewID + " for " + bookTitle + " by " + bookAuthor + " updated to the Book Club"})
}

func deleteBook(c *gin.Context) {
	// Open sqlite db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// get book_id parameter from url
	bookID := c.Param("book_id")

	// Check if book exist and get book information
	var bookTitle, bookAuthor string
	err = db.QueryRow("SELECT title, author FROM books WHERE id = ?", bookID).Scan(&bookTitle, &bookAuthor)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Delete reviews first
	_, err = db.Exec("DELETE FROM reviews WHERE book_id = ?", bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed deleting reviews from book club"})
		return
	}

	// Delete book next
	_, err = db.Exec("DELETE FROM books WHERE id = ?", bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed deleting book from book club"})
	}

	// Return 200 message
	c.IndentedJSON(http.StatusOK, gin.H{"message": bookTitle + " by " + bookAuthor + " deleted"})
}

func deleteReview(c *gin.Context) {
	// Open sqlite db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Get parameters from url
	bookID := c.Param("book_id")
	reviewID := c.Param("review_id")

	// Check to see if book exist and get book information
	var bookTitle, bookAuthor string
	err = db.QueryRow("SELECT title, author FROM books WHERE id = ?", bookID).Scan(&bookTitle, &bookAuthor)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Check to see if review exist
	var id int
	err = db.QueryRow("SELECT id FROM reviews WHERE id = ?", reviewID).Scan(&id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	// Delete review
	_, err = db.Exec("DELETE FROM reviews WHERE id = ? AND book_id = ?", reviewID, bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting review"})
		return
	}

	// Update book's average rating after deleting a review
	ratings, err := db.Query("SELECT rating FROM reviews WHERE book_id = ?", bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}
	defer ratings.Close() // ensure operation closes

	// Recalculate book rating once a review is deleted
	var totalRating float64
	var count int
	for ratings.Next() {
		var rating float64
		if err := ratings.Scan(&rating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
			return
		}
		totalRating += rating
		count++
	}

	var avgRatingStr string
	if count > 0 {
		avgRating := totalRating / float64(count)
		avgRatingStr = fmt.Sprintf("%.1f", avgRating)
	} else {
		avgRatingStr = "N/A"
	}

	// Update book rating average with the remaining reviews
	_, err = db.Exec("UPDATE books SET rating = ? WHERE id = ?", avgRatingStr, bookID)
	if err != nil {
		log.Println("Error updating book rating:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Review " + reviewID + " for " + bookTitle + " by " + bookAuthor + " deleted"})
}

func deleteAllBooks(c *gin.Context) {
	// Open sqlite db
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Handle deleting all reviews
	_, err = db.Exec("DELETE FROM reviews")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting all reviews"})
	}

	// Handle deleting all books
	_, err = db.Exec("DELETE FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting all books"})
	}

	// Return 200 message
	c.IndentedJSON(http.StatusOK, gin.H{"message": "All books have been deleted from the book club"})

}

func main() {
	initialize_db() // Initialize database for each run

	router := gin.Default() //create default router

	router.Static("/static", "./static")        //locate static files
	router.LoadHTMLFiles("templates/home.html") //load html file so it can be rendered

	// Add get methods
	router.GET("/", getHome)
	router.GET("/books", getBooks)
	router.GET("/books/read", getReadBooks)
	router.GET("/books/unread", getUnreadBooks)
	router.GET("/books/current", getCurrentBook)
	router.GET("/books/:book_id/reviews", getReviews)

	// Add post methods
	router.POST("/books", postBook)
	router.POST("/books/:book_id/reviews", postReview)

	// Add put methods
	router.PUT("/books/:book_id/status", putReadStatus)
	router.PUT("/books/:book_id/current", putCurrStatus)
	router.PUT("/books/:book_id/reviews/:review_id", putReview)

	// Add delete
	router.DELETE("/books/:book_id", deleteBook)
	router.DELETE("/books/:book_id/reviews/:review_id", deleteReview)
	router.DELETE("/books", deleteAllBooks)
	// Get port and run routers
	port := os.Getenv("PORT")
	if port != "" {
		port = "localhost:" + port
		router.Run(port)
	} else {
		router.Run()
	}
}
