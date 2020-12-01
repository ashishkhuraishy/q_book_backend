package app

import "github.com/ashishkhuraishy/q_book_backend/controller/ping"

func urlMappings() {
	router.GET("/ping", ping.Pong)
}
