/*
Copyright 2020 Clastix Labs.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	AvailableIngressClassesAnnotation       = "capsule.clastix.io/ingress-classes"
	AvailableIngressClassesRegexpAnnotation = "capsule.clastix.io/ingress-classes-regexp"
	AvailableStorageClassesAnnotation       = "capsule.clastix.io/storage-classes"
	AvailableStorageClassesRegexpAnnotation = "capsule.clastix.io/storage-classes-regexp"
)

func UsedQuotaFor(resource corev1.ResourceName) string {
	return "quota.capsule.clastix.io/used-" + resource.String()
}