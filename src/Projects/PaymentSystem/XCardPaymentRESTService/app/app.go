package app

import (
	"XCardPaymentRESTService/app/logger"
	"os"
)

func initApplication() {
	logger.InitLoggers(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
}

func Run() {
	/*
		Payment Manager TCP protocol details:
		send id as integer value of 4 bytes
		send string as "R;endpoint". String should send via length means, first send number of character as 4 bytes integer
		value, then send characters as byte array
		receive 1 byte. If its value is 1 (means success) start rest service, otherwise end process
	*/
	initApplication()
	const paymentManagerServerHost = ""
	const paymentManagerServerPort = 50540

}
