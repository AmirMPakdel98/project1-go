package helper

import (
	"c-vod/utils/types"
	"testing"
)

func TestKeygen(t *testing.T) {

	keys, err := GenerateVideoKeys()

	t.Log(keys)

	if err != nil {
		t.Error(err)
	}

	if len(*keys) != 3 {
		t.Error("keys length is not 3")
	}

	for _, k := range *keys {
		if len(k.Id) != types.VideoKeyLength {
			t.Errorf("keys.Id length is not %d; => %s", types.VideoKeyLength, k.Id)
		}
		if len(k.Value) != types.VideoKeyLength {
			t.Errorf("key.Value length is not %d; => %s", types.VideoKeyLength, k.Value)
		}
	}
}
