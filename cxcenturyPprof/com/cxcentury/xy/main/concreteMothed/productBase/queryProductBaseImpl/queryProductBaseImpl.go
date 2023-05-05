package queryProductBaseImpl

import (
	"cxcenturyPprof/com/cxcentury/xy/main/utility"
)

func (dv DataValue) QueryProductBaseImpl() *string {
	utilityDv := utility.DataValue{}
	utilityDv.Method = dv.Method
	utilityDv.Token = dv.Token

	utilityDv.Data = dv.Bodys
	utilityDv.Path = dv.Path
	bodys := string(utilityDv.MethodUtility())
	return &bodys
}
