package utils

import (
	"fmt"
	"go-basic/db"
)

func CreateLoging(ip string, method string, url string, status int, user_agent string, referer string, message string) {
	dbConnection, err := db.InitDB()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer dbConnection.Close()
	logEntry := struct {
		IPAddress     string
		RequestMethod string
		RequestURL    string
		StatusCode    int
		UserAgent     string
		Referer       string
		LogMessage    string
	}{
		IPAddress:     ip,
		RequestMethod: method,
		RequestURL:    url,
		StatusCode:    status,
		UserAgent:     user_agent,
		Referer:       referer,
		LogMessage:    message,
	}

	// Insert the log entry into the database
	logQuery := `
		INSERT INTO site_logs (ip_address, request_method, request_url, status_code, user_agent, referer, log_message)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err = dbConnection.Exec(logQuery, logEntry.IPAddress, logEntry.RequestMethod, logEntry.RequestURL, logEntry.StatusCode, logEntry.UserAgent, logEntry.Referer, logEntry.LogMessage)
	if err != nil {
		fmt.Println("Failed to insert log entry into the database:", err)
		return
	}
}
