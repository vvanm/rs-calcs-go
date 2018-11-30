package transcript

import (
	"github.com/gin-gonic/gin"
	"github.com/vvanm/rs-calcs-go/raven"
	"github.com/vvanm/rs-calcs-go/util/helpers"
	"github.com/vvanm/rs-calcs-go/util/jwt"
	"log"
)

func Create(c *gin.Context) {
	//Find claims
	claims := jwt.ClaimsFromCookie(c)

	var t Transcript
	c.BindJSON(&t)

	t.CreatedOn = helpers.CurrentEpoch()
	t.CreatedBy = claims.ID

	//create map of unique keywords
	encounteredKeywords := map[string]bool{}
	for _, v := range t.Fragments {
		for _, vv := range v.Keywords {
			encounteredKeywords[vv] = true
		}
	}
	//place all keys into cLog.Keywords slice
	for k, _ := range encounteredKeywords {
		t.Keywords = append(t.Keywords, k)
	}

	//open session
	session, err := raven.Store.OpenSession()
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	//add log
	err = session.StoreWithID(&t, "transcripts|")
	if err != nil {
		log.Println(err)
	}
	//push to raven
	err = session.SaveChanges()
	if err != nil {
		log.Println(err)
	}

}
