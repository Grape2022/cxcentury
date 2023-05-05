package uniqueProducttCharacteristic

import "cxcenturySelenium/com/cxcentury/xy/main/utility"

/*
唯一产品标识数据
*/
type DataValue struct {
	UniqueProductMap map[string]map[string]string
}

func (dv DataValue) UniqueProductCharacteristic() *map[string]map[string]string {
	dv.UniqueProductMap = map[string]map[string]string{
		"productBaseData": {
			"sku":         `Grape` + utility.StrcoveFormatintUtility(utility.RandIntUtility(10000000), 8),
			"groupID":     `500202020202020202`,
			"productName": *utility.GenerateGenRandomWorldUtility(),
			"createTime":  *utility.GetTimeUtility(),
		},
		"pdProductPurchase": {},
	}
	return &dv.UniqueProductMap
}
