package ServerCaches

import (
	"PDDiktiProxyAPI/Modules/PerguruanTinggi/DataTransferObjects"
	"encoding/json"
	"github.com/go-co-op/gocron"
	"log"
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

func universitiesWatcherTask() {
	univsSchedulerMutex.RLock()
	_started := &ptStartAt
	univsSchedulerMutex.RUnlock()
	if n := time.Now().Sub(*_started).Minutes(); n > 30 {
		cleanUniversities()
		univsScheduler.Stop()
	}
}

func universitiesWatcher() error {
	univsSchedulerMutex.Lock()
	{
		ptStartAt = time.Now()
		if univsScheduler == nil {
			univsScheduler = gocron.NewScheduler(location)
			_, err := univsScheduler.Every(1).Second().Do(universitiesWatcherTask)
			for err != nil {
				started := time.Now()
				log.Print(err)
				_, err = univsScheduler.Every(1).Second().Do(universitiesWatcherTask)
				if started.Sub(time.Now()).Seconds() > 2 {
					log.Printf("%s could not be resolved, please call administrator!", err)
					return err
				}
			}
			univsScheduler.StartAsync()
		}
	}
	univsSchedulerMutex.Unlock()
	if univsScheduler != nil && !univsScheduler.IsRunning() {
		univsSchedulerMutex.Lock()
		univsScheduler.StartAsync()
		univsSchedulerMutex.Unlock()
	}

	return nil
}

func GetUniversities() (results *[]DataTransferObjects.PerguruanTinggiDTO, err error) {
	defer func() {
		err = universitiesWatcher()
	}()
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
