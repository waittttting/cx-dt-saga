package sony_falke

import (
	"github.com/sony/sonyflake"
	"strconv"
)

func GetSequenceId() (string, error) {

	/*
		如果 SequenceID 有重复，可能是因为 node 的 IP 有重复字段 默认 的 machineID 是 IP 的低 16 位
		Default MachineID returns the lower 16 bits of the private IP address.
	*/
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	seqId, err := flake.NextID()
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(seqId, 10), nil
}
