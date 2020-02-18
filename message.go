package zcl

import (
	"fmt"
	"github.com/shimmeringbee/bytecodec/bitbuffer"
)

type Message struct {
	Header  Header
	Command interface{}
}

type AttributeDataType byte
type AttributeIdentifier uint16

type AttributeDataTypeValue struct {
	DataType AttributeDataType
	Value    interface{}
}

func (a *AttributeDataTypeValue) Marshal(bb *bitbuffer.BitBuffer) error {
	if err := bb.WriteByte(byte(a.DataType)); err != nil {
		return err
	}

	switch a.DataType {
	case TypeNull:
		return nil
	default:
		return fmt.Errorf("unsupported ZCL type to marshal: %d", a.DataType)
	}
}

func (a *AttributeDataTypeValue) Unmarshal(bb *bitbuffer.BitBuffer) error {
	if dt, err := bb.ReadByte(); err != nil {
		return err
	} else {
		a.DataType = AttributeDataType(dt)
	}

	switch a.DataType {
	case TypeNull:
		return nil
	default:
		return fmt.Errorf("unsupported ZCL type to unmarshal: %d", a.DataType)
	}
}
