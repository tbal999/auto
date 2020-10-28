package lex

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Command(item []string, lout *os.File) {
	for index := range item {
		item[index] = strings.TrimRight(item[index], "\r")
	}
	switch item[0] {
	case "":
		return
	case "pause":
		fmt.Println("Paused. Press anything to continue.")
		log.Println("Paused. Press anything to continue.")
		fmt.Scanln()
	case "run":
		if len(item) > 1 {
			log.Printf("Executing: %s\n", item[1])
			if len(item) > 2 {
				args := strings.Join(item[2:], " ")
				log.Printf("Args: %s\n", args)
				cmnd, err := exec.Command(item[1], args).Output()
				if err != nil {
					fmt.Println(err.Error())
				}
				log.Printf("%s\n", "START")
				log.Printf("\n%s", cmnd)
				log.Printf("%s\n", "END")
			} else {
				cmnd, err := exec.Command(item[1]).Output()
				if err != nil {
					fmt.Println(err.Error())
				}
				log.Printf("%s\n", "START")
				log.Printf("\n%s", cmnd)
				log.Printf("%s\n", "END")
			}
		} else {
			log.Printf("Executing: %s\n", "No Input!")
		}
	case "deletefile":
		log.Printf("Deleting File: %s\n", item[1])
		err := os.Remove(item[1])
		if err != nil {
			log.Println(err.Error())
		}
	case "copyfile":
		log.Printf("Copying File: %s -> %s\n", item[1], item[2])
		input, err := ioutil.ReadFile(item[1])
		if err != nil {
			fmt.Println(err)
			log.Println(err)
		}
		err = ioutil.WriteFile(item[2], input, 0644)
		if err != nil {
			fmt.Println(err)
			log.Println(err)
		}
	case "clearfolder":
		log.Printf("Clearing Folder: %s\n", item[1])
		inputfiles, err := ioutil.ReadDir(item[1])
		if err != nil {
			fmt.Println(err)
		}
		for _, inputfile := range inputfiles {
			switch mode := inputfile.Mode(); {
			case mode.IsDir():
				log.Printf("Skipping directory %s\n", item[1]+`\`+inputfile.Name())
			case mode.IsRegular():
				e := os.Remove(inputfile.Name())
				if e != nil {
					log.Println(e.Error())
				}
			}
		}
	case "copyfolder":
		log.Printf("Copying all Files: %s -> %s\n", item[1], item[2])
		inputfiles, err := ioutil.ReadDir(item[1])
		if err != nil {
			fmt.Println(err)
		}
		for _, inputfile := range inputfiles {
			switch mode := inputfile.Mode(); {
			case mode.IsDir():
				log.Printf("Skipping directory %s\n", item[1]+`\`+inputfile.Name())
			case mode.IsRegular():
				from := item[1] + "/" + inputfile.Name()
				to := item[2] + "/" + inputfile.Name()
				input, err := ioutil.ReadFile(from)
				if err != nil {
					fmt.Println(err)
					log.Println(err)
				} else {
					err = ioutil.WriteFile(to, input, 0644)
					if err != nil {
						fmt.Println(err)
						log.Println(err)
					}
				}
			}
		}
	default:
		log.Printf("Unknown instruction: %s\n", item[0])
	}
}
