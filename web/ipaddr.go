package web

import "fmt"

type IPAddr struct {
	Host string
	Port uint16
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("%s:%d", ip.Host, ip.Port)
}
