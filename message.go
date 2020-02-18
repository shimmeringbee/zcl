package zcl

import (
	"errors"
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
	case TypeData8:
		if data, ok := a.Value.([]byte); !ok {
			return errors.New("could not cast value")
		} else {
			return marshallData(bb, data, 1)
		}
	default:
		return fmt.Errorf("unsupported ZCL type to marshal: %d", a.DataType)
	}
}

func marshallData(bb *bitbuffer.BitBuffer, data []byte, size int) error {
	if len(data) != size {
		return fmt.Errorf("data array provided does not match output size")
	}

	for i := 0; i < size; i++ {
		if err := bb.WriteByte(data[i]); err != nil {
			return err
		}
	}

	return nil
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
	case TypeData8:
		return nil
	default:
		return fmt.Errorf("unsupported ZCL type to unmarshal: %d", a.DataType)
	}
}
