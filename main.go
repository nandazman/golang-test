package main

import (
	"os"

	"github.com/golang/Test/handler"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	envFile := os.Getenv("ENV")
	if envFile == "" {
		envFile = ".env"
	}

	// Load env data
	err := godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}

	// CREATING DATABASE

	// table := new(handler.TableXutopia)

	// err = table.CreateTable(db)
	// if err != nil {
	// 	panic(err)
	// }

	// ENDPOINT
	e.GET("/foo/:param/", handler.Bar)
	e.POST("/post/", handler.PostSomething)

	// TENTANG TABLE UTAMA
	e.POST("/tablexutopia", handler.CreateData)
	e.GET("/tablexutopia", handler.GetData)
	e.GET("/tablexutopias", handler.GetAllData)

	// Start echo on desired port
	e.Start(":" + os.Getenv("PORT"))
}
