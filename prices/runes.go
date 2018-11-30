package prices

import (
	"github.com/gin-gonic/gin"
	//"sync"
)

type Rune struct {
	Price int `json:"price"`
	Id    int
	Key   string
}

var runes = []Rune{
	Rune{
		Id:  554,
		Key: "fire",
	},
	Rune{
		Id:  555,
		Key: "water",
	},
	Rune{
		Id:  556,
		Key: "air",
	},
	Rune{
		Id:  557,
		Key: "earth",
	},
	Rune{
		Id:  558,
		Key: "mind",
	},
	Rune{
		Id:  559,
		Key: "body",
	},
	Rune{
		Id:  560,
		Key: "death",
	},
	Rune{
		Id:  561,
		Key: "nature",
	},
	Rune{
		Id:  562,
		Key: "chaos",
	},
	Rune{
		Id:  563,
		Key: "law",
	},
	Rune{
		Id:  564,
		Key: "cosmic",
	},
	Rune{
		Id:  565,
		Key: "blood",
	},
	Rune{
		Id:  566,
		Key: "soul",
	},
	Rune{
		Id:  9075,
		Key: "astral",
	},
}

func getRunesPrices(c *gin.Context) {
	output := make(map[string]int, len(runes))

	//var wg sync.WaitGroup

	for _, v := range runes {
		var p = 0
		switch v.Id {
		case 554:
			p = 5
			break
		case 555:
			p = 5
			break
		case 556:
			p = 5
			break
		case 557:
			p = 5
			break
		case 558:
			p = 5
			break
		case 559:
			p = 5
			break
		case 560:
			p = 278
			break
		case 561:
			p = 297
			break
		case 562:
			p = 102
			break
		case 563:
			p = 332
			break
		case 564:
			p = 172
			break
		case 565:
			p = 330
			break
		case 566:
			p = 146
			break
		case 9075:
			p = 217
			break
		default:

		}

		output[v.Key] = p

		/*
			wg.Add(1)
			go func(v Rune) {
			//	item := getItemPrice(v.Id)
			//	output[v.Key] = item.Item.Current.Price
				defer wg.Done()
				output[v.Key] = getPrice(v.Key)
		*/

	}

	//	wg.Wait()

	c.JSON(200, output)

}
