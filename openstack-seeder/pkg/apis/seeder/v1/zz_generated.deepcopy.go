// +build !ignore_autogenerated

/*
Copyright 2017 SAP SE

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressScopeSpec) DeepCopyInto(out *AddressScopeSpec) {
	*out = *in
	if in.Shared != nil {
		in, out := &in.Shared, &out.Shared
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.SubnetPools != nil {
		in, out := &in.SubnetPools, &out.SubnetPools
		*out = make([]SubnetPoolSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressScopeSpec.
func (in *AddressScopeSpec) DeepCopy() *AddressScopeSpec {
	if in == nil {
		return nil
	}
	out := new(AddressScopeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSQuotaSpec) DeepCopyInto(out *DNSQuotaSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSQuotaSpec.
func (in *DNSQuotaSpec) DeepCopy() *DNSQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(DNSQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSRecordsetSpec) DeepCopyInto(out *DNSRecordsetSpec) {
	*out = *in
	if in.Records != nil {
		in, out := &in.Records, &out.Records
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSRecordsetSpec.
func (in *DNSRecordsetSpec) DeepCopy() *DNSRecordsetSpec {
	if in == nil {
		return nil
	}
	out := new(DNSRecordsetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSTSIGKeySpec) DeepCopyInto(out *DNSTSIGKeySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSTSIGKeySpec.
func (in *DNSTSIGKeySpec) DeepCopy() *DNSTSIGKeySpec {
	if in == nil {
		return nil
	}
	out := new(DNSTSIGKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSZoneSpec) DeepCopyInto(out *DNSZoneSpec) {
	*out = *in
	if in.DNSRecordsets != nil {
		in, out := &in.DNSRecordsets, &out.DNSRecordsets
		*out = make([]DNSRecordsetSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSZoneSpec.
func (in *DNSZoneSpec) DeepCopy() *DNSZoneSpec {
	if in == nil {
		return nil
	}
	out := new(DNSZoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainConfigSpec) DeepCopyInto(out *DomainConfigSpec) {
	*out = *in
	if in.IdentityConfig != nil {
		in, out := &in.IdentityConfig, &out.IdentityConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(IdentityConfigSpec)
			**out = **in
		}
	}
	if in.LdapConfig != nil {
		in, out := &in.LdapConfig, &out.LdapConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(LdapConfigSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainConfigSpec.
func (in *DomainConfigSpec) DeepCopy() *DomainConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DomainConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpec) DeepCopyInto(out *DomainSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]UserSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]GroupSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Projects != nil {
		in, out := &in.Projects, &out.Projects
		*out = make([]ProjectSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]RoleSpec, len(*in))
		copy(*out, *in)
	}
	if in.RoleAssignments != nil {
		in, out := &in.RoleAssignments, &out.RoleAssignments
		*out = make([]RoleAssignmentSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		if *in == nil {
			*out = nil
		} else {
			*out = new(DomainConfigSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpec.
func (in *DomainSpec) DeepCopy() *DomainSpec {
	if in == nil {
		return nil
	}
	out := new(DomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ec2CredSpec) DeepCopyInto(out *Ec2CredSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ec2CredSpec.
func (in *Ec2CredSpec) DeepCopy() *Ec2CredSpec {
	if in == nil {
		return nil
	}
	out := new(Ec2CredSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpec) DeepCopyInto(out *EndpointSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpec.
func (in *EndpointSpec) DeepCopy() *EndpointSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalFixedIPsSpec) DeepCopyInto(out *ExternalFixedIPsSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalFixedIPsSpec.
func (in *ExternalFixedIPsSpec) DeepCopy() *ExternalFixedIPsSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalFixedIPsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalGatewayInfoSpec) DeepCopyInto(out *ExternalGatewayInfoSpec) {
	*out = *in
	if in.EnableSNAT != nil {
		in, out := &in.EnableSNAT, &out.EnableSNAT
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.ExternalFixedIPs != nil {
		in, out := &in.ExternalFixedIPs, &out.ExternalFixedIPs
		*out = make([]ExternalFixedIPsSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalGatewayInfoSpec.
func (in *ExternalGatewayInfoSpec) DeepCopy() *ExternalGatewayInfoSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalGatewayInfoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlavorSpec) DeepCopyInto(out *FlavorSpec) {
	*out = *in
	if in.IsPublic != nil {
		in, out := &in.IsPublic, &out.IsPublic
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.ExtraSpecs != nil {
		in, out := &in.ExtraSpecs, &out.ExtraSpecs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlavorSpec.
func (in *FlavorSpec) DeepCopy() *FlavorSpec {
	if in == nil {
		return nil
	}
	out := new(FlavorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupSpec) DeepCopyInto(out *GroupSpec) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RoleAssignments != nil {
		in, out := &in.RoleAssignments, &out.RoleAssignments
		*out = make([]RoleAssignmentSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupSpec.
func (in *GroupSpec) DeepCopy() *GroupSpec {
	if in == nil {
		return nil
	}
	out := new(GroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityConfigSpec) DeepCopyInto(out *IdentityConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityConfigSpec.
func (in *IdentityConfigSpec) DeepCopy() *IdentityConfigSpec {
	if in == nil {
		return nil
	}
	out := new(IdentityConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LdapConfigSpec) DeepCopyInto(out *LdapConfigSpec) {
	*out = *in
	if in.UseTLS != nil {
		in, out := &in.UseTLS, &out.UseTLS
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.UsePool != nil {
		in, out := &in.UsePool, &out.UsePool
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.UseAuthPool != nil {
		in, out := &in.UseAuthPool, &out.UseAuthPool
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.UserAllowCreate != nil {
		in, out := &in.UserAllowCreate, &out.UserAllowCreate
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.UserAllowUpdate != nil {
		in, out := &in.UserAllowUpdate, &out.UserAllowUpdate
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.UserAllowDelete != nil {
		in, out := &in.UserAllowDelete, &out.UserAllowDelete
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.GroupMembersAreIDs != nil {
		in, out := &in.GroupMembersAreIDs, &out.GroupMembersAreIDs
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.GroupAllowCreate != nil {
		in, out := &in.GroupAllowCreate, &out.GroupAllowCreate
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.GroupAllowUpdate != nil {
		in, out := &in.GroupAllowUpdate, &out.GroupAllowUpdate
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.GroupAllowDelete != nil {
		in, out := &in.GroupAllowDelete, &out.GroupAllowDelete
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LdapConfigSpec.
func (in *LdapConfigSpec) DeepCopy() *LdapConfigSpec {
	if in == nil {
		return nil
	}
	out := new(LdapConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkQuotaSpec) DeepCopyInto(out *NetworkQuotaSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkQuotaSpec.
func (in *NetworkQuotaSpec) DeepCopy() *NetworkQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	if in.AdminStateUp != nil {
		in, out := &in.AdminStateUp, &out.AdminStateUp
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.PortSecurityEnabled != nil {
		in, out := &in.PortSecurityEnabled, &out.PortSecurityEnabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.RouterExternal != nil {
		in, out := &in.RouterExternal, &out.RouterExternal
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.Shared != nil {
		in, out := &in.Shared, &out.Shared
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.VlanTransparent != nil {
		in, out := &in.VlanTransparent, &out.VlanTransparent
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]SubnetSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackSeed) DeepCopyInto(out *OpenstackSeed) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackSeed.
func (in *OpenstackSeed) DeepCopy() *OpenstackSeed {
	if in == nil {
		return nil
	}
	out := new(OpenstackSeed)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenstackSeed) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackSeedList) DeepCopyInto(out *OpenstackSeedList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenstackSeed, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackSeedList.
func (in *OpenstackSeedList) DeepCopy() *OpenstackSeedList {
	if in == nil {
		return nil
	}
	out := new(OpenstackSeedList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenstackSeedList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackSeedSpec) DeepCopyInto(out *OpenstackSeedSpec) {
	*out = *in
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]RoleSpec, len(*in))
		copy(*out, *in)
	}
	if in.RoleInferences != nil {
		in, out := &in.RoleInferences, &out.RoleInferences
		*out = make([]RoleInferenceSpec, len(*in))
		copy(*out, *in)
	}
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]RegionSpec, len(*in))
		copy(*out, *in)
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]ServiceSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Flavors != nil {
		in, out := &in.Flavors, &out.Flavors
		*out = make([]FlavorSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ShareTypes != nil {
		in, out := &in.ShareTypes, &out.ShareTypes
		*out = make([]ShareTypeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceClasses != nil {
		in, out := &in.ResourceClasses, &out.ResourceClasses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]DomainSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RBACPolicies != nil {
		in, out := &in.RBACPolicies, &out.RBACPolicies
		*out = make([]RBACPolicySpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackSeedSpec.
func (in *OpenstackSeedSpec) DeepCopy() *OpenstackSeedSpec {
	if in == nil {
		return nil
	}
	out := new(OpenstackSeedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackSeedStatus) DeepCopyInto(out *OpenstackSeedStatus) {
	*out = *in
	if in.VisitedDependencies != nil {
		in, out := &in.VisitedDependencies, &out.VisitedDependencies
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackSeedStatus.
func (in *OpenstackSeedStatus) DeepCopy() *OpenstackSeedStatus {
	if in == nil {
		return nil
	}
	out := new(OpenstackSeedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectEndpointSpec) DeepCopyInto(out *ProjectEndpointSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectEndpointSpec.
func (in *ProjectEndpointSpec) DeepCopy() *ProjectEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSpec) DeepCopyInto(out *ProjectSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.IsDomain != nil {
		in, out := &in.IsDomain, &out.IsDomain
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]ProjectEndpointSpec, len(*in))
		copy(*out, *in)
	}
	if in.RoleAssignments != nil {
		in, out := &in.RoleAssignments, &out.RoleAssignments
		*out = make([]RoleAssignmentSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Flavors != nil {
		in, out := &in.Flavors, &out.Flavors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ShareTypes != nil {
		in, out := &in.ShareTypes, &out.ShareTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AddressScopes != nil {
		in, out := &in.AddressScopes, &out.AddressScopes
		*out = make([]AddressScopeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetPools != nil {
		in, out := &in.SubnetPools, &out.SubnetPools
		*out = make([]SubnetPoolSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkQuota != nil {
		in, out := &in.NetworkQuota, &out.NetworkQuota
		if *in == nil {
			*out = nil
		} else {
			*out = new(NetworkQuotaSpec)
			**out = **in
		}
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]NetworkSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Routers != nil {
		in, out := &in.Routers, &out.Routers
		*out = make([]RouterSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Swift != nil {
		in, out := &in.Swift, &out.Swift
		if *in == nil {
			*out = nil
		} else {
			*out = new(SwiftAccountSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.DNSQuota != nil {
		in, out := &in.DNSQuota, &out.DNSQuota
		if *in == nil {
			*out = nil
		} else {
			*out = new(DNSQuotaSpec)
			**out = **in
		}
	}
	if in.DNSZones != nil {
		in, out := &in.DNSZones, &out.DNSZones
		*out = make([]DNSZoneSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DNSTSIGKeys != nil {
		in, out := &in.DNSTSIGKeys, &out.DNSTSIGKeys
		*out = make([]DNSTSIGKeySpec, len(*in))
		copy(*out, *in)
	}
	if in.Ec2Creds != nil {
		in, out := &in.Ec2Creds, &out.Ec2Creds
		*out = make([]Ec2CredSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSpec.
func (in *ProjectSpec) DeepCopy() *ProjectSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RBACPolicySpec) DeepCopyInto(out *RBACPolicySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RBACPolicySpec.
func (in *RBACPolicySpec) DeepCopy() *RBACPolicySpec {
	if in == nil {
		return nil
	}
	out := new(RBACPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionSpec) DeepCopyInto(out *RegionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionSpec.
func (in *RegionSpec) DeepCopy() *RegionSpec {
	if in == nil {
		return nil
	}
	out := new(RegionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentSpec) DeepCopyInto(out *RoleAssignmentSpec) {
	*out = *in
	if in.Inherited != nil {
		in, out := &in.Inherited, &out.Inherited
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentSpec.
func (in *RoleAssignmentSpec) DeepCopy() *RoleAssignmentSpec {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleInferenceSpec) DeepCopyInto(out *RoleInferenceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleInferenceSpec.
func (in *RoleInferenceSpec) DeepCopy() *RoleInferenceSpec {
	if in == nil {
		return nil
	}
	out := new(RoleInferenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleSpec) DeepCopyInto(out *RoleSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleSpec.
func (in *RoleSpec) DeepCopy() *RoleSpec {
	if in == nil {
		return nil
	}
	out := new(RoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouterPortSpec) DeepCopyInto(out *RouterPortSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouterPortSpec.
func (in *RouterPortSpec) DeepCopy() *RouterPortSpec {
	if in == nil {
		return nil
	}
	out := new(RouterPortSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouterRouteSpec) DeepCopyInto(out *RouterRouteSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouterRouteSpec.
func (in *RouterRouteSpec) DeepCopy() *RouterRouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouterRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouterSpec) DeepCopyInto(out *RouterSpec) {
	*out = *in
	if in.AdminStateUp != nil {
		in, out := &in.AdminStateUp, &out.AdminStateUp
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.ExternalGatewayInfo != nil {
		in, out := &in.ExternalGatewayInfo, &out.ExternalGatewayInfo
		if *in == nil {
			*out = nil
		} else {
			*out = new(ExternalGatewayInfoSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Distributed != nil {
		in, out := &in.Distributed, &out.Distributed
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.HA != nil {
		in, out := &in.HA, &out.HA
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.RouterPorts != nil {
		in, out := &in.RouterPorts, &out.RouterPorts
		*out = make([]RouterPortSpec, len(*in))
		copy(*out, *in)
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]RouterRouteSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouterSpec.
func (in *RouterSpec) DeepCopy() *RouterSpec {
	if in == nil {
		return nil
	}
	out := new(RouterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]EndpointSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareTypeSpec) DeepCopyInto(out *ShareTypeSpec) {
	*out = *in
	if in.IsPublic != nil {
		in, out := &in.IsPublic, &out.IsPublic
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.Specs != nil {
		in, out := &in.Specs, &out.Specs
		if *in == nil {
			*out = nil
		} else {
			*out = new(ShareTypeSpecifiedSpecs)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ExtraSpecs != nil {
		in, out := &in.ExtraSpecs, &out.ExtraSpecs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareTypeSpec.
func (in *ShareTypeSpec) DeepCopy() *ShareTypeSpec {
	if in == nil {
		return nil
	}
	out := new(ShareTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShareTypeSpecifiedSpecs) DeepCopyInto(out *ShareTypeSpecifiedSpecs) {
	*out = *in
	if in.DHSS != nil {
		in, out := &in.DHSS, &out.DHSS
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.SnapshotSupport != nil {
		in, out := &in.SnapshotSupport, &out.SnapshotSupport
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShareTypeSpecifiedSpecs.
func (in *ShareTypeSpecifiedSpecs) DeepCopy() *ShareTypeSpecifiedSpecs {
	if in == nil {
		return nil
	}
	out := new(ShareTypeSpecifiedSpecs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetPoolSpec) DeepCopyInto(out *SubnetPoolSpec) {
	*out = *in
	if in.Prefixes != nil {
		in, out := &in.Prefixes, &out.Prefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Shared != nil {
		in, out := &in.Shared, &out.Shared
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.IsDefault != nil {
		in, out := &in.IsDefault, &out.IsDefault
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetPoolSpec.
func (in *SubnetPoolSpec) DeepCopy() *SubnetPoolSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
	if in.EnableDHCP != nil {
		in, out := &in.EnableDHCP, &out.EnableDHCP
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.DNSNameServers != nil {
		in, out := &in.DNSNameServers, &out.DNSNameServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllocationPools != nil {
		in, out := &in.AllocationPools, &out.AllocationPools
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HostRoutes != nil {
		in, out := &in.HostRoutes, &out.HostRoutes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Prefixlen != nil {
		in, out := &in.Prefixlen, &out.Prefixlen
		if *in == nil {
			*out = nil
		} else {
			*out = new(int)
			**out = **in
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftAccountSpec) DeepCopyInto(out *SwiftAccountSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]SwiftContainerSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftAccountSpec.
func (in *SwiftAccountSpec) DeepCopy() *SwiftAccountSpec {
	if in == nil {
		return nil
	}
	out := new(SwiftAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftContainerSpec) DeepCopyInto(out *SwiftContainerSpec) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftContainerSpec.
func (in *SwiftContainerSpec) DeepCopy() *SwiftContainerSpec {
	if in == nil {
		return nil
	}
	out := new(SwiftContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpec) DeepCopyInto(out *UserSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.RoleAssignments != nil {
		in, out := &in.RoleAssignments, &out.RoleAssignments
		*out = make([]RoleAssignmentSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpec.
func (in *UserSpec) DeepCopy() *UserSpec {
	if in == nil {
		return nil
	}
	out := new(UserSpec)
	in.DeepCopyInto(out)
	return out
}
