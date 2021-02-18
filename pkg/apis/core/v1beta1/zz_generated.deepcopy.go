// +build !ignore_autogenerated

/*
Copyright 2018 The CDI Authors.

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

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDI) DeepCopyInto(out *CDI) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDI.
func (in *CDI) DeepCopy() *CDI {
	if in == nil {
		return nil
	}
	out := new(CDI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CDI) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDICertConfig) DeepCopyInto(out *CDICertConfig) {
	*out = *in
	if in.CA != nil {
		in, out := &in.CA, &out.CA
		*out = new(CertConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Server != nil {
		in, out := &in.Server, &out.Server
		*out = new(CertConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDICertConfig.
func (in *CDICertConfig) DeepCopy() *CDICertConfig {
	if in == nil {
		return nil
	}
	out := new(CDICertConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDIConfig) DeepCopyInto(out *CDIConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDIConfig.
func (in *CDIConfig) DeepCopy() *CDIConfig {
	if in == nil {
		return nil
	}
	out := new(CDIConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CDIConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDIConfigList) DeepCopyInto(out *CDIConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CDIConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDIConfigList.
func (in *CDIConfigList) DeepCopy() *CDIConfigList {
	if in == nil {
		return nil
	}
	out := new(CDIConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CDIConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDIConfigSpec) DeepCopyInto(out *CDIConfigSpec) {
	*out = *in
	if in.UploadProxyURLOverride != nil {
		in, out := &in.UploadProxyURLOverride, &out.UploadProxyURLOverride
		*out = new(string)
		**out = **in
	}
	if in.ImportProxy != nil {
		in, out := &in.ImportProxy, &out.ImportProxy
		*out = new(ImportProxy)
		(*in).DeepCopyInto(*out)
	}
	if in.ScratchSpaceStorageClass != nil {
		in, out := &in.ScratchSpaceStorageClass, &out.ScratchSpaceStorageClass
		*out = new(string)
		**out = **in
	}
	if in.PodResourceRequirements != nil {
		in, out := &in.PodResourceRequirements, &out.PodResourceRequirements
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FilesystemOverhead != nil {
		in, out := &in.FilesystemOverhead, &out.FilesystemOverhead
		*out = new(FilesystemOverhead)
		(*in).DeepCopyInto(*out)
	}
	if in.Preallocation != nil {
		in, out := &in.Preallocation, &out.Preallocation
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDIConfigSpec.
func (in *CDIConfigSpec) DeepCopy() *CDIConfigSpec {
	if in == nil {
		return nil
	}
	out := new(CDIConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDIConfigStatus) DeepCopyInto(out *CDIConfigStatus) {
	*out = *in
	if in.UploadProxyURL != nil {
		in, out := &in.UploadProxyURL, &out.UploadProxyURL
		*out = new(string)
		**out = **in
	}
	if in.ImportProxy != nil {
		in, out := &in.ImportProxy, &out.ImportProxy
		*out = new(ImportProxy)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultPodResourceRequirements != nil {
		in, out := &in.DefaultPodResourceRequirements, &out.DefaultPodResourceRequirements
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.FilesystemOverhead != nil {
		in, out := &in.FilesystemOverhead, &out.FilesystemOverhead
		*out = new(FilesystemOverhead)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDIConfigStatus.
func (in *CDIConfigStatus) DeepCopy() *CDIConfigStatus {
	if in == nil {
		return nil
	}
	out := new(CDIConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDIList) DeepCopyInto(out *CDIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CDI, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDIList.
func (in *CDIList) DeepCopy() *CDIList {
	if in == nil {
		return nil
	}
	out := new(CDIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CDIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDISpec) DeepCopyInto(out *CDISpec) {
	*out = *in
	if in.UninstallStrategy != nil {
		in, out := &in.UninstallStrategy, &out.UninstallStrategy
		*out = new(CDIUninstallStrategy)
		**out = **in
	}
	in.Infra.DeepCopyInto(&out.Infra)
	in.Workloads.DeepCopyInto(&out.Workloads)
	if in.CloneStrategyOverride != nil {
		in, out := &in.CloneStrategyOverride, &out.CloneStrategyOverride
		*out = new(CDICloneStrategy)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(CDIConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CertConfig != nil {
		in, out := &in.CertConfig, &out.CertConfig
		*out = new(CDICertConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDISpec.
func (in *CDISpec) DeepCopy() *CDISpec {
	if in == nil {
		return nil
	}
	out := new(CDISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDIStatus) DeepCopyInto(out *CDIStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDIStatus.
func (in *CDIStatus) DeepCopy() *CDIStatus {
	if in == nil {
		return nil
	}
	out := new(CDIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertConfig) DeepCopyInto(out *CertConfig) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.RenewBefore != nil {
		in, out := &in.RenewBefore, &out.RenewBefore
		*out = new(metav1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertConfig.
func (in *CertConfig) DeepCopy() *CertConfig {
	if in == nil {
		return nil
	}
	out := new(CertConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClaimPropertySet) DeepCopyInto(out *ClaimPropertySet) {
	*out = *in
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	if in.VolumeMode != nil {
		in, out := &in.VolumeMode, &out.VolumeMode
		*out = new(v1.PersistentVolumeMode)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClaimPropertySet.
func (in *ClaimPropertySet) DeepCopy() *ClaimPropertySet {
	if in == nil {
		return nil
	}
	out := new(ClaimPropertySet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolume) DeepCopyInto(out *DataVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolume.
func (in *DataVolume) DeepCopy() *DataVolume {
	if in == nil {
		return nil
	}
	out := new(DataVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeBlankImage) DeepCopyInto(out *DataVolumeBlankImage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeBlankImage.
func (in *DataVolumeBlankImage) DeepCopy() *DataVolumeBlankImage {
	if in == nil {
		return nil
	}
	out := new(DataVolumeBlankImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeCheckpoint) DeepCopyInto(out *DataVolumeCheckpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeCheckpoint.
func (in *DataVolumeCheckpoint) DeepCopy() *DataVolumeCheckpoint {
	if in == nil {
		return nil
	}
	out := new(DataVolumeCheckpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeCondition) DeepCopyInto(out *DataVolumeCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeCondition.
func (in *DataVolumeCondition) DeepCopy() *DataVolumeCondition {
	if in == nil {
		return nil
	}
	out := new(DataVolumeCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeList) DeepCopyInto(out *DataVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeList.
func (in *DataVolumeList) DeepCopy() *DataVolumeList {
	if in == nil {
		return nil
	}
	out := new(DataVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeSource) DeepCopyInto(out *DataVolumeSource) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(DataVolumeSourceHTTP)
		**out = **in
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(DataVolumeSourceS3)
		**out = **in
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(DataVolumeSourceRegistry)
		**out = **in
	}
	if in.PVC != nil {
		in, out := &in.PVC, &out.PVC
		*out = new(DataVolumeSourcePVC)
		**out = **in
	}
	if in.Upload != nil {
		in, out := &in.Upload, &out.Upload
		*out = new(DataVolumeSourceUpload)
		**out = **in
	}
	if in.Blank != nil {
		in, out := &in.Blank, &out.Blank
		*out = new(DataVolumeBlankImage)
		**out = **in
	}
	if in.Imageio != nil {
		in, out := &in.Imageio, &out.Imageio
		*out = new(DataVolumeSourceImageIO)
		**out = **in
	}
	if in.VDDK != nil {
		in, out := &in.VDDK, &out.VDDK
		*out = new(DataVolumeSourceVDDK)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeSource.
func (in *DataVolumeSource) DeepCopy() *DataVolumeSource {
	if in == nil {
		return nil
	}
	out := new(DataVolumeSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeSourceHTTP) DeepCopyInto(out *DataVolumeSourceHTTP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeSourceHTTP.
func (in *DataVolumeSourceHTTP) DeepCopy() *DataVolumeSourceHTTP {
	if in == nil {
		return nil
	}
	out := new(DataVolumeSourceHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeSourceImageIO) DeepCopyInto(out *DataVolumeSourceImageIO) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeSourceImageIO.
func (in *DataVolumeSourceImageIO) DeepCopy() *DataVolumeSourceImageIO {
	if in == nil {
		return nil
	}
	out := new(DataVolumeSourceImageIO)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeSourcePVC) DeepCopyInto(out *DataVolumeSourcePVC) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeSourcePVC.
func (in *DataVolumeSourcePVC) DeepCopy() *DataVolumeSourcePVC {
	if in == nil {
		return nil
	}
	out := new(DataVolumeSourcePVC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeSourceRegistry) DeepCopyInto(out *DataVolumeSourceRegistry) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeSourceRegistry.
func (in *DataVolumeSourceRegistry) DeepCopy() *DataVolumeSourceRegistry {
	if in == nil {
		return nil
	}
	out := new(DataVolumeSourceRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeSourceS3) DeepCopyInto(out *DataVolumeSourceS3) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeSourceS3.
func (in *DataVolumeSourceS3) DeepCopy() *DataVolumeSourceS3 {
	if in == nil {
		return nil
	}
	out := new(DataVolumeSourceS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeSourceUpload) DeepCopyInto(out *DataVolumeSourceUpload) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeSourceUpload.
func (in *DataVolumeSourceUpload) DeepCopy() *DataVolumeSourceUpload {
	if in == nil {
		return nil
	}
	out := new(DataVolumeSourceUpload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeSourceVDDK) DeepCopyInto(out *DataVolumeSourceVDDK) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeSourceVDDK.
func (in *DataVolumeSourceVDDK) DeepCopy() *DataVolumeSourceVDDK {
	if in == nil {
		return nil
	}
	out := new(DataVolumeSourceVDDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeSpec) DeepCopyInto(out *DataVolumeSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	if in.PVC != nil {
		in, out := &in.PVC, &out.PVC
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Checkpoints != nil {
		in, out := &in.Checkpoints, &out.Checkpoints
		*out = make([]DataVolumeCheckpoint, len(*in))
		copy(*out, *in)
	}
	if in.Preallocation != nil {
		in, out := &in.Preallocation, &out.Preallocation
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeSpec.
func (in *DataVolumeSpec) DeepCopy() *DataVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(DataVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataVolumeStatus) DeepCopyInto(out *DataVolumeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]DataVolumeCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataVolumeStatus.
func (in *DataVolumeStatus) DeepCopy() *DataVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(DataVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilesystemOverhead) DeepCopyInto(out *FilesystemOverhead) {
	*out = *in
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = make(map[string]Percent, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilesystemOverhead.
func (in *FilesystemOverhead) DeepCopy() *FilesystemOverhead {
	if in == nil {
		return nil
	}
	out := new(FilesystemOverhead)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportProxy) DeepCopyInto(out *ImportProxy) {
	*out = *in
	if in.HTTPProxy != nil {
		in, out := &in.HTTPProxy, &out.HTTPProxy
		*out = new(string)
		**out = **in
	}
	if in.HTTPSProxy != nil {
		in, out := &in.HTTPSProxy, &out.HTTPSProxy
		*out = new(string)
		**out = **in
	}
	if in.NoProxy != nil {
		in, out := &in.NoProxy, &out.NoProxy
		*out = new(string)
		**out = **in
	}
	if in.TrustedCAProxy != nil {
		in, out := &in.TrustedCAProxy, &out.TrustedCAProxy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportProxy.
func (in *ImportProxy) DeepCopy() *ImportProxy {
	if in == nil {
		return nil
	}
	out := new(ImportProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProfile) DeepCopyInto(out *StorageProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProfile.
func (in *StorageProfile) DeepCopy() *StorageProfile {
	if in == nil {
		return nil
	}
	out := new(StorageProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProfileList) DeepCopyInto(out *StorageProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProfileList.
func (in *StorageProfileList) DeepCopy() *StorageProfileList {
	if in == nil {
		return nil
	}
	out := new(StorageProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProfileSpec) DeepCopyInto(out *StorageProfileSpec) {
	*out = *in
	if in.ClaimPropertySets != nil {
		in, out := &in.ClaimPropertySets, &out.ClaimPropertySets
		*out = make([]ClaimPropertySet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProfileSpec.
func (in *StorageProfileSpec) DeepCopy() *StorageProfileSpec {
	if in == nil {
		return nil
	}
	out := new(StorageProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageProfileStatus) DeepCopyInto(out *StorageProfileStatus) {
	*out = *in
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	if in.Provisioner != nil {
		in, out := &in.Provisioner, &out.Provisioner
		*out = new(string)
		**out = **in
	}
	if in.ClaimPropertySets != nil {
		in, out := &in.ClaimPropertySets, &out.ClaimPropertySets
		*out = make([]ClaimPropertySet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageProfileStatus.
func (in *StorageProfileStatus) DeepCopy() *StorageProfileStatus {
	if in == nil {
		return nil
	}
	out := new(StorageProfileStatus)
	in.DeepCopyInto(out)
	return out
}
