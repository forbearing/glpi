package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Report]()
}

type Report struct {
	pkgmodel.Base
}
