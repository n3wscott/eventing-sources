/*
Copyright 2020 The Knative Authors.

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
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	duckv1alpha1 "knative.dev/eventing/pkg/apis/duck/v1alpha1"
	eventingduck "knative.dev/eventing/pkg/apis/duck/v1alpha1"
	duckv1beta1 "knative.dev/eventing/pkg/apis/duck/v1beta1"
	"knative.dev/eventing/pkg/apis/messaging/v1beta1"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	pkgduckv1alpha1 "knative.dev/pkg/apis/duck/v1alpha1"
)

// ConvertUp implements apis.Convertible
// Converts source (from v1alpha1.Channel) into v1beta1.Channel
func (source *Channel) ConvertUp(ctx context.Context, obj apis.Convertible) error {
	switch sink := obj.(type) {
	case *v1beta1.Channel:
		sink.ObjectMeta = source.ObjectMeta
		source.Status.ConvertUp(ctx, &sink.Status)
		return source.Spec.ConvertUp(ctx, &sink.Spec)
	default:
		return fmt.Errorf("unknown version, got: %T", sink)
	}
}

// ConvertChannelTemplateSpecUp converts an incoming source (duckv1alpha1.ChannelTemplateSpec
func ConvertChannelTemplateSpecUp(ctx context.Context, source *duckv1alpha1.ChannelTemplateSpec) *v1beta1.ChannelTemplateSpec {
	if source == nil {
		return nil
	}
	return &v1beta1.ChannelTemplateSpec{
		TypeMeta: source.TypeMeta,
		Spec:     source.Spec,
	}
}

// ConvertUp helps implement apis.Convertible
func (source *ChannelSpec) ConvertUp(ctx context.Context, sink *v1beta1.ChannelSpec) error {
	if source.ChannelTemplate != nil {
		sink.ChannelTemplate = ConvertChannelTemplateSpecUp(ctx, source.ChannelTemplate)
	}
	sink.ChannelableSpec.Delivery = source.Delivery
	if source.Subscribable != nil {
		sink.ChannelableSpec.Subscribers = make([]duckv1beta1.SubscriberSpec, len(source.Subscribable.Subscribers))
		for i, s := range source.Subscribable.Subscribers {
			sink.ChannelableSpec.Subscribers[i] = duckv1beta1.SubscriberSpec{
				UID:           s.UID,
				Generation:    s.Generation,
				SubscriberURI: s.SubscriberURI,
				ReplyURI:      s.ReplyURI,
			}
			// If the source has delivery, use it.
			if s.Delivery != nil {
				sink.ChannelableSpec.Subscribers[i].Delivery = s.Delivery
			} else {
				// If however, there's a Deprecated DeadLetterSinkURI, convert that up
				// to DeliverySpec.
				sink.ChannelableSpec.Subscribers[i].Delivery = &duckv1beta1.DeliverySpec{
					DeadLetterSink: &duckv1.Destination{
						URI: s.DeadLetterSinkURI,
					},
				}
			}
		}
	}
	return nil
}

// ConvertUp helps implement apis.Convertible
func (source *ChannelStatus) ConvertUp(ctx context.Context, sink *v1beta1.ChannelStatus) {
	source.Status.ConvertTo(ctx, &sink.Status)
	if source.AddressStatus.Address != nil {
		sink.AddressStatus.Address = &duckv1.Addressable{}
		source.AddressStatus.Address.ConvertUp(ctx, sink.AddressStatus.Address)
	}
	if source.SubscribableTypeStatus.SubscribableStatus != nil &&
		len(source.SubscribableTypeStatus.SubscribableStatus.Subscribers) > 0 {
		sink.SubscribableStatus.Subscribers = make([]duckv1beta1.SubscriberStatus, len(source.SubscribableTypeStatus.SubscribableStatus.Subscribers))
		for i, ss := range source.SubscribableTypeStatus.SubscribableStatus.Subscribers {
			sink.SubscribableStatus.Subscribers[i] = duckv1beta1.SubscriberStatus{
				UID:                ss.UID,
				ObservedGeneration: ss.ObservedGeneration,
				Ready:              ss.Ready,
				Message:            ss.Message,
			}
		}
	}
	if source.Channel != nil {
		sink.Channel = &duckv1.KReference{
			Kind:       source.Channel.Kind,
			APIVersion: source.Channel.APIVersion,
			Name:       source.Channel.Name,
			Namespace:  source.Channel.Namespace,
		}
	}
}

// ConvertDown implements apis.Convertible.
// Converts obj v1beta1.Channel into v1alpha1.Channel
func (sink *Channel) ConvertDown(ctx context.Context, obj apis.Convertible) error {
	switch source := obj.(type) {
	case *v1beta1.Channel:
		sink.ObjectMeta = source.ObjectMeta
		sink.Status.ConvertDown(ctx, source.Status)
		sink.Spec.ConvertDown(ctx, source.Spec)
		return nil
	default:
		return fmt.Errorf("unknown version, got: %T", source)
	}
}

// ConvertDown helps implement apis.Convertible
func (sink *ChannelSpec) ConvertDown(ctx context.Context, source v1beta1.ChannelSpec) {
	if source.ChannelTemplate != nil {
		sink.ChannelTemplate = ConvertChannelTemplateSpecDown(ctx, source.ChannelTemplate)
	}
	sink.Delivery = source.ChannelableSpec.Delivery
	if len(source.ChannelableSpec.Subscribers) > 0 {
		sink.Subscribable = &eventingduck.Subscribable{
			Subscribers: make([]eventingduck.SubscriberSpec, len(source.ChannelableSpec.Subscribers)),
		}
		for i, s := range source.ChannelableSpec.Subscribers {
			sink.Subscribable.Subscribers[i] = eventingduck.SubscriberSpec{
				UID:           s.UID,
				Generation:    s.Generation,
				SubscriberURI: s.SubscriberURI,
				ReplyURI:      s.ReplyURI,
				Delivery:      s.Delivery,
			}
		}
	}
}

// ConvertDown helps implement apis.Convertible
func (sink *ChannelStatus) ConvertDown(ctx context.Context, source v1beta1.ChannelStatus) error {
	source.Status.ConvertTo(ctx, &sink.Status)
	if source.AddressStatus.Address != nil {
		sink.AddressStatus.Address = &pkgduckv1alpha1.Addressable{}
		if err := sink.AddressStatus.Address.ConvertDown(ctx, source.AddressStatus.Address); err != nil {
			return err
		}
	}
	if len(source.SubscribableStatus.Subscribers) > 0 {
		sink.SubscribableTypeStatus.SubscribableStatus = &duckv1alpha1.SubscribableStatus{
			Subscribers: make([]duckv1alpha1.SubscriberStatus, len(source.SubscribableStatus.Subscribers)),
		}
		for i, ss := range source.SubscribableStatus.Subscribers {
			sink.SubscribableTypeStatus.SubscribableStatus.Subscribers[i] = duckv1alpha1.SubscriberStatus{
				UID:                ss.UID,
				ObservedGeneration: ss.ObservedGeneration,
				Ready:              ss.Ready,
				Message:            ss.Message,
			}
		}
	}
	if source.Channel != nil {
		sink.Channel = &corev1.ObjectReference{
			Kind:       source.Channel.Kind,
			APIVersion: source.Channel.APIVersion,
			Name:       source.Channel.Name,
			Namespace:  source.Channel.Namespace,
		}
	}
	return nil
}

// ConvertChannelTemplateSpecDown converts an incoming source (duckv1beta1.ChannelTemplateSpec)
// to v1alpha1.ChannelTemplateSpec
func ConvertChannelTemplateSpecDown(ctx context.Context, source *v1beta1.ChannelTemplateSpec) *duckv1alpha1.ChannelTemplateSpec {
	if source == nil {
		return nil
	}
	return &duckv1alpha1.ChannelTemplateSpec{
		TypeMeta: source.TypeMeta,
		Spec:     source.Spec,
	}
}
