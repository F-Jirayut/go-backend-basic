package routes

import (
	"fmt"
	"go-basic/db"
	"go-basic/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func Routes() *mux.Router {
	router := mux.NewRouter()
	// hello world
	router.HandleFunc("/", handler).Methods("GET")

	// Define CRUD routes for items
	router.HandleFunc("/book", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books", handlers.Getbooks).Methods("GET")
	router.HandleFunc("/book/{id}", handlers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", handlers.DeleteBook).Methods("DELETE")

	// Define CRUD routes for items
	router.HandleFunc("/author", handlers.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors", handlers.Getauthors).Methods("GET")
	router.HandleFunc("/author/{id}", handlers.GetAuthorByID).Methods("GET")
	router.HandleFunc("/author/{id}", handlers.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/author/{id}", handlers.DeleteAuthor).Methods("DELETE")

	return router
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// next.ServeHTTP(w, r)

		dbConnection, err := db.InitDB()
		if err != nil {
			fmt.Println("Failed to connect to the database:", err)
			return
		}
		defer dbConnection.Close()

		// logEntry := struct {
		// 	IPAddress     string
		// 	RequestMethod string
		// 	RequestURL    string
		// 	StatusCode    int
		// 	UserAgent     string
		// 	Referer       string
		// 	LogMessage    string
		// }{
		// 	IPAddress:     r.RemoteAddr,
		// 	RequestMethod: r.Method,
		// 	RequestURL:    r.URL.Path,
		// 	StatusCode:    http.StatusOK,
		// 	UserAgent:     r.UserAgent(),
		// 	Referer:       r.Referer(),
		// 	LogMessage:    fmt.Sprintf("Request: %s %s", r.Method, r.URL.Path),
		// }

		// // Insert the log entry into the database
		// logQuery := `
		// 	INSERT INTO site_logs (ip_address, request_method, request_url, status_code, user_agent, referer, log_message)
		// 	VALUES (?, ?, ?, ?, ?, ?, ?)
		// `

		// _, err = dbConnection.Exec(logQuery, logEntry.IPAddress, logEntry.RequestMethod, logEntry.RequestURL, logEntry.StatusCode, logEntry.UserAgent, logEntry.Referer, logEntry.LogMessage)
		// if err != nil {
		// 	fmt.Println("Failed to insert log entry into the database:", err)
		// 	return
		// }
	})
}
