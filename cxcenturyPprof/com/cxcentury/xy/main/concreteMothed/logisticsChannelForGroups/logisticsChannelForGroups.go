package logisticsChannelForGroups

import "cxcenturyPprof/com/cxcentury/xy/main/concreteMothed/logisticsChannelForGroups/queryLogisticsChannelForGroupsImpl"

func QueryLogisticsChannelForGroupsList(token, ChannelForGroupData string) map[string][]string {
	path := `/logisticsChannelForgroups`
	return queryLogisticsChannelForGroupsImpl.QueryLogisticsChannelForGroupsListImpl(path+`/list`, ChannelForGroupData, token)
}
