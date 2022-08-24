package database

import (
	"bookms/pkg/config"
	"bookms/pkg/model"
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	// "gorm.io/gorm"
)

var db *gorm.DB

// const DNS = "root:javed@123@/bookms_db?charset=utf8mb4&parseTime=True&loc=Local"

// func Init() {
// 	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Cannot connect to DB")
// 	}
// 	db.AutoMigrate(&model.Book{})

// }

func Init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&model.Book{})
}

//Create Book
func CreateBook(ctx context.Context, newBook model.Book) (uint, error) {
	user := model.Book{Title: newBook.Title, Author: newBook.Author}
	create := db.Create(user)
	if create.Error != nil {
		return 0, create.Error
	}

	return user.ID, nil
}

//Update Book
func UpdateBook(ctx context.Context, updateBook model.Book, id uint64) (uint, error) {

	var book model.Book

	err := db.Where("ID=?", id).Find(&book).Error
	if err != nil {
		return 0, err
	}

	err = db.Model(&book).Updates(updateBook).Error
	if err != nil {
		return 0, err
	}

	return book.ID, nil
}

//Delete Book
func DeleteBook(s string) error {
	tx := db.Delete(s)
	if tx != nil {
		return tx.Error
	}

	return nil
}

//GetAll Books
func GetAllBooks() ([]model.Book, error) {
	var listOfBooks []model.Book
	err := db.Find(&listOfBooks).Error
	if err != nil {
		return nil, err
	}

	if len(listOfBooks) == 0 {
		return nil, errors.New("empty list")
	}

	return listOfBooks, nil
}

// Search Books
func SearchBook(title string, author string) ([]model.Book, error) {

	if title == "" && author == "" {
		return nil, errors.New("nothing to search, empty argment")
	}

	if title == "" {
		err := db.Find(&author).Error
		if err != nil {
			return nil, err
		}
	}
	if author == "" {
		err := db.First(&title).Error
		if err != nil {
			return nil, err
		}
	}

	return nil, errors.New("Invalid")

}
