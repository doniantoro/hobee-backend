package config

import (
	"hobee/model"
)

func Migration() {

	db := ConnectionDatabase()
	db.Migrator().DropTable(model.User{},model.Category{},model.Day{},model.Facility{},model.Operational{},model.Order{},model.Price{},model.Review{},model.ReviewAttachment{},model.Venue{})
	db.AutoMigrate(model.User{},model.Category{},model.Day{},model.Facility{},model.Operational{},model.Order{},model.Price{},model.Review{},model.ReviewAttachment{},model.Venue{})

}