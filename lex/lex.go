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
		fmt.Printf("Paused. Press anything to continue.")
		log.Printf("Paused. Press anything to continue.")
		fmt.Scanln()
	case "run":
		if len(item) > 1 {
			fmt.Printf("Executing: %s\n", item[1])
			log.Printf("Executing: %s\n", item[1])
			if len(item) > 2 {
				args := strings.Join(item[2:], " ")
				fmt.Printf("Args: %s\n", args)
				log.Printf("Args: %s\n", args)
				cmnd, err := exec.Command(item[1], args).Output()
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Printf("%s", cmnd)
				log.Printf("%s", "START")
				log.Printf("\n%s", cmnd)
				log.Printf("%s", "END")
			} else {
				cmnd, err := exec.Command(item[1]).Output()
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Printf("%s", cmnd)
				log.Printf("%s", "START")
				log.Printf("\n%s", cmnd)
				log.Printf("%s", "END")
			}
		} else {
			fmt.Printf("Executing: %s\n", "No Input!")
			log.Printf("Executing: %s\n", "No Input!")
		}
	case "deletefile":
		fmt.Printf("Deleting File: %s\n", item[1])
		log.Printf("Deleting File: %s\n", item[1])
		err := os.Remove(item[1])
		if err != nil {
			fmt.Println(err.Error())
			log.Println(err.Error())
		}
	case "copyfile":
		fmt.Printf("Copying File: %s -> %s\n", item[1], item[2])
		log.Printf("Copying File: %s -> %s\n", item[1], item[2])
		input, err := ioutil.ReadFile(item[1])
		if err != nil {
			fmt.Println(err)
			log.Println(err)
		}
		err = ioutil.WriteFile(item[2], input, 0644)
		if err != nil {
			fmt.Println("Error creating", item[2])
			fmt.Println(err)
			log.Println(err)
		}
	case "clearfolder":
		fmt.Printf("Clearing Folder: %s\n", item[1])
		log.Printf("Clearing Folder: %s\n", item[1])
		os.RemoveAll(item[1])
	case "copyfolder":
		fmt.Printf("Copying all Files: %s -> %s\n", item[1], item[2])
		log.Printf("Copying all Files: %s -> %s\n", item[1], item[2])
		inputfiles, err := ioutil.ReadDir(item[1])
		if err != nil {
			fmt.Println(err)
		}
		for _, inputfile := range inputfiles {
			from := item[1] + "/" + inputfile.Name()
			to := item[2] + "/" + inputfile.Name()
			input, err := ioutil.ReadFile(from)
			if err != nil {
				fmt.Println(err)
				log.Println(err)
			}
			err = ioutil.WriteFile(to, input, 0644)
			if err != nil {
				fmt.Println("Error creating", to)
				fmt.Println(err)
				log.Println(err)
			}
		}
	default:
		fmt.Printf("Unknown instruction: %s", item[0])
		log.Printf("Unknown instruction: %s", item[0])
	}
	fmt.Printf("\n")
	log.Printf("\n")
}
