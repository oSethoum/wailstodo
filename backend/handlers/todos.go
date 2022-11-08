package handlers

import (
	"app/backend/db"
	"app/backend/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Query Todo
// @Tags    todos
// @Accept  json
// @Produce json
// @Param   query query    handlers.ManyQuery false "Query"
// @Success 200   {array}  models.Todo
// @Failure 400   {string} messgae
// @Failure 401   {string} messgae
// @Router  /todos [get]
func QueryTodo(c *fiber.Ctx) error {
	many := new([]models.Todo)
	q := new(ManyQuery)
	c.QueryParser(q)
	ParseManyQuery(db.Client, q).Find(many)
	return c.JSON(many)
}

// @Summary Get Todo
// @Tags    todos
// @Accept  json
// @Produce json
// @Param   id  path     int true "Todo ID"
// @Success 200 {object} models.Todo
// @Failure 400 {string} message
// @Failure 401 {string} message
// @Router  /todos/:id [get]
func OneTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	q := new(OneQuery)
	c.QueryParser(q)
	one := new(models.Todo)
	ParseOneQuery(db.Client, q).First(one, id)
	return c.JSON(one)
}

// @Summary Create One Todo
// @Tags    todos
// @Accept  json
// @Produce json
// @Param   body  body     models.Todo true "Create Todo Body"
// @Success 200 {object} models.Todo
// @Failure 400 {string} message
// @Failure 401 {string} message
// @Router  /todos/ [post]
func CreateTodo(c *fiber.Ctx) error {
	one := new(models.Todo)
	err := c.BodyParser(one)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err = db.Client.Create(one).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	Invalidate("todos")
	return c.JSON(one)
}

// @Summary Update One Todo
// @Tags    todos
// @Accept  json
// @Produce json
// @Param   id    path     int true "Todo ID"
// @Param   body  body     models.Todo true "Update Todo Body"
// @Success 200 {object} models.Todo
// @Failure 400 {string} message
// @Failure 401 {string} message
// @Router  /todos/:id [patch]
func UpdateTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	one := new(models.Todo)
	db.Client.Find(one, id)
	body := make(map[string]interface{})
	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err = db.Client.Model(one).Omit("id").Updates(body).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	Invalidate("todos")
	return c.Status(fiber.StatusOK).JSON(one)
}

// @Summary Delete One Todo
// @Tags    todos
// @Accept  json
// @Produce json
// @Param   id  path     int true "Todo ID"
// @Success 200 {object} models.Todo
// @Failure 400 {string} message
// @Failure 401 {string} message
// @Router  /todos/:id [delete]
func DeleteTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	one := new(models.Todo)
	db.Client.First(one, id)
	if err = db.Client.Delete(one, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	Invalidate("todos")
	return c.JSON(one)
}
