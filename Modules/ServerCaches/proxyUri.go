package ServerCaches

import (
	"fmt"
	"sync"
)

var (
	uriMutex   *sync.RWMutex = &sync.RWMutex{}
	baseApiUri string        = "https://api-frontend.kemdikbud.go.id"
)

func SetBaseApiUri(address string) {
	defer uriMutex.Unlock()
	uriMutex.Lock()
	baseApiUri = address
}

func GetBaseUri() (result *string) {
	defer uriMutex.RUnlock()
	uriMutex.RLock()
	result = &baseApiUri
	return
}

func GetLoadPtUrl() (result string) {
	result = fmt.Sprintf("%s/loadpt", *GetBaseUri())
	return
}

func GetListProdiUrl(id string) (result string) {
	result = fmt.Sprintf("%s/v2/detail_pt_prodi/%s", *GetBaseUri(), id)
	return
}

func GetSearchUrl(query string, hitMhs bool) (result string) {
	var hit string
	switch hitMhs {
	case true:
		hit = "hit_mhs"
	default:
		hit = "hit"
	}

	result = fmt.Sprintf("%s/%s/%s", *GetBaseUri(), hit, query)
	return
}
