package main

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
)

type River struct {
	Key        string
	Name       string
	LengthInKM int
}

func initRivers() map[string]River {

	rivers := make(map[string]River)
	rivers["rhein"] = River{Key: "rhein", Name: "Rhein", LengthInKM: 1238}
	rivers["vorderrhein"] = River{Key: "vorderrhein", Name: "Vorderrhein", LengthInKM: 67}
	rivers["hinterrhein"] = River{Key: "hinterrhein", Name: "Hinterrhein", LengthInKM: 57}
	rivers["albula"] = River{Key: "albula", Name: "Albula", LengthInKM: 36}
	rivers["alpenrhein"] = River{Key: "alpenrhein", Name: "Alpenrhein", LengthInKM: 90}
	rivers["landquart"] = River{Key: "landquart", Name: "Landquart", LengthInKM: 43}
	rivers["thur"] = River{Key: "thur", Name: "Thur", LengthInKM: 125}
	rivers["sitter"] = River{Key: "sitter", Name: "Sitter", LengthInKM: 70}
	rivers["toess"] = River{Key: "toess", Name: "Töss", LengthInKM: 57}
	rivers["birs"] = River{Key: "birs", Name: "Birs", LengthInKM: 73}
	rivers["aare"] = River{Key: "aare", Name: "Aare", LengthInKM: 295}
	rivers["kander"] = River{Key: "kander", Name: "Kander", LengthInKM: 44}
	rivers["simme"] = River{Key: "simme", Name: "Simme", LengthInKM: 53}
	rivers["saane"] = River{Key: "saane", Name: "Saane", LengthInKM: 128}
	rivers["zihl"] = River{Key: "zihl", Name: "Zihl", LengthInKM: 118}
	rivers["broye"] = River{Key: "broye", Name: "Broye", LengthInKM: 72}
	rivers["emme"] = River{Key: "emme", Name: "Emme", LengthInKM: 82}
	rivers["reuss"] = River{Key: "reuss", Name: "Reuss", LengthInKM: 164}
	rivers["kleineemme"] = River{Key: "kleineemme", Name: "KleineEmme", LengthInKM: 58}
	rivers["limmat"] = River{Key: "limmat", Name: "Limmat", LengthInKM: 36}
	rivers["linth"] = River{Key: "linth", Name: "Linth", LengthInKM: 104}
	rivers["sihl"] = River{Key: "sihl", Name: "Sihl", LengthInKM: 73}
	rivers["rhone"] = River{Key: "rhone", Name: "Rhone", LengthInKM: 264}
	rivers["vispa"] = River{Key: "vispa", Name: "Vispa", LengthInKM: 40}
	rivers["dranse"] = River{Key: "dranse", Name: "Dranse", LengthInKM: 44}
	rivers["doubs"] = River{Key: "doubs", Name: "Doubs", LengthInKM: 74}
	rivers["ticino"] = River{Key: "ticino", Name: "Ticino", LengthInKM: 91}
	rivers["maggia"] = River{Key: "maggia", Name: "Maggia", LengthInKM: 58}
	rivers["tresa"] = River{Key: "tresa", Name: "Tresa", LengthInKM: 15}
	rivers["inn"] = River{Key: "inn", Name: "Inn", LengthInKM: 104}

	return rivers
}

func main() {

	rivers := initRivers()
	http.HandleFunc("/river/", func(w http.ResponseWriter, r *http.Request) {

		minlengthStr := r.URL.Query().Get("minlength")
		maxlengthStr := r.URL.Query().Get("maxlength")

		if minlengthStr != "" && maxlengthStr != "" {
			fmt.Printf("minlength %s, maxlength %s\n", minlengthStr, maxlengthStr)
			minlength, _ := strconv.Atoi(minlengthStr)
			maxlength, _ := strconv.Atoi(maxlengthStr)

			matchingRivers := make([]River, 0)
			for _, river := range rivers {
				if river.LengthInKM > minlength && river.LengthInKM < maxlength {
					matchingRivers = append(matchingRivers, river)
				}
			}
			for _, river := range matchingRivers {
				fmt.Fprintf(w, "Name: %s, Length in km: %d\n", river.Name, river.LengthInKM)
			}
		} else {
			riverParam := path.Base(r.URL.String())
			river := rivers[riverParam]
			fmt.Fprintf(w, "Name: %s, Length in km: %d", river.Name, river.LengthInKM)
		}

	})

	http.ListenAndServe(":8088", nil)
}
