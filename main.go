package main

import (
	_ "github.com/lib/pq"
	"golang_server/database"
	"golang_server/middleware"
	"log"
	_ "net/http"
	"os"
)

/*
Documentation references:
  - go
    -> https://www.practical-go-lessons.com/chap-9-control-statements
    -> https://go.dev/tour/basics/11
  - sqlboilder https://blog.logrocket.com/introduction-sqlboiler-go-framework-orms/
  - gin https://github.com/gin-gonic/gin OR https://gin-gonic.com/docs/
*/
func main() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)

	err := database.InitDatabase()
	if err != nil {
		log.Panic(err)
	}

	r := middleware.CreateRouter()
	err = r.Run(":" + port)
	if err != nil {
		log.Panic(err)
	}
}
