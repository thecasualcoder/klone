package kubeconfig

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// ConfigFor returns kube config with given context set as CurrentContext
func ConfigFor(context, kubeCfgFile string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeCfgFile},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}
