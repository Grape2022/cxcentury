package testingErrorTestData

import (
	"fmt"
	"testing"
)

type args struct {
	Token string
	Data  string
}
type tests []struct {
	Name string

	Want map[string][]string
}

func TestIsMessageTypeMapString(t *testing.T, loginData map[string]string, testTextName string) tests {
	//token = login.Login()["token"]

	tests := tests{
		{Name: `1`,
			Want: map[string][]string{
				`msg`:  []string{`操作成功`, `查询成功`},
				`code`: []string{`200`},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			if code := loginData[`code`]; code != `200` {
				t.Errorf("%s = %v, Maps[`code`] %v", testTextName, code, tt.Want[`code`][0])
			}
			fmt.Println(loginData[`msg`] != `查询成功`)
			if msg := loginData[`msg`]; (msg != `操作成功`) && (msg != `查询成功`) {
				t.Errorf("%s = %v, Maps[`msg`] %v", testTextName, msg, tt.Want[`msg`][0], tt.Want[`msg`][0])
			}
		})
	}
	return tests
}
func TestIsMessageTypeMapArgsString(t *testing.T, data map[string][]string, TestTextName string) tests {
	//token = login.Login()["token"]

	tests := tests{
		{Name: `1`, Want: map[string][]string{
			`msg`:  []string{`操作成功`, `查询成功`},
			`code`: []string{`200`},
		}},
	}
	fmt.Println(data[`code`], data[`msg`])
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if code := data[`code`]; code[0] != `200` {
				t.Errorf("%s = %v, Maps[`code`] %v", TestTextName, code, tt.Want[`code`][0])
			}
			if msg := data[`msg`]; msg[0] != `操作成功` && msg[0] != `查询成功` {
				t.Errorf("%s = %v, Maps[`msg`] %v", TestTextName, msg, tt.Want[`msg`][0], tt.Want[`msg`][1])
			}
		})
	}
	return tests
}
