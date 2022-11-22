package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
)

type Mail struct {
	MessageID               string
	Date                    string
	From                    string
	To                      string
	Subject                 string
	Cc                      string
	MimeVersion             string
	ContentType             string
	ContentTransferEncoding string
	XFrom                   string
	XTo                     string
	Xcc                     string
	Xbcc                    string
	XFolder                 string
	XOrigin                 string
	XFileName               string
	Message                 string
}

func main() {
	searchForArchives("./enron")

}

func sendToSync(direccion string) {
	readFile, err := os.Open(direccion)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	var prueba Mail
	var countFrom = 0
	var countTo = 0
	var countSubject = 0
	var countDate = 0
	var countCc = 0
	var countMessageId = 0
	var data string
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "Message-ID: "):
			if countMessageId == 0 {
				prueba.MessageID = strings.Split(line, "Message-ID: ")[1]
			} else {
				data += line + "\n"
			}
		case strings.HasPrefix(line, "Date: "):
			countMessageId += 1
			if countDate == 0 {
				prueba.Date = strings.Split(line, "Date: ")[1]
			} else {
				data += line + "\n"
			}
		case strings.HasPrefix(line, "From: "):
			if countFrom == 0 {
				prueba.From = strings.Split(line, "From: ")[1]
			} else {
				data += line + "\n"
			}
			countDate += 1
			countFrom += 1
		case strings.HasPrefix(line, "To: "):
			if countTo == 0 {
				prueba.To = strings.Split(line, "To: ")[1]
			} else {
				data += line + "\n"
			}
		case strings.HasPrefix(line, "Subject: "):
			if countSubject == 0 {
				prueba.Subject = strings.Split(line, "Subject: ")[1]
			} else {
				data += line + "\n"
			}
			countTo += 1
		case strings.HasPrefix(line, "Cc:"):
			if countCc == 0 {
				prueba.Cc = strings.Split(line, "Cc: ")[1]
			} else {
				data += line + "\n"
			}
		case strings.HasPrefix(line, "Mime-Version: "):
			countCc += 1
			countSubject += 1
			prueba.MimeVersion = strings.Split(line, "Mime-Version: ")[1]
		case strings.HasPrefix(line, "Content-Type: "):
			prueba.ContentType = strings.Split(line, "Content-Type: ")[1]
		case strings.HasPrefix(line, "Content-Transfer-Encoding: "):
			fmt.Println(direccion)
			prueba.ContentTransferEncoding = strings.Split(line, "Content-Transfer-Encoding: ")[1]
		case strings.HasPrefix(line, "X-From: "):
			prueba.XFrom = strings.Split(line, "X-From: ")[1]
		case strings.HasPrefix(line, "X-To: "):
			prueba.XTo = strings.Split(line, "X-To: ")[1]
		case strings.HasPrefix(line, "X-cc: "):
			prueba.Xcc = strings.Split(line, "X-cc: ")[1]
		case strings.HasPrefix(line, "X-bcc: "):
			prueba.Xbcc = strings.Split(line, "X-bcc: ")[1]
		case strings.HasPrefix(line, "X-Folder: "):
			prueba.XFolder = strings.Split(line, "X-Folder: ")[1]
		case strings.HasPrefix(line, "X-Origin: "):
			prueba.XOrigin = strings.Split(line, "X-Origin: ")[1]
		case strings.HasPrefix(line, "X-FileName: "):
			prueba.XFileName = strings.Split(line, "X-FileName: ")[1]
		default:
			if countSubject == 0 {
				prueba.Subject += " " + line
			} else {
				data += line + "\n"
			}
		}
	}
	prueba.Message = data

	e, err := json.Marshal(prueba)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:4080/api/emails/_doc", strings.NewReader(string(e)))

	if err != nil {
		fmt.Println(err)
		return
	}

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func searchForArchives(direccion string) string {
	archivos, err := ioutil.ReadDir(direccion)
	var datos string
	if err != nil {
		log.Fatal(err)
	}

	for _, archivo := range archivos {
		if archivo.IsDir() {
			aux := direccion + "/" + archivo.Name()
			datos += searchForArchives(aux)
		} else {
			sendToSync(direccion + "/" + archivo.Name())
		}
	}

	return datos
}
