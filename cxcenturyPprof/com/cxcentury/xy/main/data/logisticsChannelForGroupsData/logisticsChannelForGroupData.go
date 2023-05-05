package logisticsChannelForGroupsData

import (
	"cxcenturyPprof/com/cxcentury/xy/main/data/logisticsChannelForGroupsData/queryLogisticsChannelForGroupDataImpl"
)

// 查询货代接口调用
func QueryLogisticsChannelForGroupListData(logisticsCompanyForGroupBodyList map[string][]string) string {

	return queryLogisticsChannelForGroupDataImpl.QueryLogisticsChannelForGroupsListDataImpl(logisticsCompanyForGroupBodyList)
}
