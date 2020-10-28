package main

import (
	"auto/pre"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func importfile(lout *os.File) {
	filename := "instructions.auto"
	if fileExists(filename) {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Println(err.Error())
			return
		}
		instructions := strings.Split(string(content), "\n")
		index := 0
		for index < len(instructions) {
			row := strings.Split(instructions[index], " ")
			if len(row) == 1 && row[0] == "" {
				//
			} else {
				index = pre.Process(row, index, lout)
			}
			index++
		}
		log.Println("Complete")
	} else {
		newfile, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0640)
		defer newfile.Close()
		guide := `>
'AUTO' INSTRUCTIONS
You can place multi line comments in between '>' characters. 
You can place a # to do single line comments.

Here are the commands for this application:

pause - pauses the app asking for user input
run filepath args - runs an application at a specific location i.e - "run C:\hello.exe argument1 argument2 argument3"
deletefile filepath - deletes a file at a location i.e - "deletefile C:\hello.exe"
copyfile sourcepath destinationpath - copies a file to a new location i.e - "copyfile C:\hello.exe C:\newhello.exe"
clearfolder sourcepath - clears a folder of all files in it i.e - "clearfolder C:\folder"
copyfolder sourcepath destinationpath - copies all files in a folder to new location i.e "copyfolder C:\folder C:\newfolder"

You can start a loop for so many iterations with a colon i.e - ":loop 2" (loop twice)
You can return to the loop with a goto i.e "goto :loop"

Every time you run this, a log will be generated in a folder 'autologs' alongside where the application is saved.
>

####################################  Enter instructions below ###################################################


`
		_, err := newfile.WriteString(guide)
		if err != nil {
			log.Println(err.Error())
			return
		}
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
