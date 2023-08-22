package DownloadStationS

import (
	"fmt"
	APIs "github.com/willpw/SynoAPI/API"
	"github.com/willpw/SynoAPI/HttpAPI"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

type DownloadStat struct {
	Start    bool
	TaskAPIs string `json:"TaskAPIs,omitempty"`
	DSClient *HttpAPI.NasHttpC
	API      *APIs.API
	Task     *TaskT
}

var (
	Http            = ""
	Account         = ""
	Start           *bool
	DStart          = false
	DownloadStation *DownloadStat
)

func ReErr(fun string) {
	if err := recover(); err != nil {
		switch err.(type) {

		case map[string]string:

			//网页错误
			*Start = false
			log.Println("StatusCode", err.(map[string]string)["StatusCode"])

		case map[string]int:

			//API 错误
			switch {
			case strings.Contains(fun, "Task"):

				switch err.(map[string]int)["Success"] {
				case 400:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "400 上传文件失败"))
				case 401:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "401 已达到的最大任务数"))
				case 402:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "402 目的地拒绝"))
				case 403:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "403 目标不存在"))
				case 404:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "404 无效的任务id"))
				case 405:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "405 无效的任务操作"))
				case 406:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "406 无默认目标"))
				case 407:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "407 设置目标失败"))
				case 408:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "408 文件不存在"))
				default:
					log.Println(fmt.Sprintf("fun = %s , DS Tsak ,Err = %d 未知错误类型", fun, err.(map[string]int)["Success"]))
				}

			default:

				switch err.(map[string]int)["Success"] {
				case 400:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "400 未知错误"))
				case 401:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "401 无效参数"))
				case 402:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "402 解析用户设置失败"))
				case 403:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "403 获取类别失败"))
				case 404:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "404 从数据库获取搜索结果失败"))
				case 405:
					log.Println(fmt.Sprintf("fun = %s , Err = %s", fun, "405 获取用户设置失败"))
				default:
					log.Println(fmt.Sprintf("fun = %s , default ,Err = %d 未知错误类型", fun, err.(map[string]int)["Success"]))
				}
			}

		default:

			//程序错误
			*Start = false
			log.Println(fmt.Sprintf("fun = %s , 程序错误 ,Err = %s", fun, err))

		}

	}
}

func NewDownloadStation(H, A string, APIStart *bool) *DownloadStat {
	defer ReErr("NewDownloadStat")
	DownloadStation = &DownloadStat{}
	Jar, _ := cookiejar.New(nil)
	DSC := &http.Client{Jar: Jar}
	Http, Account = H, A
	Start = APIStart
	I, B := APIs.Query(Http, Account, "SYNO.DownloadStation.Task")
	if B {

		DownloadStation.DSClient = HttpAPI.NewAPIC("SYNO.DownloadStation.Task", I.Path, H, A, I.MaxVersion, Start, DSC)
		DownloadStation.API = APIs.NewAPI(H, A, Start, DownloadStation.DSClient)
		DownloadStation.Task = NewTask(Start, HttpAPI.NewAPIC("SYNO.DownloadStation.Task", I.Path, H, A, I.MaxVersion, Start, DSC), DownloadStation.API)
	} else {
		panic("APIs.Quer 返回异常")
	}

	//创建HttpC
	DownloadStation.API.Login("SYNO.DownloadStation.Task")
	return DownloadStation
}

func (d *DownloadStat) Login() {
	defer ReErr("Login")
	d.API.Login("SYNO.DownloadStation.Task")
}

func (d *DownloadStat) Logout() {

	defer ReErr("Logout")
	//http://myds.com:5000/webapi/auth.cgi?api=SYNO.API.Auth&version=1&method=logout&session=DownloadStation
	d.API.Logout()
	//fmt.Println(string(by))
	d.Start = false

}
