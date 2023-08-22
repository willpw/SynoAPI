package SYNO

import (
	"fmt"
	APIs "github.com/willpw/SynoAPI/API"
	DownloadStationS "github.com/willpw/SynoAPI/DownloadStation"
)

var (
	Http     = "http://192.168.0.249:5000"
	Account  = "account=willpwr&passwd=231618"
	APIStart bool

	DownloadStation *DownloadStationS.DownloadStat
	API             *APIs.API
)

// SetIPAcc 设置Nas地址端口
func SetIPAcc(ip, prot, Acc, Pass string, Https bool) {
	defer ReErr("SetIPAcc")
	APIStart = true
	Http = fmt.Sprintf("http://%s:%s", ip, prot)
	if Https {
		Http = fmt.Sprintf("https://%s:%s", ip, prot)
	}
	Account = fmt.Sprintf("account=%s&passwd=%s", Acc, Pass)

	NewAPIS(Http, Account)

}

func NewAPIS(H, A string) {
	defer ReErr("NewAPIS")

	DownloadStation = DownloadStationS.NewDownloadStation(H, A, &APIStart)

	//API = APIs.NewAPI(Http, Account, APIStart)
}
