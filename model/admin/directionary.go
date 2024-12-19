package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Directionary]()
}

type Directionary struct {
	pkgmodel.Base
}
