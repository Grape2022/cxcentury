package logisticsPlansImpl

import (
	"cxcenturySelenium/com/cxcentury/xy/main/data/productData/productArcFox"
	"cxcenturySelenium/com/cxcentury/xy/main/data/productData/uniqueProducttCharacteristic"
	"cxcenturySelenium/com/cxcentury/xy/main/utility"
	"fmt"
	"github.com/tebeka/selenium"
	"math"
	"strings"
	"time"
)

type DataValue struct {
	DataMapList                         []map[string][]string
	ArcFoxDataMap                       productArcFox.DataValue
	uniqueProducttCharacteristicDataMap uniqueProducttCharacteristic.DataValue
}

func (dv DataValue) AddLogisticsPlansImpl(wd *selenium.WebDriver) {
	dv.DataMapList = make([]map[string][]string, 0)
	dv.ArcFoxDataMap.ArcFoxData = *dv.ArcFoxDataMap.ARCFOX()
	dv.uniqueProducttCharacteristicDataMap.UniqueProductMap = *dv.uniqueProducttCharacteristicDataMap.UniqueProductCharacteristic()

	//预计发货天数，当前天数开始计算
	expectedSendDate := int64(4)
	//预计签收天数，当前时间开始计算
	expectedReceiptDate := int64(24)
	//随机值返回string
	packageCount := utility.StrcoveItoaUtilityRandIntUtility(200, 5)
	packageIndex := utility.StrcoveItoaUtilityRandIntUtility(180, 5)
	totalPcs := utility.StrcoveItoaUtilityRandIntUtility(170, 5)

	cnyPerCbm := utility.StrconvFormatFloat(utility.RandFloat64tility(40), 2, 64)
	cnyPerCbmF := utility.StrconvParseFloat(cnyPerCbm, 0)

	//string 转float64
	pePackageHeight := utility.StrconvParseFloat(dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`pePackageHeight`][0], 2)
	pePackageLength := utility.StrconvParseFloat(dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`pePackageLength`][0], 2)
	pePackageWidth := utility.StrconvParseFloat(dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`pePackageWidth`][0], 2)
	peProductGrossWeight := utility.StrconvParseFloat(dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`peProductGrossWeight`][0], 2)
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

	//dv.MapData[`cnyPerCbmF`] = fmt.Sprintf("%f",utility.StrconvParseFloat(dv.MapData[`cnyPerCbm`], 0))
	dv.DataMapList = []map[string][]string{
		/*
			进入添加计划单界面
		*/
		{`LogisticsIsSecondClickCSS`: {`div.module-item:nth-child(6) > div:nth-child(1) > span:nth-child(1)`}},
		{`LogisticsPlansIsSecondClickCSS`: {`div.two-level-title:nth-child(1)`}},
		{`AddLogisticsPlansIsSecondClickCSS`: {`button.search-left-btn:nth-child(1)`}},
		/*
			基本信息
		*/
		{`AddLogisticsPlansIsSecondInputCSS`: {`div.item-hasTips:nth-child(1) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.uniqueProducttCharacteristicDataMap.UniqueProductMap[`productBaseData`][`productName`]}},
		{`AddLogisticsPlansClickCSS`: {`button.search-left-btn:nth-child(1)`,
			fmt.Sprintf("ArcFox%d", time.Now().UnixNano())}},
		//下拉框：发货地
		//=================
		//{`AddLogisticsPlansIsSecondClickCSS`: {`button.search-left-btn:nth-child(1)`}},
		//{`AddLogisticsPlansIsSecondClickCSS`: {`button.search-left-btn:nth-child(1)`}},
		//==================
		//预计发货日期
		{`ExpectedSendDateInputCSS`: {`div.is-required:nth-child(4) > div:nth-child(2) > div:nth-child(1) > div:nth-child(1) > div:nth-child(1) > div:nth-child(1) > input:nth-child(2)`,
			fmt.Sprintf("%d", utility.RuneDateUtilityToPath(expectedSendDate, 0, 10))}},
		//预计天数
		{`ExpectedDaysInputCSS`: {`div.item-hasTips:nth-child(5) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			fmt.Sprintf("%d", expectedReceiptDate)}},
		//========================
		//货件ID
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		//下拉框：运输方式
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		//下拉框：FBA
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//			`20`}},
		//下拉框：选择FBA
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//			`20`}},
		//预计签收日期
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//			`20`}},
		//退税类型
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//			`20`}},
		//======================
		/*
			装箱信息
		*/
		////快速录入信息
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		//长
		{`pePackageLengthInputCSS`: {`.package-size-cont > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			fmt.Sprintf("%f", pePackageLength)}},
		//宽
		{`pePackageWidthInputCSS`: {`.package-size-cont > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)	`,
			fmt.Sprintf("%f", pePackageWidth)}},
		//高
		{`pePackageHeightInputCSS`: {`.package-size-cont > div:nth-child(3) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			fmt.Sprintf("%f", pePackageHeight)}},
		//重量
		{`peProductGrossWeightInputCSS`: {`.package-size-cont > div:nth-child(4) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			fmt.Sprintf("%f", peProductGrossWeight)}},
		//箱数
		{`packageIndexInputCSS`: {`.bale-div-arr > div:nth-child(1) > div:nth-child(1) > span:nth-child(3) > div:nth-child(2) > input:nth-child(1)`,
			packageIndex}},
		//下拉框：箱唛
		//==============================
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	totalPcs}},
		//==============================
		//总数量
		{`totalPcsInputCSS`: {`div.totalPcs-box:nth-child(1) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			totalPcs}},
		//总重量
		{`totalWeightInputCSS`: {`div.totalPcs-box:nth-child(2) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			fmt.Sprintf("%f", totalWeight)}},
		//体积重
		{`volumeWeightInputCSS`: {`div.totalPcs-box:nth-child(3) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			fmt.Sprintf("%f", volumeWeight)}},
		/*
			费用信息
		*/
		//特殊属性
		{`ProductClassClickCSS`: {`label.el-checkbox:nth-child(2)`}},
		{`ProductClassClickCSS`: {`label.el-checkbox:nth-child(3)`}},
		{`ProductClassClickCSS`: {`label.el-checkbox:nth-child(4)`}},
		{`ProductClassClickCSS`: {`label.el-checkbox:nth-child(5)`}},
		{`ProductClassClickCSS`: {`label.el-checkbox:nth-child(6)`}},
		{`ProductClassClickCSS`: {`label.el-checkbox:nth-child(7)`}},
		//重量单价
		{`cnyPerKgInputCSS`: {`.weightVolume-unit-price-cont > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			cnyPerKg}},
		//体积单价
		{`cnyPerCbmInputCSS`: {`.weightVolume-unit-price-cont > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			cnyPerCbm}},
		//关税费用
		{`taxFeeSInputCSS`: {`.tax-fee-box > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			taxFeeS}},
		//费用备注
		{`feeMemoInputCSS`: {`.fixed-fee-box-right > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > textarea:nth-child(1)`,
			fmt.Sprintf("外观颜色：%s,\n内饰颜色：%s,\n",
				dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`外观颜色`][utility.RandIntUtility(6)],
				dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`内饰颜色`][utility.RandIntUtility(2)],
			)}},
		//重量总费用
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		//体积总费用
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			fmt.Sprintf("%f", weightTotalFee)}},
		/*
			物流信息
		*/
		//下拉框:渠道名称
		//=====================
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		//下拉框:货代公司
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		//====================
		//账期
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		//选择渠道原因
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		/*
			渠道推荐快照
		*/
		//下拉框
		//===================
		//运输方式
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		//税况
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		//搜索货代
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		//搜索渠道
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		//搜索内置渠道
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		//===========================
		//开始日期
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		//结束日期
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		//时效范围1
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		//时效范围2
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		//采用渠道
		{`AddLogisticsPlansClickCSS`: {`button.search-left-btn:nth-child(1)`}},
		/*
			其他信息
		*/
		//供应商名称
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		//提货地址
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		//提货电话
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		//退税人邮箱
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		//供应商备注
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		//下拉框
		//==========================
		//业务员
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		//仓库负责人
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		//==========================
		//操作人
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		//备注
		{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
			`20`}},
		////报价凭证上传
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		////协商凭证上传
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
		////核对凭证上传
		//{`AddLogisticsPlansInputCSS`: {`button.search-left-btn:nth-child(1)`,
		//	`20`}},
	}

	fmt.Println("开始添加计划单吧")
	for i := 0; i < len(dv.DataMapList); i++ {
		for k := range dv.DataMapList[i] {
			if strings.Contains(k, `IsSecond`) {
				time.Sleep(time.Second)
			}
			switch true {
			case strings.Contains(k, `Click`) && strings.Contains(k, `CSS`):
				utility.ElementClickCSS(wd, &dv.DataMapList[i][k][0])
			case strings.Contains(k, `Click`) && strings.Contains(k, `XPATH`):
				utility.ElementClickXPATH(wd, &dv.DataMapList[i][k][0])
			case strings.Contains(k, `Click`) && strings.Contains(k, `TagName`):
				utility.ElementClickByTagName(wd, &dv.DataMapList[i][k][0])
			case strings.Contains(k, `Input`) && strings.Contains(k, `CSS`):
				utility.ElementSendKeysCSS(wd, &dv.DataMapList[i][k][0], &dv.DataMapList[i][k][1])
			case strings.Contains(k, `Input`) && strings.Contains(k, `XPATH`):
				utility.ElementSendKeysXPATH(wd, &dv.DataMapList[i][k][0], &dv.DataMapList[i][k][1])
			case strings.Contains(k, `Input`) && strings.Contains(k, `TagName`):
				utility.ElementSendKeysByTagName(wd, &dv.DataMapList[i][k][0], &dv.DataMapList[i][k][1])
			}
		}
	}
}
