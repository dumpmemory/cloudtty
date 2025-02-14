package config

import (
	cloudshellv1alpha1 "github.com/cloudtty/cloudtty/pkg/apis/cloudshell/v1alpha1"
	"github.com/cloudtty/cloudtty/pkg/generated/clientset/versioned"
	clientset "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	componentbaseconfig "k8s.io/component-base/config"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Config struct {
	KubeClient       *clientset.Clientset
	CloudShellClient *versioned.Clientset
	Client           client.Client
	Kubeconfig       *rest.Config
	EventRecorder    record.EventRecorder
	CoreWorkerLimit  int
	MaxWorkerLimit   int
	CloudShellImage  string
	NodeSelector     map[string]string
	Resources        *cloudshellv1alpha1.ResourceSetting

	LeaderElection componentbaseconfig.LeaderElectionConfiguration
}

type completedConfig struct {
	*Config
}

// CompletedConfig same as Config, just to swap private object.
type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (c *Config) Complete() *CompletedConfig {
	cc := completedConfig{c}

	// TODO:

	return &CompletedConfig{&cc}
}
