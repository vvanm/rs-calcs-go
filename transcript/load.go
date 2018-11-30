package transcript

import (
	"github.com/gin-gonic/gin"
	"github.com/vvanm/rs-calcs-go/raven"
	"log"
)

func Load(c *gin.Context) {

	transcriptKey := "transcripts/" + c.Param("id")

	var t *Transcript

	session, err := raven.Store.OpenSession()
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	err = session.Load(&t, transcriptKey)

	if t != nil {
		c.JSON(200, t)
		return
	}

	c.JSON(404, gin.H{"errorMsg": "not found"})

}
