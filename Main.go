package main

import (
	"fmt"
	"github.com/hegedustibor/htgo-tts"
	"github.com/kbinani/screenshot"
	ps "github.com/mitchellh/go-ps"
	"image/png"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	runLoop()
}

func runLoop() {

	for {
		currentProccesses := getProcesses()
		//		numProcesses := len(currentProccesses)
		for _, info := range currentProccesses {
			for _, gameName := range gameInfo {
				if strings.Contains(info.Executable(), gameName) {
					takeScreenshot()
					time.Sleep(5 * time.Millisecond)
					proc, err := os.FindProcess(info.Pid())
					if err != nil {
						log.Println("can't kill ", info.Executable())
					}
					proc.Kill()
					speech := htgotts.Speech{Folder: "audio", Language: "en"}
					speech.Speak("Stop playing Games!")

					//				scoldWindow := makeScoldWindow()
					//				scoldWindow.ShowAndRun()
				}
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func takeScreenshot() {
	//pulled directly from the readme https://github.com/kbinani/screenshot
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", time.Now(), bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", time.Now(), bounds, fileName)
	}
}

//func makeScoldWindow() fyne.Window{
//	fyneApp := app.New()
//	fyneWindow := fyneApp.NewWindow("Games Not Yet Authorized!!!!")
//	layout := container.NewVBox()
//	style:= fyne.TextStyle{
//		Bold: true,
//	}
//	scoldText := widget.NewLabelWithStyle("GAMES ARE NOT AUTHORIZED YET!", fyne.TextAlignCenter, style)
//	layout.Add(scoldText)
//	oopsButton := widget.NewButton("Oppps I shoulda known that", func() {
//			fyneWindow.Close()
//	})
//	layout.Add(oopsButton)
//	fyneWindow.SetContent(layout)
//
//	return fyneWindow
//}

func getProcesses() []ps.Process {
	allProceeses, err := ps.Processes()
	if err != nil {
		log.Println("Error Getting Processes")
	}
	return allProceeses
}
