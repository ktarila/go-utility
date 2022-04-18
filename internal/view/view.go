package view

import (
	"github.com/gofiber/fiber/v2"
)

type Model struct {
	Name  string
	Title string
	Data  map[string]interface{}
}

func New(name, title string) *Model {
	return &Model{Name: name, Title: title}
}

func SetupRoutes(app *fiber.App) {
	app.Static("/static", "./webapp/dist")

	app.Get("/", Home)
	app.Post("/reduce-pdf", ReducePdf)
	app.Get("/reduce-pdf", ReducePdf)
	app.Get("/convert-image", ImageConvert)
	app.Post("/convert-image", ImageConvert)
}

func render(ctx *fiber.Ctx, m *Model) error {
	return ctx.Render(m.Name, fiber.Map{"Title": m.Title}, "layouts/main")

}
