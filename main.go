package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll3", "url.dll.FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Script starting")

	url := "https://stackoverflow.com/questions/56218573/how-to-replace-flash-media-server-with-ant-media-server"

	fmt.Println("Browser opening")

	openbrowser(url)

	fmt.Println("Browser ended")
}
