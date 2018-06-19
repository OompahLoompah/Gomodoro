package timer

import (
	"time"
	json "encoding/json"

	metrics "github.com/OompahLoompah/Gomodoro/pkg/metrics"
)

type timerMetric struct {
	Seconds   int
	Cancelled bool
}

func Timer(t int, notifier func(), sendMetrics bool) error {
	time.Sleep(time.Duration(t) * time.Second)
	if notifier != nil {
		notifier()
	}

	if sendMetrics {
		w := timerMetric{t, false}
		m, err := json.Marshal(w)
		if err != nil {
			return err
		}
		err = metrics.Push(m, "")
		if err != nil {
			return err
		}
	}
	return nil
}
