// +build !ignore_autogenerated

// Code generated by operator-sdk-v0.17.0. DO NOT EDIT.

package v1alpha1

import (
	status "github.com/operator-framework/operator-sdk/pkg/status"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureProvider) DeepCopyInto(out *AzureProvider) {
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
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	return
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
	return
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
	return
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
	return
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
	return
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
	return
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
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastSyncSuccessTime != nil {
		in, out := &in.LastSyncSuccessTime, &out.LastSyncSuccessTime
		*out = (*in).DeepCopy()
	}
	return
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
	return
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
func (in *Provider) DeepCopyInto(out *Provider) {
	*out = *in
	if in.ProviderType != nil {
		in, out := &in.ProviderType, &out.ProviderType
		*out = new(ProviderType)
		(*in).DeepCopyInto(*out)
	}
	return
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
	if in.Keycloak != nil {
		in, out := &in.Keycloak, &out.Keycloak
		*out = new(KeycloakProvider)
		(*in).DeepCopyInto(*out)
	}
	return
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
	return
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
