package verify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mssola/user_agent"
)

var IPCHAN = make(chan *Verify, 1024)

func SendScan(s *Verify) {
	select {
	case IPCHAN <- s:
	default:
		fmt.Println("IPCHAN is Full")
	}
}

func ExIP() {
	for {
		var s = <-IPCHAN
		s.LocationInfo = GetLocation(s.ScanIP)
		s.Update(s)
	}
}

func init() {
	go ExIP()
}

func GetLocation(ip string) *LocationInfo {

	var i IpInfo

	var url = "http://ipinfo.io/" + ip + "/json"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(body, &i)

	l := LocationInfo{}
	l.City = i.City
	l.Country = i.Country
	l.Loc = i.Loc
	l.Org = i.Org

	return &l
}

func GetDeviceInfo(r *http.Request) *DeviceInfo {
	var d = DeviceInfo{}
	ua := user_agent.New(r.UserAgent())

	name, version := ua.Browser()

	d.Browser = &Browser{}
	d.Browser.Name = name
	d.Browser.Version = version

	name, version = ua.Engine()

	d.Engine = &Engine{}
	d.Engine.Name = name
	d.Engine.Version = version

	d.OS = &OS{}
	d.OS.Name = ua.OS()
	d.OS.Platform = ua.Platform()
	d.IsBot = ua.Bot()
	d.IsMobile = ua.Mobile()

	return &d
}
