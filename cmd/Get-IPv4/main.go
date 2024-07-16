package main

import (
	"fmt"
	"os"

	"github.com/chtozamm/cli-utilities/pkg/ip"
)

func main() {
	ip, err := ip.GetIPv4()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(ip)
}
