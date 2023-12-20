package main

import (
	"time"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{ //memberikan konfigurasi
		IdleTimeout: time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout: time.Second * 5,
		Prefork: true, 
		//pre fork -> fitur menjalankan proses Fiber namun menggunakan port yg sama(kalau sudah berjalan di sisi server)
	})

	app.Use("/api", func(ctx *fiber.Ctx) error {//membuat middleware sekaligus prefix("/api")
		fmt.Println("I'm middleware before processing request")
		err := ctx.Next()//"ctx.Next()" -> untuk meneruskan request ke handler selanjutnya
		fmt.Println("I'm middleware after processing request")
		return err
	})

	app.Get("/api/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	if fiber.IsChild() {
		fmt.Println("I'm child process")
	} else {
		fmt.Println("I'm parent process")
	}

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}