//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type KnxNetIpDeviceManagement struct {
	Version uint8
	Parent  *ServiceId
}

// The corresponding interface
type IKnxNetIpDeviceManagement interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *KnxNetIpDeviceManagement) ServiceType() uint8 {
	return 0x03
}

func (m *KnxNetIpDeviceManagement) InitializeParent(parent *ServiceId) {
}

func NewKnxNetIpDeviceManagement(version uint8) *ServiceId {
	child := &KnxNetIpDeviceManagement{
		Version: version,
		Parent:  NewServiceId(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastKnxNetIpDeviceManagement(structType interface{}) *KnxNetIpDeviceManagement {
	castFunc := func(typ interface{}) *KnxNetIpDeviceManagement {
		if casted, ok := typ.(KnxNetIpDeviceManagement); ok {
			return &casted
		}
		if casted, ok := typ.(*KnxNetIpDeviceManagement); ok {
			return casted
		}
		if casted, ok := typ.(ServiceId); ok {
			return CastKnxNetIpDeviceManagement(casted.Child)
		}
		if casted, ok := typ.(*ServiceId); ok {
			return CastKnxNetIpDeviceManagement(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *KnxNetIpDeviceManagement) GetTypeName() string {
	return "KnxNetIpDeviceManagement"
}

func (m *KnxNetIpDeviceManagement) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *KnxNetIpDeviceManagement) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *KnxNetIpDeviceManagement) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func KnxNetIpDeviceManagementParse(readBuffer utils.ReadBuffer) (*ServiceId, error) {
	if pullErr := readBuffer.PullContext("KnxNetIpDeviceManagement"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (version)
	version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}

	if closeErr := readBuffer.CloseContext("KnxNetIpDeviceManagement"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &KnxNetIpDeviceManagement{
		Version: version,
		Parent:  &ServiceId{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *KnxNetIpDeviceManagement) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpDeviceManagement"); pushErr != nil {
			return pushErr
		}

		// Simple Field (version)
		version := uint8(m.Version)
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpDeviceManagement"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *KnxNetIpDeviceManagement) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
