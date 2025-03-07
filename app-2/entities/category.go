package entities

import "time"

type Category struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// import "time"

// type Category struct {
// 	Id        uint
// 	Name      string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// func NewCategory(name string) Category {
// 	currentTime := time.Now()
// 	return Category{
// 		Name:      name,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: currentTime,
// 	}
// }

// ... other methods as needed
