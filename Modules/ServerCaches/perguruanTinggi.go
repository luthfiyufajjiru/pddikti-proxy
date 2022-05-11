package ServerCaches

import (
	"PDDiktiProxyAPI/Modules/PerguruanTinggi/DataTransferObjects"
	"encoding/json"
	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	ptMutex      *sync.RWMutex
	universities []DataTransferObjects.PerguruanTinggiDTO
)

func init() {
	ptMutex = &sync.RWMutex{}
}

func fetchUniversities() (results []DataTransferObjects.PerguruanTinggiDTO) {
	url := GetLoadPtUrl()
	client, err := http.Get(url)
	if x, y := json.NewDecoder(client.Body).Decode(&results), client.Body.Close(); x != nil || y != nil || err != nil {
		results = nil
	}
	return
}

func GetUniversities() (results *[]DataTransferObjects.PerguruanTinggiDTO) {
	if universities == nil {
		_result := fetchUniversities()
		ptMutex.Lock()
		universities = _result
		ptMutex.Unlock()
		_result = nil
	}
	ptMutex.RLock()
	results = &universities
	ptMutex.RUnlock()
	return
}

func GetUniversity(input string) (result DataTransferObjects.PerguruanTinggiDTO) {
	universities := GetUniversities()
	defer func() {
		universities = nil
	}()
	for _, val := range *universities {
		if x, y := strings.ToLower(val.NamaPt), strings.ToLower(input); x == y {
			result = val
			return
		}
	}
	return
}

func cleanUniversities() {
	ptMutex.Lock()
	defer ptMutex.Unlock()
	universities = nil
}

func universitiesWatcher(app *fiber.App) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	startAt := time.Now().In(location)
	x := gocron.NewScheduler(location)
	x.Every(1).Second().Do(func() {
		if n := app.Server().GetOpenConnectionsCount(); n > 0 {
			startAt = time.Now()
		}
		if n := startAt.Sub(time.Now().In(location)).Minutes(); n > 30 {
			cleanUniversities()
			x.Stop()
		}
	})
	x.StartAsync()
}
