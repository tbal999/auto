package pre

import (
	"auto/lex"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	stop        = false
	loop        []string
	returnindex []int
	loopcount   int
)

func clean(s string) string {
	return strings.TrimSuffix(s, "\r")
}

func findloopindex(item string) int {
	for index := range loop {
		if strings.Contains(loop[index], item) == true {
			return returnindex[index]
		}
	}
	return 0
}

func Process(row []string, index int, lout *os.File) int {
	var err error
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
			log.Printf("Start Loop: %s %d\n", row[0], loopcount)
			if err != nil {
				log.Printf("Loop index failure for: %s\n", row[0])
			}
			loopcount--
		} else {
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
		log.Printf("Goto %s %d\n", clean(row[1]), loopcount)
		loopcount--
	}
	return index
}
