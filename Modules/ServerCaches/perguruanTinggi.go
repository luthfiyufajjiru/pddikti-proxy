package ServerCaches

import (
	"PDDiktiProxyAPI/Modules/PerguruanTinggi/DataTransferObjects"
	"encoding/json"
	"net/http"
	"sync"
)

var (
	ptMutex         *sync.RWMutex
	listUniversitas []DataTransferObjects.PerguruanTinggiDTO
)

func init() {
	ptMutex = &sync.RWMutex{}
}

func FetchUniversities() (result []DataTransferObjects.PerguruanTinggiDTO) {
	if x := len(listUniversitas); x == 0 {
		url := GetLoadPtUrl()
		client, err := http.Get(url)
		if err == nil {
			err = json.NewDecoder(client.Body).Decode(&result)
			client.Body.Close()
		}
	}
	return
}
