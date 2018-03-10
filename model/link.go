package model

//go:generate reform

// Link represents a row in link table.
//reform:link
type Link struct {
	ID      int64   `reform:"id,pk"`
	Type    int32   `reform:"type"`
	Address string  `reform:"address"`
	Port    int32   `reform:"port"`
	Token   string  `reform:"token"`
	Ios     int32   `reform:"ios"`
	Userid  int64   `reform:"userid"`
	Isp     *string `reform:"isp"`
	Geo     *int32  `reform:"geo"`
	Method  *string `reform:"method"`
}
