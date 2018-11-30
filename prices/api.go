package prices

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Item struct {
	Item struct {
		Current struct {
			Price int
		}
	}
}

func getItemPrice(id int) Item {

	baseUrl := "http://services.runescape.com/m=itemdb_oldschool/api/catalogue/detail.json?item="

	resp, err := http.Get(baseUrl + strconv.Itoa(id))
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	log.Println(err)

	var item Item

	_ = json.Unmarshal(body, &item)

	if item.Item.Current.Price == 0 {
		log.Println("not found", id, baseUrl)
	}

	return item

}
