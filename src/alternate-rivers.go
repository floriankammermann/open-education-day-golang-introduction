// superserver project superserver.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
)

type River struct {
	Name       string
	LengthInKm int
}

var rivers = map[string]River{
	"rhein": {
		Name:       "Rhein",
		LengthInKm: 1238},
	"vorderrhein": {
		Name:       "Voderrhein",
		LengthInKm: 67},
}

func rname(w http.ResponseWriter, r *http.Request) {
	if val, ok := rivers[path.Base(r.URL.String())]; ok {
		fmt.Fprintf(w, "%s %d km", val.Name, val.LengthInKm)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "River not found.")
	}
}

func rlen(w http.ResponseWriter, r *http.Request) {
	minlengthStr := r.URL.Query().Get("min")
	maxlengthStr := r.URL.Query().Get("max")
	minlength, _ := strconv.Atoi(minlengthStr)
	maxlength, _ := strconv.Atoi(maxlengthStr)

	for _, river := range rivers {
		if river.LengthInKm > minlength && river.LengthInKm < maxlength {
			fmt.Fprintf(w, "Name: %s, Length in km: %d\n", river.Name, river.LengthInKm)
		}
	}
}

func main() {
	http.HandleFunc("/river/", rname)
	http.HandleFunc("/length/", rlen)
	log.Fatal(http.ListenAndServe(":8181", nil))
}
