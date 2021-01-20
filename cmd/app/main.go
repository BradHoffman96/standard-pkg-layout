package main

import "log"

type App struct {
}

func main() {

	app := App{}

	if err := app.run(); err != nil {
		log.Fatal(err)
	}
}

func (app *App) run() error {

	return nil
}
