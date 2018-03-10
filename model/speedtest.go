package model

//go:generate reform

// Speedtest represents a row in speedtest table.
//reform:speedtest
type Speedtest struct {
	ID               int64  `reform:"id,pk"`
	Nodeid           int32  `reform:"nodeid"`
	Datetime         int64  `reform:"datetime"`
	Telecomping      string `reform:"telecomping"`
	Telecomeupload   string `reform:"telecomeupload"`
	Telecomedownload string `reform:"telecomedownload"`
	Unicomping       string `reform:"unicomping"`
	Unicomupload     string `reform:"unicomupload"`
	Unicomdownload   string `reform:"unicomdownload"`
	Cmccping         string `reform:"cmccping"`
	Cmccupload       string `reform:"cmccupload"`
	Cmccdownload     string `reform:"cmccdownload"`
}
