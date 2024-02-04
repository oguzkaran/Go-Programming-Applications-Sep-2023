/*
------------------------------------------------------------------------------------------------------------------------
	Aşağıdaki örneği inceleyiniz
------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"SampleGoLand/csd/console"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"os"
	"time"
)

func printMessage(message string) {
	fmt.Println(message)
}

func runApp() {
	scheduler, err := gocron.NewScheduler()

	defer func() {
		err = scheduler.Shutdown()
		if err != nil {
			fmt.Printf("Problem occurred while shut down:%s\n", err.Error())
			os.Exit(1)
		}
	}()

	if err != nil {
		fmt.Printf("Problem occurred while creating scheduler:%s\n", err.Error())
		os.Exit(1)
	}

	_, err = scheduler.NewJob(gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(time.Now().Add(10*time.Second))),
		gocron.NewTask(printMessage, "Hello, World!..."))

	if err != nil {
		fmt.Printf("Problem occurred while creating job:%s\n", err.Error())
		os.Exit(1)
	}
	scheduler.Start()

	console.ReadString("")
}

func main() {
	runApp()
}
