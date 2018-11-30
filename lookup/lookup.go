package lookup

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Stat struct {
	Rank  int
	Level int
	Xp    int
}

func ByName(c *gin.Context) {

	resp, err := http.Get("https://secure.runescape.com/m=hiscore_oldschool/index_lite.ws?player=" + c.Param("rsn"))
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	//	body, err := ioutil.ReadAll(resp.Body)
	//

	reader := csv.NewReader(resp.Body)
	reader.FieldsPerRecord = -1

	stats, err := reader.ReadAll()

	out := map[string]int{}

	for i, stat := range stats {
		if i == 0 || i > 23 {
			continue
		}

		skillName := getSkillName(i)

		out[skillName+"Level"], _ = strconv.Atoi(stat[1])
		out[skillName+"Xp"], _ = strconv.Atoi(stat[2])

	}

	c.JSON(200, gin.H{
		"skills": out,
	})

}

func getSkillName(i int) string {
	switch i {
	case 1:
		return "attack"
	case 2:
		return "defence"
	case 3:
		return "strength"
	case 4:
		return "hitpoints"
	case 5:
		return "ranged"
	case 6:
		return "prayer"
	case 7:
		return "magic"
	case 8:
		return "cooking"
	case 9:
		return "woodcutting"
	case 10:
		return "fletching"
	case 11:
		return "fishing"
	case 12:
		return "firemaking"
	case 13:
		return "crafting"
	case 14:
		return "smithing"
	case 15:
		return "mining"
	case 16:
		return "herblore"
	case 17:
		return "agility"
	case 18:
		return "thieving"
	case 19:
		return "slayer"
	case 20:
		return "farming"
	case 21:
		return "runecrafting"
	case 22:
		return "hunter"
	case 23:
		return "construction"
	default:
		return ""
	}

	return ""
}
