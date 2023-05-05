package loginData

type DataValue struct {
	PhoneNumber, Password string
}

func (dv DataValue) LoginDataStar() (string, string) {
	dv.PhoneNumber, dv.Password = `18565078921`, `YuMei000804`
	return dv.PhoneNumber, dv.Password
}
