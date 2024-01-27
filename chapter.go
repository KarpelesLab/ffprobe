package ffprobe

type Chapter struct {
	Id        int64             `json:"id"`
	TimeBase  Fraction          `json:"time_base"`
	Start     int64             `json:"start"`
	StartTime float64           `json:"start_time,string"`
	End       int64             `json:"end"`
	EndTime   float64           `json:"end_time,string"`
	Tags      map[string]string `json:"tags,omitempty"`
}
