package main

import "fmt"

// App - the struct which contains things like pointers
// to database connections
type App struct {
}

// Run - this sets up the app
func (app *App) Run() error {
	fmt.Println("Setting up our App")
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
