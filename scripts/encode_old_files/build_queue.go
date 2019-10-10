package main

import (
	"github.com/tett23/kinsro/scripts/encode_old_files/src/videotmp"
)

// VideoTmpPath VideoTmpPath
const VideoTmpPath = "/Volumes/video_tmp"

// QueuePath QueuePath
const QueuePath = "/Volumes/video1/queue.json"

func main() {
	filtered, err := videotmp.BuildVideoTmpList(VideoTmpPath)
	if err != nil {
		panic(err)
	}

	err = videotmp.WriteQueue(QueuePath, filtered)
	if err != nil {
		panic(err)
	}
}
