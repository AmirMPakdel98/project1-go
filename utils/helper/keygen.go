package helper

import (
	"c-vod/utils/types"
	"crypto/rand"
	"encoding/hex"
)

func GenerateVideoKeys() (*types.VideoKeys, error) {

	randHex := [6]string{}

	for i := 0; i < len(randHex); i++ {
		str, err := generateRandomHex(types.VideoKeyLength)
		if err != nil {
			return nil, err
		}
		randHex[i] = str
	}

	return &types.VideoKeys{
		{Id: randHex[0], Value: randHex[1]},
		{Id: randHex[2], Value: randHex[3]},
		{Id: randHex[4], Value: randHex[5]},
	}, nil
}

func generateRandomHex(length int) (string, error) {
	// Calculate the number of bytes needed to generate the desired length
	byteLength := length / 2

	// Create a byte slice to store the random bytes
	randomBytes := make([]byte, byteLength)

	// Read random bytes from the crypto/rand package
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Convert the random bytes to a hexadecimal string
	hexString := hex.EncodeToString(randomBytes)

	// Return the generated hexadecimal string
	return hexString, nil
}
