package main

import (
	"c-vod/controllers"
	"c-vod/services/clarity"
	"c-vod/utils/db"
	"c-vod/utils/globals"
	"c-vod/utils/helper"
	"c-vod/utils/storage"
	"c-vod/utils/types"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	globals.App = &types.App{}

	// setting the App.Config with .env file
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("could not load config : %v", err)
	}
	globals.App.Config = config

	// setting the App.Log
	mlog := &helper.Log{
		Log_enabled: config.Log_enabled == "true",
	}
	globals.App.Log = mlog

	// checking and creating app's necessary dirs
	err = helper.CheckAndCreateAppDirs()
	if err != nil {
		log.Fatalf("unabled to create initial directories : %v", err)
	}

	// setting the App.DB by initializing the db connection
	db, err := db.InitDB(true)
	if err != nil {
		log.Fatalf("database did not initialize : %v", err)
	}
	globals.App.DB = db

	// setting the App.Storage by initializing the minio connection
	mstoreage, err := storage.InitMinio()
	if err != nil {
		log.Fatalf("s3 storage did not initialize : %v", err)
	}
	globals.App.Storage = mstoreage

	// setting the App.Server and App.Router
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
