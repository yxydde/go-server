package models

import "os"

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
	server = Server{hostName, version}

}

type Server struct {
	HostName string
	Version  string
}

func GetServer() (s *Server) {
	return &server
}
