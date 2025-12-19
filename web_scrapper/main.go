package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {

	var siteadresi string

	var hatavarmi bool = false

	fmt.Println("Bilgisini çekmek istediğiniz sitenin adresini girin:")

	fmt.Scanln(&siteadresi)

	if siteadresi == "" {

		fmt.Println("Adres girmediniz!")
		return

	}

	if strings.HasPrefix(siteadresi, "http") == false {

		siteadresi = "https://" + siteadresi

	}

	fmt.Println("İşlem başlıyor.")

	c := colly.NewCollector()

	var linkler []string

	// 404 kısmı

	c.OnError(func(r *colly.Response, err error) {

		if r.StatusCode == 404 {

			fmt.Println("Bu sayfa bulunamadı (404)")

			hatavarmi = true

		} else {

			fmt.Println("Bir hata oluştu:", err)
			hatavarmi = true

		}

	})

	c.OnResponse(func(r *colly.Response) {

		dosya, _ := os.Create("output.html")

		dosya.Write(r.Body)
		dosya.Close()

	})

	c.OnResponse(func(r *colly.Response) {

		dosya, _ := os.Create("output.txt")

		dosya.Write(r.Body)
		dosya.Close()

	})

	c.OnHTML("a", func(e *colly.HTMLElement) {

		var link = e.Attr("href")
		if link != "" {

			linkler = append(linkler, link)

		}

	})

	c.Visit(siteadresi)

	//hata varsa programın durması için
	if hatavarmi == true {

		fmt.Println("Hata olduğu için işlem iptal edildi.")
		return

	}

	if len(linkler) > 0 {

		f, _ := os.Create("links.txt")
		for i := 0; i < len(linkler); i++ {

			f.WriteString(linkler[i] + "\n")

		}

		f.Close()

	}

	calisilan_yer, _ := os.Getwd()
	resim_yolu := calisilan_yer + "\\screenshot.png"

	var edgeYolu = "C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe"

	fmt.Println("Edge bulundu")

	cmd := exec.Command(edgeYolu,
		"--headless",
		"--disable-gpu",
		"--window-size=1920,8000",
		"--screenshot="+resim_yolu,
		siteadresi)

	cikti, hata := cmd.CombinedOutput()

	if hata != nil {

		fmt.Println("Resim çekilemedi! Hata:", hata)

		fmt.Println("Detay:", string(cikti))

	} else {

		fmt.Println("Resim başarıyla kaydedildi:", resim_yolu)
	}
}
