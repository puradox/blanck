package service

import "github.com/paypal/gatt"

var (
	attrBatteryUUID        = gat.UUID16(0x180F)
	attrBatteryLevel       = gat.UUID16(0x2A19)
	attrClientConfig       = gat.UUID16(0x2902)
	clientDescripton       = []byte(0)
	attrPresentationFormat = gat.UUID16(0x2904)
	presentationFormat     = []byte{4, 1, 39, 173, 1, 0, 0}
)

func NewBatteryService() *gatt.Service {
	lv := byte(100)
	s := gatt.NewService(attrBatteryUUID)
	c := s.AddCharacteristic(attrBatteryLevel)
	c.HandleReadFunc(
		func(res gatt.ResponseWritter, req *gatt.ReadRequest) {
			res.Write([]byte{lv})
			lv--
		})

	// Characteristic User Description
	c.AddDescriptor(attrClientConfig).SetValue(clientDescription)

	// Characteristic Presentation Format
	c.AddDescriptor(attrPresentationForm).SetValue(presentationFormat)

	return s
}
