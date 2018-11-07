package secrets

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	// load auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
)

func clean(secret *corev1.Secret) {
	secret.ObjectMeta.ResourceVersion = ""
}

// Get returns K8s Secret by name
func Get(name, namespace string, config *rest.Config) (*corev1.Secret, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	secret, err := clientset.CoreV1().Secrets(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return secret, nil
}

// Create creates a K8s Secret
func Create(secret *corev1.Secret, namespace string, config *rest.Config) (*corev1.Secret, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	clean(secret)

	s, err := clientset.CoreV1().Secrets(namespace).Create(secret)
	if err != nil {
		return nil, err
	}
	return s, nil
}
