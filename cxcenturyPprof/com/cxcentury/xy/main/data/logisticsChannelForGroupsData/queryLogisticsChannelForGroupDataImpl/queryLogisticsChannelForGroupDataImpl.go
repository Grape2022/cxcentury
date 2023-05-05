package queryLogisticsChannelForGroupDataImpl

/*
渠道查询参数
*/
func QueryLogisticsChannelForGroupsListDataImpl(logisticsCompanyForGroupBodyList map[string][]string) string {
	var logisticsCompanyId string
	//需要取出的货代ID所在数组坐标
	num := 0
	for key, value := range logisticsCompanyForGroupBodyList {
		//取出货代ID
		if key == `logisticsCompanyId` {
			logisticsCompanyId = value[num]
			break
		}
	}

	urlParse := map[string]string{
		`logisticsCompanyId`: logisticsCompanyId,
		`channelNameStr`:     ``,
	}

	return `?logisticsCompanyId=` + urlParse[`logisticsCompanyId`] + `&channelNameStr=` + urlParse[`channelNameStr`]
}
