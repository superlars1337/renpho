package main

import (
	"flag"
	"regexp"
	"strings"

	"github.com/superlars1337/go-basics/log"
	"github.com/superlars1337/go-basics/utils"
)

var (
	debug     *bool
	dryrun    *bool
	photoPath *string
	label     *string
)

func init() {
	photoPath = flag.String("p", "", "path to your photo directory")
	label = flag.String("l", "renpho", "short info label for the collection eg 'Easter Holiday'")
	debug = flag.Bool("debug", false, "debug flag")
	dryrun = flag.Bool("dryrun", false, "dryrun - dont change anything")
	flag.Parse()
	log.Init(log.LOG_PRETTY, *debug)
}

func usage() {
	log.Message("------------------------------------------------")
	log.Message("Usage: renpho -p <path-to-your-photos-directory>")
	log.Message("Possible flags:")
	flag.PrintDefaults()
	log.Message("------------------------------------------------")
	log.ExitNow("Byebye")
}

func main() {
	utils.PrintTitle("Rename Photos")

	// check if photoPath exists
	if (len(*photoPath) == 0) || (!utils.FileExists(*photoPath)) {
		log.ErrorS("Ups - Input directory not found")
		usage()
	}
	log.Info("Path : " + *photoPath)

	// check label, sanitize string
	l := *label
	if len(l) > 0 {
		reg, err := regexp.Compile("[^a-zA-Z0-9-]+")
		log.CheckErrFatal(err)
		l = strings.Replace(l, " ", "-", -1)
		l = reg.ReplaceAllString(l, "")
	}
	log.Info("Label: " + l)
	log.Info("======================================")

	// start renaming
	renamePhotos(*photoPath, l, *dryrun)
	log.Success("All done :)")
}
