package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Chef struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Country           string `json:"country"`
	YearsOfExperience int8   `json:"yearsOfExperience"`
}

type ChefsAndRecipes struct {
	Chef
	Recipes []Recipe `json:"recipes"`
}

func NewChefHandler(ctx *gin.Context) {
	var chef Chef
	if err := ctx.ShouldBindJSON(&chef); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	chef.Id = xid.New().String()
	chefs = append(chefs, chef)

	chefsIndex[chef.Id] = &chef // update index

	ctx.JSON(http.StatusCreated, chef)
}

func ListChefsHandler(ctx *gin.Context) {
	var withRecipes bool
	if wRecipes := ctx.Query("with-recipes"); wRecipes != "" {
		if wRecipes != "no" && wRecipes != "false" {
			withRecipes = true
		}
	}

	if !withRecipes {
		ctx.JSON(http.StatusOK, chefs)
		return
	}

	// List chefs and all their existing recipes
	ChefsAndTheirRecipes := make([]ChefsAndRecipes, len(chefs))

	for i, ctx := range chefs {
		chefRecipes := []Recipe{}
		for _, r := range recipes {
			if r.Chef.Id == ctx.Id {
				chefRecipes = append(chefRecipes, r)
			}
		}

		ChefsAndTheirRecipes[i] = ChefsAndRecipes{
			Chef:    ctx,
			Recipes: chefRecipes,
		}
	}

	ctx.JSON(http.StatusOK, ChefsAndTheirRecipes)
}

func UpdateChefHandler(ctx *gin.Context) {
	var chef Chef
	id := ctx.Param("chef-id")

	if err := ctx.ShouldBindJSON(&chef); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	index := -1
	for i, ctx := range chefs {
		if ctx.Id == id {
			index = i
		}
	}

	if index == -1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Chef not found",
		})
		return
	}

	chef.Id = id
	chefs[index] = chef

	ctx.JSON(http.StatusOK, chef)
}

func DeleteChefHandler(ctx *gin.Context) {
	id := ctx.Param("chef-id")

	index := -1
	for i, ctx := range chefs {
		if ctx.Id == id {
			index = i
		}
	}

	if index == -1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Chef not found",
		})
		return
	}

	if len(chefs) <= 1 {
		chefs = make([]Chef, 0)
	} else {
		chefs = append(chefs[:index], chefs[index+1:]...)
	}

	delete(chefsIndex, id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Chef deleted",
	})
}
