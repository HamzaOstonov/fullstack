package server

import (
	"log"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

/*
Start serverni to'liq ishga tushuruvchi funksiya
Agar serverning CPU yadrolari soni 4 tadan ko'p bo'ladigan bo'lsa ularni faqatgina 4 tasini ishlatish ko'rsatiladi
*/
func Start() {
	if runtime.NumCPU() > 4 {
		runtime.GOMAXPROCS(4)
	}

	//Appni e'lon qilish funksiyasi
	config := fiber.Config{ServerHeader: "AppTest server"}
	if runtime.NumCPU() > 1 {
		config.Prefork = true
	}
	app := fiber.New(config)

	// So'rovlarga yuboriladigan natijalarni compress qilish methodini qo'shish
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	//Cross site so'rovlari uchun CORS ni ochish methodini qo'shish
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,OPTIONS,PUT,DELETE",
	}))

	// Agar serverda panic hodisasi sodir bo'lsa server o'zini qayta ishga tushirish methodini qo'shish
	app.Use(recover.New())

	//Serverning routerlarini qo'shib chiqish

	log.Fatalln(app.Listen(":3000"))
}
