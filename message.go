package zcl

type ZCLFrame struct {
	Header  Header
	Command interface{}
}
