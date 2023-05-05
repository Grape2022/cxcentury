package queryLogisticsCompanyForGroupsDataImpl

// 准备查询货代数据
func QueryLogisticsCompanyForGroupDataImpl() string {
	urlParse := map[string]string{
		`companyNames`:  ``,
		`isAsc`:         `desc`,
		`orderByColumn`: `createTime`,
	}

	return `?companyNames=` + urlParse[`companyNames`] + `&isAsc=` + urlParse[`isAsc`] + `&orderByColumn=` + urlParse[`orderByColumn`]
}
