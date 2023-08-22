package APIs

import "fmt"

type AuthInfoData struct {
	B  bool
	If InfoData
}

var AInfo = &AuthInfoData{B: false}

func (A *API) Login(session string) {
	defer ReErr("Login")
	if !AInfo.B {
		In, B := Query(Http, Account, "SYNO.API.Auth")
		if B {
			AInfo.B = true
			AInfo.If = In
		}
	}
	by := A.APIClient.HttpGet("login", fmt.Sprintf("api=SYNO.API.Auth&version=%d&%s&session=%s&format=cookie", AInfo.If.MaxVersion, Account, session))
	fmt.Println(string(by))
}

func (A *API) Logout() {
	defer ReErr("Logout")

	//http://myds.com:5000/webapi/auth.cgi?api=SYNO.API.Auth&version=1&method=logout&session=DownloadStation
	A.APIClient.HttpGet("logout", "")
}
