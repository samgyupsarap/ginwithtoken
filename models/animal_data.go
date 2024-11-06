package models

type Animals struct {
	AnimalName string `json:"animal_name"`
	Age        int    `json:"age"`
	Color      string `json:"color"`
}

func (Animals) TableName() string {
    return "sample_data" // Specify the exact table name
}