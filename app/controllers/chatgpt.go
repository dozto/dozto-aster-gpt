package controllers

import (
	"github.com/dozto/dozto-aster-gpt/pkg/chatgpt"
	"github.com/gofiber/fiber/v3"
)

type GptController struct {
	token string
}

func NewGptController(token string) *GptController {
	return &GptController{token: token}
}

func (g *GptController) ChatCompletion(c fiber.Ctx) error {
	req := new(chatgpt.ChatCompletionRequest)

	if err := c.Bind().Body(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	readCloser, err := chatgpt.ChatCompletion(g.token, *req)
	if err != nil {
		return err
	}

	return c.SendStream(readCloser, -1)
}
