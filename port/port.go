package port

import (
	"net"
	"strconv"
	"time"
)

func ScanPort(protocol, hostname string, port int) bool {
	adress := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, adress, 60*time.Second)

	if err != nil {
		return false
	}

	defer conn.Close()

	return true

}
