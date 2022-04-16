package controllers

import (
	"filmes-crud/database"
	"filmes-crud/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

//Get para pegar pelo id o FILME
func ShowMoviebyID(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	db := database.Getdatabase()

	var movie model.Movie
	err = db.First(&movie, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find Movie:" + err.Error(),
		})
		return
	}
	c.JSON(200, movie)
}

//Listar todos os Filmes
func ShowMovies(c *gin.Context) {
	db := database.Getdatabase()

	var movie []model.Movie
	err := db.Find(&movie).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON:" + err.Error(),
		})
		return
	}

	c.JSON(200, movie)
}

//Criacao de Filme
func CreateMovie(c *gin.Context) {
	db := database.Getdatabase()

	var movie model.Movie

	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list Movie:" + err.Error(),
		})
		return
	}

	err = db.Create(&movie).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find JSON:" + err.Error(),
		})
		return
	}
	c.JSON(200, movie)
}

// Atualizado o movie
func UpdateMovie(c *gin.Context) {
	db := database.Getdatabase()

	var movie model.Movie

	err := c.ShouldBindJSON(&movie)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list Movie:" + err.Error(),
		})
		return
	}

	err = db.Save(&movie).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update:" + err.Error(),
		})
		return
	}
	c.JSON(200, movie)
}

//deletar Movie
func DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}
	db := database.Getdatabase()

	err = db.Delete(&model.Movie{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete movie:" + err.Error(),
		})
		return
	}
	c.Status(204)

}
