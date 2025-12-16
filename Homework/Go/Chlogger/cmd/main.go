package main

import (
	"Chlogger/pkg/chlogger"
	"flag"
	"fmt"
)

var (
	saveToCH  = flag.Bool("save", false, "Сохранить данные в ClickHouse")
	showLogs  = flag.Bool("show", false, "Показать логи из ClickHouse")
	recordID  = flag.Int("id", 1, "ID записи для сохранения")
	message   = flag.String("msg", "test", "Сообщение для сохранения")
	logsCount = flag.Int("count", 5, "Количество логов для показа")
)

func main() {
	flag.Parse()
	id := uint8(*recordID)
	count := uint8(*logsCount)
	if *saveToCH {
		fmt.Printf("Сохранение в CH: ID=%d, Message=%s\n", *recordID, *message)
		chlogger.SaveToCH(id, *message)
	}

	if *showLogs {
		fmt.Printf("Показ %d последних логов:\n", *logsCount)
		chlogger.ShowLogs(count)
	}

	if !*saveToCH && !*showLogs {
		fmt.Println("Использование:")
		fmt.Println("  --save    - сохранить данные в ClickHouse")
		fmt.Println("  --id      - ID записи (по умолчанию: 1)")
		fmt.Println("  --msg     - сообщение (по умолчанию: test)")
		fmt.Println("-------------------------------------------")
		fmt.Println("  --show    - показать логи из ClickHouse")
		fmt.Println("  --count   - количество логов (по умолчанию: 5)")
	}
}
