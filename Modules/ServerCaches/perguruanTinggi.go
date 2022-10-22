package ServerCaches

import (
	"PDDiktiProxyAPI/Modules/General/DataTransferObjects"
	"encoding/json"
	"github.com/go-co-op/gocron"
	"net/http"
	"sync"
	"time"
)

var (
	ptMutex             *sync.RWMutex
	univsSchedulerMutex *sync.RWMutex
	univsScheduler      *gocron.Scheduler
	universities        []DataTransferObjects.PerguruanTinggiDTO
	location            *time.Location
	ptStartAt           time.Time
)

func init() {
	ptMutex = &sync.RWMutex{}
	univsSchedulerMutex = &sync.RWMutex{}
	location, _ = time.LoadLocation("Asia/Jakarta")
	ptStartAt = time.Now().In(location)
	univsScheduler = gocron.NewScheduler(location)
	_, _ = univsScheduler.Every(1).Second().Do(WatcherTask(&ptStartAt, 30, cleanUniversities, univsSchedulerMutex, univsScheduler))
}

func fetchUniversities() (results []DataTransferObjects.PerguruanTinggiDTO) {
	url := GetLoadPtUrl()
	client, err := http.Get(url)
	if x, y := json.NewDecoder(client.Body).Decode(&results), client.Body.Close(); x != nil || y != nil || err != nil {
		results = nil
	}
	return
}

func cleanUniversities() {
	ptMutex.Lock()
	defer ptMutex.Unlock()
	universities = nil
}

func universitiesWatcher() error {
	err := Watcher(&ptStartAt, univsSchedulerMutex, univsScheduler)
	return err
}

func GetUniversities() (results *[]DataTransferObjects.PerguruanTinggiDTO, err error) {
	defer func() {
		err = universitiesWatcher()
	}()
	ptMutex.RLock()
	if universities == nil {
		ptMutex.RUnlock()
		ptMutex.Lock()
		universities = fetchUniversities()
		results = &universities
		ptMutex.Unlock()
	} else if universities != nil {
		results = &universities
		ptMutex.RUnlock()
	}
	return
}
