package logisticsCompanyForGroups

import (
	"cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/logisticsCompanyForGroups/querylogisticsCompanyForGroupsImpl"
)

/*
获取货代接口
*/
func QuerylogisticsCompanyForGroupsList(token, CompanyForGroup string) map[string][]string {
	path := `/logisticsCompanyForgroups`
	//调用获取货代接口
	bodyMap := querylogisticsCompanyForGroupsImpl.QueryLogisticsCompanyForGroupsListImpl(path+`/list`, CompanyForGroup, token)

	return bodyMap
}
