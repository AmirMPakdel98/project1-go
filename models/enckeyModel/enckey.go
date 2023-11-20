package enckeyModel

type Enckey struct {
	Id      uint32 `gorm:"primaryKey;autoIncrement"`
	File_id uint32
	Payload string
}
