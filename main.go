package main

import (
	"github.com/dozto/dozto-aster-gpt/app"
	"github.com/dozto/dozto-aster-gpt/pkg"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	// InitLogger
	// pkg.InitLogger()

	// LoadEnvs
	envs, err := pkg.LoadEnvs()
	if err != nil {
		panic(err)
	}

	// Connect to MongoDB
	// db, err := database.ConnecMongoDB(envs.MONGODB_URI, envs.MONGODB_NAME, 20*time.Second)
	// if err != nil {
	// 	panic(err)
	// }

	// Init Fiber Server
	fib := fiber.New()
	// app.Get("/health", healthcheck.HandlerHealthCheck)
	// gpt.InitGptRoutes(app, envs.OPENAI_API_KEY)

	app.InitRoutes(fib, envs)
	log.Infof("Server is running on port: %s", envs.PORT)
	fib.Listen(":" + envs.PORT)
}
