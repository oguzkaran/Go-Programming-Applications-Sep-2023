/*
------------------------------------------------------------------------------------------------------------------------

	Sınıf Çalışması: Aşağıdaki JSON formatına ve açıklamalaro göre ilgili programı yazınız
		{"n":10, "min":10, "max":20, "count":30, "basePath":"numbers"}
	Açıklamalar:
		- JSON formatı stdin'den alınacaktır
		- JSON formatında bulunan değerlere ilişkin bilgiler şunlardır:
			n: goroutine sayısı
			min:rassal sayının alt sınır
			max: rassal sayının üst sınırı
			count:Rassal sayıların adedi
			basePath: Taban dosya ismi
		Buna göre her bir goroutine ilgili dosyaya ilgili sayıları yazdıracaktır. basFilePath değeri numbers ve n değeri 2
		goroutine'lere ilişkin dosya isimleri şu şekilde olacaktır: numbers-1, numbers-2, numbers-3

------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"SampleGoLand/csd/err"
	"encoding/binary"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		err.ExitFailure("wrong number of arguments!...")
	}
	count, e := strconv.Atoi(os.Args[2])

	if e != nil {
		err.ExitFailure("invalid count")
	}

	n, e := strconv.Atoi(os.Args[3])

	if e != nil {
		err.ExitFailure("invalid size")
	}

	f, e := os.OpenFile(os.Args[1], os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)

	if e != nil {
		err.ExitFailureError("Open:", e)
	}

	defer func() { _ = f.Close() }()

	for i := 0; i < count; i++ {
		var val int32 = int32(rand.Intn(n))

		e = binary.Write(f, binary.LittleEndian, &val)
		if e != nil {
			err.ExitFailureError("WriteString", e)
		}
	}
}
