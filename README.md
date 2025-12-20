# Read Me
Toolu kullanamk için öncelikle go modülünü bilgisayarınıza indirmeniz gerekmektedir. https://go.dev/dl/ bu link üzerinden indirebilirsiniz. <br>
Daha sonra modülü başlatmak için <br>  
cd web_scrapper<br>
go mod init web_scrapper <br>
komutları girilmelidir.<br>
Eğer bilgisayarınızda colly kütüphanesi yoksa<br>
go get github.com/gocolly/colly/v2<br>
komutuyla indirebilirsiniz. <br>

Toolu çalıştırmak için CMD  ekranında <br>
go run main.go  <br>
komutunu girerek çalıştırabilirsiniz.


NOT: kodun default ayarlarında bir windows makinede çalıştırıldığı varsayılarak MSEDGE yolu   C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe  olarak verilmiştir. Eğer MSEDGE yolunun farklı olduğunu düşünüyorsanız 115. satırdaki kodu değiştirmeniz gerekmektedir.
