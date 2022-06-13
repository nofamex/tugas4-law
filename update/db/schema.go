package db

type User struct {
	Name		string  `gorm:"column:name"`
	NPM         string  `gorm:"column:npm"`
}