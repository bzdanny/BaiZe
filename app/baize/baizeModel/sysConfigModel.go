package baizeModel
import (
	"baize/app/common/commonModels"
	"github.com/gogf/gf/util/gconv"

)

// ==========================================================================
// date：0001-01-01 00:00:00
// author：baize
// version: v1.0
// ==========================================================================

type SysConfigVo struct {
	ConfigId    int64    `json:"configId,string" db:"config_id"`    // 参数主键
	ConfigName    string    `json:"configName" db:"config_name"`    // 参数名称
	ConfigKey    *int64    `json:"configKey" db:"config_key"`    // 参数键名
	ConfigValue    *string    `json:"configValue" db:"config_value"`    // 参数键值
	ConfigType    *string    `json:"configType" db:"config_type"`    // 系统内置（Y是 N否）
	Remark    *string    `json:"remark" db:"remark"`    // 备注

	commonModels.BaseEntity
}

type SysConfigDQL struct {
	ConfigName    string    `form:"configName" db:"config_name"`    // 参数名称
	ConfigKey    *int64    `form:"configKey" db:"config_key"`    // 参数键名
	ConfigValue    string    `form:"configValue" db:"config_value"`    // 参数键值
	ConfigType    string    `form:"configType" db:"config_type"`    // 系统内置（Y是 N否）

	commonModels.BaseEntityDQL
}

type SysConfigDML struct {
	ConfigId    int64    `json:"configId,string" db:"config_id"`    // 参数主键
	ConfigName    string    `json:"configName" db:"config_name"`    // 参数名称
	ConfigKey    *int64    `json:"configKey" db:"config_key"`    // 参数键名
	ConfigValue    string    `json:"configValue" db:"config_value"`    // 参数键值
	ConfigType    string    `json:"configType" db:"config_type"`    // 系统内置（Y是 N否）
	Remark    string    `json:"remark" db:"remark"`    // 备注

	commonModels.BaseEntityDML
}

func SysConfigListToRows(sysConfigs []*SysConfigVo) (rows [][]string) {
	rows = make([][]string, 0, len(sysConfigs)+1)
	row1 := []string{
		"参数名称",
		"参数键名",
		"参数键值",
		"系统内置（Y是 N否）",
		"备注",
	}
	rows = append(rows, row1)
	for _, data := range sysConfigs {
		row := make([]string, 0)
		row=append(row, gconv.String(data.ConfigName ))
		row=append(row, gconv.String(data.ConfigKey ))
		row=append(row, gconv.String(data.ConfigValue ))
		row=append(row, gconv.String(data.ConfigType ))
		row=append(row, gconv.String(data.Remark ))

		rows = append(rows, row)
	}
	return
}

