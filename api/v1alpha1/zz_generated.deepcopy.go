//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	legacyconfigv1 "github.com/openshift/api/legacyconfig/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureProvider) DeepCopyInto(out *AzureProvider) {
	*out = *in
	if in.BaseGroups != nil {
		in, out := &in.BaseGroups, &out.BaseGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(SecretRef)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.UserNameAttributes != nil {
		in, out := &in.UserNameAttributes, &out.UserNameAttributes
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureProvider.
func (in *AzureProvider) DeepCopy() *AzureProvider {
	if in == nil {
		return nil
	}
	out := new(AzureProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubProvider) DeepCopyInto(out *GitHubProvider) {
	*out = *in
	if in.CaSecret != nil {
		in, out := &in.CaSecret, &out.CaSecret
		*out = new(SecretRef)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(SecretRef)
		**out = **in
	}
	if in.Teams != nil {
		in, out := &in.Teams, &out.Teams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.V4URL != nil {
		in, out := &in.V4URL, &out.V4URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubProvider.
func (in *GitHubProvider) DeepCopy() *GitHubProvider {
	if in == nil {
		return nil
	}
	out := new(GitHubProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitLabProvider) DeepCopyInto(out *GitLabProvider) {
	*out = *in
	if in.CaSecret != nil {
		in, out := &in.CaSecret, &out.CaSecret
		*out = new(SecretRef)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(SecretRef)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitLabProvider.
func (in *GitLabProvider) DeepCopy() *GitLabProvider {
	if in == nil {
		return nil
	}
	out := new(GitLabProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupSync) DeepCopyInto(out *GroupSync) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupSync.
func (in *GroupSync) DeepCopy() *GroupSync {
	if in == nil {
		return nil
	}
	out := new(GroupSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupSync) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupSyncList) DeepCopyInto(out *GroupSyncList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GroupSync, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupSyncList.
func (in *GroupSyncList) DeepCopy() *GroupSyncList {
	if in == nil {
		return nil
	}
	out := new(GroupSyncList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupSyncList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupSyncSpec) DeepCopyInto(out *GroupSyncSpec) {
	*out = *in
	if in.Providers != nil {
		in, out := &in.Providers, &out.Providers
		*out = make([]Provider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupSyncSpec.
func (in *GroupSyncSpec) DeepCopy() *GroupSyncSpec {
	if in == nil {
		return nil
	}
	out := new(GroupSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupSyncStatus) DeepCopyInto(out *GroupSyncStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastSyncSuccessTime != nil {
		in, out := &in.LastSyncSuccessTime, &out.LastSyncSuccessTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupSyncStatus.
func (in *GroupSyncStatus) DeepCopy() *GroupSyncStatus {
	if in == nil {
		return nil
	}
	out := new(GroupSyncStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakProvider) DeepCopyInto(out *KeycloakProvider) {
	*out = *in
	if in.CaSecret != nil {
		in, out := &in.CaSecret, &out.CaSecret
		*out = new(SecretRef)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(SecretRef)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakProvider.
func (in *KeycloakProvider) DeepCopy() *KeycloakProvider {
	if in == nil {
		return nil
	}
	out := new(KeycloakProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LdapProvider) DeepCopyInto(out *LdapProvider) {
	*out = *in
	if in.CaSecret != nil {
		in, out := &in.CaSecret, &out.CaSecret
		*out = new(SecretRef)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(SecretRef)
		**out = **in
	}
	if in.LDAPGroupUIDToOpenShiftGroupNameMapping != nil {
		in, out := &in.LDAPGroupUIDToOpenShiftGroupNameMapping, &out.LDAPGroupUIDToOpenShiftGroupNameMapping
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RFC2307Config != nil {
		in, out := &in.RFC2307Config, &out.RFC2307Config
		*out = new(legacyconfigv1.RFC2307Config)
		(*in).DeepCopyInto(*out)
	}
	if in.ActiveDirectoryConfig != nil {
		in, out := &in.ActiveDirectoryConfig, &out.ActiveDirectoryConfig
		*out = new(legacyconfigv1.ActiveDirectoryConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AugmentedActiveDirectoryConfig != nil {
		in, out := &in.AugmentedActiveDirectoryConfig, &out.AugmentedActiveDirectoryConfig
		*out = new(legacyconfigv1.AugmentedActiveDirectoryConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Whitelist != nil {
		in, out := &in.Whitelist, &out.Whitelist
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.Blacklist != nil {
		in, out := &in.Blacklist, &out.Blacklist
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LdapProvider.
func (in *LdapProvider) DeepCopy() *LdapProvider {
	if in == nil {
		return nil
	}
	out := new(LdapProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OktaProvider) DeepCopyInto(out *OktaProvider) {
	*out = *in
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(SecretRef)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OktaProvider.
func (in *OktaProvider) DeepCopy() *OktaProvider {
	if in == nil {
		return nil
	}
	out := new(OktaProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Provider) DeepCopyInto(out *Provider) {
	*out = *in
	if in.ProviderType != nil {
		in, out := &in.ProviderType, &out.ProviderType
		*out = new(ProviderType)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Provider.
func (in *Provider) DeepCopy() *Provider {
	if in == nil {
		return nil
	}
	out := new(Provider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderType) DeepCopyInto(out *ProviderType) {
	*out = *in
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzureProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.GitHub != nil {
		in, out := &in.GitHub, &out.GitHub
		*out = new(GitHubProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.GitLab != nil {
		in, out := &in.GitLab, &out.GitLab
		*out = new(GitLabProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.Ldap != nil {
		in, out := &in.Ldap, &out.Ldap
		*out = new(LdapProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.Keycloak != nil {
		in, out := &in.Keycloak, &out.Keycloak
		*out = new(KeycloakProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.Okta != nil {
		in, out := &in.Okta, &out.Okta
		*out = new(OktaProvider)
		(*in).DeepCopyInto(*out)
	}
	if in.Static != nil {
		in, out := &in.Static, &out.Static
		*out = new(StaticProvider)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderType.
func (in *ProviderType) DeepCopy() *ProviderType {
	if in == nil {
		return nil
	}
	out := new(ProviderType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticProvider) DeepCopyInto(out *StaticProvider) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]StaticProviderGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticProvider.
func (in *StaticProvider) DeepCopy() *StaticProvider {
	if in == nil {
		return nil
	}
	out := new(StaticProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticProviderGroup) DeepCopyInto(out *StaticProviderGroup) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticProviderGroup.
func (in *StaticProviderGroup) DeepCopy() *StaticProviderGroup {
	if in == nil {
		return nil
	}
	out := new(StaticProviderGroup)
	in.DeepCopyInto(out)
	return out
}
