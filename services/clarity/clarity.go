package clarity

import (
	fileModel "c-vod/models/fileModel"
	"log"
	"time"
)

/*
Clarity will handle overall job with its queue
*/
type Clarity struct {
	sleepDuration time.Duration
	elisma        *Elisma
	funpeg        *Funpeg
	stigma        *Stigma
}

func New() *Clarity {
	Edborn = &_Edborn{}
	Trauma = &_Trauma{}
	return &Clarity{
		sleepDuration: time.Duration(5 * time.Second),
		elisma:        &Elisma{},
		funpeg:        &Funpeg{},
		stigma:        &Stigma{},
	}
}

func (cl *Clarity) Start() {

	//cl.stigma.Test_upload_file_to_s3()

	for {
		cl.checkForNewFile()
		time.Sleep(5 * time.Second)
	}
}

func (cl *Clarity) checkForNewFile() {

	log.Print("checking for new file")

	file := Edborn.findWaitingFile()

	if file == nil {
		log.Printf("no file. sleeping for %v", cl.sleepDuration)
		time.Sleep(cl.sleepDuration)
		return
	}

	/* TODO: check file exists in every steps
	file_path := fmt.Sprintf("./storage/stage1/%d.%s", dbfile.Id, dbfile.Ext)
	file_path = filepath.Join(helper.GetCurrentDir(), file_path)

	//check if file exists
	_, err := os.Stat(file_path)

	if errors.Is(err, os.ErrNotExist) {
		//delete the db record
		globals.App.DB.Delete(&dbfile)
		return
	}
	*/

	//cl.handleByStatus(dbfile)

	var err error
	/*
		+ dbfile type is Video and clarity will call funpeg
		for video standadization
	*/
	if file.Status == fileModel.UPLOADED {
		err = cl.funpeg.standardization(file)
		if err != nil {
			return
		}
		//	+ clarity will call elisma to handle this video file
		err = cl.elisma.convertToStreamableFormat(file)
		if err != nil {
			return
		}
	}

	if file.Status == fileModel.NORMALIZED {
		//	+ clarity will call elisma to handle this video file
		err = cl.elisma.convertToStreamableFormat(file)
		if err != nil {
			return
		}
	}

	// + clarity will call stigma to store file in storage
	err = cl.stigma.storeFileInStorage(file)
	if err != nil {
		return
	}
}

func (cl *Clarity) handleByStatus(file *fileModel.File) {

	switch file.Status {
	case fileModel.UPLOADED:
		cl.funpeg.standardization(file)
	case fileModel.NORMALIZED:
		cl.elisma.convertToStreamableFormat(file)
	case fileModel.READY_FOR_STORAGE:
		cl.stigma.storeFileInStorage(file)
	}
}
