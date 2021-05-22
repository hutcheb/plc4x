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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BVLCReadForeignDeviceTable struct {
	Parent *BVLC
}

// The corresponding interface
type IBVLCReadForeignDeviceTable interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCReadForeignDeviceTable) BvlcFunction() uint8 {
	return 0x06
}

func (m *BVLCReadForeignDeviceTable) InitializeParent(parent *BVLC) {
}

func NewBVLCReadForeignDeviceTable() *BVLC {
	child := &BVLCReadForeignDeviceTable{
		Parent: NewBVLC(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBVLCReadForeignDeviceTable(structType interface{}) *BVLCReadForeignDeviceTable {
	castFunc := func(typ interface{}) *BVLCReadForeignDeviceTable {
		if casted, ok := typ.(BVLCReadForeignDeviceTable); ok {
			return &casted
		}
		if casted, ok := typ.(*BVLCReadForeignDeviceTable); ok {
			return casted
		}
		if casted, ok := typ.(BVLC); ok {
			return CastBVLCReadForeignDeviceTable(casted.Child)
		}
		if casted, ok := typ.(*BVLC); ok {
			return CastBVLCReadForeignDeviceTable(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BVLCReadForeignDeviceTable) GetTypeName() string {
	return "BVLCReadForeignDeviceTable"
}

func (m *BVLCReadForeignDeviceTable) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BVLCReadForeignDeviceTable) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *BVLCReadForeignDeviceTable) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BVLCReadForeignDeviceTableParse(readBuffer utils.ReadBuffer) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLCReadForeignDeviceTable"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BVLCReadForeignDeviceTable"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCReadForeignDeviceTable{
		Parent: &BVLC{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BVLCReadForeignDeviceTable) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCReadForeignDeviceTable"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BVLCReadForeignDeviceTable"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCReadForeignDeviceTable) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
