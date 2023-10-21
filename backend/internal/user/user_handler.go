package user

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var u CreateUserReq
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.Service.CreateUser(c.Context(), &u)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(res)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var user LoginUserReq
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	u, err := h.Service.Login(c.Context(), &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    u.accessToken,
		Domain:   "localhost",
		HTTPOnly: true,
		Secure:   false,
		MaxAge:   3600,
	}

	c.Cookie(&cookie)

	res := &LoginUserRes{
		Username: u.Username,
		ID:       u.ID,
	}

	return c.JSON(res)
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Domain:   "",
		HTTPOnly: true,
		Secure:   false,
		MaxAge:   -1,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logged out successfully"})
}
