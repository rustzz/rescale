package main

import (
	"fmt"
	"github.com/rustzz/rescale"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	//url = "https://2ch.hk/makaba/templates/img/anon.jpg"
	url = "https://sun9-18.userapi.com/impg/EQ9e20QOCBfGar1CdxVl0b_XMHF9vh40DPdnFA/4PlTo_QUUdc.jpg?size=735x693&quality=96&sign=a698bde9dba7e2e3489fbf56565cbec5&type=album"
)

func GetImage(url string) (imageBytes []byte, err error) {
	resp, err := http.Get(url)
	if err != nil { return }
	defer resp.Body.Close()

	imageBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil { return }
	return
}

func main() {
	homeDir, err := os.UserHomeDir()
	file, err := os.Create(fmt.Sprintf("%s/out.png", homeDir))
	if err != nil { log.Fatal(err) }
	defer file.Close()

	srcImageBytes, err := GetImage(url)
	if err != nil { log.Fatal(err) }
	outImageBytes, err := rescale.Make(srcImageBytes, 1)

	file.Write(outImageBytes)
}
