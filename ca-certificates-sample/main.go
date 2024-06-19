package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	_, err := http.Head("https://support-lab-status.cfapps-01.slot-34.tanzu-gss-labs.vmware.com/")
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(2)
	}
	fmt.Println("SUCCESS!")
}
