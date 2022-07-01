package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	temp "go-esp8266/protobuf/pb"

	// influx "github.com/influxdata/influxdb/client/v2"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"google.golang.org/protobuf/proto"
)

var (
	addr, network          string
	client                 influxdb2.Client
	dbAddr, dbUname, dbPwd string
)

func main() {
	// setup flags
	network = "tcp"
	flag.StringVar(&addr, "e", ":10101", "service endpoint")
	flag.StringVar(&dbAddr, "r", "http://localhost:8086", "influxDB endpoint")
	flag.StringVar(&dbUname, "u", "kent", "influxDB username")
	flag.StringVar(&dbPwd, "p", "00000000", "influxDB password")
	flag.Parse()

	// attempt to connect to influxDB
	// influxDB, err := influx.NewHTTPClient(influx.HTTPConfig{
	// 	Addr:     dbAddr,
	// 	Username: dbUname,
	// 	Password: dbPwd,
	// })
	// Create a client
	// You can generate an API Token from the "API Tokens Tab" in the UI
	client = influxdb2.NewClient("http://localhost:8086", "xbA_CSSwVL19PBuWo6R3FuLjPdmzYZTTS6mIWAHfoXffrqTt2mlOB7dNpyd4b7tzxhSFBcrdeSuFJJWoaSZtuQ==")
	// always close client at the end
	defer client.Close()
	// if err != nil {
	// 	log.Println("ERROR: failed to connect to influxDB, data will not be logged: ", err)
	// }

	ln, err := net.Listen(network, addr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer ln.Close()

	log.Printf("Service started: (%s) %s\n", network, addr)

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

	fmt.Printf("{DeviceID:%d, EventID:%d, Temp:%.2f, Humidity:%.2f%%, HeatIndex:%.2f}\n",
		e.GetDeviceId(),
		e.GetEventId(),
		e.GetTempCel(),
		e.GetHumidity(),
		e.GetHeatIdxCel(),
	)

	// go func(event temp.TempEvent) {
	if err := postEvent(e); err != nil {
		log.Println("ERROR: while posting event:", err)
	}
	// }(e)
}

func postEvent(e temp.TempEvent) error {
	if client != nil {
		writeAPI := client.WriteAPI("kent", "esp8266")
		log.Println("posting temp event to influxDB")
		// Create a new point batch
		// bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		// 	Database:  "esp8266",
		// 	Precision: "s",
		// })
		// if err != nil {
		// 	return err
		// }

		tags := map[string]string{
			"deviceId": fmt.Sprintf("%d", e.GetDeviceId()),
			"eventId":  fmt.Sprintf("%d", e.GetEventId()),
		}
		fields := map[string]interface{}{
			"temp":      e.GetTempCel(),
			"humidity":  e.GetHumidity(),
			"heatIndex": e.GetHeatIdxCel(),
		}
		p := influxdb2.NewPoint("sensor-temp", tags, fields, time.Now())

		// pt, err := influx.NewPoint("sensor-temp", tags, fields, time.Now())
		// if err != nil {
		// 	return err
		// }
		// bp.AddPoint(pt)
		// write point asynchronously
		writeAPI.WritePoint(p)
		// Write the batch
		// if err := db.Write(bp); err != nil {
		// 	return err
		// }
		writeAPI.Flush()
	}
	return nil
}
