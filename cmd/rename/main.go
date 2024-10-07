package main

import (
	"fmt"
	"github.com/Mmx233/EpRename/match"
	"log"
	"os"
	"path/filepath"
)

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
			log.Fatal(err)
		}
		if err = os.Rename(
			file.Name(),
			filepath.Join(
				fileDir,
				fmt.Sprintf("S%sE%s%s", season, episode, ext),
			),
		); err != nil {
			log.Fatal(err)
		}
	}
}
