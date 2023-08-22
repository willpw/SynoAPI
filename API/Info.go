package APIs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type InfoData struct {
	MaxVersion    int    `json:"MaxVersion,omitempty"`
	MinVersion    int    `json:"MinVersion,omitempty"`
	Path          string `json:"Path,omitempty"`
	RequestFormat string `json:"RequestFormat,omitempty"`
}

type AuthAPI struct {
	Success bool                `json:"Success,omitempty"`
	Error   map[string]int      `json:"Error,omitempty"`
	Data    map[string]InfoData `json:"Data,omitempty"`
}

func Query(H, A, query string) (In InfoData, b bool) {
	defer ReErr("Query")
	Http, Account = H, A
	url := fmt.Sprintf("%s/webapi/query.cgi?api=SYNO.API.Info&version=1&method=query&query=%s", H, query)
	ret, _ := http.Get(url)
	defer func() { _ = ret.Body.Close() }()
	by, _ := io.ReadAll(ret.Body)
	var API AuthAPI
	json.Unmarshal(by, &API)
	if API.Success {
		return API.Data[query], true
	}
	panic("Query 获取失败")
	return InfoData{}, false
}
