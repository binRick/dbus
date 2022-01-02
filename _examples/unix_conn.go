package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s DBUS_TCP_ADDRESS\n", os.Args[0])
		os.Exit(1)
	}
	if err := run(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(2)
	}
	fmt.Println("ok")
}

func run(path string) error {
	conn, err := dbus.Connect(fmt.Sprintf("unix:path=%s", path))
	if err != nil {
		return err
	}
	defer conn.Close()
	fmt.Println("connected to D-Bus over socket")
	return nil
}
