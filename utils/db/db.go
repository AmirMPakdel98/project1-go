package db

import (
	"c-vod/models/enckeyModel"
	fileModel "c-vod/models/fileModel"
	uploadModel "c-vod/models/uploadModel"
	"c-vod/utils/globals"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(autoMigrate bool) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		globals.App.Config.Database_host,
		globals.App.Config.Database_username,
		globals.App.Config.Database_password,
		globals.App.Config.Database_name,
		globals.App.Config.Database_port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	if autoMigrate {
		// add new models here for db migration
		err = db.AutoMigrate(
			fileModel.File{},
			uploadModel.Upload{},
			enckeyModel.Enckey{},
		)

		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
