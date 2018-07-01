package metrics

import (
	"fmt"
	"net"
	"os"
	"time"
)

var path string = "/Users/seanheuer/.gomodoro_cache"

func push(msg []byte, srv string) error{
	fmt.Println("Message is: " + string(msg))
	if srv == "" {
		addr := os.Getenv("GOMO_METRICS_SRV")
		if addr == "" {
			//TODO: log error sensibly instead of ignoring or mixing non-fatal and fatal errs in return. 
			cache(msg)
			return nil
		}
		port := os.Getenv("GOMO_METRICS_PORT")
		if port == "" {
			//TODO: log error sensibly instead of ignoring or mixing non-fatal and fatal errs in return. 
			cache(msg)
			return nil
		}
		srv = addr + ":" + port
	}
	conn, err := net.Dial("tcp",srv)
	if err != nil {
		cache(msg)
		return err
	}
	_, err = conn.Write(msg)
	if err != nil {
		cache(msg)
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
	fmt.Println("Logged metric")
        return nil
}

func Log(measurement string, tags map[string]string, fieldSet map[string]string, t *time.Time) error{
	msg := measurement

	for key, value := range tags {
		msg = msg + "," + key + "=" + value
	}

	msg = msg + " "

	for key, value := range fieldSet {
		msg = msg + key + "=" + value + ","
	}

	msg = msg[:len(msg)-1] + " "

	if t == nil {
		*t = time.Now()
	}

	msg = fmt.Sprintf("%s%d", msg, t.UnixNano())
	fmt.Println("Message is: " + string(msg))
	err := push([]byte(msg), "")

	if err != nil{
		return err
	}
	return nil
}

func cache(msg []byte) error {

	_, err := os.Stat(path)
	if err != nil {
		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0700)
		if err != nil {
			fmt.Println("1")
			fmt.Println(err)
			return err
		}
		f.Write(msg)
	} else {
		f, err := os.OpenFile(path, os.O_RDWR, 0000) //file perms shouldn't matter here since file exists
		if err != nil {
			fmt.Println("2")
			fmt.Println(err)
			return err
		}
		_, err = f.Seek(0,2)
		if err != nil {
			return err
		}
		_, err = f.Write([]byte("\n")) //TODO really need to prepend msg with \n but not sure how to properly do it right now
		if err != nil {
			return err
		}
		_, err = f.Write(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

//func flushCache() error {
	//f, err := os.Open(path)
	//if err != nil {
	//	return err
	//}
	//msg = f

	//need a read that reads to \n and nothing more... 
//}
