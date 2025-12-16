package chlogger

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func SaveToCH(eventCode uint8, description string) {
	conn, err := connect()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var eventTmpStmp time.Time = time.Now()
	var eventDate time.Time = time.Now().Truncate(24 * time.Hour)

	err = conn.Exec(context.Background(), "INSERT INTO METRICS.LOG_MESSAGES(UUID, EVENT_ID, DESCRIPTION, UPDSTMP, DATE) VALUES(generateUUIDv4(), ?, ?, ?, ?)", eventCode, description, eventTmpStmp, eventDate)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	conn.Close()
}

func ShowLogs(top uint8) {
	conn, err := connect()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	rows, err := conn.Query(ctx, "SELECT UUID, EVENT_ID, DESCRIPTION, UPDSTMP, DATE FROM METRICS.LOG_MESSAGES LIMIT ?", top)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	for rows.Next() {
		var UUID string
		var eventId uint8
		var description string
		var updStmp time.Time
		var date time.Time
		if err := rows.Scan(&UUID, &eventId, &description, &updStmp, &date); err != nil {
			log.Fatal(err)
		}
		log.Printf("ID %s, EVENT_ID %d, DESCRIPTION %s, UPDSTMP %s, DATE %s", UUID, eventId, description, updStmp, date)
	}
	rows.Close()

}

func connect() (driver.Conn, error) {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"localhost:9000"},
			Auth: clickhouse.Auth{
				Database: "METRICS",
				Username: "default",
				Password: "123",
			},
			ClientInfo: clickhouse.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: "an-example-go-client", Version: "0.1"},
				},
			},
			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
		})
	)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}

	return conn, nil
}
