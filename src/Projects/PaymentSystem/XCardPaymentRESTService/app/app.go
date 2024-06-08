package app

func Run() {
	/*
		Payment Manager TCP protocol details:
		send id as integer value of 4 bytes
		send string as "R;endpoint". String should send via length means first send number of character as 4 bytes integer
		value, then send character as byte array
		receive 1 byte. If its value is 1 (means success) start rest service, otherwise end process
	*/
}
