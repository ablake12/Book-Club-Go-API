function getBooks() {
    fetch('/books', 
        {
            method: 'GET'
        })
        .then(response => {
            document.getElementById("all-books-http").innerText = `HTTP Status: ${response.status}`;
            document.getElementById("all-books-http").style.display = 'block';

            return response.json()
        })
        .then(data => {
            document.getElementById("all-books-json").innerText = JSON.stringify(data, null, 2);
            document.getElementById("all-books-json").style.display = 'block';
        })
        .catch(error => {
            document.getElementById("all-books-json").innerText = `Error fetching data: ${error.message}`;
            document.getElementById("all-books-json").style.display = 'block';
        });
}

function clearBooks() {
    document.getElementById("all-books-http").style.display = 'none';
    document.getElementById("all-books-json").style.display = 'none';
}

function getReadBooks() {
    fetch('/books/read', 
        { 
            method: 'GET'
        })
        .then(response => {
            document.getElementById("read-books-http").innerText = `HTTP Status: ${response.status}`;
            document.getElementById("read-books-http").style.display = 'block';

            return response.json()
        })
        .then(data => {
            document.getElementById("read-books-json").innerText = JSON.stringify(data, null, 2);
            document.getElementById("read-books-json").style.display = 'block';
        })
        .catch(error => {
            document.getElementById("read-books-json").innerText = `Error fetching data: ${error.message}`;
            document.getElementById("read-books-json").style.display = 'block';
        });   
}

function clearReadBooks() {
    document.getElementById("read-books-http").style.display = 'none';
    document.getElementById("read-books-json").style.display = 'none';
}

function getUnreadBooks() {
    fetch('/books/unread', 
        { 
            method: 'GET'
        })
        .then(response => {
            document.getElementById("unread-books-http").innerText = `HTTP Status: ${response.status}`;
            document.getElementById("unread-books-http").style.display = 'block';

            return response.json()
        })
        .then(data => {
            document.getElementById("unread-books-json").innerText = JSON.stringify(data, null, 2);
            document.getElementById("unread-books-json").style.display = 'block';
        })
        .catch(error => {
            document.getElementById("unread-books-json").innerText = `Error fetching data: ${error.message}`;
            document.getElementById("unread-books-json").style.display = 'block';
        });
}

function clearUnreadBooks() {
    document.getElementById("unread-books-http").style.display = 'none';
    document.getElementById("unread-books-json").style.display = 'none';
}

function getCurrentBook() {
    fetch('/books/current', 
        {
            method: 'GET'
        })
        .then(response => {
            document.getElementById("current-book-http").innerText = `HTTP Status: ${response.status}`;
            document.getElementById("current-book-http").style.display = 'block';

            return response.json()
        })
        .then(data => {
            document.getElementById("current-book-json").innerText = JSON.stringify(data, null, 2);
            document.getElementById("current-book-json").style.display = 'block';
        })
        .catch(error => {
            document.getElementById("current-book-json").innerText = `Error fetching data: ${error.message}`;
            document.getElementById("current-book-json").style.display = 'block';
        });
}

function clearCurrentBook() {
    document.getElementById("current-book-http").style.display = 'none';
    document.getElementById("current-book-json").style.display = 'none';
}

function getReviews(book_id) {
    fetch('/books/' + book_id + '/reviews', 
        {
            method: 'GET',
        })
        .then(response => {
            document.getElementById("reviews-http").innerText = `HTTP Status: ${response.status}`;
            document.getElementById("reviews-http").style.display = 'block';
            return response.json()
        })
        .then(data => {
            document.getElementById("reviews-json").innerText = JSON.stringify(data, null, 2);
            document.getElementById("reviews-json").style.display = 'block';
        })
        .catch(error => {
            document.getElementById("reviews-json").innerText = `Error fetching data: ${error.message}`;
            document.getElementById("reviews-json").style.display = 'block';
        }); 
}

function clearReviews() {
    document.getElementById("get-reviews-form").reset();
    document.getElementById("reviews-http").style.display = 'none';
    document.getElementById("reviews-json").style.display = 'none';
}

function addBook() {

    const title = document.getElementById("title").value;
    const author = document.getElementById("author").value;
    const genre = document.getElementById("genre").value;
    const desc = document.getElementById("desc").value;
    const read_status = document.getElementById("read_status").value;
    const curr_status = document.getElementById("curr_status").value;

    fetch('/books', {
            method: 'POST',
            body: JSON.stringify({
                    title: title,
                    author: author,
                    genre: genre,
                    description: desc,
                    read_status: read_status,
                    current_status: curr_status
                }
            )
        })
        .then(response => {
            document.getElementById("add-book-http").innerText = `HTTP Status: ${response.status}`;
            document.getElementById("add-book-http").style.display = 'block';

            return response.json()
        })
        .then(data => {
            document.getElementById("add-book-json").innerText = JSON.stringify(data, null, 2);
            document.getElementById("add-book-json").style.display = 'block';
        })
        .catch(error => {
            document.getElementById("add-book-json").innerText = `Error fetching data: ${error.message}`;
            document.getElementById("add-book-json").style.display = 'block';
        });
}
function clearAddBook() {
    document.getElementById("add-book-form").reset();
    document.getElementById("add-book-http").style.display = 'none';
    document.getElementById("add-book-json").style.display = 'none';
}

function addReview(book_id){
    const user = document.getElementById("user").value;
    const review = document.getElementById("review").value;
    const rating = document.getElementById("rating").value;
    fetch('/books/' + book_id + '/reviews', {
        method: 'POST',
        body: JSON.stringify({
                user: user,
                review: review,
                rating: rating
            }
        )
    })
    .then(response => {
        document.getElementById("add-review-http").innerText = `HTTP Status: ${response.status}`;
        document.getElementById("add-review-http").style.display = 'block';

        return response.json()
    })
    .then(data => {
        document.getElementById("add-review-json").innerText = JSON.stringify(data, null, 2);
        document.getElementById("add-review-json").style.display = 'block';
    })
    .catch(error => {
        document.getElementById("add-review-json").innerText = `Error fetching data: ${error.message}`;
        document.getElementById("add-review-json").style.display = 'block';
    });
}

function clearAddReview() {
    document.getElementById("add-review-form").reset();
    document.getElementById("add-review-http").style.display = 'none';
    document.getElementById("add-review-json").style.display = 'none';
}

function updateReadStatus(book_id) {
    fetch('/books/' + book_id + '/status', 
        {
            method: 'PUT'
        })
        .then(response => {
            document.getElementById("update-read-http").innerText = `HTTP Status: ${response.status}`;
            document.getElementById("update-read-http").style.display = 'block';

            return response.json()
        })
        .then(data => {
            document.getElementById("update-read-json").innerText = JSON.stringify(data, null, 2);
            document.getElementById("update-read-json").style.display = 'block';
        })
        .catch(error => {
            document.getElementById("update-read-json").innerText = `Error fetching data: ${error.message}`;
            document.getElementById("update-read-json").style.display = 'block';
        }); 
}

function clearUpdateReadStatus() {
    document.getElementById("update-read-form").reset();
    document.getElementById("update-read-http").style.display = 'none';
    document.getElementById("update-read-json").style.display = 'none';
}

function updateCurrStatus(book_id) {
    fetch('/books/' + book_id + '/current', 
        {
            method: 'PUT'
        })
        .then(response => {
            document.getElementById("update-curr-http").innerText = `HTTP Status: ${response.status}`;
            document.getElementById("update-curr-http").style.display = 'block';

            return response.json()
        })
        .then(data => {
            document.getElementById("update-curr-json").innerText = JSON.stringify(data, null, 2);
            document.getElementById("update-curr-json").style.display = 'block';
        })
        .catch(error => {
            document.getElementById("update-curr-json").innerText = `Error fetching data: ${error.message}`;
            document.getElementById("update-curr-json").style.display = 'block';
        });
}

function clearUpdateCurrStatus() {
    document.getElementById("update-curr-form").reset();
    document.getElementById("update-curr-http").style.display = 'none';
    document.getElementById("update-curr-json").style.display = 'none';
}

function updateReview(book_id, review_id) {
    const review = document.getElementById("update_review").value;
    const rating = document.getElementById("update_rating").value;
    fetch("/books/" + book_id + "/reviews/" + review_id, {
        method: 'PUT',
        body: JSON.stringify({
                review: review,
                rating: rating
            }
        )
    })
    .then(response => {
        document.getElementById("update-review-http").innerText = `HTTP Status: ${response.status}`;
        document.getElementById("update-review-http").style.display = 'block';
    
        return response.json()
    })
    .then(data => {
        document.getElementById("update-review-json").innerText = JSON.stringify(data, null, 2);
        document.getElementById("update-review-json").style.display = 'block';
    })
    .catch(error => {
        document.getElementById("update-review-json").innerText = `Error fetching data: ${error.message}`;
        document.getElementById("update-review-json").style.display = 'block';
    });
}

function clearUpdateReview() {
    document.getElementById("update-review-form").reset();
    document.getElementById("update-review-http").style.display = 'none';
    document.getElementById("update-review-json").style.display = 'none';
}

function deleteBook(book_id) {
    fetch('/books/' + book_id, 
        {
            method: 'DELETE'
        })
    .then(response => {
        document.getElementById("delete-book-http").innerText = `HTTP Status: ${response.status}`;
        document.getElementById("delete-book-http").style.display = 'block';

        return response.json()
    })
    .then(data => {
        document.getElementById("delete-book-json").innerText = JSON.stringify(data, null, 2);
        document.getElementById("delete-book-json").style.display = 'block';
    })
    .catch(error => {
        document.getElementById("delete-book-json").innerText = `Error fetching data: ${error.message}`;
        document.getElementById("delete-book-json").style.display = 'block';
    });
}

function clearDeleteBook() {
    document.getElementById("delete-book-form").reset();
    document.getElementById("delete-book-http").style.display = 'none';
    document.getElementById("delete-book-json").style.display = 'none';
}

function deleteReview(book_id, review_id) {
    fetch('/books/' + book_id + '/reviews/' + review_id, 
        {
            method: 'DELETE'
        })
    .then(response => {
        document.getElementById("delete-review-http").innerText = `HTTP Status: ${response.status}`;
        document.getElementById("delete-review-http").style.display = 'block';

        return response.json()
    })
    .then(data => {
        document.getElementById("delete-review-json").innerText = JSON.stringify(data, null, 2);
        document.getElementById("delete-review-json").style.display = 'block';
    })
    .catch(error => {
        document.getElementById("delete-review-json").innerText = `Error fetching data: ${error.message}`;
        document.getElementById("delete-review-json").style.display = 'block';
    });
}

function clearDeleteReview() {
    document.getElementById("delete-review-form").reset();
    document.getElementById("delete-review-http").style.display = 'none';
    document.getElementById("delete-review-json").style.display = 'none';
}

function deleteAllBooks() {
    fetch('/books', 
        {
            method: 'DELETE'
        })
    .then(response => {
        document.getElementById("delete-all-http").innerText = `HTTP Status: ${response.status}`;
        document.getElementById("delete-all-http").style.display = 'block';

        return response.json()
    })
    .then(data => {
        document.getElementById("delete-all-json").innerText = JSON.stringify(data, null, 2);
        document.getElementById("delete-all-json").style.display = 'block';
    })
    .catch(error => {
        document.getElementById("delete-all-json").innerText = `Error fetching data: ${error.message}`;
        document.getElementById("delete-all-json").style.display = 'block';
    });
}

function clearDeleteBooks() {
    document.getElementById("delete-all-http").style.display = 'none';
    document.getElementById("delete-all-json").style.display = 'none';
}