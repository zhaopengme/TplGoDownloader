package main
import (
	"github.com/magiconair/properties"
	"github.com/lealife/leacrawler"
	"fmt"
)

func main() {
	config := properties.MustLoadFile("config.properties", properties.UTF8)
	url := config.MustGetString("url")
	path := config.MustGetString("path")
	fmt.Println("url:",url)
	fmt.Println("path:",path)

	leacrawler.Fetch(url,path)
}