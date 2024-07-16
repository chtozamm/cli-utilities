package main

import (
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"

	"github.com/chtozamm/cli-utilities/pkg/ip"
)

func main() {
	ipData, err := ip.GetIPInfo()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Initialize tabwriter to print aligned columns
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Use reflection to iterate over fields
	v := reflect.ValueOf(ipData)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name
		fieldValue := field.String()
		if fieldValue == "" {
			continue
		}
		// Print formatted output
		fmt.Fprintf(w, "%s\t%v\n", fieldName, fieldValue)
	}

	// Flush buffer
	if err := w.Flush(); err != nil {
		fmt.Println("Error flushing tabwriter:", err)
		os.Exit(1)
	}
}
