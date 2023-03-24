package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"time"
)

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
}

type Recipe struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Chef         Chef      `json:"chef"`
	Keywords     []string  `json:"keywords"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

type Chef struct {
	Name              string `json:"name"`
	Country           string `json:"country"`
	YearsOfExperience string `json:"yearsOfExperience"`
}

func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("recipe-id")
	index := -1

	for i := 0; i < len(recipes); i++ {
		if recipes[i].Id == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found",
		})
		return
	}

	recipes = append(recipes[:index], recipes[index+1:]...)

	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe deleted",
	})
}

func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("recipe-id")

	var recipe Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	index := -1

	for i := 0; i < len(recipes); i++ {
		if recipes[i].Id == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found",
		})
		return
	}

	recipe.Id = id
	recipes[index] = recipe

	c.JSON(http.StatusOK, recipe)
}

func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if recipe.Chef.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe must have a Chef name"})
		return
	}
	recipe.Id = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.PUT("recipes/:recipe-id", UpdateRecipeHandler)
	router.DELETE("recipes/:recipe-id", DeleteRecipeHandler)
	router.Run()
}
