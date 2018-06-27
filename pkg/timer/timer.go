package timer

import (
	"fmt"
	"time"

	metrics "github.com/OompahLoompah/Gomodoro/pkg/metrics"
)

func Timer(seconds int, notifier func(), sendMetrics bool, userTag string) error {
	time.Sleep(time.Duration(seconds) * time.Second)
	if notifier != nil {
		notifier()
	}

	if sendMetrics {
		measurement := "pomodoro"
		tag := make(map[string]string)
		tag["SessionType"] = userTag

		metric := make(map[string]string)
		metric["Seconds"] = fmt.Sprintf("%d", seconds)
		metric["Completed"] = "true" //This is a bad assumption but we're using it for now
		t := time.Now()
		metrics.Log(measurement, tag, metric, &t)
	}
	return nil
}
