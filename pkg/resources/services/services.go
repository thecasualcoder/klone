package services

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	// load auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
)

func clean(service *corev1.Service) {
	service.ObjectMeta.ResourceVersion = ""
	service.Spec.ClusterIP = ""
}

// Get returns K8s Service by name
func Get(name, namespace string, config *rest.Config) (*corev1.Service, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	service, err := clientset.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return service, nil
}

// Create creates a K8s Service
func Create(service *corev1.Service, namespace string, config *rest.Config) (*corev1.Service, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	clean(service)

	s, err := clientset.CoreV1().Services(namespace).Create(service)
	if err != nil {
		return nil, err
	}
	return s, nil
}
