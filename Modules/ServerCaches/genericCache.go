package ServerCaches

import (
	"github.com/go-co-op/gocron"
	"log"
	"sync"
	"time"
)

func WatcherTask(referenceTime *time.Time, limit float64, Clean func(), mutex *sync.RWMutex, scheduler *gocron.Scheduler) func() {
	return func() {
		mutex.RLock()
		_started := referenceTime
		mutex.RUnlock()
		if n := time.Now().Sub(*_started).Minutes(); n > limit {
			log.Print("cleaning caches")
			Clean()
			log.Print("clean was successful")
			scheduler.Stop()
		}
	}
}

func Watcher(referenceTime *time.Time, mutex *sync.RWMutex, scheduler *gocron.Scheduler) error {
	mutex.Lock()
	*referenceTime = time.Now()
	if scheduler != nil && scheduler.IsRunning() {
		scheduler.Stop()
		scheduler.StartAsync()
	} else if scheduler != nil && !scheduler.IsRunning() {
		scheduler.StartAsync()
	}
	mutex.Unlock()
	return nil
}
