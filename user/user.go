package user

import (
	"errors"
	"fmt"
	"github.com/vvanm/rs-calcs-go/raven"
	"github.com/vvanm/rs-calcs-go/util/auth"
	"github.com/vvanm/rs-calcs-go/util/helpers"
)

type User struct {
	ID             string
	Name           string `json:"name,omitempty"`
	Name_          string `json:"name_,omitempty"`
	Password       string `json:"password,omitempty"`
	UserType       string `json:"userType"`
	RepeatPassword string `json:"repeatPassword,omitempty"`
	Email          string `json:"email,omitempty"`
}

func GetUser(q string) (u *User, err error) {
	session, err := raven.Store.OpenSession()
	if err != nil {
		return
	}
	defer session.Close()

	var users []*User
	err = session.Advanced().RawQuery(q).ToList(&users)

	if err != nil {
		return
	}

	if len(users) == 1 {
		u = users[0]
	}

	return

}

func (u *User) Create() error {
	if u.Name == "" || u.Email == "" || u.Password == "" || u.RepeatPassword == "" {
		return errors.New("Empty inputs")
	}

	if u.Password != u.RepeatPassword {
		return errors.New("Passwords dont match")
	}

	u.RepeatPassword = ""
	u.UserType = "user"
	u.Name_ = helpers.SaniString(u.Name)

	//open session
	session, err := raven.Store.OpenSession()
	if err != nil {
		return err
	}
	defer session.Close()

	//Verify unique username and email
	var users []*User
	err = session.Advanced().RawQuery(fmt.Sprintf("from Users where email == '%s' OR name_ == '%s'", u.Email, u.Name_)).ToList(&users)
	if err != nil {
		return err
	}

	if len(users) > 0 {
		return errors.New("Email or username is in use")
	}

	//Process the password
	u.Password = auth.SecurePassword(u.Password)

	//add user
	err = session.StoreWithID(u, "users|")
	if err != nil {
		return err
	}

	//push to raven
	err = session.SaveChanges()
	if err != nil {
		return err
	}

	return nil

}
