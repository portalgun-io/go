// Copyright 2015-2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style license described in the
// LICENSE file.

package diag

import (
        "fmt"
        "net"
	"time"
	"github.com/platinasystems/goes/optional/gpio"
        "github.com/platinasystems/goes/i2c"
	"github.com/tatsushid/go-fastping"
)

func diagPing (address string, count int) bool {

	result := false
	dest := address
        pinger := fastping.NewPinger()
        pinger.Size = 64
        da, err := net.ResolveIPAddr("ip4:icmp", dest)
        if err != nil {
                if debug {
			fmt.Printf("Cannot resolve IP\n")
		}
        }
        pinger.AddIPAddr(da)
        pinger.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
                if debug {fmt.Printf("%d bytes from %s in %s\n", pinger.Size, addr.String(), rtt.String())}
		result = true
	}
        pinger.OnIdle = func() {}
	for i :=0; i < count; i++ {
		pinger.Run()
	}

	return result
}


// performs a read to i2c device address a on bus number b, returns true if read is successful 
func diagI2cPing (b uint8, a uint8, c uint8, count int) (bool, uint8) {
       var (
                bus i2c.Bus
                sd i2c.SMBusData
        )
        rw := i2c.Read
        op := i2c.ByteData

        err := bus.Open(int(b))
        if err != nil {
                if debug {fmt.Println(err)}
                return false,sd[0]
        }
        defer bus.Close()

        err = bus.ForceSlaveAddress(int(a))
        if err != nil {
                if debug {fmt.Println(err)}
                return false,sd[0]
        }
	for i :=0; i < count; i++ {
		err = bus.Do(rw, c, op, &sd)
		if err != nil {
		if debug {fmt.Println(err)}
			return false,sd[0]
		}
		if debug {fmt.Printf("%x.%02x.%02x = %02x\n", b, a, c, sd[0])}
	}
        return true,sd[0]
}

// write 1byte to bus b device address a (i.e. set mux channel)
func diagI2cWrite1Byte (b uint8, a uint8, c uint8) {
        var (
                bus i2c.Bus
                sd i2c.SMBusData
        )
        rw := i2c.Write
        op := i2c.Byte

        err := bus.Open(int(b))
        if err != nil {
                if debug {fmt.Println(err)}
        }
        defer bus.Close()

        err = bus.ForceSlaveAddress(int(a))
        if err != nil {
                if debug {fmt.Println(err)}
        }

        err = bus.Do(rw, c, op, &sd)
        if err != nil {
                if debug {fmt.Println(err)}
        }
        //if debug {fmt.Printf("%x.%02x.%02x = %02x\n", b, a, c, sd[0])}
}

func diagI2cWriteOffsetByte (b uint8, a uint8, c uint8, d uint8) {
        var (
                bus i2c.Bus
                sd i2c.SMBusData
        )
        rw := i2c.Write
        op := i2c.ByteData
	sd[0] = d

        err := bus.Open(int(b))
        if err != nil {
                if debug {fmt.Println(err)}
        }
        defer bus.Close()

        err = bus.ForceSlaveAddress(int(a))
        if err != nil {
                if debug {fmt.Println(err)}
        }

        err = bus.Do(rw, c, op, &sd)
        if err != nil {
                if debug {fmt.Println(err)}
        }
        //if debug {fmt.Printf("%x.%02x.%02x = %02x\n", b, a, c, sd[0])}
}

// returns value of BMC gpio name
func gpioGet (name string) bool {
        pv := gpio.PinValue{Name: name}
	pv.PinNum = gpio.Gpios[pv.Name]
        v, err := pv.PinNum.Value()
        if err != nil {
                fmt.Println(err)
        }
        return v
}

// sets BMC gpio name to value
func gpioSet (name string, value bool) {
        pv := gpio.PinValue{Name: name}
        pv.PinNum = gpio.Gpios[pv.Name]
        pv.PinNum.SetValue(value)
        return
}

// return true if test result r is within min and max limits
func CheckPass(r float32, min float32, max float32) bool {
        if r >= min && r <=max {
                return true
        } else {
                return false
        }
}