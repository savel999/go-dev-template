// main.go
package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	appPort := os.Getenv("APP_PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Docker 4545:port" + appPort))
		fmt.Println("Hello from Docker")
	})

	fmt.Println("Listening on :" + appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}
