package models

import (
	"jira-clone/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Issue ..
type Issue struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Reporter    string `json:"reporter"`
	Timelog     int    `json:"timelog"`
	Estimated   int    `json:"estimated"`
	Progress    string `json:"progress"`
	Priority    string `json:"priority"`
	UserID      uint
	Comments    []Comment
}

func GetIssues(c *fiber.Ctx) {
	db := database.DBConn
	var issues []Issue
	db.Preload("Comments").Find(&issues)
	c.JSON(issues)
}

func GetIssue(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var issue Issue
	db.Preload("Comments").Find(&issue, id)
	c.JSON(issue)
}

func NewIssue(c *fiber.Ctx) {
	db := database.DBConn
	issue := new(Issue)
	if err := c.BodyParser(issue); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&issue)
	c.JSON(issue)
}

func DeleteIssue(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var issue Issue
	db.First(&issue, id)
	if issue.Title == "" {
		c.Status(500).Send("No issue Found with ID")
		return
	}
	db.Delete(&issue)
	c.Send("Issue Successfully deleted")
}
