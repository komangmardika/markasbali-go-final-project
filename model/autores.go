package model

import "gorm.io/gorm"

type AutoRes struct {
	Model
	Id              int    `gorm:"primarykey;autoIncrement" json:"uuid"`
	Nama_Database   string `gorm:"primarykey" json:"nama_database"`
	Nama_File_Backup string `gorm:"not null" json:"nama_file_backup"`
}

func (cr *AutoRes) Create(db *gorm.DB) (*AutoRes, error) {
	err := db.
		Model(AutoRes{}).
		Create(cr).
		Error

	if err != nil {
		return nil, err
	}

	return cr, nil
}

func (cr *AutoRes) Save(db *gorm.DB) error {
	// Assuming `id` is the primary key of the Car struct
	err := db.Save(&cr).Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *AutoRes) GetAll(db *gorm.DB) ([]AutoRes, error) {
	res := [] AutoRes{}
	
	err := db.Model(AutoRes{}).Find(&res).Error

	if err != nil{
		return []AutoRes{}, err
	}

	return res, nil
}

func (cr *AutoRes) GetAllByDbName(db *gorm.DB) ([]AutoRes, error) {
	res := [] AutoRes{}
	
	err := db.Model(AutoRes{}).Where("nama_database = ?",cr.Nama_Database).Find(&res).Error

	if err != nil{
		return []AutoRes{}, err
	}

	return res, nil
}


func (cr *AutoRes) GetDistinct(db *gorm.DB) ([]string, error) {
    var res [] string
    
    // Membuat query SQL mentah
    query := "SELECT DISTINCT nama_database from `auto_res`"

    // Mengeksekusi query menggunakan Raw SQL di GORM
    if err := db.Raw(query).Scan(&res).Error; err != nil {
        return []string{}, err
    }

    return res, nil
}

func (cr *AutoRes) GetLatestByDBName(db *gorm.DB) (AutoRes, error) {
    res := AutoRes{}
	
	// Membuat query SQL mentah
    query := "SELECT * FROM auto_res " +
             "WHERE nama_database = ? " +
             "AND created_at = (SELECT MAX(created_at) FROM auto_res WHERE nama_database = ?)"

    // Mengeksekusi query menggunakan Raw SQL di GORM
    if err := db.Raw(query, cr.Nama_Database, cr.Nama_Database).Scan(&res).Error; err != nil {
        return AutoRes{}, err
    }

	return res, nil
}
