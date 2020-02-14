package model

import (
	"fmt"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var _db *gorm.DB

func db() (db *gorm.DB, err error) {
	if _db == nil {
		db, err := gorm.Open("mysql", "root:1234@/shop?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			return db, err
		}
		_db = db
	}
	return _db, nil
}
func Add(m interface{}) {
	db, err := db()
	if err != nil {
		panic("failed to connect database")
	}

	if ret := db.Table("product_group").Create(m); ret.Error != nil {
		fmt.Println(ret.Error)
	}
}
func Get(m interface{}, query interface{}, args ...interface{}) (ok bool) {
	db, err := db()
	if err != nil {
		panic("failed to connect database")
		return false
	}
	// var group ProductGroup
	// db = db.Table("product_group")
	if query != nil {
		db.Where(query, args...)
	} else {
		v := reflect.ValueOf(m)
		db.Where("id = ?", v.Elem().FieldByName("ID").Uint())
	}
	if db.First(m).RecordNotFound() {
		return false
	}
	return true
	// db.First(&product, "code = ?", "L1212") // find product with code l1212
}
func Update(m interface{}) {

}
func Del(query interface{}, args ...interface{}) {
	// db, err := gorm.Open("mysql", "root:1234@/shop?charset=utf8&parseTime=True&loc=Local")
	// db.Where(query, args).Delete()
}
