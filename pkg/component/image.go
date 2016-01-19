package component

import (
	"rsprd.com/spread/pkg/deploy"
	"rsprd.com/spread/pkg/image"
)

// Image represents a Docker image in the Redspread hierarchy. It wraps image.Image.
type Image struct {
	ComponentBase
	image *image.Image
}

func NewImage(image *image.Image, source string, objects ...deploy.KubeObject) (*Image, error) {
	base, err := newComponentBase(ComponentImage, source, objects)
	if err != nil {
		return nil, err
	} else {
		return &Image{ComponentBase: base, image: image}, nil
	}
}

func (c Image) Deployment() deploy.Deployment {
	return deploy.Deployment{}
}

func (c Image) Images() []*image.Image {
	return []*image.Image{
		c.image,
	}
}

// Kubernetes representation of image
func (c Image) kube() string {
	return c.image.DockerName()
}
