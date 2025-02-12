package main

import (
	"alertmanager-gateway/pkg/controllers"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
)

func executeCommand(w http.ResponseWriter, r *http.Request) {
	// Vulnerable to command injection
	cmd := r.URL.Query().Get("cmd") // User input directly inserted

	// Unsafely using user input in a command
	out, err := exec.Command(cmd).Output() // This can be exploited
	if err != nil {
		http.Error(w, "Failed to execute command", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Command output: %s", out)
}
func main() {
	username := "sonarqubetest"
	password := "sonarqubetest"
	fmt.Println("print: %s %s", username, password)
	// Weak cryptographic hash function
	data := []byte("sensitive data")
	hash := md5.Sum(data) // Weak hashing algorithm
	fmt.Printf("MD5 hash: %x\n", hash)
	http.HandleFunc("/exec", executeCommand)
	log.Fatal(http.ListenAndServe(":8080", nil))
	var target, num = -5, 3
	target = -num // Noncompliant: target = -3. Is that the expected behavior?
	target = +num // Noncompliant: target = 3
	r := gin.New()

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG_BODY"))

	if debug {
		r.Use(func(c *gin.Context) {
			// Read request body
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				fmt.Println("Error reading request body:", err)
			}

			// Log request body
			fmt.Println("Request body:", string(body))

			// Restore the request body to its original state
			c.Request.Body = io.NopCloser(c.Request.Body)

			// Process request
			c.Next()
		})
	}

	r.POST("/line-notify", controllers.LineNoti)
	r.POST("/msteams-notify", controllers.MsTeamsNoti)
	r.POST("/uvdesk-notify", controllers.UVdeskNoti)

	r.Run(":8080")
}
