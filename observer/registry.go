package observer

import (
	"fmt"
	"time"
)

type Observer interface {
	UpdateShaOrAddImage(name, tag, sha string)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Register(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Update(name, tag, sha string) {
	for _, observer := range s.observers {
		observer.UpdateShaOrAddImage(name, tag, sha)
	}
}

type Registry struct {
	Images []Image
}

type Image struct {
	Name    string
	Tag     string
	Sha     string
	Updated time.Time
}

func (r Registry) UpdateShaOrAddImage(name, tag, sha string) {
	updated := false
	for i, image := range r.Images {
		if image.Name == name && image.Tag == tag {
			r.Images[i].Sha = sha
			r.Images[i].Updated = time.Now()
			updated = true
		}
	}
	if !updated {
		fmt.Println("Image was not found in registry")
	}
}

func UpdateImageInRegistry(imageName, imageTag, imageSha string) {
	subject := Subject{}
	registry := Registry{
		Images: []Image{
			{Name: "etzba/etz", Tag: "development", Sha: "sha256:f42d7f87fb6388933ba346f14f800bc7550fa5e1df336f609e18a9415d592f3d", Updated: time.Date(2011, time.October, 10, 12, 23, 12, 1233241342, time.UTC)},
			{Name: "etzba/gopu", Tag: "latest", Sha: "sha256:8f4b2676dad4e9be0fe9a2fc2ad611b0d9fc81cf426d3e5593f7b2f11834a6b1", Updated: time.Date(2003, time.October, 12, 9, 54, 44, 987876576, time.UTC)},
			{Name: "etzba/pggo", Tag: "latest", Sha: "sha256:d135a04a2ac74466c7c01747daee7d4efae7d09e457b0e8d54bc510f2be1408a", Updated: time.Date(1999, time.October, 21, 1, 32, 33, 654332, time.UTC)},
		},
	}

	subject.Register(registry)
	subject.Update(imageName, imageTag, imageSha)

	fmt.Println("Images in registry\n##################")
	for i, img := range registry.Images {
		fmt.Println(i+1, ".", img)
	}
}
