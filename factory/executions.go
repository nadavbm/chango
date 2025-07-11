package factory

import (
	"fmt"

	"github.com/nadavbm/chango/decorator"
	"github.com/nadavbm/chango/fanio"
	"github.com/nadavbm/chango/hello"
	"github.com/nadavbm/chango/messaging"
	"github.com/nadavbm/chango/observer"
	"github.com/nadavbm/chango/pipeline"
	"github.com/nadavbm/chango/pubsub"
	"github.com/nadavbm/chango/repository"
	"github.com/nadavbm/chango/singleton"
	"github.com/nadavbm/chango/workerpool"
)

type Execution interface {
	Execute()
}

type Hello struct {
}

func (h Hello) Execute() {
	hello.SayHello()
}

type Observer struct {
	logger decorator.Logger
	image  observer.Image
}

func (o Observer) Execute() {
	o.logger.Info("image registry update args " + o.image.Name + " " + o.image.Tag + " " + o.image.Sha)
	observer.UpdateImageInRegistry(o.image.Name, o.image.Tag, o.image.Sha)
}

type Messaging struct {
	logger decorator.Logger
	config *singleton.Config
}

func (m Messaging) Execute() {
	messaging.SendReceive(m.logger, m.config)
}

type WorkerPool struct {
	logger decorator.Logger
	config *singleton.Config
}

func (m WorkerPool) Execute() {
	workerpool.WorkInSupermarket(m.logger, m.config)
}

type FanOutFanIn struct {
	logger decorator.Logger
	config *singleton.Config
}

func (m FanOutFanIn) Execute() {
	fanio.MathClass(m.logger, m.config)
}

type Repository struct {
	logger decorator.Logger
	config *singleton.Config
}

func (m Repository) Execute() {
	repository.DynamicPhoneBook(m.logger, m.config)
}

type PipeLine struct {
}

func (m PipeLine) Execute() {
	pipeline.LuckySupermarket()
}

type PubSub struct {
}

func (m PubSub) Execute() {
	pubsub.Lottery()
}

func ExecutionFactory(logger decorator.Logger, config *singleton.Config, image observer.Image, pattern string) Execution {
	switch pattern {
	case "hello":
		return Hello{}
	case "observer":
		return Observer{
			logger: logger,
			image:  image,
		}
	case "messaging":
		return Messaging{
			logger: logger,
			config: config,
		}
	case "workerpool":
		return WorkerPool{
			logger: logger,
			config: config,
		}
	case "fanio":
		return FanOutFanIn{
			logger: logger,
			config: config,
		}
	case "repository":
		return Repository{
			logger: logger,
			config: config,
		}
	case "pipeline":
		return PipeLine{}
	case "pubsub":
		return PubSub{}
	}
	fmt.Println("Choose relevant pattern by using -pattern=<pattern_name>. Details in README.md")
	return nil
}
