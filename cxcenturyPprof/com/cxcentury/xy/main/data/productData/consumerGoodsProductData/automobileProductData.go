package consumerGoodsProductData

/*
ArcFox产品参数
*/
func ARCFOX() (ArcFoxData map[string]map[string][]string) {
	ArcFoxData = map[string]map[string][]string{
		"productInfo": {
			`peProductName`:        []string{`阿尔法S 525S`, `阿尔法S 525S+`, `阿尔法S 603H`, `阿尔法S 708S+`, `阿尔法S 708S`, `阿尔法S 708S+ 小王子联名款`, `阿尔法S 603H 小王子联名款 `},
			`pePrice`:              []string{`223800`, `253800`, `350300`, `287300`, `267300`, `293300`, `356300`},
			`peProductLength`:      []string{`493`},
			`peProductWidth`:       []string{`194`},
			`peProductHeight`:      []string{`159.9`},
			`peProductNetWeight`:   []string{`2245000`},
			`peProductGrossWeight`: []string{`2245000`},
			`peProductMaterial`:    []string{`锂电池`},
			`0-100km/h加速时间`:        []string{`7.7s`, `8.3s`, `4.2s`},
			`最大功率`:                 []string{`160kW`, `175kW`, `320kW`},
			`NEDC综合工况续航里程`:         []string{`525km`, `708km`, `603km`},
			`外观颜色`:                 []string{`薄荷银`, `冰河蓝`, `冰灰`, `极狐白`, `极夜黑`, `星钻黑`},
			`内饰颜色`:                 []string{`矅石黑`, `深海蓝`},
			`充电时间`:                 []string{`0.5h`, `0.6h`},
			`轮胎尺寸`:                 []string{`235/50 R19`, `245/45 R20`},
			`pePackageLength`:      []string{`493`},
			`pePackageWidth`:       []string{`194`},
			`pePackageHeight`:      []string{`159.9`},
			`peBoxLength`:          []string{`493`},
			`peBoxWidth`:           []string{`194`},
			`peBoxHeight`:          []string{`493`},
			`peBoxWeight`:          []string{`2245000`},
			`peBoxPcs`:             []string{`1000`},
		},
		`profuctInfoDeclarationData`: {
			`declaredNameEn`:    []string{``},
			`customsCode`:       []string{`212425`},
			`productMaterial`:   []string{`汽车`},
			`productUsage`:      []string{`代步`},
			`declaredBrandName`: []string{`ArcFox`},
			"isDeclared":        []string{`1`},
		},
		`profuctInfoShipmentTypes`: {
			`shipmentTypes`:   []string{`2,3,4,11,12,13`},
			`declaredPicture`: []string{`https://www.arcfox.com.cn/static/web/img2022/carImage/s/prince.png`},
		},
	}

	return
}
