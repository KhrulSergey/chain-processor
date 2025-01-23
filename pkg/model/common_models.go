package model

import (
	"net"
)

type SenderInfo struct {
	IP     net.IP
	Device string
}
