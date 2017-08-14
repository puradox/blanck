package main

import (
	"fmt"
	"log"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/linux/cmd"
	"github.com/puradox/blanck/service"
)

var DefaultServerOptions = []gatt.Option{
	gatt.LnxMaxConnections(1),
	gatt.LnxDeviceID(-1, true),
	gatt.LnxSetAdvertisingParameters(&cmd.LESetAdvertisingParameters{
		AdvertisingIntervalMin: 0x00f4,
		AdvertisingIntervalMax: 0x00f4,
		AdvertisingChannelMap:  0x7,
	}),
}

func main() {
	device, err := gatt.NewDevice(DefaultServerOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s", err)
	}

	// Regiser optional handlers.
	device.Handle(
		gatt.CentralConnected(func(c gatt.Central) { fmt.Println("Connect: ", c.ID()) }),
		gatt.CentralDisconnected(func(c gatt.Central) { fmt.Println("Disconnect: ", c.ID()) }),
	)

	// A mandatory handler for monitoring device state.
	device.Init(onStateChanged)
	select {}
}

func onStateChanged(device gatt.Device, state gatt.State) {
	fmt.Printf("State: %s\n", s)
	switch s {
	case gatt.StatePoweredOn:
		// Setup GATT service for Linux implementation.
		// OS X doesn't export the access of these services.
		device.AddService(service.NewGattService())

		// Human Interaction Device (HID) service.
		hid := service.NewHidService()
		device.AddService(hid)

		// A fake battery service for demo.
		battery := service.NewBatteryService()
		device.AddService(battery)

		// Device information service.
		info := service.NewInfoService()
		device.AddService(info)

		// Advertise device anem and service's UUIDs.
		device.AdvertiseNameAndServices("Blanck", []gatt.UUID{
			hid.UUID(),
			battery.UUID(),
			info.UUID(),
		})

		// Advertise as an OpenBeacon iBeacon.
		d.AdvertiseIBeacon(gatt.MustParseUUID("AA6062F098CA42118EC4193EB73CCEB6"), 1, 2, -59)

	default:
		// Nothing.
	}
}
