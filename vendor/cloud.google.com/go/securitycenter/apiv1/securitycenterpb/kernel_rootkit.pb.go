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
// source: google/cloud/securitycenter/v1/kernel_rootkit.proto

package securitycenterpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Kernel mode rootkit signatures.
type KernelRootkit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rootkit name when available.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// True when unexpected modifications of kernel code memory are present.
	UnexpectedCodeModification bool `protobuf:"varint,2,opt,name=unexpected_code_modification,json=unexpectedCodeModification,proto3" json:"unexpected_code_modification,omitempty"`
	// True when unexpected modifications of kernel read-only data memory are
	// present.
	UnexpectedReadOnlyDataModification bool `protobuf:"varint,3,opt,name=unexpected_read_only_data_modification,json=unexpectedReadOnlyDataModification,proto3" json:"unexpected_read_only_data_modification,omitempty"`
	// True when `ftrace` points are present with callbacks pointing to regions
	// that are not in the expected kernel or module code range.
	UnexpectedFtraceHandler bool `protobuf:"varint,4,opt,name=unexpected_ftrace_handler,json=unexpectedFtraceHandler,proto3" json:"unexpected_ftrace_handler,omitempty"`
	// True when `kprobe` points are present with callbacks pointing to regions
	// that are not in the expected kernel or module code range.
	UnexpectedKprobeHandler bool `protobuf:"varint,5,opt,name=unexpected_kprobe_handler,json=unexpectedKprobeHandler,proto3" json:"unexpected_kprobe_handler,omitempty"`
	// True when kernel code pages that are not in the expected kernel or module
	// code regions are present.
	UnexpectedKernelCodePages bool `protobuf:"varint,6,opt,name=unexpected_kernel_code_pages,json=unexpectedKernelCodePages,proto3" json:"unexpected_kernel_code_pages,omitempty"`
	// True when system call handlers that are are not in the expected kernel or
	// module code regions are present.
	UnexpectedSystemCallHandler bool `protobuf:"varint,7,opt,name=unexpected_system_call_handler,json=unexpectedSystemCallHandler,proto3" json:"unexpected_system_call_handler,omitempty"`
	// True when interrupt handlers that are are not in the expected kernel or
	// module code regions are present.
	UnexpectedInterruptHandler bool `protobuf:"varint,8,opt,name=unexpected_interrupt_handler,json=unexpectedInterruptHandler,proto3" json:"unexpected_interrupt_handler,omitempty"`
	// True when unexpected processes in the scheduler run queue are present. Such
	// processes are in the run queue, but not in the process task list.
	UnexpectedProcessesInRunqueue bool `protobuf:"varint,9,opt,name=unexpected_processes_in_runqueue,json=unexpectedProcessesInRunqueue,proto3" json:"unexpected_processes_in_runqueue,omitempty"`
}

func (x *KernelRootkit) Reset() {
	*x = KernelRootkit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_kernel_rootkit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KernelRootkit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KernelRootkit) ProtoMessage() {}

func (x *KernelRootkit) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_kernel_rootkit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KernelRootkit.ProtoReflect.Descriptor instead.
func (*KernelRootkit) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_kernel_rootkit_proto_rawDescGZIP(), []int{0}
}

func (x *KernelRootkit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KernelRootkit) GetUnexpectedCodeModification() bool {
	if x != nil {
		return x.UnexpectedCodeModification
	}
	return false
}

func (x *KernelRootkit) GetUnexpectedReadOnlyDataModification() bool {
	if x != nil {
		return x.UnexpectedReadOnlyDataModification
	}
	return false
}

func (x *KernelRootkit) GetUnexpectedFtraceHandler() bool {
	if x != nil {
		return x.UnexpectedFtraceHandler
	}
	return false
}

func (x *KernelRootkit) GetUnexpectedKprobeHandler() bool {
	if x != nil {
		return x.UnexpectedKprobeHandler
	}
	return false
}

func (x *KernelRootkit) GetUnexpectedKernelCodePages() bool {
	if x != nil {
		return x.UnexpectedKernelCodePages
	}
	return false
}

func (x *KernelRootkit) GetUnexpectedSystemCallHandler() bool {
	if x != nil {
		return x.UnexpectedSystemCallHandler
	}
	return false
}

func (x *KernelRootkit) GetUnexpectedInterruptHandler() bool {
	if x != nil {
		return x.UnexpectedInterruptHandler
	}
	return false
}

func (x *KernelRootkit) GetUnexpectedProcessesInRunqueue() bool {
	if x != nil {
		return x.UnexpectedProcessesInRunqueue
	}
	return false
}

var File_google_cloud_securitycenter_v1_kernel_rootkit_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_kernel_rootkit_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x6b, 0x69, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xc2, 0x04, 0x0a, 0x0d, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c,
	0x52, 0x6f, 0x6f, 0x74, 0x6b, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x1c, 0x75,
	0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x1a, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x64,
	0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x52, 0x0a,
	0x26, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x22, 0x75,
	0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3a, 0x0a, 0x19, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x66, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x46, 0x74, 0x72, 0x61, 0x63, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x3a, 0x0a,
	0x19, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6b, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x17, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4b, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x1c, 0x75, 0x6e, 0x65,
	0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x19, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4b, 0x65, 0x72, 0x6e, 0x65,
	0x6c, 0x43, 0x6f, 0x64, 0x65, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x1e, 0x75, 0x6e,
	0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x1b, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x43, 0x61, 0x6c, 0x6c, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12,
	0x40, 0x0a, 0x1c, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x72, 0x75, 0x70, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x12, 0x47, 0x0a, 0x20, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x72, 0x75, 0x6e,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x75, 0x6e, 0x65,
	0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x49, 0x6e, 0x52, 0x75, 0x6e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x42, 0xec, 0x01, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x12, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x52, 0x6f, 0x6f, 0x74, 0x6b, 0x69, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x70, 0x62, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_securitycenter_v1_kernel_rootkit_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_kernel_rootkit_proto_rawDescData = file_google_cloud_securitycenter_v1_kernel_rootkit_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_kernel_rootkit_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_kernel_rootkit_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_kernel_rootkit_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_kernel_rootkit_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_kernel_rootkit_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_kernel_rootkit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_v1_kernel_rootkit_proto_goTypes = []interface{}{
	(*KernelRootkit)(nil), // 0: google.cloud.securitycenter.v1.KernelRootkit
}
var file_google_cloud_securitycenter_v1_kernel_rootkit_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_kernel_rootkit_proto_init() }
func file_google_cloud_securitycenter_v1_kernel_rootkit_proto_init() {
	if File_google_cloud_securitycenter_v1_kernel_rootkit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1_kernel_rootkit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KernelRootkit); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_securitycenter_v1_kernel_rootkit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_kernel_rootkit_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_kernel_rootkit_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v1_kernel_rootkit_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_kernel_rootkit_proto = out.File
	file_google_cloud_securitycenter_v1_kernel_rootkit_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_kernel_rootkit_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_kernel_rootkit_proto_depIdxs = nil
}
