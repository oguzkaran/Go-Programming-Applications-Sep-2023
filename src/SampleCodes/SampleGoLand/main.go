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

func printDateTime() {
	fmt.Printf("%s\r", time.Now().Format("02/01/2006 15.04.05"))
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

	_, err = scheduler.NewJob(gocron.DurationJob(1*time.Second),
		gocron.NewTask(printDateTime))

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
