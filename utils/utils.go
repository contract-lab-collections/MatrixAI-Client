package utils

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func ParseUUID(uuidStr string) ([16]byte, error) {
	var uuid [16]byte

	uuidParts := strings.Split(uuidStr, "-")
	if len(uuidParts) != 5 {
		return uuid, fmt.Errorf("invalid UUID format")
	}

	partLengths := []int{8, 4, 4, 4, 12}
	dstIndex := 0

	for i, part := range uuidParts {
		data, err := hex.DecodeString(part)
		if err != nil || len(data) != partLengths[i]/2 {
			return uuid, fmt.Errorf("invalid UUID format")
		}

		copy(uuid[dstIndex:], data)
		dstIndex += len(data)
	}

	return uuid, nil
}
