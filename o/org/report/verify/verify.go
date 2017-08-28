package verify

import (
	"qrcode-bulk/qrcode-bulk-generator/x/db/mgo"
	"qrcode-bulk/qrcode-bulk-generator/x/mlog"
)

var objreportLoging = mlog.NewTagLog("obj_Report")

type Verify struct {
	mgo.BaseModel `bson:",inline"`
	QrCodeID      string        `bson:"qrcode_id" json:"qrcode_id"`
	UserID        string        `bson:"user_id" json:"user_id"`
	Name          string        `bson:"name" json:"name"`
	ScanIP        string        `bson:"scan_ip" json:"scan_ip"`
	DeviceInfo    *DeviceInfo   `bson:"device_info" json:"device_info"`
	LocationInfo  *LocationInfo `bson:"location_info" json:"location_info"`
}

type DeviceInfo struct {
	Browser  *Browser `bson:"browser" json:"browser"`
	OS       *OS      `bson:"os" json:"os"`
	Engine   *Engine  `bson:"engine" json:"engine"`
	IsMobile bool     `bson:"is_mobile" json:"is_mobile"`
	IsBot    bool     `bson:"is_bot" json:"is_bot"`
}

type Browser struct {
	Name    string
	Version string
}

type OS struct {
	Name     string
	Platform string `bson:"platform" json:"platform"`
}

type Engine struct {
	Name    string
	Version string
}

type LocationInfo struct {
	City    string `bson:"city" json:"city"`
	Country string `bson:"country" json:"country"`
	Loc     string `bson:"loc" json:"loc"`
	Org     string `bson:"org" json:"org"`
}

type IpInfo struct {
	IP       string `bson:"ip" json:"ip"`
	Hostname string `bson:"hostname" json:"hostname"`
	City     string `bson:"city" json:"city"`
	Region   string `bson:"region" json:"region"`
	Country  string `bson:"country" json:"country"`
	Loc      string `bson:"loc" json:"loc"`
	Org      string `bson:"org" json:"org"`
}

// type IPAPI struct {
// 	Status      string `json:"status"`
// 	Country     string `json:"country"`
// 	CountryCode string `json:"countryCode"`
// 	Region      string `json:"region"`
// 	RegionName  string `json:"regionName"`
// 	City        string `json:"city"`
// 	Zip         string `json:"zip"`
// 	Lat         string `json:"lat"`
// 	Lon         string `json:"lon"`
// 	Timezone    string `json:"timezone"`
// 	Isp         string `json:"isp"`
// 	Org         string `json:"org"`
// 	As          string `json:"as"`
// 	Query       string `json:"query"`
// }
