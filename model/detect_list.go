package model

//go:generate reform

// DetectList represents a row in detect_list table.
//reform:detect_list
type DetectList struct {
	ID    int64  `reform:"id,pk"`
	Name  string `reform:"name"`
	Text  string `reform:"text"`
	Regex string `reform:"regex"`
	Type  int32  `reform:"type"`
}
