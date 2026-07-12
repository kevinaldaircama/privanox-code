package main

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

func main() {
	var m tele.ReplyMarkup
	btn := m.WebApp("Ver Anuncio", &tele.WebApp{URL: "https://example.com"})
	fmt.Printf("%T\n", btn)
}
