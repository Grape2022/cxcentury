package queryLogisticsChannelForGroupsImpl

import (
	"cxcenturyPprof/com/cxcentury/xy/main/utility"
	"github.com/tidwall/gjson"
)

/*
渠道查询接口实现
*/
func QueryLogisticsChannelForGroupsListImpl(path, ChannelForGroupData, token string) map[string][]string {
	var bodyMap = make(map[string][]string)

	body := utility.MethodGetUtility(path, ChannelForGroupData, token, `QueryLogisticsChannelForGroupsListUrl`)

	for _, value := range gjson.Get(string(body), `data`).Array() {
		bodyMap[`logisticsChannelId`] = append(bodyMap[`logisticsChannelId`], gjson.Get(value.String(), `logisticsChannelId`).String())
		bodyMap[`logisticsChannelForgroupId`] = append(bodyMap[`logisticsChannelForgroupId`], gjson.Get(value.String(), `logisticsChannelForgroupId`).String())
		bodyMap[`logisticsCompanyId`] = append(bodyMap[`logisticsCompanyId`], gjson.Get(value.String(), `logisticsCompanyId`).String())
		bodyMap[`companyName`] = append(bodyMap[`companyName`], gjson.Get(value.String(), `companyName`).String())
		bodyMap[`companyShortName`] = append(bodyMap[`companyShortName`], gjson.Get(value.String(), `companyShortName`).String())
		bodyMap[`sendProv`] = append(bodyMap[`sendProv`], gjson.Get(value.String(), `sendProv`).String())
		bodyMap[`sendCity`] = append(bodyMap[`sendCity`], gjson.Get(value.String(), `sendCity`).String())
		bodyMap[`destinationCountry`] = append(bodyMap[`destinationCountry`], gjson.Get(value.String(), `destinationCountry`).String())
		bodyMap[`channelName`] = append(bodyMap[`channelName`], gjson.Get(value.String(), `channelName`).String())
		bodyMap[`transportType`] = append(bodyMap[`transportType`], gjson.Get(value.String(), `transportType`).String())
		bodyMap[`sendDay`] = append(bodyMap[`sendDay`], gjson.Get(value.String(), `sendDay`).String())
		bodyMap[`volumeWeightRadio`] = append(bodyMap[`volumeWeightRadio`], gjson.Get(value.String(), `volumeWeightRadio`).String())
		bodyMap[`shipmentName`] = append(bodyMap[`shipmentName`], gjson.Get(value.String(), `shipmentName`).String())
		bodyMap[`feeMemo`] = append(bodyMap[`feeMemo`], gjson.Get(value.String(), `feeMemo`).String())
		bodyMap[`customsFees`] = append(bodyMap[`customsFees`], gjson.Get(value.String(), `customsFees`).String())
		bodyMap[`logisticsChannelAliasId`] = append(bodyMap[`logisticsChannelAliasId`], gjson.Get(value.String(), `logisticsChannelAliasId`).String())
		bodyMap[`channelAliasName`] = append(bodyMap[`channelAliasName`], gjson.Get(value.String(), `channelAliasName`).String())
		bodyMap[`companyChannelName`] = append(bodyMap[`companyChannelName`], gjson.Get(value.String(), `companyChannelName`).String())
	}
	bodyMap[`msg`] = append(bodyMap[`msg`], gjson.Get(string(body), `msg`).String())
	bodyMap[`code`] = append(bodyMap[`code`], gjson.Get(string(body), `code`).String())
	//fmt.Println(string(body))
	return bodyMap
}
