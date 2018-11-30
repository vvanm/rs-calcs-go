package transcript

import (
	"github.com/gin-gonic/gin"
	"github.com/vvanm/rs-calcs-go/raven"
	"log"
	"strconv"
	"strings"
)

func handleGetParam(key string, c *gin.Context) (string, bool) {
	params, found := c.GetQuery(key)
	return strings.Join(strings.Split(params, ","), `','`), found

}

func Search(c *gin.Context) {
	isAdmin := c.Query("admin")

	//open session
	session, err := raven.Store.OpenSession()
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	var transcripts []*Transcript

	var query string

	if isAdmin == "true" {
		query = `
				declare function projection(entry,ID){
					let createdBy = load(entry.createdBy)
					return {
						ID,
						title : entry.title,
						type : entry.title,
						twitchVOD : entry.twitchVOD,
						createdOn : entry.createdOn,
						createdByName : createdBy.name,
						fragmentsCount : entry.fragments.length,
						keywordsCount : entry.keywords.length,
						modsCount : entry.mods.length
					}
				}
				from index 'transcripts/search' as entry
				select projection(entry,Id())
		`
	} else {

		//collect filters
		keywords, keywordsFound := handleGetParam("keywords", c)

		//build query
		qWhere := ""
		if keywordsFound {
			qWhere += "where "
		}
		if keywordsFound {
			qWhere += `entry.keywords in ('` + keywords + `')`
		}

		query = `
			declare function projection(entry){
			fragments = entry.fragments
			if(` + strconv.FormatBool(keywordsFound) + `){
				fragments = fragments.filter(f => f.keywords !== null && f.keywords.indexOf('` + keywords + `') !== -1)
			}	

			let createdBy = load(entry.createdBy)

			return {
				title : entry.title,
				type : entry.type,
				twitchVOD : entry.twitchVOD,
				createdOn : entry.createdOn,
				createdByName : createdBy.name,
				keywords : entry.keywords,
				mods : entry.mods,
				fragments,
			}
		}

		from index 'transcripts/search' as entry
		` + qWhere + `
		order by entry.createdOn desc
		select projection(entry)`
	}

	err = session.Advanced().RawQuery(query).ToList(&transcripts)
	if err != nil {
		log.Println(err)
	}

	c.JSON(200, transcripts)

}
