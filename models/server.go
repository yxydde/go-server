package models

import (
	"os"
	"time"
)

var server Server

func init() {

	hostName, err := os.Hostname()
	if err != nil {
		hostName = "Unknown"
	}
	version := os.Getenv("Version")
	if version == "" {
		version = "V1.0"
	}
	server = Server{hostName, version, time.Now().Format("2006-01-02 15:04:05")}

}

type Server struct {
	HostName string
	Version  string
	NowTime  string
}

func GetServer() (s *Server) {
	return &server
}
