package communicator

import (
	"context"
	"github.com/shimmeringbee/zcl"
	"github.com/shimmeringbee/zcl/commands/global"
	"github.com/shimmeringbee/zigbee"
)

type Communicator interface {
	RegisterMatch(match Match)
	UnregisterMatch(match Match)

	ProcessIncomingMessage(msg zigbee.NodeIncomingMessageEvent) error

	Request(ctx context.Context, address zigbee.IEEEAddress, requireAck bool, message zcl.Message) error
	RequestResponse(ctx context.Context, address zigbee.IEEEAddress, requireAck bool, message zcl.Message) (zcl.Message, error)

	ReadAttributes(ctx context.Context, ieeeAddress zigbee.IEEEAddress, requireAck bool, cluster zigbee.ClusterID, code zigbee.ManufacturerCode, sourceEndpoint zigbee.Endpoint, destEndpoint zigbee.Endpoint, transactionSequence uint8, attributes []zcl.AttributeID) ([]global.ReadAttributeResponseRecord, error)
	WriteAttributes(ctx context.Context, ieeeAddress zigbee.IEEEAddress, requireAck bool, cluster zigbee.ClusterID, code zigbee.ManufacturerCode, sourceEndpoint zigbee.Endpoint, destEndpoint zigbee.Endpoint, transactionSequence uint8, attributes map[zcl.AttributeID]zcl.AttributeDataTypeValue) ([]global.WriteAttributesResponseRecord, error)
	ConfigureReporting(ctx context.Context, ieeeAddress zigbee.IEEEAddress, requireAck bool, cluster zigbee.ClusterID, code zigbee.ManufacturerCode, sourceEndpoint zigbee.Endpoint, destEndpoint zigbee.Endpoint, transactionSequence uint8, attributeId zcl.AttributeID, dataType zcl.AttributeDataType, minimumReportingInterval uint16, maximumReportingInterval uint16, reportableChange interface{}) error
}
