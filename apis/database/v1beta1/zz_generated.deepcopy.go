// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLEntry) DeepCopyInto(out *ACLEntry) {
	*out = *in
	if in.ExpirationTime != nil {
		in, out := &in.ExpirationTime, &out.ExpirationTime
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLEntry.
func (in *ACLEntry) DeepCopy() *ACLEntry {
	if in == nil {
		return nil
	}
	out := new(ACLEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupConfiguration) DeepCopyInto(out *BackupConfiguration) {
	*out = *in
	if in.BinaryLogEnabled != nil {
		in, out := &in.BinaryLogEnabled, &out.BinaryLogEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ReplicationLogArchivingEnabled != nil {
		in, out := &in.ReplicationLogArchivingEnabled, &out.ReplicationLogArchivingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupConfiguration.
func (in *BackupConfiguration) DeepCopy() *BackupConfiguration {
	if in == nil {
		return nil
	}
	out := new(BackupConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstance) DeepCopyInto(out *CloudsqlInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstance.
func (in *CloudsqlInstance) DeepCopy() *CloudsqlInstance {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudsqlInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceClass) DeepCopyInto(out *CloudsqlInstanceClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.SpecTemplate.DeepCopyInto(&out.SpecTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceClass.
func (in *CloudsqlInstanceClass) DeepCopy() *CloudsqlInstanceClass {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudsqlInstanceClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceClassList) DeepCopyInto(out *CloudsqlInstanceClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudsqlInstanceClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceClassList.
func (in *CloudsqlInstanceClassList) DeepCopy() *CloudsqlInstanceClassList {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudsqlInstanceClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceClassSpecTemplate) DeepCopyInto(out *CloudsqlInstanceClassSpecTemplate) {
	*out = *in
	in.ClassSpecTemplate.DeepCopyInto(&out.ClassSpecTemplate)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceClassSpecTemplate.
func (in *CloudsqlInstanceClassSpecTemplate) DeepCopy() *CloudsqlInstanceClassSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceClassSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceList) DeepCopyInto(out *CloudsqlInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudsqlInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceList.
func (in *CloudsqlInstanceList) DeepCopy() *CloudsqlInstanceList {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudsqlInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceObservation) DeepCopyInto(out *CloudsqlInstanceObservation) {
	*out = *in
	if in.DiskEncryptionStatus != nil {
		in, out := &in.DiskEncryptionStatus, &out.DiskEncryptionStatus
		*out = new(DiskEncryptionStatus)
		**out = **in
	}
	if in.FailoverReplica != nil {
		in, out := &in.FailoverReplica, &out.FailoverReplica
		*out = new(DatabaseInstanceFailoverReplicaStatus)
		**out = **in
	}
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]*IPMapping, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(IPMapping)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceObservation.
func (in *CloudsqlInstanceObservation) DeepCopy() *CloudsqlInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceParameters) DeepCopyInto(out *CloudsqlInstanceParameters) {
	*out = *in
	in.Settings.DeepCopyInto(&out.Settings)
	if in.DatabaseVersion != nil {
		in, out := &in.DatabaseVersion, &out.DatabaseVersion
		*out = new(string)
		**out = **in
	}
	if in.MasterInstanceName != nil {
		in, out := &in.MasterInstanceName, &out.MasterInstanceName
		*out = new(string)
		**out = **in
	}
	if in.DiskEncryptionConfiguration != nil {
		in, out := &in.DiskEncryptionConfiguration, &out.DiskEncryptionConfiguration
		*out = new(DiskEncryptionConfiguration)
		**out = **in
	}
	if in.FailoverReplica != nil {
		in, out := &in.FailoverReplica, &out.FailoverReplica
		*out = new(DatabaseInstanceFailoverReplicaSpec)
		**out = **in
	}
	if in.GceZone != nil {
		in, out := &in.GceZone, &out.GceZone
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.MaxDiskSize != nil {
		in, out := &in.MaxDiskSize, &out.MaxDiskSize
		*out = new(int64)
		**out = **in
	}
	if in.OnPremisesConfiguration != nil {
		in, out := &in.OnPremisesConfiguration, &out.OnPremisesConfiguration
		*out = new(OnPremisesConfiguration)
		**out = **in
	}
	if in.ReplicaNames != nil {
		in, out := &in.ReplicaNames, &out.ReplicaNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SuspensionReason != nil {
		in, out := &in.SuspensionReason, &out.SuspensionReason
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceParameters.
func (in *CloudsqlInstanceParameters) DeepCopy() *CloudsqlInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceSpec) DeepCopyInto(out *CloudsqlInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceSpec.
func (in *CloudsqlInstanceSpec) DeepCopy() *CloudsqlInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudsqlInstanceStatus) DeepCopyInto(out *CloudsqlInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudsqlInstanceStatus.
func (in *CloudsqlInstanceStatus) DeepCopy() *CloudsqlInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(CloudsqlInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseFlags) DeepCopyInto(out *DatabaseFlags) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseFlags.
func (in *DatabaseFlags) DeepCopy() *DatabaseFlags {
	if in == nil {
		return nil
	}
	out := new(DatabaseFlags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseInstanceFailoverReplicaSpec) DeepCopyInto(out *DatabaseInstanceFailoverReplicaSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseInstanceFailoverReplicaSpec.
func (in *DatabaseInstanceFailoverReplicaSpec) DeepCopy() *DatabaseInstanceFailoverReplicaSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseInstanceFailoverReplicaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseInstanceFailoverReplicaStatus) DeepCopyInto(out *DatabaseInstanceFailoverReplicaStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseInstanceFailoverReplicaStatus.
func (in *DatabaseInstanceFailoverReplicaStatus) DeepCopy() *DatabaseInstanceFailoverReplicaStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseInstanceFailoverReplicaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryptionConfiguration) DeepCopyInto(out *DiskEncryptionConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryptionConfiguration.
func (in *DiskEncryptionConfiguration) DeepCopy() *DiskEncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(DiskEncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryptionStatus) DeepCopyInto(out *DiskEncryptionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryptionStatus.
func (in *DiskEncryptionStatus) DeepCopy() *DiskEncryptionStatus {
	if in == nil {
		return nil
	}
	out := new(DiskEncryptionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPConfiguration) DeepCopyInto(out *IPConfiguration) {
	*out = *in
	if in.AuthorizedNetworks != nil {
		in, out := &in.AuthorizedNetworks, &out.AuthorizedNetworks
		*out = make([]*ACLEntry, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ACLEntry)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Ipv4Enabled != nil {
		in, out := &in.Ipv4Enabled, &out.Ipv4Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PrivateNetwork != nil {
		in, out := &in.PrivateNetwork, &out.PrivateNetwork
		*out = new(string)
		**out = **in
	}
	if in.RequireSsl != nil {
		in, out := &in.RequireSsl, &out.RequireSsl
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPConfiguration.
func (in *IPConfiguration) DeepCopy() *IPConfiguration {
	if in == nil {
		return nil
	}
	out := new(IPConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPMapping) DeepCopyInto(out *IPMapping) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPMapping.
func (in *IPMapping) DeepCopy() *IPMapping {
	if in == nil {
		return nil
	}
	out := new(IPMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocationPreference) DeepCopyInto(out *LocationPreference) {
	*out = *in
	if in.FollowGaeApplication != nil {
		in, out := &in.FollowGaeApplication, &out.FollowGaeApplication
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocationPreference.
func (in *LocationPreference) DeepCopy() *LocationPreference {
	if in == nil {
		return nil
	}
	out := new(LocationPreference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindow) DeepCopyInto(out *MaintenanceWindow) {
	*out = *in
	if in.Day != nil {
		in, out := &in.Day, &out.Day
		*out = new(int64)
		**out = **in
	}
	if in.Hour != nil {
		in, out := &in.Hour, &out.Hour
		*out = new(int64)
		**out = **in
	}
	if in.UpdateTrack != nil {
		in, out := &in.UpdateTrack, &out.UpdateTrack
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindow.
func (in *MaintenanceWindow) DeepCopy() *MaintenanceWindow {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnPremisesConfiguration) DeepCopyInto(out *OnPremisesConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnPremisesConfiguration.
func (in *OnPremisesConfiguration) DeepCopy() *OnPremisesConfiguration {
	if in == nil {
		return nil
	}
	out := new(OnPremisesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Settings) DeepCopyInto(out *Settings) {
	*out = *in
	if in.ActivationPolicy != nil {
		in, out := &in.ActivationPolicy, &out.ActivationPolicy
		*out = new(string)
		**out = **in
	}
	if in.AuthorizedGaeApplications != nil {
		in, out := &in.AuthorizedGaeApplications, &out.AuthorizedGaeApplications
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AvailabilityType != nil {
		in, out := &in.AvailabilityType, &out.AvailabilityType
		*out = new(string)
		**out = **in
	}
	if in.CrashSafeReplicationEnabled != nil {
		in, out := &in.CrashSafeReplicationEnabled, &out.CrashSafeReplicationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.StorageAutoResize != nil {
		in, out := &in.StorageAutoResize, &out.StorageAutoResize
		*out = new(bool)
		**out = **in
	}
	if in.DataDiskType != nil {
		in, out := &in.DataDiskType, &out.DataDiskType
		*out = new(string)
		**out = **in
	}
	if in.PricingPlan != nil {
		in, out := &in.PricingPlan, &out.PricingPlan
		*out = new(string)
		**out = **in
	}
	if in.ReplicationType != nil {
		in, out := &in.ReplicationType, &out.ReplicationType
		*out = new(string)
		**out = **in
	}
	if in.UserLabels != nil {
		in, out := &in.UserLabels, &out.UserLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DatabaseFlags != nil {
		in, out := &in.DatabaseFlags, &out.DatabaseFlags
		*out = make([]*DatabaseFlags, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(DatabaseFlags)
				**out = **in
			}
		}
	}
	if in.BackupConfiguration != nil {
		in, out := &in.BackupConfiguration, &out.BackupConfiguration
		*out = new(BackupConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.IPConfiguration != nil {
		in, out := &in.IPConfiguration, &out.IPConfiguration
		*out = new(IPConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.LocationPreference != nil {
		in, out := &in.LocationPreference, &out.LocationPreference
		*out = new(LocationPreference)
		(*in).DeepCopyInto(*out)
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindow)
		(*in).DeepCopyInto(*out)
	}
	if in.DataDiskSizeGb != nil {
		in, out := &in.DataDiskSizeGb, &out.DataDiskSizeGb
		*out = new(int64)
		**out = **in
	}
	if in.DatabaseReplicationEnabled != nil {
		in, out := &in.DatabaseReplicationEnabled, &out.DatabaseReplicationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.StorageAutoResizeLimit != nil {
		in, out := &in.StorageAutoResizeLimit, &out.StorageAutoResizeLimit
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Settings.
func (in *Settings) DeepCopy() *Settings {
	if in == nil {
		return nil
	}
	out := new(Settings)
	in.DeepCopyInto(out)
	return out
}