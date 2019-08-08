package api_test

import (
	"github.com/mattermost/mattermost-cloud/internal/tools/k8s"
	"github.com/mattermost/mattermost-cloud/model"
)

type mockSupervisor struct {
}

func (s *mockSupervisor) Do() error {
	return nil
}

type mockProvisioner struct {
}

func (s *mockProvisioner) ExecMattermostCLI(*model.Cluster, *model.ClusterInstallation, ...string) ([]byte, error) {
	return []byte(`{"ServiceSettings":{"SiteURL":"http://test.example.com"}}`), nil
}

func (s *mockProvisioner) GetClusterResources(*model.Cluster) (*k8s.ClusterResources, error) {
	return nil, nil
}

func sToP(s string) *string {
	return &s
}
