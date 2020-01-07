// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1beta1 "knative.dev/pkg/apis/duck/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaLimitsSpec) DeepCopyInto(out *KafkaLimitsSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaLimitsSpec.
func (in *KafkaLimitsSpec) DeepCopy() *KafkaLimitsSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaLimitsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaRequestsSpec) DeepCopyInto(out *KafkaRequestsSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaRequestsSpec.
func (in *KafkaRequestsSpec) DeepCopy() *KafkaRequestsSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaRequestsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaResourceSpec) DeepCopyInto(out *KafkaResourceSpec) {
	*out = *in
	out.Requests = in.Requests
	out.Limits = in.Limits
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaResourceSpec.
func (in *KafkaResourceSpec) DeepCopy() *KafkaResourceSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSource) DeepCopyInto(out *KafkaSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSource.
func (in *KafkaSource) DeepCopy() *KafkaSource {
	if in == nil {
		return nil
	}
	out := new(KafkaSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSourceList) DeepCopyInto(out *KafkaSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KafkaSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSourceList.
func (in *KafkaSourceList) DeepCopy() *KafkaSourceList {
	if in == nil {
		return nil
	}
	out := new(KafkaSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSourceNetSpec) DeepCopyInto(out *KafkaSourceNetSpec) {
	*out = *in
	in.SASL.DeepCopyInto(&out.SASL)
	in.TLS.DeepCopyInto(&out.TLS)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSourceNetSpec.
func (in *KafkaSourceNetSpec) DeepCopy() *KafkaSourceNetSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaSourceNetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSourceSASLSpec) DeepCopyInto(out *KafkaSourceSASLSpec) {
	*out = *in
	in.User.DeepCopyInto(&out.User)
	in.Password.DeepCopyInto(&out.Password)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSourceSASLSpec.
func (in *KafkaSourceSASLSpec) DeepCopy() *KafkaSourceSASLSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaSourceSASLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSourceSpec) DeepCopyInto(out *KafkaSourceSpec) {
	*out = *in
	in.Net.DeepCopyInto(&out.Net)
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = new(v1beta1.Destination)
		(*in).DeepCopyInto(*out)
	}
	out.Resources = in.Resources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSourceSpec.
func (in *KafkaSourceSpec) DeepCopy() *KafkaSourceSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSourceStatus) DeepCopyInto(out *KafkaSourceStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSourceStatus.
func (in *KafkaSourceStatus) DeepCopy() *KafkaSourceStatus {
	if in == nil {
		return nil
	}
	out := new(KafkaSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSourceTLSSpec) DeepCopyInto(out *KafkaSourceTLSSpec) {
	*out = *in
	in.Cert.DeepCopyInto(&out.Cert)
	in.Key.DeepCopyInto(&out.Key)
	in.CACert.DeepCopyInto(&out.CACert)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSourceTLSSpec.
func (in *KafkaSourceTLSSpec) DeepCopy() *KafkaSourceTLSSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaSourceTLSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretValueFromSource) DeepCopyInto(out *SecretValueFromSource) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretValueFromSource.
func (in *SecretValueFromSource) DeepCopy() *SecretValueFromSource {
	if in == nil {
		return nil
	}
	out := new(SecretValueFromSource)
	in.DeepCopyInto(out)
	return out
}
