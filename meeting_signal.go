package main

import (
	"fmt"
	"log"
	"os/exec"
	"html/template"
	"net/http"
	"strings"
)

type MeetingSignalData struct {
	LightStatus bool
}

func main() {
	tmpl := template.Must(template.ParseFiles("meeting_signal.html"))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		formErr := r.ParseForm()
    if formErr != nil {
      log.Fatal(formErr)
    }
		
		switchAction := r.FormValue("switch")
		
		if switchAction == "on" {
			// fmt.Println("Switching light to on!")
			
			_, onCmdErr := exec.Command("sh", "-c", "sudo ykushcmd ykushxs -u 0").Output()
			
			if onCmdErr != nil {
				log.Fatal(onCmdErr)
			}
		} else if switchAction == "off" {
			// fmt.Println("Switching light to off!")
			
			_, offCmdErr := exec.Command("sh", "-c", "sudo ykushcmd ykushxs -d 0").Output()
			
			if offCmdErr != nil {
				log.Fatal(offCmdErr)
			}
		}
		
		out, cmdErr := exec.Command("sh", "-c", "sudo ykushcmd ykushxs -g").Output()

		if cmdErr != nil {
			log.Fatal(cmdErr)
		}
		
		// Convert bytes array to a string
		responseOutput := string(out[:])
		
		response := strings.TrimSpace(responseOutput)
		status := response == "Downstream port is ON"
		
		data := MeetingSignalData{
			LightStatus: status,
		}
		
		tmpl.Execute(w, data)
	})

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
