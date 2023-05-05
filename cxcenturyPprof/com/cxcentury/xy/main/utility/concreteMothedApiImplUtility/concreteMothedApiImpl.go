package concreteMothedApiImplUtility

import "cxcenturyPprof/com/cxcentury/xy/main/utility"

var utilityDv = utility.DataValue{}

func (dv DataValue) MothedApiImpl() (data *string) {
	utilityDv.Method = dv.Method
	utilityDv.Token = dv.Token

	utilityDv.Data = dv.Bodys
	utilityDv.Path = dv.Path
	dv.Data = string(utilityDv.MethodUtility())
	data = &dv.Data

	return
}
