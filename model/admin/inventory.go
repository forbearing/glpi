package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Inventory]()
}

type Inventory struct {
	pkgmodel.Base
}
