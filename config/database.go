package config

type SomeDatabase struct {
	Point int
}

var database *SomeDatabase

func Init() {
	// Do stuff like configuring the database

	database = new(SomeDatabase)

	database.Point = 10
}

func GetDatabase() *SomeDatabase {
	return database
}