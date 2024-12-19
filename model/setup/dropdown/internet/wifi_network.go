package model // WifiNetwork 无线网络配置
import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*WifiNetwork]()
}

type WifiNetwork struct {
	model.NameEntity

	// WiFi配置
	ESSID *string          `json:"essid"` // 扩展服务集标识(网络名称)
	Mode  *WifiNetworkMode `json:"mode"`  // 工作模式
}

func (*WifiNetwork) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*WifiNetwork]())
}
