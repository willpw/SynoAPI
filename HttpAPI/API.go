package HttpAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type NasHttpC struct {
	Start   *bool
	Version int
	APIName string
	Path    string
	Http    string
	Account string
	Client  *http.Client
}

type AuthAPI struct {
	Success bool           `json:"Success,omitempty"`
	Error   map[string]int `json:"Error,omitempty"`
}

var Start *bool

func NewAPIC(AN, P, H, AC string, V int, S *bool, C *http.Client) *NasHttpC {
	defer ReErr("NewAPIC")
	Start = S
	return &NasHttpC{APIName: AN, Path: P, Http: H, Account: AC, Version: V, Start: S, Client: C}
}

func (N *NasHttpC) HttpGet(Method, url string) []byte {

	if !*N.Start {
		log.Panicln("HttpGet N.Start=", *N.Start, N.Start)
		return nil
	}

	Url := fmt.Sprintf("%s/webapi/%s?api=%s&version=%d&method=%s&%s", N.Http, N.Path, N.APIName, N.Version, Method, url)

	switch Method {
	case "login":
		Url = fmt.Sprintf("%s/webapi/entry.cgi?method=login&%s", N.Http, url)
	case "logout":
		//http://myds.com:5000/webapi/auth.cgi?api=SYNO.API.Auth&version=1&method=logout&session=DownloadStation
		Url = fmt.Sprintf("%s/webapi/auth.cgi?api=SYNO.API.Auth&version=1&method=logout&session=%s", N.Http, N.APIName)
	}

	fmt.Println(Method, Url)
	ret, _ := N.Client.Get(Url)
	defer func() { _ = ret.Body.Close() }()
	body, _ := io.ReadAll(ret.Body)
	if ret.StatusCode != 200 {
		panic(map[string]string{"StatusCode": ret.Status})
	}

	var RetAPI AuthAPI
	//fmt.Println(string(body))
	_ = json.Unmarshal(body, &RetAPI)
	if !RetAPI.Success {
		panic(map[string]int{"Success": RetAPI.Error["code"]})
	}

	return body
}

func (N *NasHttpC) HttpPost(Method string, data url.Values) []byte {
	if !*N.Start {
		return nil
	}

	Url := fmt.Sprintf("%s/webapi/%s?api=%s&version=%d&method=%s", N.Http, N.Path, N.APIName, N.Version, Method)

	switch Method {
	case "CreateFile":
		Url = fmt.Sprintf("%s/webapi/DownloadStation/task.cgi", N.Http)
	}

	fmt.Println(Url)
	ret, err := N.Client.PostForm(Url, data)
	if err != nil {
		panic(err)
	}
	defer func() { _ = ret.Body.Close() }()
	body, _ := io.ReadAll(ret.Body)
	if ret.StatusCode != 200 {
		panic(map[string]string{"StatusCode": ret.Status})
	}

	var RetAPI AuthAPI

	_ = json.Unmarshal(body, &RetAPI)
	if !RetAPI.Success {
		panic(map[string]int{"Success": RetAPI.Error["code"]})
	}

	return body
}
