package handler

import (
	"SimpleWebProject/internal/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type UserService interface {
	GetAll() []*model.User
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (handler *UserHandler) InitRoutes(app *fiber.App) {
	app.Get("/users", handler.GetAll)
	app.Get("/delay", handler.GetAllWithDelay)
	app.Get("/error", handler.Error)
}

// GetAll godoc
// @Summary 	Get Users
// @Description Get list of users
// @Tags 		users
// @Produce 	json
// @Success 	200 {object} 		[]model.User
// @Router 		/users [get]
func (handler *UserHandler) GetAll(ctx *fiber.Ctx) error {
	users := handler.service.GetAll()

	return ctx.Status(http.StatusOK).JSON(
		fiber.Map{
			"users": users,
		})
}

func (handler *UserHandler) GetAllWithDelay(ctx *fiber.Ctx) error {
	time.Sleep(6 * time.Second)
	users := handler.service.GetAll()

	return ctx.Status(http.StatusOK).JSON(
		fiber.Map{
			"users": users,
		})
}

func (handler *UserHandler) Error(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"description": "Error",
	})
}
