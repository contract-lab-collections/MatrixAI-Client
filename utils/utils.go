package utils

import (
	"MatrixAI-Client/pattern"
	"encoding/hex"
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"strings"
)

func ParseMachineUUID(uuidStr string) (pattern.MachineUUID, error) {
	var uuid [16]byte
	var machineUUID pattern.MachineUUID

	uuidParts := strings.Split(uuidStr, "-")
	if len(uuidParts) != 5 {
		return machineUUID, fmt.Errorf("invalid UUID format")
	}

	partLengths := []int{8, 4, 4, 4, 12}
	dstIndex := 0

	for i, part := range uuidParts {
		data, err := hex.DecodeString(part)
		if err != nil || len(data) != partLengths[i]/2 {
			return machineUUID, fmt.Errorf("invalid UUID format")
		}

		copy(uuid[dstIndex:], data)
		dstIndex += len(data)
	}

	for i := 0; i < len(machineUUID); i++ {
		machineUUID[i] = types.U8(uuid[i])
	}

	return machineUUID, nil
}
