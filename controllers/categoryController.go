package controllers

import (
	"net/http"
	"social-journal/initializers"
	"social-journal/models"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var body struct {
		Title    string
		Journals []models.Journal
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Get the user's id to add the category with
	user, _ := c.Get("user")

	category := models.Category{
		Title:  body.Title,
		UserID: user.(models.User).ID,
	}

	result := initializers.DB.Create(&category)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"New Category": category,
	})
}

// Get all categories

func GetAllCategories(c *gin.Context) {
	var categories []models.Category

	result := initializers.DB.Preload("Journals").Find(&categories)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Categories": categories,
	})
}

// Get category by ID
func GetCategoryByID(c *gin.Context) {
	// Get ID of the category
	id := c.Param("id")

	var category models.Category

	result := initializers.DB.Preload("Journals").First(&category, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Category": category,
	})
}

func UpdateCategory(c *gin.Context) {
	// Get the id of the post
	id := c.Param("id")

	// Get the data of the req body
	var body struct {
		Title string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Find the article we are updating
	var category models.Category
	initializers.DB.First(&category, id)
	user, _ := c.Get("user")

	if category.UserID != user.(models.User).ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden to edit other people's categories..."})
		return
	}

	// Update attributes with `struct`, will only update non-zero fields
	initializers.DB.Model(&category).Updates(models.Category{
		Title:  body.Title,
		UserID: user.(models.User).ID,
	})

	c.JSON(http.StatusOK, category)
}

// Delete category
func DeleteCategory(c *gin.Context) {
	// Get category id to delete
	id := c.Param("id")

	// Get the task itself
	var category models.Category

	result := initializers.DB.First(&category, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Get the owner of the category
	user, _ := c.Get("user")

	// Delete on if it belongs to the user.
	if category.UserID != user.(models.User).ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden to delete other peoples categories"})
		return
	}

	initializers.DB.Delete(&models.Category{}, id)

	// Respond
	c.Status(http.StatusNoContent)
	c.JSON(200, gin.H{
		"post": "A category was deleted...",
	})
}
