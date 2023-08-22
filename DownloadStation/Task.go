package DownloadStationS

import (
	"encoding/json"
	"fmt"
	APIs "github.com/willpw/SynoAPI/API"
	"github.com/willpw/SynoAPI/HttpAPI"
	"net/url"
)

type TaskT struct {
	Start    *bool
	TaClient *HttpAPI.NasHttpC
	API      *APIs.API
}

type TaskTasks struct {
	Size     int    `json:"size,omitempty"`
	Id       string `json:"id,omitempty"`
	Status   string `json:"status,omitempty"`
	Title    string `json:"title,omitempty"`
	Type     string `json:"type,omitempty"`
	Username string `json:"username,omitempty"`
}

type TaskList struct {
	Offset int         `json:"offset,omitempty"`
	Total  int         `json:"total,omitempty"`
	Tasks  []TaskTasks `json:"tasks,omitempty"`
}

type DSTaskListS struct {
	Success bool     `json:"success,omitempty"`
	Data    TaskList `json:"data"`
}

var (
	Task *TaskT
)

func NewTask(B *bool, DSClient *HttpAPI.NasHttpC, API *APIs.API) *TaskT {
	defer ReErr("NewTask")
	Task = &TaskT{Start: B, TaClient: DSClient, API: API}
	return Task
}

// List 获取列表
func (T *TaskT) List(url string, L *DSTaskListS) {

	defer ReErr("TaskList")
	//http://myds.com:5000/webapi/DownloadStation/task.cgi?api=SYNO.DownloadStation.Task&version=1&method=list

	//GET /webapi/DownloadStation/task.cgi?api=SYNO.DownloadStation.Task&version=1&method=list&additional=detail,file
	fmt.Println("List=", url, !*T.Start)
	if !*T.Start {

		T.API.Login("SYNO.DownloadStation.Task")

	}

	by := T.TaClient.HttpGet("list", url)

	//fmt.Println(API.Data)

	json.Unmarshal(by, L)
	//fmt.Println(string(by), L)

}

// GetInfo 获取信息
func (T *TaskT) GetInfo(url string) {

	defer ReErr("TaskGetInfo")
	//GET /webapi/DownloadStation/task.cgi? api=SYNO.DownloadStation.Task&version=1&method=getinfo&id=dbid_001,dbid_002&additional=detail
	if !*T.Start {
		T.API.Login("SYNO.DownloadStation.Task")
	}

}

// CreateUrl 创建URL任务
func (T *TaskT) CreateUrl(url, Dir string) {

	defer ReErr("TaskCreateUrl")
	//POST                     /webapi/DownloadStation/task.cgi?api=SYNO.DownloadStation.Task&version=1&method=create&uri=ftps://192.0.0.1:21/test/test.zip&username=admin&password=123
	//http://192.168.0.249:5000/webapi/DownloadStation/task.cgi?api=SYNO.DownloadStation.Task&version=3&method=create&uri=ftps://192.0.0.1:21/test/test.zip
	if !*T.Start {
		T.API.Login("SYNO.DownloadStation.Task")
	}
	if Dir != "" {
		url += "&destination=" + Dir
	}
	byt := T.TaClient.HttpGet("create", "uri="+url)
	_ = byt
	fmt.Println(string(byt))
}

// Create 创建任务
func (T *TaskT) Create(ur, Dir string) {
	defer ReErr("TaskCreateFile")
	//POST /webapi/DownloadStation/task.cgi
	//
	//api=SYNO.DownloadStation.Task&version=1&method=create&uri=ftps://192.0.0.1:2 1/test/test.zip&username=admin&password=123

	if !*T.Start {
		T.API.Login("SYNO.DownloadStation.Task")
	}

	var Files = url.Values{"api": {"SYNO.DownloadStation.Task"}, "version": {"3"}, "method": {"create"}, "uri": {ur}, "destination": {Dir}}

	byt := T.TaClient.HttpPost("CreateFile", Files)
	_ = byt
	fmt.Println(string(byt))
}

// Delete 删除任务
func (T *TaskT) Delete(url string) {
	defer ReErr("TaskDelete")
	//GET /webapi/DownloadStation/task.cgi?api=SYNO.DownloadStation.Task&version=1&method=delete&id=dbid_001,dbid_002&force_complete=true
	if !*T.Start {
		T.API.Login("SYNO.DownloadStation.Task")
	}
	by := T.TaClient.HttpGet("list", url)
	_ = by
}

// Pause 暂停任务
func (T *TaskT) Pause(url string) {
	defer ReErr("TaskPause")
	//GET /webapi/DownloadStation/task.cgi?api=SYNO.DownloadStation.Task&version=1&method=pause&id=dbid_001,dbid_002
	if !*T.Start {
		T.API.Login("SYNO.DownloadStation.Task")
	}

}

// Resume 重新开始
func (T *TaskT) kResume(url string) {
	defer ReErr("TaskResume")
	//GET /webapi/DownloadStation/task.cgi?api=SYNO.DownloadStation.Task&version=1&method=resume&id=dbid_001,dbid_002
	if !*T.Start {
		T.API.Login("SYNO.DownloadStation.Task")
	}

}

// Edit  编辑
func (T *TaskT) Edit(url string) {
	defer ReErr("TaskEdit")
	//GET /webapi/DownloadStation/task.cgi?api=SYNO.DownloadStation.Task&version=1&method=edit&id=dbid_001,dbid_002&destinatio n=sharedfolder
	if !*T.Start {
		T.API.Login("SYNO.DownloadStation.Task")
	}

}
