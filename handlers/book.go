package handlers

import (
	"encoding/json"
	"fmt"
	"go-basic/db"
	"go-basic/utils"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {

	// Parse the request body to get book data
	type Book struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		AuthorID    int    `json:"author_id"`
	}
	var book = Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data.", http.StatusBadRequest)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusBadRequest, r.UserAgent(),
			r.Referer(), "Invalid request data.")
		return
	}

	// Insert the book into the database
	dbConnection, err := db.InitDB()
	defer dbConnection.Close()

	// Define the SQL query to insert a new book
	query := `
        INSERT INTO books (name, description, author_id)
        VALUES (?, ?, ?)
    `
	_, err = dbConnection.Exec(query, book.Name, book.Description, book.AuthorID) // Don't assign to result

	if err != nil {
		http.Error(w, "Failed to create the book.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to create the book.")
		return
	}

	// Send a success response
	successMsg := struct {
		Message string `json:"message"`
	}{
		Message: "Book created successfully.",
	}
	utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusOK, r.UserAgent(),
		r.Referer(), "Book created successfully.")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(successMsg)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	// Extract the book ID from the request parameters
	vars := mux.Vars(r)
	bookID := vars["id"]
	// Initialize the database connection
	dbConnection, err := db.InitDB()
	if err != nil {
		http.Error(w, "Failed to connect to the database.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to connect to the database.")
		return
	}
	defer dbConnection.Close()

	// Define the SQL query to select a book by its ID
	query := "SELECT id, name, description, author_id, created_at, updated_at FROM books WHERE id = ?"
	type Book struct {
		BookID      int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		AuthorID    int    `json:"author_id"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
	// Execute the query and scan the result into a Book struct
	var book Book
	err = dbConnection.QueryRow(query, bookID).Scan(
		&book.BookID,
		&book.Name,
		&book.Description,
		&book.AuthorID,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		http.Error(w, "Book not found.", http.StatusNotFound)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusNotFound, r.UserAgent(),
			r.Referer(), "Book not found.")
		return
	}

	// Serialize the book as JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusOK, r.UserAgent(),
		r.Referer(), "Get book.")
	json.NewEncoder(w).Encode(book)
}

func Getbooks(w http.ResponseWriter, r *http.Request) {
	dbConnection, err := db.InitDB()
	if err != nil {
		http.Error(w, "Failed to connect to the database.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to connect to the database.")
		return
	}
	defer dbConnection.Close()

	query := "SELECT id, name, description, author_id, created_at, updated_at FROM books"

	rows, err := dbConnection.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch books.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to fetch books.")
		return
	}
	defer rows.Close()

	type Book struct {
		BookID      int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		AuthorID    int    `json:"author_id"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}

	var books []Book

	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.BookID,
			&book.Name,
			&book.Description,
			&book.AuthorID,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			http.Error(w, "Failed to scan books.", http.StatusInternalServerError)
			utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
				r.Referer(), "Failed to fetch books.")
			return
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Failed to iterate over books.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to iterate over books.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
		r.Referer(), "Get books.")
	json.NewEncoder(w).Encode(books)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	var updatedBook struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		AuthorID    int    `json:"author_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid request data.", http.StatusBadRequest)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusBadRequest, r.UserAgent(),
			r.Referer(), "Invalid request data.")
		return
	}

	dbConnection, err := db.InitDB()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to connect to the database.")
		return
	}
	defer dbConnection.Close()

	query := "UPDATE books SET"
	updateFields := []string{}
	values := []interface{}{}

	if updatedBook.Name != "" {
		updateFields = append(updateFields, "name = ?")
		values = append(values, updatedBook.Name)
	}
	if updatedBook.Description != "" {
		updateFields = append(updateFields, "description = ?")
		values = append(values, updatedBook.Description)
	}
	if updatedBook.AuthorID != 0 {
		updateFields = append(updateFields, "author_id = ?")
		values = append(values, updatedBook.AuthorID)
	}

	if len(updateFields) == 0 {
		http.Error(w, "No fields to update.", http.StatusBadRequest)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusBadRequest, r.UserAgent(),
			r.Referer(), "No fields to update.")
		return
	}

	query += " " + strings.Join(updateFields, ", ") + " WHERE id = ?"
	values = append(values, bookID)

	result, err := dbConnection.Exec(query, values...)
	if err != nil {
		http.Error(w, "Failed to update the book.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to update the book.")
		return
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected > 0 {
		successMsg := struct {
			Message string `json:"message"`
		}{
			Message: "Book updated successfully.",
		}
		w.Header().Set("Content-Type", "application/json")
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusOK, r.UserAgent(),
			r.Referer(), "Book updated successfully.")
		json.NewEncoder(w).Encode(successMsg)
	} else {
		http.Error(w, "No matching book found for the update.", http.StatusNotFound)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusNotFound, r.UserAgent(),
			r.Referer(), "No matching book found for the update.")
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	dbConnection, err := db.InitDB()
	if err != nil {
		http.Error(w, "Failed to connect to the database.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to connect to the database.")
		return
	}
	defer dbConnection.Close()

	query := "DELETE FROM books WHERE id = ?"

	result, err := dbConnection.Exec(query, bookID)
	if err != nil {
		http.Error(w, "Failed to delete the book.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to delete the book.")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to get the number of rows affected.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to get the number of rows affected.")
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Book not found or already deleted.", http.StatusNotFound)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusNotFound, r.UserAgent(),
			r.Referer(), "Book not found or already deleted.")
		return
	}

	successMsg := struct {
		Message string `json:"message"`
	}{
		Message: "Book deleted successfully.",
	}
	w.Header().Set("Content-Type", "application/json")
	utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusOK, r.UserAgent(),
		r.Referer(), "Book deleted successfully.")
	json.NewEncoder(w).Encode(successMsg)
}
