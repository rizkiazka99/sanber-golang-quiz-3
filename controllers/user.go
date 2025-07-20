package controllers

import (
	"fmt"
	"net/http"
	"quiz3/middleware"
	"quiz3/models"
	"quiz3/repository"
	"quiz3/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else if user.Username == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Username and/or password fields cannot be empty",
		})
	} else {
		user.Id = utils.IDGenerator()
		strId := strconv.Itoa(user.Id)
		user.Password, _ = middleware.HashPassword(user.Password)
		user.CreatedAt = time.Now()
		user.CreatedBy = strId
		user.ModifiedAt = time.Now()
		user.ModifiedBy = strId

		fmt.Println(user)

		repository.CreateUser(user)
		ctx.JSON(http.StatusCreated, gin.H{
			"user": user,
		})
	}
}

func Login(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else if user.Username == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Username and/or password fields cannot be empty",
		})
	} else {
		userData, err := repository.Login(user.Username, user.Password)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			var data models.LoggedIn
			idStr := strconv.Itoa(userData.Id)

			accessToken, e := middleware.GenerateJwt(idStr)

			if e != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": e.Error(),
				})
			} else {
				data.Id = userData.Id
				data.Username = userData.Username
				data.AccessToken = accessToken
				data.TokenExpirationTime = time.Now().Add(time.Hour * 1)

				ctx.JSON(http.StatusOK, gin.H{
					"data": data,
				})
			}
		}
	}
}
