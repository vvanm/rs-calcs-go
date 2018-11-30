package prices

import (
	"github.com/gin-gonic/gin"
)

func Controller(c *gin.Context) {
	category := c.Param("category")

	switch category {
	case "runes":
		getRunesPrices(c)
	}

}
