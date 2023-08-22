package main

import (
	"fmt"
	SYNO "github.com/willpw/SynoAPI"
	DownloadStationS "github.com/willpw/SynoAPI/DownloadStation"
)

func main() {

	SYNO.SetIPAcc("192.168.0.249", "5000", "willpwr", "231618", false)

	var List = &DownloadStationS.DSTaskListS{}
	SYNO.DownloadStation.Task.List("", List)
	fmt.Println(List.Data.Tasks)

	SYNO.DownloadStation.Task.Create("http://192.168.0.1:8080/5.torrent", "XZ/视频/NAS下载/20230815")
	SYNO.DownloadStation.Logout()

}
