package backend

import (
	"app/backend/db"
	_ "app/backend/docs"
	"app/backend/routes"
	"log"
)

// @title       App
// @version     1.0
// @description This is an API Application

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
func Serve() {
	db.Connect()
	app := routes.New()
	log.Fatalln(app.Listen(":5000"))
}
