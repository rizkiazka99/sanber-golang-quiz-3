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

func PostBook(ctx *gin.Context) {
	var book models.Book

	userId, accessTokenValidation := middleware.ValidateAccessToken(ctx)

	if accessTokenValidation != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": accessTokenValidation,
		})
	} else {
		if err := ctx.ShouldBindJSON(&book); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			releaseYearValidation := utils.ReleaseYearValidator(book.ReleaseYear)

			if releaseYearValidation != "" {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": releaseYearValidation,
				})
			} else {
				fmt.Println("User ID:", userId)
				book.Id = utils.IDGenerator()
				book.Thickness = utils.ThicknessChecker(book.TotalPage)
				book.CreatedAt = time.Now()
				book.CreatedBy = userId
				book.ModifiedAt = time.Now()
				book.ModifiedBy = userId
				fmt.Println(book)

				repository.CreateBook(book)
				ctx.JSON(http.StatusCreated, gin.H{
					"Book": book,
				})
			}
		}
	}
}

func GetBooks(ctx *gin.Context) {
	Books, err := repository.GetBooks()

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
				"Books": Books,
			})
		}
	}
}

func GetBookById(ctx *gin.Context) {
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
			Book, err := repository.GetBookById(id)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, gin.H{
						"error": "Book doesn't exist",
					})
				} else {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
				}
				return
			} else {
				ctx.JSON(http.StatusOK, Book)
			}
		}
	}
}

func UpdateBook(ctx *gin.Context) {
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
			var input models.Book

			if err := ctx.ShouldBindJSON(&input); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error":   "Invalid input",
					"details": err.Error(),
				})
				return
			} else {
				releaseYearValidation := utils.ReleaseYearValidator(input.ReleaseYear)

				if releaseYearValidation != "" {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"error": releaseYearValidation,
					})
				} else {
					input.Thickness = utils.ThicknessChecker(input.TotalPage)
					input.ModifiedBy = userId
					input.ModifiedAt = time.Now()
					rowsUpdated, err := repository.UpdateBook(id, input)

					if err != nil {
						ctx.JSON(http.StatusInternalServerError, gin.H{
							"error":   "Failed to update Book",
							"details": err.Error(),
						})
						return
					} else if rowsUpdated == 0 {
						ctx.JSON(http.StatusNotFound, gin.H{
							"message": "No Book found with the given ID",
						})
						return
					} else {
						ctx.JSON(http.StatusOK, gin.H{
							"message": "Book was successfully updated",
							"rows":    rowsUpdated,
						})
					}
				}
			}
		}
	}
}

func DeleteBook(ctx *gin.Context) {
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
			rowsDeleted, err := repository.DeleteBook(id)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error":  "Failed to delete Book",
					"detais": err.Error(),
				})
				return
			} else if rowsDeleted == 0 {
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": "No Book found with the given ID",
				})
				return
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Book has been deleted",
					"rows":    rowsDeleted,
				})
			}
		}
	}
}
