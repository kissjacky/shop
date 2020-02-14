package product

import (
	"shop-gin/model"
)

type ProductGroup struct {
	model.BaseModel
	Name         string
	Richtext     string
	DisplayOrder int
}

func AddGroup() {
	m := &ProductGroup{Name: "L1212", DisplayOrder: 1000}
	model.Add(m)
}
