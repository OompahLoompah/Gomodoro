package metrics

import (
	"os"
	"net"
	"time"
)

func push(message []byte, srv string) error{
	if srv == "" {
		addr := os.Getenv("GOMO_METRICS_SRV")
		if addr == "" {
			//TODO: log error sensibly instead of ignoring or mixing non-fatal and fatal errs in return. 
			return nil
		}
		port := os.Getenv("GOMO_METRICS_PORT")
		if port == "" {
			//TODO: log error sensibly instead of ignoring or mixing non-fatal and fatal errs in return. 
			return nil
		}
		srv = addr + ":" + port
	}
	conn, err := net.Dial("tcp",srv)
	if err != nil {
		return err
	}
	_, err = conn.Write(message)
	if err != nil {
		err = conn.Close()
		if err != nil{
			return err
		}
		return err
	}
	err = conn.Close()
	if err != nil {
		return err
	}
        return nil
}

func Log(measurement string, tags map[]string, fieldSet map[]string, t time.Time) error{
	msg := measurement

	for key, value := range tags {
		msg = msg + "," + key + "=" + value
	}

	msg = msg + " "

	for key, value := range fieldSet {
		msg = msg + key + "=" + value + ","
	}

	msg = msg[:len(msg)-1]

	if t == nil {
		t = time.Now()
	}

	msg = fmt.Sprintf("%s%d", msg, t.UnixNano())
	push(msg, "")
}
