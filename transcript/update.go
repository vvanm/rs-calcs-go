package transcript

import (
	"github.com/gin-gonic/gin"
	"github.com/ravendb/ravendb-go-client"
	"github.com/vvanm/rs-calcs-go/raven"
)

func Update(c *gin.Context) {
	transcriptKey := "transcripts/" + c.Param("id")

	var t Transcript
	c.BindJSON(&t)

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

	patch := ravendb.PatchRequest_forScript(`

	this.title = args.title;
	this.type = args.type;
	this.publishDate = args.publishDate;
	this.twitchVOD = args.twitchVOD;
	this.patchNotesLink = args.patchNotesLink;
	this.mods = args.mods;
	this.keywords = args.keywords
	this.fragments = args.fragments
`)

	//add values
	patch.SetValues(map[string]interface{}{
		"title":          t.Title,
		"type":           t.Type,
		"publishDate":    t.PublishDate,
		"twitchVOD":      t.TwitchVOD,
		"patchNotesLink": t.PatchNotesLink,
		"mods":           t.Mods,
		"keywords":       t.Keywords,
		"fragments":      t.Fragments,
	},
	)

	patchOp := ravendb.NewPatchOperation(transcriptKey, nil, patch, nil, false)
	err := raven.Store.Operations().Send(patchOp)

	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{})

}
