package productImpl

import (
	"cxcenturySelenium/com/cxcentury/xy/main/data/productData/productArcFox"
	"cxcenturySelenium/com/cxcentury/xy/main/data/productData/uniqueProducttCharacteristic"
	"cxcenturySelenium/com/cxcentury/xy/main/utility"
	"github.com/tebeka/selenium"
	"strings"
	"time"
)

type DataValue struct {
	DataMapList                         []map[string][]string
	ArcFoxDataMap                       productArcFox.DataValue
	uniqueProducttCharacteristicDataMap uniqueProducttCharacteristic.DataValue
}

//type DataMap struct {
//	CSSSelectorPath string
//	Value           string
//}

func (dv DataValue) AddProductImpl(wd *selenium.WebDriver) {
	dv.DataMapList = make([]map[string][]string, 0)
	dv.ArcFoxDataMap.ArcFoxData = *dv.ArcFoxDataMap.ARCFOX()
	dv.uniqueProducttCharacteristicDataMap.UniqueProductMap = *dv.uniqueProducttCharacteristicDataMap.UniqueProductCharacteristic()

	dv.DataMapList = []map[string][]string{

		/*
			进入添加产品界面
		*/
		{`ProductIsSecondClickCSS`: {`.icon-chanpinguanli`}},
		{`ProductManageIsSecondClickCSS`: {`div.two-level-title:nth-child(1)`}},
		/*
			基本信息
		*/
		{`AddProductIsSecondClickCSS`: {`.add-btn-box > button:nth-child(3)`}},
		//=======================
		{`ProductNameInputCSS`: {
			`div.one-line:nth-child(1) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`peProductName`][5]}},
		{`ProductSKUIsSecondInputXPATH`: {
			`/html/body/div[3]/div/div[2]/div/div[2]/form/div[1]/div/div[2]/div[1]/div/div[1]/input`,
			dv.uniqueProducttCharacteristicDataMap.UniqueProductMap[`productBaseData`][`sku`]}},
		{`ProductModelInputCSS`: {
			`.base-form > div:nth-child(2) > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`model`][0]}},
		{`CompanyInputCSS`: {
			`.base-form > div:nth-child(2) > div:nth-child(3) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfoDeclarationData`][`declaredBrandName`][0]}},
		//===============选择框================================
		{`ProductStatusClickCSS`: {`.base-form > div:nth-child(2) > div:nth-child(4) > div:nth-child(2) > div:nth-child(1) > div:nth-child(1) > input:nth-child(1)`}},
		//{`ProductStatusIsSecondClickTagName`: {`<span>开发中</span>`}},
		//{`ProductTypeTagNameClickXPATH`: {`/html/body/div[4]/div/div[2]/div/div[2]/form/div[1]/div/div[2]/div[5]/div/div/div/input`}},
		//{`ProductStatusIsSecondClickTagName`: {`<span>分类</span>`}},
		//{`ProductPlansIsSecondClickXPATH`: {`/html/body/div[4]/div/div[2]/div/div[2]/form/div[1]/div/div[2]/div[6]/div/div/div/input`}},
		//{`ProductStatusIsSecondClickTagName`: {`<span>品牌</span>`}},
		//===============================================
		{`MenoInputCSS`: {`.el-textarea__inner`, `这是产品备注`}},
		/*
			采购信息
		*/
		{`ProductPurchaseClickCSS`: {`.tab-list > span:nth-child(2)`}},
		//{`PurchaseUserIsSecondClickCSS`: {`.is-focus > input:nth-child(1)`}},
		//==============选择框=================================================
		{`PurchaseUserIsSecondClickXPATH`: {`/html/body/div[3]/div/div[2]/div/div[2]/form/div[2]/div[1]/div[1]/div/div/div/input`}},
		//{`PurchaseUserIsSecondClickXPATH`: {`采购员`}},
		//===========================
		{`PurchaseDeliverInputCSS`: {
			`div.form-purchase:nth-child(2) > div:nth-child(2) > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			utility.StrcoveItoaUtilityRandIntUtility(80, 2)}},

		{`PurchasePriceInputCSS`: {
			`div.form-purchase:nth-child(2) > div:nth-child(2) > div:nth-child(3) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`pePrice`][5]}},
		//
		{`PurchaseProductLengthInputCSS`: {
			`div.input-specs:nth-child(4) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`peProductLength`][0]}},
		{`PurchaseProductWidthInputCSS`: {
			`div.input-specs:nth-child(4) > div:nth-child(2) > div:nth-child(2) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`peProductWidth`][0]}},
		{`PurchaseProductHeightInputCSS`: {
			`div.input-specs:nth-child(4) > div:nth-child(2) > div:nth-child(3) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`peProductHeight`][0]}},
		//
		{`PurchasePackageLengthInputCSS`: {
			`div.is-unit:nth-child(5) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`pePackageLength`][0]}},
		{`PurchasePackageWidthInputCSS`: {
			`div.is-unit:nth-child(5) > div:nth-child(2) > div:nth-child(2) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`pePackageWidth`][0]}},
		{`PurchasePackageHeightInputCSS`: {
			`div.is-unit:nth-child(5) > div:nth-child(2) > div:nth-child(3) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`pePackageHeight`][0]}},
		//
		{`PurchaseBoxLengthInputCSS`: {
			`div.is-unit:nth-child(6) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`peBoxLength`][0]}},
		{`PurchaseBoxWidthInputCSS`: {
			`div.is-unit:nth-child(6) > div:nth-child(2) > div:nth-child(2) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`peBoxWidth`][0]}},
		{`PurchaseBoxHeightInputCSS`: {
			`div.is-unit:nth-child(6) > div:nth-child(2) > div:nth-child(3) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`peBoxHeight`][0]}},
		//
		{`PurchaseProductNetWeightInputCSS`: {
			`div.is-unit:nth-child(1) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`peProductNetWeight`][0]}},
		{`PurchaseProductGrossWeightInputCSS`: {
			`div.form-tree-column:nth-child(7) > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`peProductGrossWeight`][0]}},
		{`PurchaseBoxWeightInputCSS`: {
			`div.form-tree-column:nth-child(7) > div:nth-child(3) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`peBoxWeight`][0]}},
		//
		{`PurchaseBoxPcInputCSS`: {
			`div.form-tree-column:nth-child(7) > div:nth-child(4) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`peBoxPcs`][0]}},
		{`PurchaseProductMaterialInputCSS`: {
			`div.form-tree-column:nth-child(7) > div:nth-child(5) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`peProductMaterial`][0]}},
		/*
			申报信息
		*/
		{`DeclarationInfoClickCSS`: {`.tab-list > span:nth-child(3)`}},
		//============选择框==================
		//{`DeclarationInfoIsSecondClickXPATH`: {`/html/body/div[8]/div/div[2]/div/div[2]/form/div[4]/div/div[1]/form/div/div[1]/div[1]/div/div/div/input`}},
		//==============================
		{`ChinesDeclarationNameIsSecondInputCSS`: {
			`div.form-purchase:nth-child(1) > div:nth-child(3) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`chinesDeclarationName`][0]}},
		{`EnglishDeclarationNameIsSecondInputCSS`: {
			`div.form-purchase:nth-child(1) > div:nth-child(4) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`englishDeclarationName`][0]}},
		//特殊属性多选
		{`DeclarationClassClickCSS`: {`label.el-checkbox:nth-child(2) > span:nth-child(2)`}},
		{`DeclarationClassClickCSS`: {`label.el-checkbox:nth-child(3) > span:nth-child(2)`}},
		{`DeclarationClassClickCSS`: {`label.el-checkbox:nth-child(4) > span:nth-child(2)`}},
		{`DeclarationClassClickCSS`: {`label.el-checkbox:nth-child(11) > span:nth-child(2)`}},
		{`DeclarationClassClickCSS`: {`label.el-checkbox:nth-child(12) > span:nth-child(2)`}},
		{`DeclarationClassClickCSS`: {`label.el-checkbox:nth-child(13) > span:nth-child(2)`}},
		//
		{`DeclarationPriceUSDInputCSS`: {
			`div.form-tree-column:nth-child(6) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			utility.GetExRateString(dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`pePrice`][0], 64, "USD")}},
		{`DeclarationPriceEURInputCSS`: {
			`div.form-tree-column:nth-child(6) > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			utility.GetExRateString(dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`pePrice`][0], 64, "EUR")}},
		{`DeclarationPriceGBPInputCSS`: {
			`div.form-tree-column:nth-child(6) > div:nth-child(3) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			utility.GetExRateString(dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`pePrice`][0], 64, "GBP")}},
		//
		{`DeclarationProductMaterialInputCSS`: {
			`div.form-purchase:nth-child(1) > div:nth-child(7) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfoDeclarationData`][`productMaterial`][0]}},
		{`DeclarationProductUsageInputCSS`: {
			`div.form-purchase:nth-child(1) > div:nth-child(7) > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfoDeclarationData`][`productUsage`][0]}},
		{`DeclarationProductModelInputCSS`: {
			`div.form-purchase:nth-child(1) > div:nth-child(7) > div:nth-child(3) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`model`][0]}},
		//
		{`DeclarationProductMateRiaEnglishInputCSS`: {
			`div.form-tree-column:nth-child(8) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfoDeclarationData`][`productMaterial`][0]}},
		{`DeclarationProductUsageEnglishInputCSS`: {
			`div.form-tree-column:nth-child(8) > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfoDeclarationData`][`productUsage`][0]}},
		{`DeclarationProductModelEnglishInputCSS`: {
			`div.form-tree-column:nth-child(8) > div:nth-child(3) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfo`][`model`][0]}},
		//
		{`DeclarationProductBrandNameInputCSS`: {
			`div.form-tree-column:nth-child(9) > div:nth-child(1) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfoDeclarationData`][`declaredBrandName`][0]}},
		{`DeclarationProductBrandNameEnglishInputCSS`: {
			`div.form-tree-column:nth-child(9) > div:nth-child(2) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfoDeclarationData`][`declaredBrandName`][0]}},
		//品牌类型下拉框
		//{`DeclarationBrandTypeClickXPATH`: {
		//	`/html/body/div[9]/div/div[2]/div/div[2]/form/div[4]/div/div[1]/form/div/div[8]/div[3]/div/div/div/input`}},
		//
		{`DeclarationCusTomsCodeInputCSS`: {
			`div.form-tree-column:nth-child(9) > div:nth-child(4) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			utility.StrcoveItoaUtilityRandIntUtility(9999999, 8)}},
		//是否优惠关税
		{`DeclarationIsFavouredRateOfDutyClickCSS`: {
			`div.form-tree-column:nth-child(9) > div:nth-child(5) > div:nth-child(2) > div:nth-child(1) > div:nth-child(1) > input:nth-child(1)`}},
		{`DeclarationTariffRateInputCSS`: {
			`div.form-tree-column:nth-child(9) > div:nth-child(6) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			utility.StrcoveItoaUtilityRandIntUtility(5, 2)}},
		//链接
		{`DeclarationSalesLinkInputCSS`: {
			`div.el-form-item:nth-child(7) > div:nth-child(2) > div:nth-child(1) > input:nth-child(1)`,
			dv.ArcFoxDataMap.ArcFoxData[`productInfoShipmentTypes`][`ProductUrl`][0]}},
		//上传产品基本图片
		//{`UpLoadPictureIsSecondInputCSS`: {
		//	`.el-upload > i:nth-child(1)`,
		//	dv.ArcFoxDataMap.ArcFoxData[`productInfoShipmentTypes`][`declaredPicture`][0],
		//}},
		//申报图片
		//申报图片按钮
		/*{`UpLoadPictureClickCSS`: {
		`.el-upload > i:nth-child(1)`}},*/
		//确定按钮
		{`DeclarationSendAddProductInfoClickCSS`: {
			`div.el-dialog__footer:nth-child(3) > span:nth-child(1) > button:nth-child(2) > span:nth-child(1)`}},
	}

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
