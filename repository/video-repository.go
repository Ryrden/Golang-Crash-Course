package repository

import (
	"github.com/glebarez/sqlite" // not based on CGO (Linux only)
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	// err := db.connection.Close()
	// if err != nil {
	// 	panic("Failed to close database")
	// }
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Video) {
	//FIXME: If update internal object, this will not work
	db.connection.Preload("Author").Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	//auto_preload means that it will also fetch the author details
	db.connection.Preload("Author").Find(&videos)
	return videos
}
