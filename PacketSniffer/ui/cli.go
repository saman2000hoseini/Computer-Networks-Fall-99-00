package ui

import (
	"fmt"
	"github.com/saman2000hoseini/Computer-Networks-Fall-99-00/PacketSniffer/sniffer"
)

type CLI struct {
	PacketSniffer *sniffer.Sniffer
	Path          string
}

func (c *CLI) Start() {

	for {
		printMainMenu()

		var command byte
		fmt.Scanf("%d", &command)

		switch command {
		case 1:
			c.PacketSniffer.Finished = new(bool)
			*c.PacketSniffer.Finished = false

			dev := c.getDevice()
			if dev != nil {
				go c.PacketSniffer.Capture(c.Path, *dev)
			} else {
				break
			}
			fmt.Print("Enter any key to finish capturing")
			fmt.Scanf("%s")
			*c.PacketSniffer.Finished = true
			break
		case 2:
			break
		case 3:
			break
		default:
			return
		}
	}
}

func printMainMenu() {
	fmt.Print("1) Start Packet Sniffing\n" +
		"4) Exit\n")
}

func (c *CLI) getDevice() *string {
	devs := c.PacketSniffer.GetDevices()
	if devs != nil {
		for {
			for index, dev := range devs {
				fmt.Printf("%d) %s: %s\n", index+1, dev.Name, dev.Description)
			}
			fmt.Print("Please enter device number or 0 to return to main menu: ")

			var d int
			fmt.Scanf("%d", &d)

			if d <= len(devs) && d > 0 {
				return &devs[d-1].Name
			}

			if d == 0 {
				return nil
			}

			fmt.Println("Please enter valid command")
		}
	}

	fmt.Println("Couldn't fetch network interfaces")
	return nil
}
