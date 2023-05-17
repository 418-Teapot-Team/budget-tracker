package budget

const categoriesTable = "categories"

func (Categories) TableName() string {
	return categoriesTable

}

type Categories struct {
	Id        int    `json:"-" gorm:"column:id"`
	Name      string `json:"category" gorm:"column:name"`
	ColorHash string `json:"hash" gorm:"column:color_hash"`
	ImageLink string `json:"imageUrl" gorm:"column:image_link"`
}
