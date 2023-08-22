package APIs

import (
	"fmt"
	"github.com/willpw/SynoAPI/HttpAPI"
	"log"
	"strings"
)

type API struct {
	APIClient *HttpAPI.NasHttpC
}

var (
	Http    = ""
	Account = ""
	Start   *bool
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

func NewAPI(H, A string, APIStart *bool, Client *HttpAPI.NasHttpC) *API {
	defer ReErr("NewAPI")
	Http, Account = H, A
	Start = APIStart
	AP := &API{APIClient: Client}
	return AP
}
