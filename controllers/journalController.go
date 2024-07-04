package controllers

import (
	"net/http"
	"social-journal/initializers"
	"social-journal/models"

	"github.com/gin-gonic/gin"
)

func PostJournal(c *gin.Context) {

	categoryID := c.Param("id")
	var category models.Category

	myCategory := initializers.DB.Preload("Journals").First(&category, categoryID)

	if myCategory.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Category not found",
		})
		return
	}

	var body struct {
		Title    string
		Content  string
		ImageUrl string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read request body",
		})
		return
	}

	user, _ := c.Get("user")

	journal := models.Journal{
		Title:      body.Title,
		Content:    body.Content,
		ImageUrl:   body.ImageUrl,
		CategoryID: category.ID,
		UserID:     user.(models.User).ID,
	}

	result := initializers.DB.Create(&journal)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"New JOurnal": journal,
	})
}

func GetAllJournals(c *gin.Context) {
	// Get owner of the category
	userID, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var journals []models.Journal

	result := initializers.DB.Where("user_id = ?", userID.(models.User).ID).Find(&journals)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Journals": journals,
	})
}

func GetAllJournalsByCategory(c *gin.Context) {
	// Get owner of the category
	userID, exists := c.Get("user")

	categoryID := c.Param("id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var journals []models.Journal

	result := initializers.DB.Where("user_id = ? AND category_id = ?", userID.(models.User).ID, categoryID).Find(&journals)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Journals": journals,
	})
}

func GetJournalByID(c *gin.Context) {
	// Get owner of the category
	userID, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	// Get ID of the journal
	id := c.Param("id")

	var journal models.Journal

	result := initializers.DB.Where("user_id = ?", userID.(models.User).ID).First(&journal, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Journal": journal,
	})
}

func UpdateJournal(c *gin.Context) {
	// Get the id of the post
	id := c.Param("id")

	// Get the data of the req body
	var body struct {
		Title    string
		Content  string
		ImageUrl string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Find the journal we are updating
	var journal models.Journal
	initializers.DB.First(&journal, id)
	user, _ := c.Get("user")

	if journal.UserID != user.(models.User).ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden to edit other people's journals..."})
		return
	}

	// Update attributes with `struct`, will only update non-zero fields
	initializers.DB.Model(&journal).Updates(models.Journal{
		Title:    body.Title,
		Content:  body.Content,
		ImageUrl: body.ImageUrl,
		UserID:   user.(models.User).ID,
	})

	c.JSON(http.StatusOK, journal)
}

func DeleteJournal(c *gin.Context) {
	// Get journal id to delete
	id := c.Param("id")

	// Get the task itself
	var journal models.Journal

	result := initializers.DB.First(&journal, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Journal not found"})
		return
	}

	// Get the owner of the category
	user, _ := c.Get("user")

	// Delete on if it belongs to the user.
	if journal.UserID != user.(models.User).ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden to delete other peoples journals"})
		return
	}

	initializers.DB.Delete(&models.Journal{}, id)

	// Respond
	c.Status(http.StatusNoContent)
	c.JSON(200, gin.H{
		"message": "A journal was deleted...",
	})
}
