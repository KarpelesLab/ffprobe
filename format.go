package ffprobe

type Format struct {
	Filename       string            `json:"filename"`
	FormatName     string            `json:"format_name"`       // mov,mp4,m4a,3gp,3g2,mj2
	FormatLongName string            `json:"format_long_name"`  // QuickTime / MOV
	StartTime      float64           `json:"start_time,string"` // "start_time": "0.000000",
	Duration       float64           `json:"duration,string"`   // "duration": "240.048000",
	Size           int64             `json:"size,string"`       // "size": "73087904",
	BitRate        int               `json:"bit_rate,string"`   // "bit_rate": "2435776",
	Tags           map[string]string `json:"tags"`
}
