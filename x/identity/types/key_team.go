package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TeamKeyPrefix is the prefix to retrieve all Team
	TeamKeyPrefix = "Team/value/"
)

// TeamKey returns the store key to retrieve a Team from the index fields
func TeamKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
