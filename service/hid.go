package service

import (
	"log"

	"github.com/paypal/gatt"
	"github.com/tarm/serial"
)

var (
	attrHidUUID           = gatt.UUID16(0x1812)
	attrReportMap         = gatt.UUID16(0x2A4B)
	attrKeyboardInReport  = gatt.UUID16(0x2A22)
	attrKeyboardOutReport = gatt.UUID16(0x2A32)
	attrHidInformation    = gatt.UUID16(0x2A4A)

	attrHidReportType = gatt.UUID16(0x22)

	txd = rpio.Pin(14)
	rxd = rpio.Pin(15)
)

func NewHidService() *gatt.Service {
	service := gatt.NewService(HidDeviceServiceUUID)

	// Open the UART serial port
	config := &serial.Config{Name: "dev/ttyAMA0", Baud: 9600}
	serial, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}

	//service.AddCharacteristic(attrReportMap).SetValue(0)
	service.AddCharacteristic(attrKeyboardInReport).HandleReadFunc(
		func(res gatt.ResponseWriter, req *gatt.ReadRequest) {
			// Read from TXD
			buf := make([]byte, 16)
			n, err = serial.Read(buf)
			if err != nil {
				log.Fatal(err)
			}

			// Log for testing
			log.Println(buf)

			// Send over BlueTooth
			n, err = res.Write(buf)
			if err != nil {
				log.Fatal(err)
			}
		})
	//service.AddCharacteristic(attrKeyboardOutReport).HandleReadFunc(
	//service.AddCharacteristic(attrHidInformation).etValue(0)

	return service
}
