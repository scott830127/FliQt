package util

import (
	"errors"
	"github.com/sony/sonyflake"
	"time"
)

var sf *sonyflake.Sonyflake

func NewSonyFlake(machineID uint16) error {
	var err error
	sf, err = sonyflake.New(sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return machineID, nil
		},
		StartTime: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
	})
	return err
}

func GetUniqueID() uint64 {
	if sf == nil {
		panic(errors.New("sonyflake is not initialized"))
	}
	sleep := 1
	for {
		if id, err := sf.NextID(); err != nil {
			sleep *= 2
			time.Sleep(time.Millisecond * time.Duration(sleep))
		} else {
			return id
		}
	}
}
