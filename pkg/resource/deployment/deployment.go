package deployment

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	// load auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
)

func clean(deployment *appsv1.Deployment) *appsv1.Deployment {
	deployment.ObjectMeta.ResourceVersion = ""
	return deployment
}

// Get returns K8s Deployment by name
func Get(name, namespace string, config *rest.Config) (*appsv1.Deployment, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	deployment, err := clientset.AppsV1().Deployments(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return deployment, nil
}

// Create creates a K8s Deployment
func Create(deployment *appsv1.Deployment, namespace string, config *rest.Config) (*appsv1.Deployment, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	d, err := clientset.AppsV1().Deployments(namespace).Create(clean(deployment))
	if err != nil {
		return nil, err
	}
	return d, nil
}
