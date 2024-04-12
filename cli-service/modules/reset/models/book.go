package models

import (
	"errors"
	"gorm.io/gorm"
)

type Book struct {
	Model
	ISBN    string `json:"isbn"`
	Penulis string `json:"penulis"`
	Tahun   uint   `json:"tahun"`
	Judul   string `json:"judul"`
	Gambar  string `json:"gambar"`
	Stok    uint   `json:"stok"`
}

func (book *Book) Upsert(db *gorm.DB) error {
	var b Book
	result := db.Model(Book{}).Where("isbn = ?", book.ISBN).Take(&b)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result := db.Create(book)
		if result.Error != nil {
			return result.Error
		}
	} else {
		result := db.Model(&b).Updates(book)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
