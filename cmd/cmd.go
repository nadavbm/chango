package cmd

import (
	"flag"
	"time"

	"github.com/nadavbm/chango/decorator"
	"github.com/nadavbm/chango/factory"
	"github.com/nadavbm/chango/observer"
	"github.com/nadavbm/chango/singleton"
)

var pattern, filename, imageName, imageTag, imageSha string
var duration, interval time.Duration
var workers int

// ParseCommandArgs
func ParseCommandArgs() {
	flag.StringVar(&pattern, "pattern", "observer", "start observer pattern")
	flag.StringVar(&filename, "filename", "data.json", "file name location")
	flag.StringVar(&imageName, "name", "etzba/etz", "image name from docker registry")
	flag.StringVar(&imageTag, "tag", "latest", "image tag")
	flag.StringVar(&imageSha, "sha", "sha256:d135a04a2ac74466c7c01747daee7d4efae7d09e457b0e8d54bc510f2be1408a", "image sha digest")
	flag.IntVar(&workers, "workers", 20, "add duration to command execution")
	flag.DurationVar(&duration, "duration", 20*time.Second, "add duration to command execution")
	flag.DurationVar(&interval, "interval", 3*time.Second, "add interval to command execution")
	flag.Parse()
}

func Execute() {
	logger := decorator.Log{}
	log := decorator.WithMessage(logger)
	config := singleton.GetConfig(filename, workers, duration)
	image := observer.Image{Name: imageName, Sha: imageSha, Tag: imageTag}
	log.Info("start executing pattern " + pattern)
	exec := factory.ExecutionFactory(logger, config, image, pattern)
	exec.Execute()
}
