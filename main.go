package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// User represents the user entity.
type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Games    []Game `gorm:"many2many:user_games"` // Many-to-Many relationship with games
}

// Game represents the game entity.
type Game struct {
	gorm.Model
	Title       string
	Description string
	AdminID     uint // Foreign Key to the User who created the game
	Questions   []Question
	Users       []User `gorm:"many2many:user_games"` // Many-to-Many relationship with users
}

// Question represents the question entity.
type Question struct {
	gorm.Model
	Text            string
	CorrectAnswerID uint // Foreign Key to Answer
	GameID          uint // Foreign Key to Game
	Answers         []Answer
}

// Answer represents the answer entity.
type Answer struct {
	gorm.Model
	Text       string
	QuestionID uint // Foreign Key to Question
}

// UserGame represents the many-to-many relationship between User and Game.
type UserGame struct {
	gorm.Model
	UserID uint
	GameID uint
	Score  int
}

var db *gorm.DB

func initDB() {
	// Replace with your MySQL database connection details
	dsn := "root:root@tcp(localhost:3306)/golang-game?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Enable foreign key constraints for MySQL
	db.Exec("SET foreign_key_checks = 1")

	// AutoMigrate will create or update the database tables based on the struct definitions.
	err = db.AutoMigrate(&User{}, &Game{}, &Question{}, &Answer{})
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	initDB()

	// You can now use GORM to interact with the database, create, read, update, and delete records.
}
