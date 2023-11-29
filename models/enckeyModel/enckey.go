package enckeyModel

import (
	"c-vod/utils/types"
	"errors"
	"strings"
)

type Enckey struct {
	Id      uint32 `gorm:"primaryKey;autoIncrement"`
	File_id uint32
	Payload string
}

func (ek *Enckey) ParseToVideoKeys() (*types.VideoKeys, error) {

	ss := strings.Split(ek.Payload, ",")

	if len(ss) != 6 {
		return nil, errors.New("invalid payload")
	}

	keys := &types.VideoKeys{
		{Id: ss[0], Value: ss[1]},
		{Id: ss[2], Value: ss[3]},
		{Id: ss[4], Value: ss[5]},
	}

	return keys, nil
}
