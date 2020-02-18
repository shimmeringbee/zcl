package zcl

type Message struct {
	Header  Header
	Command interface{}
}
