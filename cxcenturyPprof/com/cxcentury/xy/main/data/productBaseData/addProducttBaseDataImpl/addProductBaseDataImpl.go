package addProducttBaseDataImpl

import (
	"cxcenturyPprof/com/cxcentury/xy/main/utility"
)

/*
添加产品数据模板
*/
func AddProductBaseDataImpl(productDataList map[string]map[string][]string, uniqueProductCList map[string]map[string]string) (data string) {

	//productInfo,productInfoDeclarationData,productInfoShipmentTypes:= make(map[string][]string),make(map[string][]string),make(map[string][]string)
	productInfo, profuctInfoDeclarationData, profuctInfoShipmentTypes := make(map[string][]string), make(map[string][]string), make(map[string][]string)
	productData := productDataList
	dataList := uniqueProductCList

	productInfo = productData[`productInfo`]
	profuctInfoDeclarationData = productData[`profuctInfoDeclarationData`]
	profuctInfoShipmentTypes = productData[`profuctInfoShipmentTypes`]

	//fmt.Println(`sku:`, dataList["productBaseData"]["sku"])
	//fmt.Println(`productName:`, dataList["productBaseData"]["productName"])
	//产品
	//` + productData[`profuctInfoDeclarationData`][`declaredBrandName`][0] + `
	data = `{
   "pdProductBase": {
       "productId": "",
       "groupId": "` + dataList["productBaseData"]["groupID"] + `",
       "productName": "` + productInfo[`peProductName`][0] + `",
       "productLogo": "/2023/03/06/35b69430-5574-4759-8252-656fbd2a7fa0.jpg",
       "productLogoUrl": "http://192.168.10.101/statics//2023/03/06/35b69430-5574-4759-8252-656fbd2a7fa0.jpg",
       "sku": "` + dataList["productBaseData"]["sku"] + `",
       "model": "G230306",
       "unit": "G03306",
       "categoryId": null,
       "categoryName": null,
       "brandId": null,
       "brandName": "",
       "spuId": null,
       "spu": "",
       "spuName": null,
       "productDesc": "DSDFSD",
       "productDeveloperUserid": "710",
       "productDeveloper": "",
       "productCreatorUserid": "710",
       "productCreator": "",
       "status": 1,
       "isCombo": 0,
       "isDelete": 0,
       "createTime": "` + dataList["productBaseData"]["createTime"] + `",
       "updateTime": "` + dataList["productBaseData"]["createTime"] + `",
       "createBy": "710",
       "updateBy": "710",
       "remark": "",
       "attributeValue": null,
       "isSpu": null,
       "attributeValueList": null,
       "productAttributeValueList": null,
       "pdProductImportId": null,
       "statusStr": null,
       "supplierName": null,
       "purchaseUrl": null,
       "skus": null,
       "categoryTwo": null,
       "categoryThree": null,
       "productMaterial": null,
       "buyDay": null,
       "buyPrice": null,
       "productLength": null,
       "productWidth": null,
       "productHeight": null,
       "packageLength": null,
       "packageWidth": null,
       "packageHeight": null,
       "boxLength": null,
       "boxWidth": null,
       "boxHeight": null,
       "productNetWeight": null,
       "productGrossWeight": null,
       "boxWeight": null,
       "boxPcs": null,
       "buyCurrency": null,
       "isTax": null,
       "tax": null,
       "minCount": null,
       "purchaseTaxPrice": null,
       "peUserId": null,
       "hasDeclaration": null,
       "shipmentTypesName": null,
       "erpType": 0,
       "erpProductId": null,
       "sellerId": null,
       "sellerSku": null,
       "trueEndTime": null
   },
  "pdProductPurchase": {
       "productId": "85636",
       "groupId": null,
       "peUserId": 710,
       "peUserName": "Star",
       "peDelivery":25,
       "pePrice": ` + productInfo[`pePrice`][0] + `,
       "peProductLength": ` + productInfo[`peProductLength`][0] + `,
       "peProductLengthIn": null,
       "peProductWidth": ` + productInfo[`peProductWidth`][0] + `,
       "peProductWidthIn": null,
       "peProductHeight": ` + productInfo[`peProductHeight`][0] + `,
       "peProductHeightIn": null,
       "pePackageLength": ` + productInfo[`pePackageLength`][0] + `,
       "pePackageLengthIn": null,
       "pePackageWidth": ` + productInfo[`pePackageWidth`][0] + `,
       "pePackageWidthIn": null,
       "pePackageHeight": ` + productInfo[`pePackageHeight`][0] + `,
       "pePackageHeightIn": null,
       "peBoxLength": ` + productInfo[`peBoxLength`][0] + `,
       "peBoxLengthIn": null,
       "peBoxWidth": ` + productInfo[`peBoxWidth`][0] + `,
       "peBoxWidthIn": null,
       "peBoxHeight": ` + productInfo[`peBoxHeight`][0] + `,
       "peBoxHeightIn": null,
       "peProductNetWeight": ` + productInfo[`peProductNetWeight`][0] + `,
       "peProductNetWeightIn": null,
       "peProductGrossWeight": ` + productInfo[`peProductGrossWeight`][0] + `,
       "peProductGrossWeightIn": null,
       "peBoxWeight": ` + productInfo[`peBoxWeight`][0] + `,
       "peBoxWeightIn": null,
       "peBoxPcs": ` + productInfo[`peBoxPcs`][0] + `,
       "peProductMaterial": "` + productInfo[`peProductMaterial`][0] + `",
       "createTime": "` + dataList["productBaseData"]["createTime"] + `",
       "updateTime": "` + dataList["productBaseData"]["createTime"] + `",
       "createBy": "710",
       "updateBy": "710",
       "remark": "",
       "erpType": 0,
       "erpProductId": null
   },
   "pdDeclarationInfo": {
       "productId": "85636",
       "templateId": null,
       "declaredNameCn": "红葡萄",
       "declaredNameEn": "RedGrape",
       "customsCode": "` + dataList["productBaseData"]["customsCode"] + `",
       "declaredWeight": null,
       "declaredValueUsd": ` + utility.GetExRateString(productInfo[`pePrice`][0], 64, "USD") + `,
       "declaredValueEur": ` + utility.GetExRateString(productInfo[`pePrice`][0], 64, "EUR") + `,
       "declaredValueGbp": ` + utility.GetExRateString(productInfo[`pePrice`][0], 64, "GBP") + `,
       "productMaterial": "` + profuctInfoDeclarationData[`productMaterial`][0] + `",
       "productUsage": "` + profuctInfoDeclarationData[`productUsage`][0] + `",
       "productModel": "RG23",
       "declaredBrandName": "` + profuctInfoDeclarationData[`declaredBrandName`][0] + `",
       "productMaterialEn": "fruit",
       "productUsageEn": "edible",
       "productModelEn": "RG23",
       "declaredBrandNameEn": "Mr.Grape",
       "brandType": 1,
       "salesLink": "https://p4psearch.1688.com/hamlet.html?scene=2&keywords=%E8%91%A1%E8%90%84&cosite=baidujj&trackid=88564237933482014796346&format=normal&location=landing_t4&keywordid=479809693174&adposition=cr1&pagenum=1&creative=63874981741&matchtype=1",
       "declaredPicture": "/2023/02/25/0932870d-78f1-4ee6-b1e4-662d347c140c.jpg",
       "isDeclared": ` + profuctInfoDeclarationData[`isDeclared`][0] + `,
       "createTime": "` + dataList["productBaseData"]["createTime"] + `",
       "updateTime": "` + dataList["productBaseData"]["createTime"] + `",
       "createBy": "710",
       "updateBy": "710",
       "shipmentTypes": [` + profuctInfoShipmentTypes[`shipmentTypes`][0] + `],
       "sku": null,
       "declaredPictureUrl": "http://192.168.10.101/statics//2023/02/25/0932870d-78f1-4ee6-b1e4-662d347c140c.jpg",
       "amzListingsDeclarationInfoBoList": null,
       "pdDeclarationInfoSalesLinkList": []
   },
   "lstPurchaseSupplierProduct": []
}`
	return
}
