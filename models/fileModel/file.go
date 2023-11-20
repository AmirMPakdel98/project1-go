package fileModel

type Type int
type Status int

const (
	VIDEO Type = iota
	OTHER
)

const (
	COMPLETED Status = iota
	DELETED
	UPLOADED
	NORMALIZED
	READY_TO_STORE
)

type File struct {
	Id        uint32
	Type      Type
	Ext       string
	Status    Status
	Size      int64
	Duration  int
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
	CreatedAt int64 `gorm:"autoCreateTime:milli"`
}
