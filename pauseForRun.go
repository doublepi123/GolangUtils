package GolangUtils

import "time"

func PauseForRun() {
	for {
		time.Sleep(time.Second)
	}
}
