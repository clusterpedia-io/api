//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGroupResources) DeepCopyInto(out *ClusterGroupResources) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGroupResources.
func (in *ClusterGroupResources) DeepCopy() *ClusterGroupResources {
	if in == nil {
		return nil
	}
	out := new(ClusterGroupResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGroupResourcesStatus) DeepCopyInto(out *ClusterGroupResourcesStatus) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ClusterResourceStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGroupResourcesStatus.
func (in *ClusterGroupResourcesStatus) DeepCopy() *ClusterGroupResourcesStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterGroupResourcesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceStatus) DeepCopyInto(out *ClusterResourceStatus) {
	*out = *in
	if in.SyncConditions != nil {
		in, out := &in.SyncConditions, &out.SyncConditions
		*out = make([]ClusterResourceSyncCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceStatus.
func (in *ClusterResourceStatus) DeepCopy() *ClusterResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceSyncCondition) DeepCopyInto(out *ClusterResourceSyncCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceSyncCondition.
func (in *ClusterResourceSyncCondition) DeepCopy() *ClusterResourceSyncCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceSyncCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.Kubeconfig != nil {
		in, out := &in.Kubeconfig, &out.Kubeconfig
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.TokenData != nil {
		in, out := &in.TokenData, &out.TokenData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CAData != nil {
		in, out := &in.CAData, &out.CAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CertData != nil {
		in, out := &in.CertData, &out.CertData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.KeyData != nil {
		in, out := &in.KeyData, &out.KeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.SyncResources != nil {
		in, out := &in.SyncResources, &out.SyncResources
		*out = make([]ClusterGroupResources, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SyncResources != nil {
		in, out := &in.SyncResources, &out.SyncResources
		*out = make([]ClusterGroupResourcesStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ShardingName != nil {
		in, out := &in.ShardingName, &out.ShardingName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSyncResources) DeepCopyInto(out *ClusterSyncResources) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSyncResources.
func (in *ClusterSyncResources) DeepCopy() *ClusterSyncResources {
	if in == nil {
		return nil
	}
	out := new(ClusterSyncResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSyncResources) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSyncResourcesList) DeepCopyInto(out *ClusterSyncResourcesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSyncResources, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSyncResourcesList.
func (in *ClusterSyncResourcesList) DeepCopy() *ClusterSyncResourcesList {
	if in == nil {
		return nil
	}
	out := new(ClusterSyncResourcesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSyncResourcesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSyncResourcesSpec) DeepCopyInto(out *ClusterSyncResourcesSpec) {
	*out = *in
	if in.SyncResources != nil {
		in, out := &in.SyncResources, &out.SyncResources
		*out = make([]ClusterGroupResources, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSyncResourcesSpec.
func (in *ClusterSyncResourcesSpec) DeepCopy() *ClusterSyncResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSyncResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PediaCluster) DeepCopyInto(out *PediaCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PediaCluster.
func (in *PediaCluster) DeepCopy() *PediaCluster {
	if in == nil {
		return nil
	}
	out := new(PediaCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PediaCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PediaClusterList) DeepCopyInto(out *PediaClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PediaCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PediaClusterList.
func (in *PediaClusterList) DeepCopy() *PediaClusterList {
	if in == nil {
		return nil
	}
	out := new(PediaClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PediaClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
