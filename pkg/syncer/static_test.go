package syncer_test

import (
	"testing"

	userv1 "github.com/openshift/api/user/v1"
	"github.com/stretchr/testify/require"

	"github.com/redhat-cop/group-sync-operator/api/v1alpha1"
	"github.com/redhat-cop/group-sync-operator/pkg/syncer"
)

func Test_StaticSyncer_Sync(t *testing.T) {
	groups := []v1alpha1.StaticProviderGroup{
		{Name: "group-1", Users: []string{"xerxes.xoe"}},
		{Name: "group-2", Users: []string{"mermes.moe", "dave.devel"}},
	}
	subject := &syncer.StaticSyncer{
		Provider: &v1alpha1.StaticProvider{
			Groups: groups,
		},
	}

	synced, err := subject.Sync()
	require.NoError(t, err)
	require.Len(t, synced, len(groups))
	for i := range groups {
		require.Equal(t, groups[i].Name, synced[i].Name)
		require.Equal(t, userv1.OptionalNames(groups[i].Users), synced[i].Users)
	}
}
