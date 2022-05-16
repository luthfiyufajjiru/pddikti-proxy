package ServerCaches

import (
	"PDDiktiProxyAPI/Modules/General/DataTransferObjects"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"sync"
)

type ProdiMap struct {
	IdSp        string
	DaftarProdi []DataTransferObjects.ProdiDTO
}

var (
	prodiMutex    *sync.RWMutex
	daftarProdi   []ProdiMap
	maxProdiQueue int
)

func init() {
	prodiMutex = &sync.RWMutex{}
	maxProdiQueue = 25
	daftarProdi = make([]ProdiMap, 0, maxProdiQueue)
}

func fetchProdi(id string) (results []DataTransferObjects.ProdiDTO) {
	id = base64.StdEncoding.EncodeToString([]byte(id))
	url := GetListProdiUrl(id)
	client, err := http.Get(url)
	if x, y := json.NewDecoder(client.Body).Decode(&results), client.Body.Close(); x != nil || y != nil || err != nil {
		results = nil
	}
	return
}

func GetProdi(id string) (results []DataTransferObjects.ProdiDTO) {
	prodiMutex.Lock()
	for _, val := range daftarProdi {
		if val.IdSp == id {
			results = val.DaftarProdi
		}
	}
	if results == nil {
		data := fetchProdi(id)
		if len(daftarProdi) >= maxProdiQueue {
			daftarProdi = daftarProdi[1:]
		}
		daftarProdi = append(daftarProdi, ProdiMap{IdSp: id, DaftarProdi: data})
		results = data
		data = nil
	}
	prodiMutex.Unlock()
	return
}
