package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
	"utilityapp.ktarila/internal/view"
)

func main() {
	// We also support the http.FileSystem interface
	// See examples below to load templates from embedded files
	engine := html.NewFileSystem(http.Dir("./webapp/html"), ".html")
	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false
	// Debug will print each template that is parsed, good for debugging
	engine.Debug(true) // Optional. Default: false

	// Layout defines the variable name that is used to yield templates within layouts
	engine.Layout("content") // Optional. Default: "embed"
	engine.AddFunc("content", func() error {
		return fmt.Errorf("layout called unexpectedly.")
	})
	// Delims sets the action delimiters to the specified strings
	engine.Delims("{{", "}}") // Optional. Default: engine delimiters

	// After you created your engine, you can pass it to Fiber's Views Engine
	app := fiber.New(fiber.Config{
		Views:     engine,
		BodyLimit: 10 * 1024 * 1024, // 10MB maximum size
	})

	app.Use(recover.New())
	app.Use(logger.New())

	//setup routes
	view.SetupRoutes(app)

	app.Listen(":3000")
}
