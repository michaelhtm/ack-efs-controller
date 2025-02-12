// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package file_system

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	customPreCompare(delta, a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.AvailabilityZoneName, b.ko.Spec.AvailabilityZoneName) {
		delta.Add("Spec.AvailabilityZoneName", a.ko.Spec.AvailabilityZoneName, b.ko.Spec.AvailabilityZoneName)
	} else if a.ko.Spec.AvailabilityZoneName != nil && b.ko.Spec.AvailabilityZoneName != nil {
		if *a.ko.Spec.AvailabilityZoneName != *b.ko.Spec.AvailabilityZoneName {
			delta.Add("Spec.AvailabilityZoneName", a.ko.Spec.AvailabilityZoneName, b.ko.Spec.AvailabilityZoneName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.BackupPolicy, b.ko.Spec.BackupPolicy) {
		delta.Add("Spec.BackupPolicy", a.ko.Spec.BackupPolicy, b.ko.Spec.BackupPolicy)
	} else if a.ko.Spec.BackupPolicy != nil && b.ko.Spec.BackupPolicy != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.BackupPolicy.Status, b.ko.Spec.BackupPolicy.Status) {
			delta.Add("Spec.BackupPolicy.Status", a.ko.Spec.BackupPolicy.Status, b.ko.Spec.BackupPolicy.Status)
		} else if a.ko.Spec.BackupPolicy.Status != nil && b.ko.Spec.BackupPolicy.Status != nil {
			if *a.ko.Spec.BackupPolicy.Status != *b.ko.Spec.BackupPolicy.Status {
				delta.Add("Spec.BackupPolicy.Status", a.ko.Spec.BackupPolicy.Status, b.ko.Spec.BackupPolicy.Status)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Encrypted, b.ko.Spec.Encrypted) {
		delta.Add("Spec.Encrypted", a.ko.Spec.Encrypted, b.ko.Spec.Encrypted)
	} else if a.ko.Spec.Encrypted != nil && b.ko.Spec.Encrypted != nil {
		if *a.ko.Spec.Encrypted != *b.ko.Spec.Encrypted {
			delta.Add("Spec.Encrypted", a.ko.Spec.Encrypted, b.ko.Spec.Encrypted)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.FileSystemProtection, b.ko.Spec.FileSystemProtection) {
		delta.Add("Spec.FileSystemProtection", a.ko.Spec.FileSystemProtection, b.ko.Spec.FileSystemProtection)
	} else if a.ko.Spec.FileSystemProtection != nil && b.ko.Spec.FileSystemProtection != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.FileSystemProtection.ReplicationOverwriteProtection, b.ko.Spec.FileSystemProtection.ReplicationOverwriteProtection) {
			delta.Add("Spec.FileSystemProtection.ReplicationOverwriteProtection", a.ko.Spec.FileSystemProtection.ReplicationOverwriteProtection, b.ko.Spec.FileSystemProtection.ReplicationOverwriteProtection)
		} else if a.ko.Spec.FileSystemProtection.ReplicationOverwriteProtection != nil && b.ko.Spec.FileSystemProtection.ReplicationOverwriteProtection != nil {
			if *a.ko.Spec.FileSystemProtection.ReplicationOverwriteProtection != *b.ko.Spec.FileSystemProtection.ReplicationOverwriteProtection {
				delta.Add("Spec.FileSystemProtection.ReplicationOverwriteProtection", a.ko.Spec.FileSystemProtection.ReplicationOverwriteProtection, b.ko.Spec.FileSystemProtection.ReplicationOverwriteProtection)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID) {
		delta.Add("Spec.KMSKeyID", a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID)
	} else if a.ko.Spec.KMSKeyID != nil && b.ko.Spec.KMSKeyID != nil {
		if *a.ko.Spec.KMSKeyID != *b.ko.Spec.KMSKeyID {
			delta.Add("Spec.KMSKeyID", a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.KMSKeyRef, b.ko.Spec.KMSKeyRef) {
		delta.Add("Spec.KMSKeyRef", a.ko.Spec.KMSKeyRef, b.ko.Spec.KMSKeyRef)
	}
	if len(a.ko.Spec.LifecyclePolicies) != len(b.ko.Spec.LifecyclePolicies) {
		delta.Add("Spec.LifecyclePolicies", a.ko.Spec.LifecyclePolicies, b.ko.Spec.LifecyclePolicies)
	} else if len(a.ko.Spec.LifecyclePolicies) > 0 {
		if !reflect.DeepEqual(a.ko.Spec.LifecyclePolicies, b.ko.Spec.LifecyclePolicies) {
			delta.Add("Spec.LifecyclePolicies", a.ko.Spec.LifecyclePolicies, b.ko.Spec.LifecyclePolicies)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PerformanceMode, b.ko.Spec.PerformanceMode) {
		delta.Add("Spec.PerformanceMode", a.ko.Spec.PerformanceMode, b.ko.Spec.PerformanceMode)
	} else if a.ko.Spec.PerformanceMode != nil && b.ko.Spec.PerformanceMode != nil {
		if *a.ko.Spec.PerformanceMode != *b.ko.Spec.PerformanceMode {
			delta.Add("Spec.PerformanceMode", a.ko.Spec.PerformanceMode, b.ko.Spec.PerformanceMode)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Policy, b.ko.Spec.Policy) {
		delta.Add("Spec.Policy", a.ko.Spec.Policy, b.ko.Spec.Policy)
	} else if a.ko.Spec.Policy != nil && b.ko.Spec.Policy != nil {
		if *a.ko.Spec.Policy != *b.ko.Spec.Policy {
			delta.Add("Spec.Policy", a.ko.Spec.Policy, b.ko.Spec.Policy)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ProvisionedThroughputInMiBps, b.ko.Spec.ProvisionedThroughputInMiBps) {
		delta.Add("Spec.ProvisionedThroughputInMiBps", a.ko.Spec.ProvisionedThroughputInMiBps, b.ko.Spec.ProvisionedThroughputInMiBps)
	} else if a.ko.Spec.ProvisionedThroughputInMiBps != nil && b.ko.Spec.ProvisionedThroughputInMiBps != nil {
		if *a.ko.Spec.ProvisionedThroughputInMiBps != *b.ko.Spec.ProvisionedThroughputInMiBps {
			delta.Add("Spec.ProvisionedThroughputInMiBps", a.ko.Spec.ProvisionedThroughputInMiBps, b.ko.Spec.ProvisionedThroughputInMiBps)
		}
	}
	if !ackcompare.MapStringStringEqual(ToACKTags(a.ko.Spec.Tags), ToACKTags(b.ko.Spec.Tags)) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ThroughputMode, b.ko.Spec.ThroughputMode) {
		delta.Add("Spec.ThroughputMode", a.ko.Spec.ThroughputMode, b.ko.Spec.ThroughputMode)
	} else if a.ko.Spec.ThroughputMode != nil && b.ko.Spec.ThroughputMode != nil {
		if *a.ko.Spec.ThroughputMode != *b.ko.Spec.ThroughputMode {
			delta.Add("Spec.ThroughputMode", a.ko.Spec.ThroughputMode, b.ko.Spec.ThroughputMode)
		}
	}

	return delta
}
