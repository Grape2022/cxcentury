package productArcFox

import (
	"fmt"
	"time"
)

/*
ArcFox产品参数
*/
type DataValue struct {
	ArcFoxData map[string]map[string][]string
}

func (dv DataValue) ARCFOX() *map[string]map[string][]string {
	dv.ArcFoxData = map[string]map[string][]string{}
	dv.ArcFoxData = map[string]map[string][]string{
		"productInfo": {
			`peProductName`:          {`阿尔法S 525S`, `阿尔法S 525S+`, `阿尔法S 603H`, `阿尔法S 708S+`, `阿尔法S 708S`, `阿尔法S 708S+ 小王子联名款`, `阿尔法S 603H 小王子联名款 `},
			`pePrice`:                {`223800`, `253800`, `350300`, `287300`, `267300`, `293300`, `356300`},
			`peProductLength`:        {`493`},
			`peProductWidth`:         {`194`},
			`peProductHeight`:        {`159.9`},
			`peProductNetWeight`:     {`2245000`},
			`peProductGrossWeight`:   {`2245000`},
			`peProductMaterial`:      {`锂电池`},
			`0-100km/h加速时间`:      {`7.7s`, `8.3s`, `4.2s`},
			`最大功率`:               {`160kW`, `175kW`, `320kW`},
			`NEDC综合工况续航里程`:   {`525km`, `708km`, `603km`},
			`外观颜色`:               {`薄荷银`, `冰河蓝`, `冰灰`, `极狐白`, `极夜黑`, `星钻黑`},
			`内饰颜色`:               {`矅石黑`, `深海蓝`},
			`充电时间`:               {`0.5h`, `0.6h`},
			`轮胎尺寸`:               {`235/50 R19`, `245/45 R20`},
			`pePackageLength`:        {`493`},
			`pePackageWidth`:         {`194`},
			`pePackageHeight`:        {`159.9`},
			`peBoxLength`:            {`493`},
			`peBoxWidth`:             {`194`},
			`peBoxHeight`:            {`159.9`},
			`peBoxWeight`:            {`2245000`},
			`peBoxPcs`:               {`1000`},
			`model`:                  {`ArcFox` + fmt.Sprintf("%d", time.Now().UnixNano())},
			`chinesDeclarationName`:  {`新能源汽车`},
			`englishDeclarationName`: {`New energy vehicles`},
		},
		`productInfoDeclarationData`: {
			`declaredNameEn`:    {``},
			`customsCode`:       {`212425`},
			`productMaterial`:   {`汽车`},
			`productUsage`:      {`代步`},
			`declaredBrandName`: {`ArcFox`},
			"isDeclared":        {`1`},
		},
		`productInfoShipmentTypes`: {
			`shipmentTypes`:   {`2,3,4,11,12,13`},
			`declaredPicture`: {`https://www.arcfox.com.cn/static/web/img2022/carImage/s/prince.png`},
			`ProductUrl`:      {`https://www.arcfox.com.cn/forestS/forestS.html`},
		},
	}

	return &dv.ArcFoxData
}
