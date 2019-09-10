package main
import (
	"fmt"
	"github.com/zhaopengme/tpldownloader/lib/github.com/lealife/leacrawler"
	"github.com/zhaopengme/tpldownloader/lib/github.com/magiconair/properties"
)

func main() {
	config := properties.MustLoadFile("config.properties", properties.UTF8)
	url := config.MustGetString("url")
	path := config.MustGetString("path")
	fmt.Println("url:",url)
	fmt.Println("path:",path)

	leacrawler.Fetch(url,path)
}