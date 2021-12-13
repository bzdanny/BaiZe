package excelDictionaries

import (
	"baize/app/system/service/systemService"
	"baize/app/system/service/systemService/systemServiceImpl"
)

var iDictData systemService.IDictDataService = systemServiceImpl.GetDictDataService()

func ValueToLabel(dictType string,value string)string  {
	sysDictDataList := iDictData.SelectDictDataByType(dictType)
	for _, dictData := range sysDictDataList {
		if dictData. DictValue==value{
			return dictData.DictLabel
		}
	}
	return ""
}