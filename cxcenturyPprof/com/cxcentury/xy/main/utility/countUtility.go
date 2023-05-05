package utility

/*
string 转int除n 返回string
*/
func CountAVGVarString(str string, num int64, length int) string {
	return StrcoveFormatintUtility(StrcoveParseInt(str)/num, length)
}

/*
*
体积重计算方法
*/
func CountVolumeWeight(length, width, height, count, volumeWeightPercent float64) float64 {
	var value float64

	value = length * width * height * count / volumeWeightPercent
	return value
}

/*
*
计费重计算方法
*/
func CountTotalWeight(length, width, height, count, volumeWeightPercent float64) float64 {
	var value float64
	value = length * width * height * count / volumeWeightPercent
	return value
}
