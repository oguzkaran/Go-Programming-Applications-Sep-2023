/*
------------------------------------------------------------------------------------------------------------------------

    Sınıf Çalışması: Aşağıda açıklanan programı yazınız
    Açıklamalar:
        - Program tarih zaman ve isim bilgisini aşağıdaki formatta komut satırı argümanından alacaktır
            ./eqapp 17/08/1999 03:02:00 Gölcük

        - Alınan deprem tarihinden itibaren kaç yıl geçtiği float64 olarak kullanıcıya verilecektir.

        - Alınan bilgiler aşağıdaki gibi tasarlanmış bir veritabanına eklenecektir. Veritabanında deprem ismi
        case-insensitive olarak primary key biçiminde tutulacaktır. Her sorgulama bilgisi tarih zamana göre kaydedilecektir.
            - earthquakes
                 - name
                 - date_time
            - earthquake_info
                - earthquake_info_id
                - name (foreign key)
                - query_date_time

        - Geçersiz bilgiler durumunda uygun mesaj verilecektir.

        - Uygulama katmanlı mimari kullanılacaktır. Katmanlar üstten alta doğru şu şekilde olacaktır:
            - Application
            - Data Service
            - Data Access Layer
            - Repository

        - Repository katmanında daha önceden yazmış olduğumuz CrudRepository interface'i kullanılacaktır ve yalnızca
        gereken metotlar impelemente edilecektir.

        - Tüm katmanlardaki yazılan fonksiyonlar ayrı ayrı testing paketi kullanılarak test edilecektir edilecektir
------------------------------------------------------------------------------------------------------------------------
*/

package main

import "EarthquakeApp/app"

func main() {
	app.Run()
}
