package ffprobe

import (
	"os"
	"path/filepath"

	"github.com/KarpelesLab/runutil"
)

type File struct {
	Streams  []*Stream  `json:"streams"`
	Chapters []*Chapter `json:"chapters"`
	Format   *Format    `json:"format"`
	filename string
}

func Probe(fn string) (*File, error) {
	var info *File
	err := runutil.RunJson(&info, exe("ffprobe"), "-print_format", "json", "-hide_banner", "-loglevel", "warning", "-show_format", "-show_streams", "-show_chapters", fn)
	if err != nil {
		return nil, err
	}
	info.filename = fn

	return info, nil
}

// exe returns the path for the executable if found in azusa media-video/ffmpeg, and will just return the parameter if not found
func exe(n string) string {
	// /pkg/main/media-video.ffmpeg.core/bin/...
	p := filepath.Join("/pkg/main/media-video.ffmpeg.core/bin", n)
	if _, err := os.Stat(p); err == nil {
		return p
	}

	// by default exec will perform path lookup, so just return n
	return n
}
