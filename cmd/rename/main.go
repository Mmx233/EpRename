package main

import (
	"fmt"
	"github.com/Mmx233/EpRename/match"
	"os"
	"path/filepath"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	season, err := match.GetSeasonSerial(pwd)
	if err != nil {
		log.Fatal(err)
	}

	dir, err := os.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dir {
		fileDir, fileName := filepath.Split(file.Name())
		ext := filepath.Ext(fileName)
		episode, err := match.GetEpisodeSerial(fileName)
		if err != nil {
			log.Warnf("rename '%s' failed: %v", fileName, err)
			continue
		}
		episodeName := fmt.Sprintf("S%sE%s%s", season, episode, ext)
		if err = os.Rename(
			file.Name(),
			filepath.Join(fileDir, episodeName),
		); err != nil {
			log.Errorf("rename '%s' failed: %v", fileName, err)
			continue
		}

		log.Infof("successfully renamed '%s' to '%s'", fileName, episodeName)
	}
}
