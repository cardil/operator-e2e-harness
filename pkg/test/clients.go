package test

import (
	olmversioned "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned"
	"k8s.io/client-go/dynamic"
	k8sv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"knative.dev/pkg/test"
	"path"
	"testing"
)

// Clients holds instances of interfaces for making requests
type Clients struct {
	KubeClient *test.KubeClient
	Dynamic    dynamic.Interface
	Config     *rest.Config
	OLM        olmversioned.Interface
	Apps       *k8sv1.AppsV1Client
}

// Context holds objects related to test execution
type Context struct {
	Name        string
	T           *testing.T
	Clients     *Clients
	CleanupList []CleanupFunc
}

// CleanupFunc defines a function that is called when the respective resource
// should be deleted. When creating resources the user should also create a CleanupFunc
// and register with the Context
type CleanupFunc func() error

// NewClients instantiates and returns several clientsets required for making
// request to the cluster specified by the combination of clusterName
// and configPath.
func NewClients(configPath string, clusterName string) (*Clients, error) {
	clients := &Clients{}
	cfg, err := buildClientConfig(configPath, clusterName)
	if err != nil {
		return nil, err
	}

	// We poll, so set our limits high.
	cfg.QPS = 100
	cfg.Burst = 200

	clients.KubeClient, err = test.NewKubeClient(configPath, clusterName)
	if err != nil {
		return nil, err
	}

	clients.Dynamic, err = dynamic.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	clients.Apps, err = k8sv1.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	clients.Config = cfg
	return clients, nil
}

func NewContext(t *testing.T, role string, idx int) *Context {
	setupContextsOnce(t)
	return contextAtIndex(idx, role, t)
}

var contexts []*Context

func buildClientConfig(kubeConfigPath string, clusterName string) (*rest.Config, error) {
	overrides := clientcmd.ConfigOverrides{}
	// Override the cluster name if provided.
	if clusterName != "" {
		overrides.Context.Cluster = clusterName
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath},
		&overrides).ClientConfig()
}

func contextAtIndex(i int, role string, t *testing.T) *Context {
	if len(contexts) < i+1 {
		t.Fatalf("kubeconfig for user with %s role not present", role)
	}
	ctx := contexts[i]
	ctx.Name = role
	return ctx
}

func setupContextsOnce(t *testing.T) {
	if len(contexts) == 0 {
		for _, cfg := range Kubeconfigs {
			clients, err := NewClients(cfg, "")
			if err != nil {
				t.Fatalf("Couldn't initialize clients for config %s: %v", cfg, err)
			}
			ctx := &Context{
				T:       t,
				Clients: clients,
			}
			contexts = append(contexts, ctx)
		}
	}
}
