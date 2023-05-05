package queryProductBaseDataImpl

import "fmt"

// Product Prand
func (dv DataValue) QueryProductBaseListDataPath() *DataValue {
	dv.PageNum = 1
	dv.PageSize = 20
	dv.IsAsc = `asc`
	dv.OrderByColumn = `createTime`
	dv.StockTypeTab = `sku`
	dv.OrderByKey = `createTime`
	dv.OrderByType = `2`
	dv.QueryValue = fmt.Sprintf("?pageNum=%d&pageSize=%d&isAsc=%s&orderByColumn=%s&isCombo=%s&status=%s&brandId=%s&searchColumn=%s&searchValue=%s&jointQuery=%s&stockTypeTab=%s&brandName=%s&categoryName=%s&productName=%s&orderByKey=%s&orderByType=%s",
		dv.PageNum,
		dv.PageSize,
		dv.IsAsc,
		dv.OrderByColumn,
		dv.IsCombo,
		dv.Status,
		dv.BrandId,
		dv.SearchColumn,
		dv.SearchValue,
		dv.JointQuery,
		dv.StockTypeTab,
		dv.BrandName,
		dv.CategoryName,
		dv.ProductName,
		dv.OrderByKey,
		dv.OrderByType,
	)
	return &dv
}

// Product
func (dv DataValue) QueryProductBaseListData() DataValue {
	dv.PageNum = 1
	dv.PageSize = 20
	dv.IsAsc = `asc`
	dv.OrderByColumn = `createTime`
	dv.StockTypeTab = `sku`
	dv.OrderByKey = `createTime`
	dv.OrderByType = `2`
	dv.QueryValue = fmt.Sprintf("?pageNum=%d&pageSize=%d&isAsc=%s&orderByColumn=%s&isCombo=%s&status=%s&brandId=%s&searchColumn=%s&searchValue=%s&jointQuery=%s&stockTypeTab=%s&brandName=%s&categoryName=%s&productName=%s&orderByKey=%s&orderByType=%s",
		dv.PageNum,
		dv.PageSize,
		dv.IsAsc,
		dv.OrderByColumn,
		dv.IsCombo,
		dv.Status,
		dv.BrandId,
		dv.SearchColumn,
		dv.SearchValue,
		dv.JointQuery,
		dv.StockTypeTab,
		dv.BrandName,
		dv.CategoryName,
		dv.ProductName,
		dv.OrderByKey,
		dv.OrderByType,
	)
	return dv
}

func (dv DataValue) QueryProducBaseDataPresstStatus() DataValue {
	dv.PageNum = 1
	dv.PageSize = 20
	dv.IsAsc = `asc`
	dv.OrderByColumn = `createTime`
	dv.StockTypeTab = `sku`
	dv.QueryValue = fmt.Sprintf("?pageNum=%d&pageSize=%d&isAsc=%s&orderByColumn=%s&isCombo=%s&status=%s&brandId=%s&searchColumn=%s&searchValue=%s&jointQuery=%s&stockTypeTab=%s&brandName=%s&categoryName=%s&productName=%s",
		dv.PageNum,
		dv.PageSize,
		dv.IsAsc,
		dv.OrderByColumn,
		dv.IsCombo,
		dv.Status,
		dv.BrandId,
		dv.SearchColumn,
		dv.SearchValue,
		dv.JointQuery,
		dv.StockTypeTab,
		dv.BrandName,
		dv.CategoryName,
		dv.ProductName,
	)
	return dv
}
