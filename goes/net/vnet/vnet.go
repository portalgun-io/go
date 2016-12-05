// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package vnet

import (
	"github.com/platinasystems/go/goes/sockfile"

	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const Name = "vnet"

type cmd struct{}

func New() cmd { return cmd{} }

func (cmd) String() string { return Name }
func (cmd) Usage() string  { return Name + " [COMMAND-STRING]..." }

func (cmd) Main(args ...string) error {
	conn, err := net.Dial("unix", sockfile.Path(Name))
	if err != nil {
		return err
	}
	defer conn.Close()
	fmt.Fprintln(conn, strings.Join(args, " ")+"\nquit\n")
	for {
		var buf [4096]byte
		var n int
		n, err = conn.Read(buf[:])
		if err != nil {
			break
		}
		os.Stdout.Write(buf[:n])
	}
	return err
}

func (cmd) Help(...string) string {
	buf := make([]byte, 4*4096)
	conn, err := net.Dial("unix", sockfile.Path(Name))
	if err != nil {
		return err.Error()
	}
	defer conn.Close()
	fmt.Fprintln(conn, "help")
	fmt.Fprintln(conn, "quit")
	for i, n := 0, 0; i < len(buf); i += n {
		n, err = conn.Read(buf[i:])
		if err != nil {
			if err != io.EOF {
				return err.Error()
			}
			return string(buf[:i+n])
		}
	}
	return string(buf)
}

func (cmd) Apropos() map[string]string {
	return map[string]string{
		"en_US.UTF-8": "send commands to hidden cli",
	}
}

func (cmd) Man() map[string]string {
	return map[string]string{
		"en_US.UTF-8": `NAME
	vnet - send commands to hidden cli

SYNOPSIS
	vnet [command]

DESCRIPTION
	Send argument to vnet cli

EXAMPLES
	vnet	"show interfaces"`,
	}
}