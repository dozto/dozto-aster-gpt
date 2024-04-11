package app

import (
	"github.com/dozto/dozto-aster-gpt/app/controllers"
	"github.com/dozto/dozto-aster-gpt/pkg"
	"github.com/gofiber/fiber/v3"
)

func InitRoutes(a *fiber.App, envs pkg.EnvVars) {
	generalRoute := a.Group("/")

	generalRoute.Get("/health", controllers.HealthCheck)

	gptRoute := a.Group("/gpt")
	gptController := controllers.NewGptController(envs.OPENAI_API_KEY)

	gptRoute.Post("/chat", gptController.ChatCompletion)
}
