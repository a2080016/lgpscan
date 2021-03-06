package lgpparser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/a2080016/lgpscan/internal/cfg"
	"github.com/a2080016/lgpscan/internal/logger"
	"github.com/a2080016/lgpscan/internal/status"
	"github.com/a2080016/lgpscan/pkg/lgfparser"
)

var LgfInfo lgfparser.LgfInfoType
var CurrInfoBase string

type event struct {
	datetime   time.Time // 1. Дата и время события в формате "20060102150405"
	tr_status  string    // 2. Статус транзакции
	tr_time    string    // 3-1. Время транзакции
	tr_offset  string    // 3-2. Смещение транзакции
	user       string    // 4. Пользователь
	pc         string    // 5. Компьютер
	app        string    // 6. Приложение
	conn       int       // 7. Соединение
	event      string    // 8. Событие
	importance string    // 9. Важность
	comment    string    // 10. Комментарий
	metadata   string    // 11. Метаданные
	data       string    // 12. Данные
	dataView   string    // 13. Представление данных
	server     string    // 14. Сервер
	mainPort   int       // 15. Основной порт
	secondPort int       // 16. Вспомогательный порт
	session    int       // 17. Сеанс
}

func ScanLgpFiles(infoBase string) {

	CurrInfoBase = infoBase
	lgpPath := cfg.Config.InfoBases[infoBase].LgpPath

	currFile := ""
	currPos := int64(0)

	status.ReadStatus()

	_, exists := status.CurrStatus.InfoBases[infoBase]

	if exists == true {
		currFile = status.CurrStatus.InfoBases[infoBase].CurrentFile
		currPos = status.CurrStatus.InfoBases[infoBase].CurrentPos
	} else {

		lgpFiles, err := ioutil.ReadDir(lgpPath)

		if err != nil {
			logger.ErrLog.Fatal(err)
		}

		for _, lgpFile := range lgpFiles {

			lgpFileName := lgpFile.Name()

			if filepath.Ext(lgpFileName) == ".lgp" {

				currFile = lgpFile.Name()
				currPos = 0
				break
			}

		}

		status.CurrStatus.InfoBases[infoBase] = struct {
			CurrentFile string "yaml:\"CurrentFile\""
			CurrentPos  int64  "yaml:\"CurrentPos\""
		}{
			CurrentFile: currFile,
			CurrentPos:  currPos,
		}

		status.WriteStatus()

	}

	// Парсим LGF
	LgfInfo = lgfparser.ParseLgf(lgpPath + `\1Cv8.lgf`)
	lgpFiles, err := ioutil.ReadDir(lgpPath)

	if err != nil {
		logger.ErrLog.Fatal(err)
	}

	for _, lgpFile := range lgpFiles {

		lgpFileName := lgpFile.Name()

		if filepath.Ext(lgpFileName) == ".lgp" && lgpFileName == currFile {

			logger.InfLog.Println(lgpFile.Name(), lgpFile.IsDir())
			ScanLgpFile(lgpPath+`\`+lgpFile.Name(), currPos)
		}

	}

}

func ScanLgpFile(lgpFileName string, pos int64) {

	lgpFile, err := os.Open(lgpFileName)
	if err != nil {
		panic(err)
	}
	defer lgpFile.Close()

	eventBuffer := bytes.Buffer{}
	eventsList := []event{}

	eventBegin := false
	eventEnd := false

	lgpFile.Seek(pos, os.SEEK_SET)

	lgpReader := bufio.NewReader(lgpFile)

	var line []byte
	fPos := int(pos) // or saved position

	for i := 1; ; i++ {
		line, err = lgpReader.ReadBytes('\n')

		if err != nil {
			break
		}
		fPos += len(line)

		NowString := string(line[0 : len(line)-2])

		if len(NowString) > 0 && NowString[0] == '{' {
			eventBegin = true
		} else if NowString == "}," {
			eventEnd = true
		} else {
		}

		if eventBegin {
			eventBuffer.WriteString(NowString)
		}

		if eventBegin && eventEnd {

			eventsList = append(eventsList, parseEventString(eventBuffer.String()))

			eventBuffer.Reset()
			eventBegin = false
			eventEnd = false
		}

	}

	status.CurrStatus.InfoBases[CurrInfoBase] = struct {
		CurrentFile string "yaml:\"CurrentFile\""
		CurrentPos  int64  "yaml:\"CurrentPos\""
	}{
		CurrentFile: filepath.Base(lgpFileName),
		CurrentPos:  int64(fPos),
	}

	status.WriteStatus()

	if err != io.EOF {
		log.Fatal(err)
	}

}

func parseEventString(eventString string) event {

	var e event

	newBlock := false

	currBlock := 1
	currBlockBegin := 1
	currBlockEnd := 15

	openBrackets := 0

	for i := 15; i < len(eventString); i++ {

		if newBlock {
			newBlock = false
			currBlock++
			currBlockBegin = i
		}

		if eventString[i] == ',' && openBrackets == 0 {
			currBlockEnd = i
			e = addBlock(e, currBlock, eventString[currBlockBegin:currBlockEnd])
			newBlock = true
		} else if eventString[i] == '{' {
			openBrackets++
		} else if eventString[i] == '}' {
			openBrackets--
		}

	}

	if cfg.Config.InfoBases["ds_estate"].Show {
		printEvent(e)
	}

	return e
}

func addBlock(e event, num_block int, textBlock string) event {

	switch num_block {
	case 1:
		e.datetime = textBlockToDatetime(textBlock)
	case 2:
		e.tr_status = textBlockToTrStatus(textBlock)
	case 31:
		e.tr_time = textBlock
	case 32:
		e.tr_offset = textBlock
	case 4:
		e.user = textBlockToUser(textBlock)
	case 5:
		e.pc = textBlockToPC(textBlock)
	case 6:
		e.app = textBlockToApp(textBlock)
	case 7:
		e.conn = textBlockToConn(textBlock)
	case 8:
		e.event = textBlockToEvent(textBlock)
	case 9:
		e.importance = textBlockToImportance(textBlock)
	case 10:
		e.comment = textBlockToComment(textBlock)
	case 11:
		e.metadata = textBlockToMetadata(textBlock)
	case 12:
		e.data = textBlockToData(textBlock)
	case 13:
		e.dataView = textBlockToDataView(textBlock)
	case 14:
		e.server = textBlockToServer(textBlock)
	case 15:
		e.mainPort = textBlockToPort(textBlock)
	case 16:
		e.secondPort = textBlockToPort(textBlock)
	case 17:
		e.session = textBlockToSession(textBlock)
	case 180:
		e.tr_status = textBlock
	case 190:
		e.tr_status = textBlock
	}

	return e
}

func textBlockToDatetime(textBlock string) time.Time {
	dt, err := time.Parse("20060102150405", textBlock)
	if err != nil {
		panic(err)
	}
	return dt
}

func textBlockToTrStatus(textBlock string) string {

	status := ""

	switch textBlock {
	case "N":
		status = "Отсутствует"
	case "U":
		status = "Не завершена"
	case "R":
		status = "Отменена"
	case "C":
		status = "Зафиксирована"
	}

	return status
}

func textBlockToUser(textBlock string) string {
	return LgfInfo.Users[textBlock]
}

func textBlockToPC(textBlock string) string {

	return LgfInfo.Computers[textBlock]
}

func textBlockToApp(textBlock string) string {

	return LgfInfo.Apps[textBlock]
}

func textBlockToConn(textBlock string) int {

	conn := 0
	conn, err := strconv.Atoi(textBlock)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func textBlockToEvent(textBlock string) string {

	return LgfInfo.Events[textBlock]

}

func textBlockToImportance(textBlock string) string {

	importance := ""

	switch textBlock {
	case "I":
		importance = "Информация"
	case "E":
		importance = "Ошибки"
	case "W":
		importance = "Предупреждения"
	case "N":
		importance = "Примечания"
	}
	return importance
}

func textBlockToComment(textBlock string) string {

	return textBlock
}

func textBlockToMetadata(textBlock string) string {

	return LgfInfo.Metadata[textBlock]
}

func textBlockToData(textBlock string) string {

	return textBlock
}

func textBlockToDataView(textBlock string) string {

	return textBlock
}

func textBlockToServer(textBlock string) string {

	return LgfInfo.Servers[textBlock]
}

func textBlockToPort(textBlock string) int {

	port := 0
	port, err := strconv.Atoi(textBlock)
	if err != nil {
		log.Fatal(err)
	}

	return port
}

func textBlockToSession(textBlock string) int {

	sess := 0
	sess, err := strconv.Atoi(textBlock)
	if err != nil {
		log.Fatal(err)
	}
	return sess
}

func printEvent(e event) {

	fmt.Println("******************************************************************")
	fmt.Println("1.   datetime: ", e.datetime)
	fmt.Println("2.   tr_status: ", e.tr_status)
	fmt.Println("3-1. tr_time: ", e.tr_time)
	fmt.Println("3-2. tr_offset: ", e.tr_offset)
	fmt.Println("4.   user: ", e.user)
	fmt.Println("5.   pc: ", e.pc)
	fmt.Println("6.   app: ", e.app)
	fmt.Println("7.   conn: ", e.conn)
	fmt.Println("8.   event: ", e.event)
	fmt.Println("9.   importance: ", e.importance)
	fmt.Println("10.  comment: ", e.comment)
	fmt.Println("11.  metadata: ", e.metadata)
	fmt.Println("12.  data: ", e.data)
	fmt.Println("13.  data_pres: ", e.dataView)
	fmt.Println("14.  server: ", e.server)
	fmt.Println("15.  port_main: ", e.mainPort)
	fmt.Println("16.  port_second: ", e.secondPort)
	fmt.Println("17.  session: ", e.session)

}
