package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Unmanaged]()
}

type Unmanaged struct {
	pkgmodel.Base
}
