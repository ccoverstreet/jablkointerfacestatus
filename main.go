// HEllo
package jablkointerfacestatus

import (
	"fmt"
	"net/http"
)

func Initialize(input string) {
	fmt.Printf("INTERFACE STATUS\n")
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("High\n")
}
