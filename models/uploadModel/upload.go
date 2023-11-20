package uploadModel

import fileModel "c-vod/models/fileModel"

type Status int

const (
	NOT_UPLOADED Status = iota
	UPLOADED
	Expired
)

type Upload struct {
	Id        uint32 `gorm:"primaryKey;autoIncrement"`
	Token     string
	Type      fileModel.Type
	Ext       string
	Size      int64
	Status    Status
	ExpiresAt int64
}
