package syncer

import (
	userv1 "github.com/openshift/api/user/v1"
	redhatcopv1alpha1 "github.com/redhat-cop/group-sync-operator/api/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StaticSyncer struct {
	Name string

	Provider *redhatcopv1alpha1.StaticProvider
}

func (g *StaticSyncer) Init() bool {
	return false
}

func (g *StaticSyncer) Validate() error {
	return nil
}

func (g *StaticSyncer) Bind() error {
	return nil
}

func (g *StaticSyncer) Sync() ([]userv1.Group, error) {
	ocpGroups := make([]userv1.Group, 0, len(g.Provider.Groups))

	for _, group := range g.Provider.Groups {
		ocpGroups = append(ocpGroups, userv1.Group{
			TypeMeta: v1.TypeMeta{
				Kind:       "Group",
				APIVersion: userv1.GroupVersion.String(),
			},
			ObjectMeta: v1.ObjectMeta{
				Name:   group.Name,
				Labels: map[string]string{},
			},
			Users: group.Users,
		})
	}

	return ocpGroups, nil
}

func (g *StaticSyncer) GetProviderName() string {
	return g.Name
}
