package vagrant

import (
	"github.com/mitchellh/packer/packer"
	"github.com/mitchellh/packer/post-processor/vagrant"
)

type PostProcessor struct {
	vagrant.PostProcessor
}

func (p *PostProcessor) PostProcess(ui packer.Ui, artifact packer.Artifact) (packer.Artifact, bool, error) {
	if artifact.BuilderId() != "mindjiver.cloudstack" {
		return p.PostProcessor.PostProcess(ui, artifact)
	}

	return p.PostProcessor.PostProcessProvider("cloudstack", new(CloudStackProvider), ui, artifact)
}
