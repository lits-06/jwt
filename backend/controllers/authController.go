package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/lits-06/go-auth/database"
	"github.com/lits-06/go-auth/models"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data map[string]string

		if err := c.Bind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

		user := models.User{
			Name: data["name"],
			Email: data["email"],
			Password: password,
		}

		database.DB.Create(&user)

		c.JSON(http.StatusOK, user)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data map[string]string

		if err := c.Bind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		var user models.User

		database.DB.Where("email = ?", data["email"]).First(&user)

		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}

		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect password"})
			return
		}

		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer: strconv.Itoa(int(user.ID)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		})

		token, err := claims.SignedString([]byte(SecretKey))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "could not login"})
			return
		}

		c.SetCookie("jwt", token, 3600 * 24, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func User() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		cookie, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusOK, user)
			return
		}

		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthenticated",
			})
			return
		}

		claims := token.Claims.(*jwt.StandardClaims)

		database.DB.Where("id = ?", claims.Issuer).First(&user)

		c.JSON(http.StatusOK, user)
	}
}

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("jwt", "", -3600, "/", "localhost", false, true)

		_, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to logout",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
