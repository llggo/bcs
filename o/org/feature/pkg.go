package feature

type Pkg struct {
	Code        string  `bson:"code" json:"code"`
	Name        string  `bson:"name_package" json:"name"`
	QrcodeLimit int     `bson:"qrcode_limit" json:"qrcode_limit"`
	TimeLife    int     `bson:"time_life" json:"time_life"`
	Feature     Feature `bson:"feature" json:"feature"`
}

type Message struct {
	Access  bool   `bson:"access" json:"access"`
	Message string `bson:"message" json:"message"`
}

func (p *Pkg) CheckAccess(name FeatureName, action FeatureAction) (bool, *Message) {
	if _, ok := p.Feature[name]; !ok {
		return false, &Message{
			Access:  false,
			Message: "Access is denied. Check your subcrition",
		}
	}
	if val, ok := p.Feature[name][action]; ok {
		if !val {
			return false, &Message{
				Access:  false,
				Message: "Access is denied. Check your subcrition",
			}
		}
	}
	return true, nil
}

func GetPkg(code string) *Pkg {
	for _, v := range Pkgs {
		if v.Code == code {
			return &v
		}
	}
	return nil
}

var Pkgs = []Pkg{
	Free,
	Basic,
	Standard,
	Advance,
	Premium,
}

var Free = Pkg{
	Code:        "pkg_free",
	Name:        "Gói miễn phí",
	QrcodeLimit: 9,
	TimeLife:    30,
	Feature: Feature{
		Qrcode: Action{
			Create: true,
			List:   true,
			Delete: true,
		},
	},
}

var Basic = Pkg{
	Code:        "pkg_basic",
	Name:        "Gói cơ bản",
	QrcodeLimit: 99,
	TimeLife:    30,
	Feature: Feature{
		Qrcode: Action{
			Create: true,
			List:   true,
			Delete: true,
		},
	},
}

var Standard = Pkg{
	Code:        "pkg_standard",
	Name:        "Gói tiêu chuẩn",
	QrcodeLimit: 999,
	TimeLife:    30,
	Feature: Feature{
		Qrcode: Action{
			Create: true,
			List:   true,
			Delete: true,
		},
	},
}

var Advance = Pkg{
	Code:        "pkg_advance",
	Name:        "Gói nâng cao",
	QrcodeLimit: 9999,
	TimeLife:    30,
	Feature: Feature{
		Qrcode: Action{
			Create: true,
			List:   true,
			Delete: true,
		},
	},
}

var Premium = Pkg{
	Code:        "pkg_premium",
	Name:        "Gói chất lượng cao",
	QrcodeLimit: 99999,
	TimeLife:    30,
	Feature: Feature{
		Qrcode: Action{
			Create: true,
			List:   true,
			Delete: true,
		},
	},
}
