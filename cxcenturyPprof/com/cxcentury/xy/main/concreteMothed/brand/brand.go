package brand

import (
	"cxcenturyPprof/com/cxcentury/xy/main/utility/concreteMothedApiImplUtility"
)

func (dv DataValue) QueryBrandList() *string {
	MothedApiUtilityDv := concreteMothedApiImplUtility.DataValue{}
	MothedApiUtilityDv.Token = dv.Token
	MothedApiUtilityDv.Path = `/brand/list`
	MothedApiUtilityDv.Method = `GET`
	return MothedApiUtilityDv.MothedApiImpl()
}
