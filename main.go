package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	temp "go-esp8266/protobuf/pb"

	"google.golang.org/protobuf/proto"
)

var (
	addr, network string
	// db influx.Client
	dbAddr, dbUname, dbPwd string
)

func main() {
	// setup flags
	network = "tcp"
	flag.StringVar(&addr, "e", ":10101", "service endpoint")
	// flag.StringVar(&dbAddr, "r", "http://localhost:8086", "influxDB endpoint")
	// flag.StringVar(&dbUname, "u", "admin", "influxDB username")
	// flag.StringVar(&dbPwd, "p", "admin", "influxDB password")
	flag.Parse()

	// // attempt to connect to influxdb
	// influxDB, err := influx.NewHTTPClient(influx.HTTPConfig{
	// 	Addr:     dbAddr,
	// 	Username: dbUname,
	// 	Password: dbPwd,
	// })
	// defer influxDB.Close()
	// db = influxDB

	ln, err := net.Listen(network, addr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer ln.Close()

	log.Printf("Temperator Service started: (%s) %s\n", network, addr)
	// connection loop
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			conn.Close()
			continue
		}
		log.Println("Connected to ", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		log.Println("INFO: closing connection")
		if err := conn.Close(); err != nil {
			log.Println("error closing connection", err)
		}
	}()

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	if n <= 0 {
		log.Println("no data received")
		return
	}

	var e temp.TempEvent
	if err := proto.Unmarshal(buf[:n], &e); err != nil {
		log.Println("failed to unmarshal:", err)
		return
	}

	fmt.Printf("{DeviceID:%d, EventID:%d, Temp:%.2f%%, HeatIndex:%.2f}\n",
		e.GetDeviceId(),
		e.GetEventId(),
		e.GetTempCel(),
		e.GetHeatIdxCel(),
	)

	// go func (event temp.TempEvent){}
	if err := postEvent(e); err != nil {
		log.Println("ERROR: while posting event:", err)
	}

}

func postEvent(e temp.TempEvent) error {
	tags := map[string]string{
		"deviceID": fmt.Sprintf("%d", e.GetDeviceId()),
		"eventID":  fmt.Sprintf("%d", e.GetEventId()),
	}
	fields := map[string]interface{}{
		"temp":      e.GetTempCel(),
		"humidity":  e.GetHumidity(),
		"heatIndex": e.GetHeatIdxCel(),
	}

	fmt.Printf("Tags:%v \nFields:%v", tags, fields)
	return nil
}
