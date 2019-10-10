package videotmp

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
)

type maybeVideoTmp struct {
	TSFileName      *string
	ProgramFileName *string
	ErrorFileName   *string
}

// VideoTmp VideoTmp
type VideoTmp struct {
	TSFileName      string `json:"ts"`
	ProgramFileName string `json:"program"`
	ErrorFileName   string `json:"error"`
}

// BuildVideoTmpList BuildVideoTmpList
func BuildVideoTmpList(videoTmpDir string) ([]VideoTmp, error) {
	maybeVideoTmp, err := loadTemporaryVideoList(videoTmpDir)
	if err != nil {
		return nil, nil
	}

	return filterVideoTmp(maybeVideoTmp), nil
}

func loadTemporaryVideoList(videoTmpDir string) ([]maybeVideoTmp, error) {
	tsFile := regexp.MustCompile("\\.ts$")
	programFile := regexp.MustCompile("\\.program\\.txt$")
	errFile := regexp.MustCompile("\\.err$")

	items, err := ioutil.ReadDir(videoTmpDir)
	if err != nil {
		return nil, err
	}

	var tmpMap = map[string]maybeVideoTmp{}
	for _, item := range items {
		name := item.Name()
		if len(name) < 20 {
			continue
		}

		key := name[:18]
		if _, ok := tmpMap[key]; !ok {
			tmpMap[key] = maybeVideoTmp{
				TSFileName:      nil,
				ErrorFileName:   nil,
				ProgramFileName: nil,
			}
		}

		if tsFile.MatchString(name) {
			tmp := tmpMap[key]
			tmp.TSFileName = &name
			tmpMap[key] = tmp
			continue
		}
		if programFile.MatchString(name) {
			tmp := tmpMap[key]
			tmp.ProgramFileName = &name
			tmpMap[key] = tmp
			continue
		}
		if errFile.MatchString(name) {
			tmp := tmpMap[key]
			tmp.ErrorFileName = &name
			tmpMap[key] = tmp
			continue
		}
	}

	ret := []maybeVideoTmp{}
	for _, item := range tmpMap {
		ret = append(ret, item)
	}
	return ret, nil
}

func filterVideoTmp(list []maybeVideoTmp) []VideoTmp {
	ret := []VideoTmp{}
	for _, item := range list {
		if item.TSFileName == nil || item.ProgramFileName == nil || item.ErrorFileName == nil {
			continue
		}

		ret = append(ret, VideoTmp{
			TSFileName:      *item.TSFileName,
			ProgramFileName: *item.ProgramFileName,
			ErrorFileName:   *item.ErrorFileName,
		})
	}

	return ret
}

// WriteQueue WriteQueue
func WriteQueue(path string, list []VideoTmp) error {
	bytes, err := json.Marshal(list)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, bytes, 0644)
}

// ReadQueue WriteQueue
func ReadQueue(path string) (EncodeQueue, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var ret []VideoTmp
	err = json.Unmarshal(bytes, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// EncodeQueue EncodeQueue
type EncodeQueue []VideoTmp

func NewEncodeQueue(path string, videoTmps []VideoTmp) (EncodeQueue, error) {
	queue, err := ReadQueue(path)
	if err != nil {
		return nil, err
	}
}

// Deque Deque
func (encodeQueue EncodeQueue) Deque() *VideoTmp {
	return &encodeQueue[0]
}
