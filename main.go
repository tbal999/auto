package main

import (
	"auto/lex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	stop        = false
	loop        []string
	returnindex []int
	loopcount   int
)

func findloopindex(item string) int {
	for index := range loop {
		if strings.Contains(loop[index], item) == true {
			return returnindex[index]
		}
	}
	return 0
}

func clean(s string) string {
	return strings.TrimSuffix(s, "\r")
}

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
		if row[0][0] == '>' {
			if stop == false {
				stop = true
			} else {
				stop = false
			}
		}
		if row[0][0] == ':' {
			if len(row) == 2 {
				loop = append(loop, row[0])
				returnindex = append(returnindex, index)
				cleaned := clean(row[1])
				loopcount, err = strconv.Atoi(cleaned)
				fmt.Printf("Start Loop: %s %d\n", row[0], loopcount)
				log.Printf("Start Loop: %s %d\n", row[0], loopcount)
				if err != nil {
					fmt.Println(err.Error())
					fmt.Printf("Loop index failure for: %s\n", row[0])
					log.Printf("Loop index failure for: %s\n", row[0])
				}
				loopcount--
			} else {
				fmt.Printf("Loop index not present for: %s\n", row[0])
				log.Printf("Loop index not present for: %s\n", row[0])
			}
		}
		if row[0][0] != '>' && row[0][0] != ':' && row[0][0] != '#' && row[0] != "goto" {
			if stop == false {
				lex.Command(row, lout)
			}
		}
		if row[0] == "goto" && loopcount > 0 {
			index = findloopindex(clean(row[1]))
			fmt.Printf("Goto %s %d\n", clean(row[1]), loopcount)
			log.Printf("Goto %s %d\n", clean(row[1]), loopcount)
			loopcount--
		}
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
