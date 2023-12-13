package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"vlfr/config"
	"vlfr/match"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flagConfigPath := flag.String("config-path", "/etc/vlfr/vlfr.yaml", "Path to the config file")
	flagPort := flag.String("port", "8080", "Port to listen on")
	flagMetricsPort := flag.String("metrics-port", "8081", "Port to listen on for metrics")
	flagLogPath := flag.String("log-path", "/var/log/vlfr.log", "Path to the log file")
	flagLogLevel := flag.String("log-level", "info", "Log level to use")
	flag.Parse()

	cfg, err := config.GetConfig(*flagConfigPath)
	if err != nil {
		log.Fatalf("Error getting config: %v", err)
	}

	if *flagPort != "" {
		cfg.Port = *flagPort
	}
	if *flagMetricsPort != "" {
		cfg.MetricsPort = *flagMetricsPort
	}
	if *flagLogPath != "" {
		cfg.LogPath = *flagLogPath
	}
	if *flagLogLevel != "" {
		cfg.LogLevel = *flagLogLevel
	}

	if cfg.Port != "" {
		go func() {
			ln, err := net.Listen("tcp", ":"+cfg.Port)
			if err != nil {
				log.Fatalf("Error listening on port: %v", err)
			}
			defer ln.Close()

			for {
				conn, err := ln.Accept()
				if err != nil {
					log.Fatalf("Error accepting connection: %v", err)
				}

				go handleConnection(conn)
			}
		}()
	}

	go handleMetrics(cfg.MetricsPort)
	select {}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("Error reading from connection: %v", err)
	}

	log.Printf("Received %d bytes: %s", n, buf)
	// match the log line against the matchers
	matches, err := match.NewMatches()
	if err != nil {
		log.Fatalf("Error creating matches: %v", err)
	}

	category := matches.MatchLogLine(string(buf))
	log.Printf("Category: %s", category)

	// send the log line to the notifiers

}

func handleMetrics(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error listening on port: %v", err)
	}
	defer ln.Close()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.Serve(ln, nil))
}
