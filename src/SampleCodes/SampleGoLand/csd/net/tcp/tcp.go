package tcp

import (
	"encoding/binary"
	"net"
)

func Send(conn net.Conn, buf []byte) (int, error) {
	count := len(buf)
	for {
		n, err := conn.Write(buf)
		if err != nil {
			return 0, err
		}
		count -= n
		if count == 0 {
			break
		}
	}

	return len(buf), nil
}

func Receive(conn net.Conn, buf []byte) (int, error) {
	size := len(buf)
	count := 0
	data := make([]byte, size)

	for {
		n, err := conn.Read(data)
		if err != nil {
			return 0, err
		}

		for i := 0; i < len(data); i++ {
			buf[i+count] = data[i]
		}
		count += n

		if count == size {
			break
		}
	}

	return count, nil
}

func SendInt32(conn net.Conn, val int32) error {
	buf := make([]byte, 4)
	binary.NativeEndian.PutUint32(buf, uint32(val))
	_, err := Send(conn, buf)

	return err
}

func SendStringViaLength(conn net.Conn, str string) error {
	err := SendInt32(conn, int32(len(str)))
	if err != nil {
		return err
	}
	_, err = Send(conn, []byte(str))

	if err != nil {
		return err
	}
	return nil
}

//...
