package main

import (
	"auto/pre"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func importfile(lout *os.File) {
	filename := "instructions.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
		return
	}
	instructions := strings.Split(string(content), "\n")
	index := 0
	for index < len(instructions) {
		row := strings.Split(instructions[index], " ")
		index = pre.Process(row, index, lout)
		index++
	}
}

func timestamp() string {
	t := time.Now()
	var x = t.Format("2006-01-02 1504")
	return x + ".log"
}

func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModeDir)
	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func main() {
	ensureDir("./autologs")
	logf, err := os.OpenFile("./autologs/"+timestamp(), os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logf)
	importfile(logf)
}
