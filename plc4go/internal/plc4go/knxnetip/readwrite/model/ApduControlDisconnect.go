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
type ApduControlDisconnect struct {
	Parent *ApduControl
}

// The corresponding interface
type IApduControlDisconnect interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduControlDisconnect) ControlType() uint8 {
	return 0x1
}

func (m *ApduControlDisconnect) InitializeParent(parent *ApduControl) {
}

func NewApduControlDisconnect() *ApduControl {
	child := &ApduControlDisconnect{
		Parent: NewApduControl(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduControlDisconnect(structType interface{}) *ApduControlDisconnect {
	castFunc := func(typ interface{}) *ApduControlDisconnect {
		if casted, ok := typ.(ApduControlDisconnect); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduControlDisconnect); ok {
			return casted
		}
		if casted, ok := typ.(ApduControl); ok {
			return CastApduControlDisconnect(casted.Child)
		}
		if casted, ok := typ.(*ApduControl); ok {
			return CastApduControlDisconnect(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduControlDisconnect) GetTypeName() string {
	return "ApduControlDisconnect"
}

func (m *ApduControlDisconnect) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduControlDisconnect) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduControlDisconnect) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduControlDisconnectParse(readBuffer utils.ReadBuffer) (*ApduControl, error) {
	if pullErr := readBuffer.PullContext("ApduControlDisconnect"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduControlDisconnect"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduControlDisconnect{
		Parent: &ApduControl{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduControlDisconnect) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlDisconnect"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduControlDisconnect"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduControlDisconnect) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}