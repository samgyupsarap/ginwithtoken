package models

// User represents the user model with only the necessary fields
type Sample struct {
    ID   uint   `json:"id" gorm:"primaryKey"` // You can define ID as primary key
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func (Sample) TableName() string {
    return "sample" // Specify the exact table name
}