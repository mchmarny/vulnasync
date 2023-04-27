// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: google/cloud/securitycenter/v1/notification_config.proto

package securitycenterpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Cloud Security Command Center (Cloud SCC) notification configs.
//
// A notification config is a Cloud SCC resource that contains the configuration
// to send notifications for create/update events of findings, assets and etc.
type NotificationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The relative resource name of this notification config. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/{organization_id}/notificationConfigs/notify_public_bucket",
	// "folders/{folder_id}/notificationConfigs/notify_public_bucket",
	// or "projects/{project_id}/notificationConfigs/notify_public_bucket".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the notification config (max of 1024 characters).
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The Pub/Sub topic to send notifications to. Its format is
	// "projects/[project_id]/topics/[topic]".
	PubsubTopic string `protobuf:"bytes,3,opt,name=pubsub_topic,json=pubsubTopic,proto3" json:"pubsub_topic,omitempty"`
	// Output only. The service account that needs "pubsub.topics.publish"
	// permission to publish to the Pub/Sub topic.
	ServiceAccount string `protobuf:"bytes,4,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	// The config for triggering notifications.
	//
	// Types that are assignable to NotifyConfig:
	//	*NotificationConfig_StreamingConfig_
	NotifyConfig isNotificationConfig_NotifyConfig `protobuf_oneof:"notify_config"`
}

func (x *NotificationConfig) Reset() {
	*x = NotificationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_notification_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationConfig) ProtoMessage() {}

func (x *NotificationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_notification_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationConfig.ProtoReflect.Descriptor instead.
func (*NotificationConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_notification_config_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NotificationConfig) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NotificationConfig) GetPubsubTopic() string {
	if x != nil {
		return x.PubsubTopic
	}
	return ""
}

func (x *NotificationConfig) GetServiceAccount() string {
	if x != nil {
		return x.ServiceAccount
	}
	return ""
}

func (m *NotificationConfig) GetNotifyConfig() isNotificationConfig_NotifyConfig {
	if m != nil {
		return m.NotifyConfig
	}
	return nil
}

func (x *NotificationConfig) GetStreamingConfig() *NotificationConfig_StreamingConfig {
	if x, ok := x.GetNotifyConfig().(*NotificationConfig_StreamingConfig_); ok {
		return x.StreamingConfig
	}
	return nil
}

type isNotificationConfig_NotifyConfig interface {
	isNotificationConfig_NotifyConfig()
}

type NotificationConfig_StreamingConfig_ struct {
	// The config for triggering streaming-based notifications.
	StreamingConfig *NotificationConfig_StreamingConfig `protobuf:"bytes,5,opt,name=streaming_config,json=streamingConfig,proto3,oneof"`
}

func (*NotificationConfig_StreamingConfig_) isNotificationConfig_NotifyConfig() {}

// The config for streaming-based notifications, which send each event as soon
// as it is detected.
type NotificationConfig_StreamingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Expression that defines the filter to apply across create/update events
	// of assets or findings as specified by the event type. The expression is a
	// list of zero or more restrictions combined via logical operators `AND`
	// and `OR`. Parentheses are supported, and `OR` has higher precedence than
	// `AND`.
	//
	// Restrictions have the form `<field> <operator> <value>` and may have a
	// `-` character in front of them to indicate negation. The fields map to
	// those defined in the corresponding resource.
	//
	// The supported operators are:
	//
	// * `=` for all value types.
	// * `>`, `<`, `>=`, `<=` for integer values.
	// * `:`, meaning substring matching, for strings.
	//
	// The supported value types are:
	//
	// * string literals in quotes.
	// * integer literals without quotes.
	// * boolean literals `true` and `false` without quotes.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *NotificationConfig_StreamingConfig) Reset() {
	*x = NotificationConfig_StreamingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_notification_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationConfig_StreamingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationConfig_StreamingConfig) ProtoMessage() {}

func (x *NotificationConfig_StreamingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_notification_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationConfig_StreamingConfig.ProtoReflect.Descriptor instead.
func (*NotificationConfig_StreamingConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_notification_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NotificationConfig_StreamingConfig) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

var File_google_cloud_securitycenter_v1_notification_config_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_notification_config_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x04, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0xfa, 0x41, 0x1d, 0x0a, 0x1b,
	0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x0b, 0x70, 0x75, 0x62,
	0x73, 0x75, 0x62, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x6f, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x29, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x3a, 0xf8, 0x01, 0xea, 0x41, 0xf4, 0x01, 0x0a, 0x30, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x46, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x7b, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x7d, 0x12, 0x3a, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x7d, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x7b, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d, 0x12,
	0x3c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x7d, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x7b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d, 0x42, 0x0f, 0x0a,
	0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xb4,
	0x02, 0xea, 0x41, 0x40, 0x0a, 0x1b, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x21, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2f, 0x7b, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x7d, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_notification_config_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_notification_config_proto_rawDescData = file_google_cloud_securitycenter_v1_notification_config_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_notification_config_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_notification_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_notification_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_notification_config_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_notification_config_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_notification_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_securitycenter_v1_notification_config_proto_goTypes = []interface{}{
	(*NotificationConfig)(nil),                 // 0: google.cloud.securitycenter.v1.NotificationConfig
	(*NotificationConfig_StreamingConfig)(nil), // 1: google.cloud.securitycenter.v1.NotificationConfig.StreamingConfig
}
var file_google_cloud_securitycenter_v1_notification_config_proto_depIdxs = []int32{
	1, // 0: google.cloud.securitycenter.v1.NotificationConfig.streaming_config:type_name -> google.cloud.securitycenter.v1.NotificationConfig.StreamingConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_notification_config_proto_init() }
func file_google_cloud_securitycenter_v1_notification_config_proto_init() {
	if File_google_cloud_securitycenter_v1_notification_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1_notification_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_securitycenter_v1_notification_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationConfig_StreamingConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_cloud_securitycenter_v1_notification_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*NotificationConfig_StreamingConfig_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securitycenter_v1_notification_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_notification_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_notification_config_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v1_notification_config_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_notification_config_proto = out.File
	file_google_cloud_securitycenter_v1_notification_config_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_notification_config_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_notification_config_proto_depIdxs = nil
}
