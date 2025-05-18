package SalseTracker

import "time"

func AutoSyncData() {
	for {
		lHrs := time.Now().Hour()
		if lHrs ==8 { // daily auto sync on mornung 8 am  only once
			SyncData()
			time.Sleep(time.Duration(60-time.Now().Minute()) * time.Minute)
		}
		time.Sleep(20* time.Minute)
	}

}	