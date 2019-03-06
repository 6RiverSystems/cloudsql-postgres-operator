/*
Copyright 2019 The cloudsql-operator Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/travelaudience/cloudsql-operator/pkg/apis/cloudsql"
)

var (
	AddToScheme        = SchemeBuilder.AddToScheme
	SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
	SchemeGroupVersion = schema.GroupVersion{
		Group:   cloudsql.GroupName,
		Version: "v1alpha1",
	}
)

func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, &PostgresqlInstance{}, &PostgresqlInstanceList{})
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
