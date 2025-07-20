package controllers

import (
	"database/sql"
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

func PostCategory(ctx *gin.Context) {
	var category models.Category

	userId, accessTokenValidation := middleware.ValidateAccessToken(ctx)

	if accessTokenValidation != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": accessTokenValidation,
		})
	} else {
		if err := ctx.ShouldBindJSON(&category); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if category.Name == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Name field cannot be empty",
			})
		} else {
			fmt.Println("User ID:", userId)
			category.Id = utils.IDGenerator()
			category.CreatedAt = time.Now()
			category.CreatedBy = userId
			category.ModifiedAt = time.Now()
			category.ModifiedBy = userId
			fmt.Println(category)

			repository.CreateCategory(category)
			ctx.JSON(http.StatusCreated, gin.H{
				"category": category,
			})
		}
	}
}

func GetCategories(ctx *gin.Context) {
	categories, err := repository.GetCategories()

	_, accessTokenValidation := middleware.ValidateAccessToken(ctx)

	if accessTokenValidation != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": accessTokenValidation,
		})
	} else {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"categories": categories,
			})
		}
	}
}

func GetCategoryById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	_, accessTokenValidation := middleware.ValidateAccessToken(ctx)

	if accessTokenValidation != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": accessTokenValidation,
		})
	} else {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
		} else {
			category, err := repository.GetCategoryById(id)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, gin.H{
						"error": "category doesn't exist",
					})
				} else {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
				}
				return
			} else {
				ctx.JSON(http.StatusOK, category)
			}
		}
	}
}

func GetBooksByCategoryId(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	_, accessTokenValidation := middleware.ValidateAccessToken(ctx)

	if accessTokenValidation != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": accessTokenValidation,
		})
	} else {
		books, err := repository.GetBooksByCategoryId(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"books": books,
			})
		}
	}
}

func UpdateCategory(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	userId, accessTokenValidation := middleware.ValidateAccessToken(ctx)

	if accessTokenValidation != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": accessTokenValidation,
		})
	} else {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
			return
		} else {
			var input models.Category

			if err := ctx.ShouldBindJSON(&input); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error":   "Invalid input",
					"details": err.Error(),
				})
				return
			} else {
				input.ModifiedBy = userId
				input.ModifiedAt = time.Now()
				rowsUpdated, err := repository.UpdateCategory(id, input)

				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"error":   "Failed to update category",
						"details": err.Error(),
					})
					return
				} else if rowsUpdated == 0 {
					ctx.JSON(http.StatusNotFound, gin.H{
						"message": "No category found with the given ID",
					})
					return
				} else {
					ctx.JSON(http.StatusOK, gin.H{
						"message": "category was successfully updated",
						"rows":    rowsUpdated,
					})
				}
			}
		}
	}
}

func DeleteCategory(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	_, accessTokenValidation := middleware.ValidateAccessToken(ctx)

	if accessTokenValidation != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": accessTokenValidation,
		})
	} else {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
			return
		} else {
			rowsDeleted, err := repository.DeleteCategory(id)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error":  "Failed to delete category",
					"detais": err.Error(),
				})
				return
			} else if rowsDeleted == 0 {
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": "No category found with the given ID",
				})
				return
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "category has been deleted",
					"rows":    rowsDeleted,
				})
			}
		}
	}
}
