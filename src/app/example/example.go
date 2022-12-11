package example

type Example struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	example1 string `gorm:"type:varchar(300)" json:"example"`
	example2 string `gorm:"type:text" json:"deskripsi"`
}
