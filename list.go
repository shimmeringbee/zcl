package zcl

import "github.com/shimmeringbee/zigbee"

/*
 * Zigbee Cluster List, as per ZCL Revision 6 (14 January 2016).
 * Downloaded From: https://zigbeealliance.org/developer_resources/zigbee-cluster-library/
 */

var ClusterList = map[zigbee.ClusterID]string{
	/* ZCL6: General */
	0x0000: "Basic",
	0x0001: "Power Configuration",
	0x0002: "Device Temperature Configuration",
	0x0003: "Identify",
	0x0004: "Groups",
	0x0005: "Scenes",
	0x0006: "On/Off",
	0x0007: "On/Off Switch Configuration",
	0x0008: "Level Control",
	0x0009: "Alarms",
	0x000a: "Time",
	0x000b: "RSSI Location",
	0x0b05: "Diagnostics",
	0x0020: "Poll Control",
	0x001a: "Power Profile",
	0x0b01: "Meter Identification",
	0x000c: "Analog Input (Basic)",
	0x000d: "Analog Output (Basic)",
	0x000e: "Analog Valve (Basic)",
	0x000f: "Binary Input (Basic)",
	0x0010: "Binary Output (Basic)",
	0x0011: "Binary Value (Basic)",
	0x0012: "Multistate Input (Basic)",
	0x0013: "Multistate Output (Basic)",
	0x0014: "Multistate Value (Basic)",
	/* ZCL6: Measurement and Sensing */
	0x0400: "Illuminance Measurement",
	0x0401: "Illuminance Level Sensing",
	0x0402: "Temperature Measurement",
	0x0403: "Pressure Measurement",
	0x0404: "Flow Measurement",
	0x0405: "Relative Humidity Measurement",
	0x0406: "Occupancy Sensing",
	0x0b04: "Electrical Measurement",
	/* ZCL6: Lighting */
	0x0300: "Color Control",
	0x0301: "Ballast Configuration",
	/* ZCL6: HVAC */
	0x0200: "Pump Configuration and Control",
	0x0201: "Thermostat",
	0x0202: "Fan Control",
	0x0203: "Dehumidification Control",
	0x0204: "Thermostat User Interface Configuration",
	/* ZCL6: Closures */
	0x0100: "Shade Configuration",
	0x0101: "Door Lock",
	0x0102: "Window Covering",
	/* ZCL6: Security and Safety */
	0x0500: "IAS Zone",
	0x0501: "IAS Ancillary Control Equipment",
	0x0502: "IAS Warning Devices",
	/* ZCL6: Protocol Interfaces */
	0x0016: "Partition",
	0x0600: "Generic Tunnel",
	0x0601: "BACnet Protocol Tunnel",
	0x0602: "Analog Input (BACnet Regular)",
	0x0603: "Analog Input (BACnet Extended)",
	0x0604: "Analog Output (BACnet Regular)",
	0x0605: "Analog Output (BACnet Extended)",
	0x0606: "Analog Value (BACnet Regular)",
	0x0607: "Analog Value (BACnet Extended)",
	0x0608: "Binary Input (BACnet Regular)",
	0x0609: "Binary Input (BACnet Extended)",
	0x060a: "Binary Output (BACnet Regular)",
	0x060b: "Binary Output (BACnet Extended)",
	0x060c: "Binary Value (BACnet Regular)",
	0x060d: "Binary Value (BACnet Extended)",
	0x060e: "Multistate Input (BACnet Regular)",
	0x060f: "Multistate Input (BACnet Extended)",
	0x0610: "Multistate Output (BACnet Regular)",
	0x0611: "Multistate Output (BACnet Extended)",
	0x0612: "Multistate Value (BACnet Regular)",
	0x0613: "Multistate Value (BACnet Extended)",
	0x0614: "11073 Protocol Tunnel",
	0x0615: "ISO7816 Tunnel",
	/* ZCL6: Smart Energy */
	0x0700: "Price",
	0x0701: "Demand Response and Load Control",
	0x0702: "Metering",
	0x0703: "Messaging",
	0x0704: "Tunneling",
	0x0800: "Key Establishment",
	/* ZCL6: Over-The-Air Upgrading */
	0x0019: "OTA Upgrade",
	/* ZCL6: Telecommunications */
	0x0900: "Information",
	0x0905: "Chatting",
	0x0904: "Voice Over ZigBee",
	/* ZCL6: Commissioning */
	0x0015: "Commissioning",
	0x1000: "Touchlink",
	/* ZCL6: Retail */
	0x0617: "Retail Tunnel Cluster",
	0x0022: "Mobile Device Configuration Cluster",
	0x0023: "Neighbor Cleaning Cluster",
	0x0024: "Nearest Gateway Cluster",
	/* ZCL6: Appliances */
	0x001b: "EN50523 Appliance Control",
	0x0b00: "EN50523 Appliance Identification",
	0x0b02: "EN50523 Appliance Events and Alerts",
	0x0b03: "EN50523 Appliance Statistics",
}
