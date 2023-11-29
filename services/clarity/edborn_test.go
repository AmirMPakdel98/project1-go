package clarity

import (
	"c-vod/utils/helper"
	"testing"
)

func TestInsertVideoKeys(t *testing.T) {

	var err error

	file_id := uint32(1)

	keys, _ := helper.GenerateVideoKeys()

	err = Edborn.insertVideoKeys(keys, file_id)

	if err != nil {
		t.Error(err)
	}
}
