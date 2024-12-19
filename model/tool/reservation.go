package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Reservation]()
}

type Reservation struct {
	pkgmodel.Base
}
