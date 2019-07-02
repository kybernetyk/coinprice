//TODO: https://play.golang.org/p/KNPxDL1yqL
package main

import (
	"os"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
	"fmt"
)

type APIResponse struct {
	Timestamp int64
	Success bool
	Error string

	Ticker struct {
		Base string
		Target string
		Price float64 `json:",string"`
		Volume float64 `json:",string"`
		Change float64 `json:",string"`
	}
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Wrong number of argumentas!")
	}
	base := os.Args[1]
	target := os.Args[2]

	url := fmt.Sprintf("https://api.cryptonator.com/api/ticker/%s-%s", base, target)
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.47 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp := APIResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		//the API sometimes sets volume to "" (empty string)
		//which blows up go's json decoding. so we just ignore all errors
		//and hope for the best. yay go
		//log.Print("Error while unmarshaling JSON: ", err)
	}
	if resp.Success {
		fmt.Printf("%.4f\n", resp.Ticker.Price)
	} else {
		fmt.Println(resp.Error)
	}
}
