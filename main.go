package main

import (
	"c-vod/controllers"
	"c-vod/services/clarity"
	"c-vod/utils/db"
	"c-vod/utils/globals"
	"c-vod/utils/storage"
	"c-vod/utils/types"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	globals.App = &types.App{}

	config, err := LoadConfig()

	if err != nil {
		log.Fatalf("could not load config : %v", err)
	}

	globals.App.Config = config

	db, err := db.InitDB(true)

	if err != nil {
		log.Fatalf("database did not initialize : %v", err)
	}

	globals.App.DB = db

	mstoreage, err := storage.InitMinio()

	if err != nil {
		log.Fatalf("s3 storage did not initialize : %v", err)
	}

	globals.App.Storage = mstoreage

	globals.App.Server = fiber.New(fiber.Config{
		BodyLimit: 900 * 1024 * 1024,
	})

	globals.App.Router = globals.App.Server.Group("/vod/api/v1")

	globals.App.Router.Post("/file/create", controllers.CreateFile)

	globals.App.Router.Post("/file/upload", controllers.UploadFile)

	globals.App.Router.Get("/file/:id", controllers.GetFile)

	globals.App.Router.Delete("/file/:id", controllers.DeleteFile)

	//create and start clarity service
	clatiry := clarity.New()

	go clatiry.Start()

	err = globals.App.Server.Listen(":" + globals.App.Config.App_port)

	if err != nil {
		log.Fatalf("server did not run : %v", err)
	}
}
