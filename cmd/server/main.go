package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/staticfrost/td-go-rest-api/internal/transport/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct {
}

// Run - this sets up the app
func (app *App) Run() error {
	fmt.Println("Setting up our App")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error start up REST API")
		fmt.Println(err)
	}
}
