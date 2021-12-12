package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/superlars1337/go-basics/log"
	"github.com/superlars1337/go-basics/utils"
)

// Renpho functions

type exifData struct {
	camera    string
	date      time.Time
	timestamp string
}

// helper: read exif info from file
func getExifDataFromFile(filename string) exifData {
	f, err := os.Open(filename)
	log.CheckErrFatal(err)
	defer f.Close()

	x, err := exif.Decode(f)
	log.CheckErrFatal(err)

	camera := "unknown"
	camModel, err := x.Get(exif.Model)
	if err != nil {
		log.Error(err, "in file "+filepath.Base(filename)+" - Defaulting to '"+camera+"' ")
	} else {
		camera, _ = camModel.StringVal()
		camera = strings.Replace(camera, " ", "-", -1)
	}

	tm, err := x.DateTime()
	if err != nil {
		log.Error(err, "in file "+filepath.Base(filename)+" - Defaulting to 'TODAY' ")
		tm = time.Now()
	}
	// 2017-10-05 14:09:33 +0200 CEST
	timestamp := fmt.Sprintf("%d-%02d-%02d-%02dh%02dm%02d",
		tm.Year(), tm.Month(), tm.Day(),
		tm.Hour(), tm.Minute(), tm.Second())

	return exifData{
		camera:    camera,
		date:      tm,
		timestamp: timestamp,
	}
}

// main function
func renamePhotos(photoPath string, label string, dryrun bool) {
	// search for JPGs
	filter := fmt.Sprintf("%s/*.JPG", photoPath)
	files, err := filepath.Glob(filter)
	log.CheckErrFatal(err)
	if len(files) < 1 {
		log.ExitNow("UPS - Sorry - No JPG files found :(  (*.JPG)")
	}

	// dryrun mode
	if dryrun {
		log.Success("Dryrun - won`t change anything")
	}

	// create new JPG directory
	today := utils.GetDateTimestamp()
	newJpgDir := fmt.Sprintf("%s/%s_%s_JPG", filepath.Dir(photoPath), today, label)
	newRawDir := fmt.Sprintf("%s/%s_%s_RAW", filepath.Dir(photoPath), today, label)
	log.Debug("new JPG directory: " + newJpgDir)
	log.Debug("new RAW directory: " + newRawDir)
	if !dryrun {
		err = os.MkdirAll(newJpgDir, 0755)
		log.CheckErrFatal(err)
	}

	// process JPGs
	for _, filename := range files {
		// get exif data
		exifData := getExifDataFromFile(filename)
		log.Debug("--")
		log.Debug("file     : " + filename)
		log.Debug("camera   : " + exifData.camera)
		log.Debug("timestamp: " + exifData.timestamp)

		// move and rename files
		newFilename := fmt.Sprintf("%s/%s_%s_%s", newJpgDir, exifData.timestamp, exifData.camera, filepath.Base(filename))
		log.Info("moving " + filepath.Base(filename) + " -> " + newFilename)
		if !dryrun {
			err = os.Rename(filename, newFilename)
			log.CheckErrFatal(err)
			err = os.Chtimes(newFilename, exifData.date, exifData.date)
			log.CheckErrFatal(err)
		}
	}

	// rename input dir to newRawDir
	if !dryrun {
		err = os.Rename(photoPath, newRawDir)
		log.CheckErrFatal(err)
	}
}
