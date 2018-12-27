package handler

import (
	"github.com/jinzhu/gorm"
)

// TableXutopia -
type TableXutopia struct {
	gorm.Model
	Nama   string `gorm:"type:varchar(100)"`
	Umur   int
	Alamat string
}

// CreateTable -
func (t TableXutopia) CreateTable(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&TableXutopia{}).Error

	return
}

// InsertData -
func (t *TableXutopia) InsertData(db *gorm.DB) (err error) {
	err = db.Create(&t).Error

	return
}

// UpdateData -
func (t TableXutopia) UpdateData(db *gorm.DB) (err error) {
	data := new(TableXutopia)
	db.Where("id = ?", t.ID).First(&data)

	data.Nama = t.Nama
	data.Alamat = t.Alamat
	data.Umur = t.Umur

	err = db.Save(&data).Error

	return
}

// DeleteData -
func (t TableXutopia) DeleteData(db *gorm.DB) (err error) {
	err = db.Delete(&t).Error

	return

}

// GetDataByName -
func (t TableXutopia) GetDataByName(db *gorm.DB) (res TableXutopia, err error) {

	err = db.Where("nama = ?", t.Nama).First(&res).Error

	return
}

// GetAllData -
func (t TableXutopia) GetAllData(db *gorm.DB) (res []TableXutopia, err error) {

	err = db.Find(&res).Error

	return
}
