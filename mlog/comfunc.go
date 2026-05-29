package mlog

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func CheckErr(err error, mes string) {
	if err != nil {
		fmt.Println(mes)
		panic(err)
	}
}

func InfoMyConn(filename string) (string, string, string) {
	host, user, pass := "", "", ""
	file, err := os.Open(filename)
	CheckErr(err, "Не могу открыть файл "+filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s != "" {
			ar := strings.Split(s, ":")
			host, user, pass = ar[0], ar[1], ar[2]
			break
		}
	}
	err = scanner.Err()
	CheckErr(err, "Не могу считать строку в файле "+filename)
	return host, user, pass
}

func MyReadFile(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
