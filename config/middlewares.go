package config

import (
	"context"
	"errors"
	"fmt"
	"hmis/helpers"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	redisStorage "github.com/gofiber/storage/redis"
)

var limiterStorage = redisStorage.New()

func SetupMiddlewares(app *fiber.App) {

	client, ctx, cancel, err := Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer Close(client, ctx, cancel)
	Ping(client, ctx)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     helpers.Ternary(os.Getenv("SERVER_ENV") == "development", "*", "*"),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Use(limiter.New(limiter.Config{
		Storage:    limiterStorage,
		Max:        50,
		Expiration: 1 * time.Minute,
	}))

	app.Use(logger.New(logger.Config{
		Format: "${time} ${locals:requestid} ${status} ${method}${path}  ${latency}\n",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(recover.New(recover.Config{
		EnableStackTrace:  os.Getenv("SERVER_ENV") == "development",
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {},
	}))

	app.Use(func(c *fiber.Ctx) error {
		helpers.SetLocal(c, "db", client) // add db connection to context
		helpers.SetLocal(c, "ctx", ctx)   // add context to context
		return c.Next()
	})
}

var ErrorTimeOut = errors.New("server taking too long to respond")

func SleepWithContext(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	select {
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C
		}
		return ErrorTimeOut
	case <-timer.C:
	}
	return nil
}

func ServerTimeout(c *fiber.Ctx) error {
	sleepTime, _ := time.ParseDuration(c.Params("sleepTime") + "ms")
	if err := SleepWithContext(c.UserContext(), sleepTime); err != nil {
		return fmt.Errorf("%w: execution error", err)
	}
	return nil
}
