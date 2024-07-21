package concurrencyUtils

//TODO: move this to internal?

import (
	"log"
	"time"
)

// TODO: come up with a more description nave
type DataState struct {
	ElementIdentifier string
	IsDone            bool
}

func NewDataState(elementIdentifier string) *DataState {
	return &DataState{elementIdentifier, false}
}

// StateMonitor() maintains a map of [state:T] = completionState:K, and logs the map at every logInterval

func StateMonitor(logInterval time.Duration) <-chan DataState {
	updates := make(chan DataState)
	stateMap := make(map[string]bool)

	ticker := time.NewTicker(logInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				LogState(stateMap)
			case s := <-updates: //if updates channel receives an update
				stateMap[s.ElementIdentifier] = s.IsDone
			}
		}
	}()
	return updates
}

func LogState(stateMap map[string]bool) {
	log.Println("*** STATE MONITOR ***")
	for k, v := range stateMap {
		log.Printf("%s: %t\n", k, v)
	}
}
