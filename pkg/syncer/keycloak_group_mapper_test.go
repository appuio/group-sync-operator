package syncer_test

import (
	"sort"
	"testing"

	"github.com/Nerzal/gocloak/v5"
	"github.com/google/uuid"
	userv1 "github.com/openshift/api/user/v1"
	"github.com/stretchr/testify/require"

	redhatcopv1alpha1 "github.com/redhat-cop/group-sync-operator/api/v1alpha1"
	. "github.com/redhat-cop/group-sync-operator/pkg/syncer"
)

func Test_KeycloakGroupMapper_Map_ScopeOne(t *testing.T) {
	groups, users := groupHierarchy()

	mapper := KeycloakGroupMapper{
		GetGroupMembers: getGroupMembersFunc(users),

		Scope: redhatcopv1alpha1.OneSyncScope,
	}

	mapped, err := mapper.Map(groups)
	require.NoError(t, err)
	require.ElementsMatch(t, groupMembers(mapped), []groupMember{
		{"top-level-1", []string{"top-level-1.user"}},
		{"top-level-2", []string{"top-level-2.user"}},
	})
}

func Test_KeycloakGroupMapper_Map_ScopeSub(t *testing.T) {
	groups, users := groupHierarchy()

	mapper := KeycloakGroupMapper{
		GetGroupMembers: getGroupMembersFunc(users),

		Scope: redhatcopv1alpha1.SubSyncScope,
	}

	mapped, err := mapper.Map(groups)
	require.NoError(t, err)
	// Note how the "child-1" group exists twice.
	// Members are only passed up one level.
	require.ElementsMatch(t, groupMembers(mapped), []groupMember{
		{"top-level-1", []string{"top-level-1.child-1.user", "top-level-1.user"}},
		{"child-1", []string{"top-level-1.child-1.deeply-nested-1.user", "top-level-1.child-1.user"}},
		{"deeply-nested-1", []string{"top-level-1.child-1.deeply-nested-1.user"}},
		{"top-level-2", []string{"top-level-2.child-1.user", "top-level-2.child-2.user", "top-level-2.user"}},
		{"child-1", []string{"top-level-2.child-1.user"}},
		{"child-2", []string{"top-level-2.child-2.user"}},
	})
}

func Test_KeycloakGroupMapper_Map_ScopeSub_JoinSubGroupProcessing(t *testing.T) {
	groups, users := groupHierarchy()

	mapper := KeycloakGroupMapper{
		GetGroupMembers: getGroupMembersFunc(users),

		Scope:                 redhatcopv1alpha1.SubSyncScope,
		SubGroupProcessing:    redhatcopv1alpha1.JoinSubGroupProcessing,
		SubGroupJoinSeparator: ":",
	}

	mapped, err := mapper.Map(groups)
	require.NoError(t, err)
	require.ElementsMatch(t, groupMembers(mapped), []groupMember{
		{"top-level-1", []string{"top-level-1.user"}},
		{"top-level-1:child-1", []string{"top-level-1.child-1.user"}},
		{"top-level-1:child-1:deeply-nested-1", []string{"top-level-1.child-1.deeply-nested-1.user"}},
		{"top-level-2", []string{"top-level-2.user"}},
		{"top-level-2:child-1", []string{"top-level-2.child-1.user"}},
		{"top-level-2:child-2", []string{"top-level-2.child-2.user"}},
	})
}

func Test_KeycloakGroupMapper_Map_ScopeSub_JoinSubGroupProcessing_FilterAndStripRoot(t *testing.T) {
	groups, _ := groupHierarchy()

	mapper := KeycloakGroupMapper{
		GetGroupMembers: noopGetGroupMembersFunc,

		SubGroupJoinStripRootGroups: []string{"top-level-1"},
		AllowedGroups:               []string{"top-level-1"},
		Scope:                       redhatcopv1alpha1.SubSyncScope,
		SubGroupProcessing:          redhatcopv1alpha1.JoinSubGroupProcessing,
		SubGroupJoinSeparator:       ":",
	}

	mapped, err := mapper.Map(groups)
	require.NoError(t, err)
	require.ElementsMatch(t, groupNames(mapped), []string{
		"child-1",
		"child-1:deeply-nested-1",
	})
}

func Test_KeycloakGroupMapper_Map_ScopeSub_JoinSubGroupProcessing_StripRoot_NotStripSub(t *testing.T) {
	// Sub groups with the same name should not be stripped
	groups := []*gocloak.Group{
		newKeycloakGroup("top",
			newKeycloakGroup("my-org",
				newKeycloakGroup("top",
					newKeycloakGroup("team")))),
	}

	mapper := KeycloakGroupMapper{
		GetGroupMembers: noopGetGroupMembersFunc,

		SubGroupJoinStripRootGroups: []string{"top"},
		Scope:                       redhatcopv1alpha1.SubSyncScope,
		SubGroupProcessing:          redhatcopv1alpha1.JoinSubGroupProcessing,
		SubGroupJoinSeparator:       ":",
	}

	mapped, err := mapper.Map(groups)
	require.NoError(t, err)
	require.ElementsMatch(t, groupNames(mapped), []string{
		"my-org",
		"my-org:top",
		"my-org:top:team",
	})
}

func Test_KeycloakGroupMapper_Map_ScopeSub_JoinSubGroupProcessing_IgnoreGroupsWithSeparator(t *testing.T) {
	groups := []*gocloak.Group{
		newKeycloakGroup("top-level:1"),
		newKeycloakGroup("top-level-2"),
	}

	mapper := KeycloakGroupMapper{
		GetGroupMembers: noopGetGroupMembersFunc,

		Scope:                 redhatcopv1alpha1.SubSyncScope,
		SubGroupProcessing:    redhatcopv1alpha1.JoinSubGroupProcessing,
		SubGroupJoinSeparator: ":",
	}

	mapped, err := mapper.Map(groups)
	require.NoError(t, err)
	require.ElementsMatch(t, groupNames(mapped), []string{
		"top-level-2",
	})
}

func noopGetGroupMembersFunc(string) ([]*gocloak.User, error) {
	return nil, nil
}

func getGroupMembersFunc(users map[string][]*gocloak.User) func(groupID string) ([]*gocloak.User, error) {
	return func(groupID string) ([]*gocloak.User, error) {
		return users[groupID], nil
	}
}

func newKeycloakGroup(name string, sub ...*gocloak.Group) *gocloak.Group {
	id := uuid.New().String()
	return &gocloak.Group{ID: &id, Name: &name, SubGroups: sub}
}

func newKeycloakUser(name string) *gocloak.User {
	id := uuid.New().String()
	return &gocloak.User{ID: &id, Username: &name}
}

func groupNames(groups []userv1.Group) []string {
	out := make([]string, len(groups))
	for i, group := range groups {
		out[i] = group.Name
	}
	return out
}

type groupMember struct {
	GroupName string
	Usernames []string
}

func groupMembers(groups []userv1.Group) []groupMember {
	out := make([]groupMember, len(groups))
	for i, group := range groups {
		users := make([]string, len(group.Users))
		copy(users, group.Users)
		sort.Strings(users)

		out[i].GroupName = group.Name
		out[i].Usernames = users
	}
	return out
}

func groupHierarchy() ([]*gocloak.Group, map[string][]*gocloak.User) {
	h := []*gocloak.Group{
		newKeycloakGroup("top-level-1",
			newKeycloakGroup("child-1",
				newKeycloakGroup("deeply-nested-1"))),
		newKeycloakGroup("top-level-2",
			newKeycloakGroup("child-1"), newKeycloakGroup("child-2")),
	}
	u := map[string][]*gocloak.User{
		*h[0].ID: {newKeycloakUser("top-level-1.user")},
		*h[1].ID: {newKeycloakUser("top-level-2.user")},

		*h[0].SubGroups[0].ID:              {newKeycloakUser("top-level-1.child-1.user")},
		*h[0].SubGroups[0].SubGroups[0].ID: {newKeycloakUser("top-level-1.child-1.deeply-nested-1.user")},

		*h[1].SubGroups[0].ID: {newKeycloakUser("top-level-2.child-1.user")},
		*h[1].SubGroups[1].ID: {newKeycloakUser("top-level-2.child-2.user")},
	}
	return h, u
}
