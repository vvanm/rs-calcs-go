package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vvanm/rs-calcs-go/util/auth"
	"github.com/vvanm/rs-calcs-go/util/helpers"
	"github.com/vvanm/rs-calcs-go/util/jwt"
	"net/http"
	"time"
)

var projectionSession = `
	declare function projection(user,ID){
		return {
			ID,
			password : user.password,
			name : user.name,
			userType : user.userType
		}
	}
`

func Auth(c *gin.Context) {
	//Find claims
	claims := jwt.ClaimsFromCookie(c)

	if claims == nil {
		c.JSON(404, gin.H{"errorMsg": "noSession"})
		return
	}

	u, err := GetUser(fmt.Sprintf(projectionSession+`
	from index 'users/search' as entry
	where entry.ID == '%s'
	select projection(entry,Id())`, claims.ID,
	))

	if err != nil {
		c.JSON(400, gin.H{"errorMsg": err.Error()})
		return
	}

	if u == nil {
		c.JSON(400, gin.H{"errorMsg": "no user found"})
		return
	}

	authLoginReturnUser(c, u)

}

func Register(c *gin.Context) {
	var u User
	c.BindJSON(&u)

	err := u.Create()

	if err != nil {
		c.JSON(400, gin.H{"errorMsg": err.Error()})
		return
	}

	authLoginReturnUser(c, &u)

}

func Login(c *gin.Context) {
	//Receive post data
	var postU User
	c.BindJSON(&postU)

	u, err := GetUser(fmt.Sprintf(projectionSession+`
from index 'users/search' as entry
where entry.name_ == '%s'
select projection(entry,Id())`, helpers.SaniString(postU.Name)),
	)

	if err != nil {
		c.JSON(400, gin.H{"errorMsg": err.Error()})
		return
	}

	if u == nil {
		c.JSON(400, gin.H{"errorMsg": "no user found"})
		return
	}

	//Check password match
	if !auth.PasswordMatch(postU.Password, u.Password) {
		c.JSON(400, gin.H{"errorMsg": "failAuth"})
		return
	}

	//Create claims
	claims := jwt.NewClaims(u.ID)

	//Create token
	token, _ := jwt.NewToken(claims)

	//Create cookie
	cookie := jwt.NewCookie(token)

	//Set cookie
	http.SetCookie(c.Writer, cookie)

	authLoginReturnUser(c, u)

}

func Logout(c *gin.Context) {
	//Set cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "rs-calcs",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
}

func authLoginReturnUser(c *gin.Context, u *User) {
	c.JSON(200, gin.H{
		"userType": u.UserType,
		"name":     u.Name,
	})
}
