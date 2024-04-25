// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package main implements a simple TURN server
package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/pion/stun/v2"
	"github.com/pion/turn/v3"
)

type stunLogger struct {
	net.PacketConn
}

func (s *stunLogger) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	if n, addr, err = s.PacketConn.ReadFrom(p); err == nil && stun.IsMessage(p) {
		msg := &stun.Message{Raw: p}
		if err = msg.Decode(); err != nil {
			return
		}

		log.Printf("Inbound[%v] STUN: %s \n", addr, msg.String())
	}
	return
}

func (s *stunLogger) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	if n, err = s.PacketConn.WriteTo(p, addr); err == nil && stun.IsMessage(p) {
		msg := &stun.Message{Raw: p}
		if err = msg.Decode(); err != nil {
			return
		}

		log.Printf("Outbound[%v] STUN: %s \n", addr,  msg.String())
	}

	return
}

func main() {
	publicIP := flag.String("public-ip", "", "IP Address that STUN can be contacted by.")
	port := flag.Int("port", 3478, "Listening port.")
	flag.Parse()

	if len(*publicIP) == 0 {
		log.Fatalf("'public-ip' is required")
	}

	// Create a UDP listener to pass into pion/turn
	// pion/turn itself doesn't allocate any UDP sockets, but lets the user pass them in
	// this allows us to add logging, storage or modify inbound/outbound traffic
	udpListener, err := net.ListenPacket("udp4", "0.0.0.0:"+strconv.Itoa(*port))
	if err != nil {
		log.Panicf("Failed to create STUN server listener: %s", err)
	}
	log.Printf("STUN server started %s:%d\n", *publicIP, *port)
	s, err := turn.NewServer(turn.ServerConfig{
		// PacketConnConfigs is a list of UDP Listeners and the configuration around them
		PacketConnConfigs: []turn.PacketConnConfig{
			{
				PacketConn: &stunLogger{udpListener},
			},
		},
	})
	if err != nil {
		log.Panic(err)
	}

	// Block until user sends SIGINT or SIGTERM
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	if err = s.Close(); err != nil {
		log.Panic(err)
	}
}
