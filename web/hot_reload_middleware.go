package web

import (
	"context"
	"gofinance/web/view/devtools"
	"strings"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Message struct {
	Id string `json:"id"`
}

func WithHotReload(app *fiber.App) {
	id := uuid.New().String()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/hotreload", websocket.New(func(c *websocket.Conn) {
		for {
			var message Message
			c.ReadJSON(&message)

			if message.Id != id {
				var builder = &strings.Builder{}
				reload := message.Id != ""

				if err := devtools.HotReload(id, reload).Render(context.Background(), builder); err != nil {
					break
				}

				if err := c.WriteMessage(websocket.TextMessage, []byte(builder.String())); err != nil {
					break
				}
			}

		}
	}))

}
