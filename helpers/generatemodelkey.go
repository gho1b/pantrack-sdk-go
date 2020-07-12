package helpers

import (
	"crypto/sha512"
	"hash/crc64"

	"github.com/martinlindhe/base36"
)

// GenerateModelKey returns new identity model key
func GenerateModelKey(txID string, data []byte) (string, error) {
	hashSum := base36.DecodeToBytes(txID)
	tSize := base36.Decode(txID)
	tbl := crc64.MakeTable(tSize)
	crc := crc64.New(tbl)
	sha := sha512.New()
	sha.Write(data)

	hash := sha.Sum(hashSum)

	_, err := crc.Write(hash)
	if err != nil {
		return "", err
	}

	return string(crc.Sum(hashSum)), nil
}
