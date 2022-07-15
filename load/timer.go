package load

import (
	"sync"
	"time"
)

var pillTimer *time.Timer
var pillMx sync.Mutex

func ProcessPill() {
	pillMx.Lock()
	UpdateEnemies(Enemies, EnemyStatusBlue)
	if pillTimer != nil {
		pillTimer.Stop()
	}
	pillTimer = time.NewTimer(time.Second * cfg.PillDuration)
	pillMx.Unlock()
	<-pillTimer.C

	pillMx.Lock()
	pillTimer.Stop()
	UpdateEnemies(Enemies, EnemyStatusNormal)
	pillMx.Unlock()
}

var enemiesStatusMx sync.RWMutex

func UpdateEnemies(Enemies []*Enemy, EnemyStatus EnemyStatus) {
	enemiesStatusMx.Lock()
	defer enemiesStatusMx.Unlock()
	for _, e := range Enemies {
		e.status = EnemyStatus
	}
}