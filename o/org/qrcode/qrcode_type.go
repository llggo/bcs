package qrcode

/********Text --- Email --- Youtube*******/
type TPEY struct {
	Content string `json:"content" bson:"content"`
}
type Phone struct {
	Number string `json:"number" bson:"number"`
}

//URL Type
type URL struct {
	URL string `json:"url" bson:"url"`
}

//Urls Type
type Urls struct {
	Name string `json:"name" bson:"name"`
	URL  string `json:"url" bson:"url"`
}

//URLs Type

type MultiUrl struct {
	LinkDefaul string `json:"linkdf" bson:"linkdf"`
	URLs       []url
}
type url struct {
	Link   string `json:"link" bson:"link"`
	Filter filter
}
type filter struct {
	Local     string `json:"local" bson:"local"`
	Device    string `json:"device" bson:"device"`
	TimeStart int64
	TimeEnd   int64
}

//File Type
type File struct {
	Name string `json:"name" bson:"name"`
	Path string `json:"path" bson:"path"`
}

//type SMS - MMS
type MS struct {
	SMSTo   string `json:"smsto" bson:"smsto"`
	Content string `json:"content" bson:"content"`
}

// BEGIN:VCARD
// VERSION:4.0
// N:Forrest;Gump;;Mr.;
// FN:Forrest Gump
// ORG:Bubba Gump Shrimp Co.
// TITLE:Shrimp Man
// PHOTO;MEDIATYPE=image/gif:http://www.example.com/dir_photos/my_photo.gif
// TEL;TYPE=work,voice;VALUE=uri:tel:+1-111-555-1212
// TEL;TYPE=home,voice;VALUE=uri:tel:+1-404-555-1212
// ADR;TYPE=WORK,PREF:;;100 Waters Edge;Baytown;LA;30314;United States of Amer
//  ica
// LABEL;TYPE=WORK,PREF:100 Waters Edge\nBaytown\, LA 30314\nUnited States of
//  America
// ADR;TYPE=HOME:;;42 Plantation St.;Baytown;LA;30314;United States of America
// LABEL;TYPE=HOME:42 Plantation St.\nBaytown\, LA 30314\nUnited States of Ame
//  rica
// EMAIL:forrestgump@example.com
// REV:20080424T195243Z
// END:VCARD

type VCard struct {
}

// BEGIN:VCALENDAR
// VERSION:1.0
// BEGIN:VEVENT
// CATEGORIES:MEETING
// STATUS:TENTATIVE
// DTSTART:19960401T033000Z
// DTEND:19960401T043000Z
// SUMMARY:Your Proposal Review
// DESCRIPTION:Steve and John to review newest proposal material
// CLASS:PRIVATE
// END:VEVENT
// END:VCALENDAR
//BIZCARD:N:Sean;X:Owen;T:Software Engineer;C:Google;A:76 9th Avenue, New York, NY 10011;B:+12125551212;E:srowen@google.com;;
type VCalendar struct {
	Categpries string `json:"categories" bson:"categories"`
	Status     string `json:"status" bson:"status"`
	DtStart    string `json:"dtstart" bson:"dtstart"`
	DtEnd      string `json:"dtend" bson:"dtend"`
	Summary    string `json:"summary" bson:"summary"`
	Desciption string `json:"description" bson:"description"`
}

type VEVENT struct {
	Summary string `json:"summary" bson:"summary"`
	DtStart string `json:"dtstart" bson:"dtstart"`
	DtEnd   string `json:"dtend" bson:"dtend"`
}
type GEO struct {
	Longitude string `json:"longitude" bson:"longitude"` //kinh do
	Latitude  string `json:"latitude" bson:"latitude"`   // vi do
}
type Wifi struct {
	Name string `json:"namewifi" bson:"namewifi"` //kinh do
	Pass string `json:"password" bson:"password"` // vi do
}

/********Image --- PDF --- Audio*******/
type IPA struct {
	Path string `json:"path" bson:"path"`
}
