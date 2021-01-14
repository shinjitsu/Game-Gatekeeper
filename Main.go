package main

import (
	"fmt"
	"github.com/hegedustibor/htgo-tts"
	ps "github.com/mitchellh/go-ps"
	"log"
	"os"
	"strings"
)

func main() {
	runLoop()
}

func runLoop() {
	//	for{
	currentProccesses := getProcesses()
	//		numProcesses := len(currentProccesses)
	for _, info := range currentProccesses {
		for _, gameName := range gameInfo {
			if strings.Contains(info.Executable(), gameName) {
				fmt.Println(info.Executable())
				proc, err := os.FindProcess(info.Pid())
				if err != nil {
					log.Println("can't kill ", info.Executable())
				}
				proc.Kill()
				speech := htgotts.Speech{Folder: "audio", Language: "en"}
				speech.Speak("Stop playing Games!")
			}
		}
	}
	//	time.Sleep(5000)
	//	}
}

func getProcesses() []ps.Process {
	allProceeses, err := ps.Processes()
	if err != nil {
		log.Println("Error Getting Processes")
	}
	return allProceeses
}
