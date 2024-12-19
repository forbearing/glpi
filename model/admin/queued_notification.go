package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*QueuedNotification]()
}

type QueuedNotification struct {
	pkgmodel.Base
}
