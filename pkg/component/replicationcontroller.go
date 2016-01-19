package component

import (
	"rsprd.com/spread/pkg/deploy"
	"rsprd.com/spread/pkg/image"

	"k8s.io/kubernetes/pkg/api"
)

// ReplicationController represents api.ReplicationController in the Redspread hierarchy.
type ReplicationController struct {
	ComponentBase
	rc  *api.ReplicationController
	pod *Pod
}

func NewReplicationController(kubeRC *api.ReplicationController, source string, objects ...deploy.KubeObject) (*ReplicationController, error) {
	base, err := newComponentBase(ComponentReplicationController, source, objects)
	if err != nil {
		return nil, err
	}

	rc := ReplicationController{ComponentBase: base}
	if kubeRC.Spec.Template != nil {
		rc.pod, err = NewPodFromPodSpec(kubeRC.Spec.Template.Spec, source)
		if err != nil {
			return nil, err
		}
		kubeRC.Spec.Template = nil
	}
	return &rc, nil
}

func (c ReplicationController) Deployment() deploy.Deployment {
	return deploy.Deployment{}
}

func (c ReplicationController) Images() (images []*image.Image) {
	return c.pod.Images()
}

func (c ReplicationController) kube() *api.ReplicationController {
	c.rc.Spec.Template.Spec = c.pod.kube().Spec
	return c.rc
}
