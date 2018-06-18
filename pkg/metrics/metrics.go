package metrics

import (
	"os"
	"net"
)

func Push(message []byte, srv string) error{
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
