package logisticsPlansDataImpl

import (
	"cxcenturyPprof/com/cxcentury/xy/main/utility"
	"math"
	"strings"
)

func AddLogisticsPlansDataImpl(logisticsChannelForGroupBodyList map[string][]string, uniqueProductCList map[string]map[string]string, productDataList map[string]map[string][]string) string {

	//控制参数坐标
	logisticsChannelForGroupBodyListNum := 0

	//遍历shipmentTypes列
	shipmentTypes := strings.Split(productDataList[`profuctInfoShipmentTypes`][`shipmentTypes`][0], `,`)
	var logisticsPlanShipmentTypeList string
	for key, value := range shipmentTypes {
		str := `{
            "logisticsPlanId": 0,
            "shipmentType": ` + value + `
        }`
		if key < len(shipmentTypes)-1 {
			str = str + `,`
		}
		logisticsPlanShipmentTypeList = logisticsPlanShipmentTypeList + str
	}
	//预计发货天数，当前天数开始计算
	expectedSendDate := int64(4)
	//预计签收天数，当前时间开始计算
	expectedReceiptDate := int64(24)
	//随机值返回string
	packageCount := utility.StrcoveItoaUtilityRandIntUtility(200)
	packageIndex := utility.StrcoveItoaUtilityRandIntUtility(180)
	totalPcs := utility.StrcoveItoaUtilityRandIntUtility(170)

	cnyPerCbm := utility.StrconvFormatFloat(utility.RandFloat64tility(40), 2, 64)
	cnyPerCbmF := utility.StrconvParseFloat(cnyPerCbm, 0)

	//string 转float64
	pePackageHeight := utility.StrconvParseFloat(productDataList[`productInfo`][`pePackageHeight`][0], 2)
	pePackageLength := utility.StrconvParseFloat(productDataList[`productInfo`][`pePackageLength`][0], 2)
	pePackageWidth := utility.StrconvParseFloat(productDataList[`productInfo`][`pePackageWidth`][0], 2)
	peProductGrossWeight := utility.StrconvParseFloat(utility.CountAVGVarString(productDataList[`productInfo`][`peProductGrossWeight`][0], 1000, 10), 2)
	packageCountF := utility.StrconvParseFloat(packageCount, 2)

	//取随机值
	cnyPerKg := utility.StrconvFormatFloat(utility.RandFloat64tility(45), 4, 64)
	cnyPerKgF := utility.StrconvParseFloat(cnyPerKg, 0)
	//实重
	totalWeight := utility.StrconvParseFloat(packageCount, 2) * peProductGrossWeight

	//体积重
	volumeWeight := utility.StrconvParseFloat(utility.StrconvFormatFloat(utility.CountVolumeWeight(pePackageLength, pePackageWidth, pePackageHeight, packageCountF, 6000), 4, 64), 4)
	//fmt.Println(`volumeWeight`, volumeWeight, packageCount, peProductGrossWeight)

	var weightTotalFee float64
	if totalWeight >= volumeWeight {
		weightTotalFee = math.Ceil(totalWeight) * cnyPerKgF
	} else {
		weightTotalFee = math.Ceil(volumeWeight) * cnyPerKgF
	}
	//fmt.Println(`cnyPerKg:`, cnyPerKg, math.Ceil(totalWeight)*cnyPerKgF, math.Ceil(volumeWeight)*cnyPerKgF)
	//fmt.Println(`totalWeight:`, math.Ceil(totalWeight), `volumeWeight:`, math.Ceil(volumeWeight), `weightTotalFee1:`, weightTotalFee)
	//fmt.Println(`cnyPerKg:`, cnyPerKgF)
	//体积
	volume := pePackageHeight * pePackageLength * pePackageWidth * packageCountF / (100 * 100 * 100)
	//关税费
	taxFee := utility.RandFloat64tility(40)
	taxFeeS := utility.StrconvFormatFloat(taxFee, 4, 64)

	taxFeeCurrency := `USD`
	var USDtaxFee float64
	if taxFeeCurrency == `USD` {
		USDtaxFee = taxFee * 6.9362
	}
	//体积总费用
	volumeTotalFee := volume*cnyPerCbmF + USDtaxFee
	//fmt.Println(`volumeTotalFee:`, volumeTotalFee, `volume*cnyPerCbmF:`, volume*cnyPerCbmF, `cnyPerCbmF:`, cnyPerCbmF, `volume:`, volume, `USDtaxFee:`, USDtaxFee)

	data := `{
    "lineConsolidationOrder": 0,
    "logisticsPlan": {
        "shipmentName": "` + uniqueProductCList[`productBaseData`][`productName`] + `",
        "referenceId": "",
        "sendProv": ` + logisticsChannelForGroupBodyList[`sendProv`][logisticsChannelForGroupBodyListNum] + `,
        "sendCity": ` + logisticsChannelForGroupBodyList[`sendCity`][logisticsChannelForGroupBodyListNum] + `,
        "expectedSendDate": "` + utility.RuneDateTimeUtility(expectedSendDate) + `",
        "expectedReceiptDate": "` + utility.RuneDateTimeUtility(expectedReceiptDate) + `",
        "channelName": "` + logisticsChannelForGroupBodyList[`channelName`][logisticsChannelForGroupBodyListNum] + `",
        "companyName": "` + logisticsChannelForGroupBodyList[`companyName`][logisticsChannelForGroupBodyListNum] + `",
        "destinationContinent": "",
        "destinationCountry": ` + logisticsChannelForGroupBodyList[`destinationCountry`][logisticsChannelForGroupBodyListNum] + `,
        "fbaPlace": "",
        "isFba": 0,
        "expectedDays": 1,
        "transportType": 6,
        "logisticsAccount": "",
        "weixinNotice": 1,
        "smsNotice": 0,
        "totalPcs": ` + totalPcs + `,
        "totalPackageCount": 20,
        "cnyPerKg": "` + cnyPerKg + `",
        "totalWeight": "` + utility.StrconvFormatFloat(totalWeight, 4, 64) + `",
        "cnyPerCbm": "` + cnyPerCbm + `",
        "totalVolume": "0.1600",
        "customsFees": 0,
        "additionalFee": 0,
        "otherFee": 0,
        "feeIsTax": 1,
        "finalDeliveryFee": 0,
        "taxFee": "` + taxFeeS + `",
        "weightTotalFee": ` + utility.StrconvFormatFloat(weightTotalFee, 4, 64) + `,
        "volumeTotalFee": ` + utility.StrconvFormatFloat(volumeTotalFee, 4, 64) + `,
        "volumeWeight": "` + utility.StrconvFormatFloat(volumeWeight, 10, 64) + `",
        "salesman": "",
        "salesmanId": 0,
        "warehouseManager": "",
        "warehouseManagerId": 0,
        "shopAccount": "",
        "sellerId": null,
        "volumeWeightPercent": 6000,
        "receiptStatus": "",
        "volumeEstimatedFee": "3.20",
        "weightEstimatedFee": "8000.00",
        "weightSharedExpenses": 6.94,
        "volumeSharedExpenses": 6.94,
        "logisticsCompany": "",
        "paymentDays": 20,
        "logisticsCompanyMemo": "2020",
        "logisticsMemo": "",
        "remark": "",
        "companyType": 1,
        "logisticsCompanyId": "` + logisticsChannelForGroupBodyList[`logisticsCompanyId`][logisticsChannelForGroupBodyListNum] + `",
        "logisticsCompanyIdInfo": "",
        "logisticsChannelId": "` + logisticsChannelForGroupBodyList[`logisticsChannelId`][logisticsChannelForGroupBodyListNum] + `",
        "logisticsChannelAliasId": ` + logisticsChannelForGroupBodyList[`logisticsChannelAliasId`][logisticsChannelForGroupBodyListNum] + `,
        "logisticsChannelIdInfo": "",
        "overseaCompanyId": "",
        "overseaChannelId": "",
        "overseaCompanyIdInfo": "",
        "overseaChannelIdInfo": "",
        "redisKeyRandom": "4nrdv3c4",
        "taxFeeCurrency": "` + taxFeeCurrency + `",
        "exchangeRate": 6.9362,
        "taxFeeRmb": "138.72",
        "feeMemo": "` + logisticsChannelForGroupBodyList[`feeMemo`][logisticsChannelForGroupBodyListNum] + `",
        "totalAdditionalFee": 0,
        "isOrder": 0,
        "hasDestinationStr": true,
        "originTotalVolume": 160000,
        "forwardingChannelId": "",
        "additionalWeightFee": 0,
        "totalAdditionalWeightFee": 0,
        "additionalVolumeFee": 0,
        "totalAdditionalVolumeFee": 0,
        "updateReason": "",
        "destinationWarehouseId": "",
        "packageWeight": 0,
        "shipmentId": ""
    },
    "erpInbondShipmentId": [],
    "logisticsPlanFileList": [],
    "logisticsPlanShipmentBoList": [
        {
            "logisticsPlanId": null,
            "logisticsPlanPackageList": [
                {
                    "logisticsPlanId": null,
                    "logisticsPlanPackageId": null,
                    "logisticsPlanPackageItemList": [
                        {
                            "logisticsPlanPackageId": null,
                            "logisticsPlanPackageItemId": null,
                            "msku": "",
                            "pcs": "",
                            "productId": "",
                            "productName": "",
                            "hasDeclaration": ""
                        }
                    ],

                    "logisticsPlanPackageName": null,
                    "logisticsPlanShipmentId": null,
                    "packageCount": "` + packageCount + `",
                    "packageHeight": "` + productDataList[`productInfo`][`pePackageHeight`][0] + `",
                    "packageLength": "` + productDataList[`productInfo`][`pePackageLength`][0] + `",
                    "packageWeight": "` + utility.CountAVGVarString(productDataList[`productInfo`][`peProductGrossWeight`][0], 1000, 10) + `",
                    "packageWidth": "` + productDataList[`productInfo`][`pePackageWidth`][0] + `",
                    "packageIndex": ` + packageIndex + `,
                    "shipmentId": "",
                    "shipmentName": null,
                    "shipmentsId": 0, 
                    "sellerId": null,
                    "packageListLength": 1,
                    "_X_ROW_KEY": "row_77"
                }
            ],
            "logisticsPlanShipmentId": null,
            "shipmentId": "",
            "shipmentName": null,
            "shipmentsId": 0,
            "sellerId": null
        }
    ],
    "logisticsPlanPackageList": [
        {
            "packageLength": "",
            "packageWidth": "",
            "packageHeight": "",
            "packageCount": "",
            "packageWeight": ""
        }
    ],
    "logisticsPlanShipmentTypeList": [
        ` + logisticsPlanShipmentTypeList + `
    ],
    "consolidationOrders": {
        "addressInfo": "",
        "bookingStatus": 0,
        "chargedVolume": 0,
        "chargedWeight": 0,
        "customsClearanceMethod": null,
        "customsDeclarationFee": 0,
        "customsDeclarationMethod": null,
        "deliveryMethod": 1,
        "eoriNumber": "",
        "fbaCode": "",
        "forwardingChannelId": 0,
        "internalRemark": "",
        "linker": "",
        "linkerPhone": "",
        "ordersNumber": "",
        "otherFee": 0,
        "ownOrderNumber": "",
        "paymentStatus": 0,
        "pickupArea": "",
        "pickupCity": "",
        "pickupDetailAddress": "",
        "destinationCountry": "",
        "pickupProv": "",
        "pickupDate": "",
        "pickupTimeEnd": "",
        "pickupTimeStart": "",
        "deliveryToWarehouseDate": "",
        "deliveryToWarehouseTimeEnd": "",
        "deliveryToWarehouseTimeStart": "",
        "deliveryToWarehouseLinker": "",
        "deliveryToWarehouseLinkerPhone": "",
        "postCode": "",
        "price": 0,
        "remark": "",
        "status": 0,
        "surcharge": 0,
        "totalAmount": 0,
        "totalPackageCount": 0,
        "totalVolume": "",
        "totalWeight": "",
        "tradeType": 1,
        "vatNumber": "",
        "volumeWeight": 0,
        "weightVolume": "",
        "logisticsFee": 0
    },
    "logisticsPlanOtherFees": [],
    "inquiryLogisticsChannelForgroupIdList": []
}`
	return data
}
