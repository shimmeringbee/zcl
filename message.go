package zcl

type Message struct {
	Header  Header
	Command interface{}
}

type AttributeIdentifier uint16