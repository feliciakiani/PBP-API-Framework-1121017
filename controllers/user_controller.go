package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db := connect()
	defer db.Close()

	var user User

	if err := c.Bind(&user); err != nil {
		fmt.Print(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	name_input := c.Query("Name")
	pasword_input := c.Query("Password")
	fmt.Println("Function Login, Name User = ", name_input, " , pass = ", pasword_input)

	row, err := db.Query("SELECT * FROM users WHERE Name=? AND Password=?", name_input, pasword_input)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the user Login query"})
		return
	}

	for row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Type); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "user not found"})
		} else {

			SetCookie(c, user)

			c.IndentedJSON(http.StatusOK, gin.H{"message ": "Login success", "data": user})
		}
	}
}

func Logout(c *gin.Context) {

	ResetCookie(c)

	c.IndentedJSON(http.StatusOK, gin.H{"message ": "Logout success"})
}

func SetCookie(c *gin.Context, user User) {

	cookie := http.Cookie{
		Name:     "user",
		Value:    encodeUser(user),
		Expires:  time.Now().Add(5 * time.Minute),
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, &cookie)

	// syntax singkatnya : (tetapi value hanya bisa string)
	// c.SetCookie("user", user, 3600, "/", "localhost", false, true)
}

func encodeUser(user User) string {
	return "encoded_user_string"
}

func decodeUser(str string) (User, error) {
	return User{}, nil
}

func ResetCookie(c *gin.Context) {

	cookie := http.Cookie{
		Name:     "user",
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, &cookie)

}
