package client

import (
	"net"
	"strings"

	externalip "github.com/glendc/go-external-ip"
)

func cleanNames() {
	frameworkName = strings.ToLower(frameworkName)
	frameworkVersion = strings.ToLower(frameworkVersion)
	modelName = strings.ToLower(modelName)
	modelVersion = strings.ToLower(modelVersion)
}

func getTracerServerAddress(addr string) string {
	trimPrefix := func(s string) string {
		s = strings.TrimSpace(s)
		if strings.HasPrefix(s, "http://") {
			return strings.TrimPrefix(s, "http://")
		}
		if strings.HasPrefix(s, "https://") {
			return strings.TrimPrefix(s, "https://")
		}
		return s
	}
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return trimPrefix(addr)
	}
	host = trimPrefix(host)

	if host == "localhost" {
		return GetOutboundIP().String()
	}
	return host
}

// This does not give the outbound ip of this machine
func getHostIP() string {
	consensus := externalip.DefaultConsensus(nil, nil)
	ip, err := consensus.ExternalIP()
	if err != nil {
		return ""
	}
	return ip.String()
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

