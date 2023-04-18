package main

import (
    "log"

    "bank_api/pkg/commons/config"
    "bank_api/pkg/commons/db"
    "bank_api/pkg/users"

    "github.com/gofiber/fiber/v2"

    swagger "github.com/arsmn/fiber-swagger/v2"
    _ "bank_api/docs"
)

// @title Bank API
// @version 0.0.1
// @description This is a sample bank api
// @contact.email fbazmitsuishi@gmail.com
// @host localhost:3000
// @BasePath /
func main() {
    configs, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed to load configs", err)
    }

    database := db.Init(&configs)
    app := fiber.New()

    app.Get("/swagger/*", swagger.HandlerDefault) // default
    // app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
	// 	DeepLinking: false,
	// 	// Expand ("list") or Collapse ("none") tag groups by default
	// 	DocExpansion: "none",
	// 	// Prefill OAuth ClientId on Authorize popup
	// 	OAuth: &swagger.OAuthConfig{
	// 		AppName:  "OAuth Provider",
	// 		ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
	// 	},
	// 	// Ability to change OAuth2 redirect uri location
	// 	OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	// }))

    users.RegisterRoutes(app, database)

    app.Listen(configs.Port)
}