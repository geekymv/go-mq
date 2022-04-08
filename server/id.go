package server

import "github.com/sony/sonyflake"

type ID struct {
	sf *sonyflake.Sonyflake
}

func NewID(nodeID uint16) *ID {
	st := sonyflake.Settings{MachineID: func() (uint16, error) {
		return nodeID, nil
	}}
	return &ID{sf: sonyflake.NewSonyflake(st)}
}

func (i *ID) NextID() (uint64, error) {
	return i.sf.NextID()
}
