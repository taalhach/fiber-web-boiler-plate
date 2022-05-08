package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberrecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/cobra"
)

const port = 3000

var serveApiCmd = &cobra.Command{
	Use:   "serve_api",
	Short: "servers api",
	Long:  fmt.Sprintf("servers  api on localhost port %v", port),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := loadConfigs()
		if err != nil {
			fmt.Printf("failed to load configs: %v", err)
			os.Exit(1)
		}

		appConfigs := fiber.Config{
			ErrorHandler: customErrHandler,
			AppName:      "fiber-boiler-plate",
		}
		app := fiber.New(appConfigs)

		// logger middleware
		app.Use(logger.New(logger.Config{
			Format: fmt.Sprintf("[${time}] method=${method} uri=${path} status=${status} time=${latency}\n"),
		}))

		// middleware to recover from panics calls ErrorHandler
		app.Use(fiberrecover.New(fiberrecover.Config{
			EnableStackTrace: true,
		}))

		app.Use(compress.New(compress.Config{
			Level: compress.LevelBestCompression,
		}))

		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Go fiber says hello")
		})

		if err = app.Listen(fmt.Sprintf(":%d", port)); err != nil {
			fmt.Printf("failed to start server: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCommand.AddCommand(serveApiCmd)
}
