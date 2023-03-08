package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type CreateRecipeBody struct {
	Name         string   `json:"name"`
	Keywords     []string `json:"keywords"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	ChefId       string   `json:"chefId"`
}

type Recipe struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Keywords     []string  `json:"keywords"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
	Chef         *Chef     `json:"chef"`
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
	var body CreateRecipeBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if body.ChefId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing required field 'chefId'",
		})
		return
	}

	chef, e := chefsIndex[body.ChefId]
	if !e {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Specified chef not found",
		})
		return
	}

	recipe := Recipe{
		Id:           xid.New().String(),
		PublishedAt:  time.Now(),
		Name:         body.Name,
		Keywords:     body.Keywords,
		Ingredients:  body.Ingredients,
		Instructions: body.Instructions,
		Chef:         chef,
	}
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}
