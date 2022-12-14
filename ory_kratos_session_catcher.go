package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type configuration struct {
	port                    string
	redirectPath            string
	redirectSessionParamKey string
	sessionCookieKey        string
}

func (c *configuration) load() {
	c.port = os.Getenv("PORT")
	c.redirectPath = os.Getenv("REDIRECT_PATH")
	c.redirectSessionParamKey = os.Getenv("REDIRECT_SESSION_PARAM_KEY")
	c.sessionCookieKey = os.Getenv("SESSION_COOKIE_KEY")
}

func (c *configuration) display() {
	fmt.Println(
		"Configuration:\n##############",
		"\nPORT: ", c.port,
		"\nREDIRECT_PATH: ", c.redirectPath,
		"\nREDIRECT_SESSION_PARAM_KEY: ", c.redirectSessionParamKey,
		"\nSESSION_COOKIE_KEY: ", c.sessionCookieKey,
	)
}

func main() {
	c := &configuration{}
	c.load()
	c.display()

	app := fiber.New(fiber.Config{
		AppName: "Ory Kratos Session Catcher",
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		var sessionToken string = ctx.Cookies(c.sessionCookieKey)
		if sessionToken == "" {
			return ctx.SendStatus(400)
		}
		return ctx.Redirect(fmt.Sprintf("%s?%s=%s", c.redirectPath, c.redirectSessionParamKey, sessionToken))
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", c.port)))
}
