package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type pageData struct {
	Messages map[string]string
	Selected string
}

func getContentByLines(filePath string) []string {

	f, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return lines

}

func handler(w http.ResponseWriter, r *http.Request) {

	logFilePath := os.Args[1]
	tpl, _ := template.ParseFiles("./index.html")
	messageMap := map[string]string{}
	lines := getContentByLines(logFilePath)

	for _, line := range lines {
		timestamp := line[0:30]
		message := strings.TrimLeft(line[30:], "")
		messageMap[timestamp] = message
	}

	tpl.Execute(w, pageData{
		Messages: messageMap,
		Selected: r.FormValue("timestamp"),
	})
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:9090",
	}
	http.HandleFunc("/", handler)
	server.ListenAndServe()

}
