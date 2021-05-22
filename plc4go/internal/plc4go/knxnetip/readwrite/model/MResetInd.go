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
type MResetInd struct {
	Parent *CEMI
}

// The corresponding interface
type IMResetInd interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *MResetInd) MessageCode() uint8 {
	return 0xF0
}

func (m *MResetInd) InitializeParent(parent *CEMI) {
}

func NewMResetInd() *CEMI {
	child := &MResetInd{
		Parent: NewCEMI(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastMResetInd(structType interface{}) *MResetInd {
	castFunc := func(typ interface{}) *MResetInd {
		if casted, ok := typ.(MResetInd); ok {
			return &casted
		}
		if casted, ok := typ.(*MResetInd); ok {
			return casted
		}
		if casted, ok := typ.(CEMI); ok {
			return CastMResetInd(casted.Child)
		}
		if casted, ok := typ.(*CEMI); ok {
			return CastMResetInd(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MResetInd) GetTypeName() string {
	return "MResetInd"
}

func (m *MResetInd) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *MResetInd) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *MResetInd) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func MResetIndParse(readBuffer utils.ReadBuffer) (*CEMI, error) {
	if pullErr := readBuffer.PullContext("MResetInd"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("MResetInd"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MResetInd{
		Parent: &CEMI{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *MResetInd) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MResetInd"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("MResetInd"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *MResetInd) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
