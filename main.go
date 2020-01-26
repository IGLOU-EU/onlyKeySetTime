// Implementation of Official OnlyKey "SetTime"
// From official python script at https://github.com/trustcrypto/python-onlykey

package main

import (
	"encoding/binary"
	"flag"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/karalabe/usb"
)

var (
	wg sync.WaitGroup

	flagVid = flag.String("vid", "", "VendorID of the device to which to connect.")
	flagPid = flag.String("pid", "", "ProductID of the device to which to connect.")
)

func onlyKeySetTime(devInfo usb.DeviceInfo) {
	dev, err := devInfo.Open()
	if err != nil {
		log.Fatalf("Failed to connect OpenDevices(): %v", err)
	}

	_, err = dev.Write(buildTimeData())
	if err != nil {
		log.Fatalf("Failed write Write(): %v", err)
	}

	dev.Close()
	defer wg.Done()
}

func buildTimeData() []byte {
	data := make([]byte, 64)

	messageSettime := []byte{228}
	messageHeader := []byte{255, 255, 255, 255}
	messageTime := make([]byte, 8)
	binary.BigEndian.PutUint64(messageTime, uint64(time.Now().Unix()))

	copy(data[0:], messageHeader)
	copy(data[4:], messageSettime)
	copy(data[5:], messageTime[4:])

	return data
}

func main() {
	//
	flag.Parse()

	//
	if *flagVid == "" || *flagPid == "" {
		log.Println("VID and PID cannot be empty !\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	//
	vid, _ := strconv.ParseUint(*flagVid, 16, 16)
	pid, _ := strconv.ParseUint(*flagPid, 16, 16)

	//
	devsInfo, err := usb.EnumerateHid(uint16(vid), uint16(pid))
	if err != nil {
		log.Fatalf("Failed to list HID devices which match the vendor and product id: %v", err)
	}
	if len(devsInfo) == 0 {
		log.Fatalf("No devices found matching VID %x and PID %x", vid, pid)
	}

	//
	for _, devInfo := range devsInfo {
		if devInfo.Serial == "1000000000" {
			if devInfo.UsagePage == uint16(65451) || devInfo.Interface == 2 {
				wg.Add(1)
				go onlyKeySetTime(devInfo)
			}
		} else {
			if devInfo.UsagePage == uint16(61904) || devInfo.Interface == 1 {
				wg.Add(1)
				go onlyKeySetTime(devInfo)
			}
		}
	}

	wg.Wait()
}
