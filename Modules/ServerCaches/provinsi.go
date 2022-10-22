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
	provinsiMutex          *sync.RWMutex
	provinsiStartAt        time.Time
	provinsiScheduler      *gocron.Scheduler
	provinsiSchedulerMutex *sync.RWMutex
	daftarProvinsi         []DataTransferObjects.ProvinsiDTO
)

func init() {
	provinsiMutex = &sync.RWMutex{}
	provinsiSchedulerMutex = &sync.RWMutex{}
	provinsiStartAt = time.Now().In(location)
	provinsiScheduler = gocron.NewScheduler(location)
	_, _ = provinsiScheduler.Every(1).Second().Do(WatcherTask(&ptStartAt, 30, cleanDaftarProvinsi, provinsiSchedulerMutex, provinsiScheduler))
}

func fetchProvinsi() (results []DataTransferObjects.ProvinsiDTO) {
	url := GetProvinsiUrl()
	client, err := http.Get(url)
	if x, y := json.NewDecoder(client.Body).Decode(&results), client.Body.Close(); x != nil || y != nil || err != nil {
		results = nil
	}
	return
}

func cleanDaftarProvinsi() {
	provinsiMutex.Lock()
	defer provinsiMutex.Unlock()
	daftarProvinsi = nil
}

func provinsiWatcher() error {
	err := Watcher(&provinsiStartAt, provinsiSchedulerMutex, provinsiScheduler)
	return err
}

func GetProvinsi() (results *[]DataTransferObjects.ProvinsiDTO, err error) {
	defer func() {
		_ = provinsiWatcher()
	}()
	provinsiMutex.RLock()
	if daftarProvinsi == nil {
		provinsiMutex.RUnlock()
		provinsiMutex.Lock()
		daftarProvinsi = fetchProvinsi()
		results = &daftarProvinsi
		provinsiMutex.Unlock()
	} else if daftarProvinsi != nil {

	}
	return
}
