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

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	type Author struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Birthdate string `json:"birthdate"`
		Biography string `json:"biography"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	var author Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, "Invalid request data.", http.StatusBadRequest)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusBadRequest, r.UserAgent(),
			r.Referer(), "Invalid request data.")
		return
	}

	dbConnection, err := db.InitDB()
	if err != nil {
		http.Error(w, "Failed to connect to the database.", http.StatusBadRequest)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusBadRequest, r.UserAgent(),
			r.Referer(), "Failed to connect to the database.")
		return
	}
	defer dbConnection.Close()

	query := `
        INSERT INTO authors (first_name, last_name, birthdate, biography)
        VALUES (?, ?, ?, ?)
    `
	_, err = dbConnection.Exec(query, author.FirstName, author.LastName, author.Birthdate, author.Biography)
	if err != nil {
		http.Error(w, "Failed to create the author.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to create the author.")
		return
	}

	successMsg := struct {
		Message string `json:"message"`
	}{
		Message: "Author created successfully.",
	}
	w.Header().Set("Content-Type", "application/json")
	utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusOK, r.UserAgent(),
		r.Referer(), "Author created successfully.")
	json.NewEncoder(w).Encode(successMsg)
}

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorID := vars["id"]
	fmt.Println(vars)
	fmt.Println(authorID)

	dbConnection, err := db.InitDB()
	if err != nil {
		http.Error(w, "Failed to connect to the database.", http.StatusBadRequest)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusBadRequest, r.UserAgent(),
			r.Referer(), "Failed to connect to the database.")
		return
	}
	defer dbConnection.Close()

	query := "SELECT id, first_name, last_name, birthdate, biography, created_at, updated_at FROM authors WHERE id = ?"
	type Author struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Birthdate string `json:"birthdate"`
		Biography string `json:"biography"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	var author Author
	err = dbConnection.QueryRow(query, authorID).Scan(
		&author.ID,
		&author.FirstName,
		&author.LastName,
		&author.Birthdate,
		&author.Biography,
		&author.CreatedAt,
		&author.UpdatedAt,
	)

	if err != nil {
		http.Error(w, "Author not found.", http.StatusNotFound)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusNotFound, r.UserAgent(),
			r.Referer(), "Author not found.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusBadRequest, r.UserAgent(),
		r.Referer(), "Get author.")
	json.NewEncoder(w).Encode(author)
}

func Getauthors(w http.ResponseWriter, r *http.Request) {
	dbConnection, err := db.InitDB()
	if err != nil {
		http.Error(w, "Failed to connect to the database.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to connect to the database.")
		return
	}
	defer dbConnection.Close()

	query := "SELECT id, first_name, last_name, birthdate, biography, created_at, updated_at FROM authors"

	rows, err := dbConnection.Query(query)
	if err != nil {
		http.Error(w, "Failed to fetch authors.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to fetch authors.")
		return
	}
	defer rows.Close()

	type Author struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Birthdate string `json:"birthdate"`
		Biography string `json:"biography"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	var authors []Author

	for rows.Next() {
		var author Author
		err := rows.Scan(
			&author.ID,
			&author.FirstName,
			&author.LastName,
			&author.Birthdate,
			&author.Biography,
			&author.CreatedAt,
			&author.UpdatedAt,
		)
		if err != nil {
			http.Error(w, "Failed to scan authors.", http.StatusInternalServerError)
			utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
				r.Referer(), "Failed to scan authors.")
			return
		}
		authors = append(authors, author)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, "Failed to iterate over authors.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to iterate over authors.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusOK, r.UserAgent(),
		r.Referer(), "Get authors.")
	json.NewEncoder(w).Encode(authors)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorID := vars["id"]

	var updatedAuthor struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Birthdate string `json:"birthdate"`
		Biography string `json:"biography"`
	}

	err := json.NewDecoder(r.Body).Decode(&updatedAuthor)
	if err != nil {
		http.Error(w, "Invalid request data.", http.StatusBadRequest)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusBadRequest, r.UserAgent(),
			r.Referer(), "Invalid request data.")
		return
	}

	dbConnection, err := db.InitDB()
	if err != nil {
		http.Error(w, "Failed to connect to the database.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to connect to the database.")
		return
	}
	defer dbConnection.Close()

	query := "UPDATE authors SET"
	updateFields := []string{}
	values := []interface{}{}

	if updatedAuthor.FirstName != "" {
		updateFields = append(updateFields, "first_name = ?")
		values = append(values, updatedAuthor.FirstName)
	}
	if updatedAuthor.LastName != "" {
		updateFields = append(updateFields, "last_name = ?")
		values = append(values, updatedAuthor.LastName)
	}
	if updatedAuthor.Birthdate != "" {
		updateFields = append(updateFields, "birthdate = ?")
		values = append(values, updatedAuthor.Birthdate)
	}
	if updatedAuthor.Biography != "" {
		updateFields = append(updateFields, "biography = ?")
		values = append(values, updatedAuthor.Biography)
	}

	if len(updateFields) == 0 {
		http.Error(w, "No fields to update.", http.StatusBadRequest)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusBadRequest, r.UserAgent(),
			r.Referer(), "No fields to update.")
		return
	}

	query += " " + strings.Join(updateFields, ", ") + " WHERE id = ?"
	values = append(values, authorID)

	result, err := dbConnection.Exec(query, values...)
	if err != nil {
		http.Error(w, "Failed to update the author.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to update the author.")
		return
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected > 0 {
		successMsg := struct {
			Message string `json:"message"`
		}{
			Message: "Author updated successfully.",
		}
		w.Header().Set("Content-Type", "application/json")
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusOK, r.UserAgent(),
			r.Referer(), "Author updated successfully.")
		json.NewEncoder(w).Encode(successMsg)
	} else {
		http.Error(w, "No matching author found for the update.", http.StatusNotFound)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusNotFound, r.UserAgent(),
			r.Referer(), "No matching author found for the update.")
	}
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorID := vars["id"]

	dbConnection, err := db.InitDB()
	if err != nil {
		http.Error(w, "Failed to connect to the database.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to connect to the database.")
		return
	}
	defer dbConnection.Close()

	query := "DELETE FROM authors WHERE id = ?"

	result, err := dbConnection.Exec(query, authorID)
	if err != nil {
		http.Error(w, "Failed to delete the author.", http.StatusInternalServerError)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusInternalServerError, r.UserAgent(),
			r.Referer(), "Failed to delete the author.")
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
		http.Error(w, "Author not found or already deleted.", http.StatusNotFound)
		utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusNotFound, r.UserAgent(),
			r.Referer(), "Author not found or already deleted.")
		return
	}

	successMsg := struct {
		Message string `json:"message"`
	}{
		Message: "Author deleted successfully.",
	}
	w.Header().Set("Content-Type", "application/json")
	utils.CreateLoging(r.RemoteAddr, r.Method, r.URL.Path, http.StatusOK, r.UserAgent(),
		r.Referer(), "Author deleted successfully.")
	json.NewEncoder(w).Encode(successMsg)
}
