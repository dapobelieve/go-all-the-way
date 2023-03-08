package main

import (
	"github.com/gin-gonic/gin"
)

var chefs []Chef
var chefsIndex map[string]*Chef
var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
	chefs = make([]Chef, 0)
	chefsIndex = make(map[string]*Chef)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.PUT("recipes/:recipe-id", UpdateRecipeHandler)
	router.DELETE("recipes/:recipe-id", DeleteRecipeHandler)

	router.POST("/chefs", NewChefHandler)
	router.GET("/chefs", ListChefsHandler)
	router.PUT("chefs/:chef-id", UpdateChefHandler)
	router.DELETE("chefs/:chef-id", DeleteChefHandler)
	// GET /chefs?with-recipes=true
	router.Run()
}
