package service

import "github.com/paypal/gatt"

var (
	attrDeviceInfoUUID = gatt.UUID16(0x180A)

	attrManufacturerNameUUID = gatt.UUID16(0x2A29)
	attrModelNumberUUID      = gatt.UUID16(0x2A24)
)

// https://developer.bluetooth.org/gatt/characteristics/Pages/CharacteristicViewer.aspx?u=org.bluetooth.characteristic.gap.appearance.xml
var gapCharAppearanceGenericComputer = []byte{0x00, 0x80}

// NOTE: OS X provides GAP and GATT services, and they can't be customized.
// For Linux/Embedded, however, this is something we want to fully control.
func NewInfoService() *gatt.Service {
	s := gatt.NewService(attrDeviceInfoUUID)
	s.AddCharacteristic(attrManufacturerNameUUID).SetValue([]byte("OLBK"))
	s.AddCharacteristic(attrModelNumberUUID).SetValue([]byte("Planck Keyboard"))
	return s
}
