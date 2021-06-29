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
type ApduDataUserMessage struct {
	Parent *ApduData
}

// The corresponding interface
type IApduDataUserMessage interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataUserMessage) ApciType() uint8 {
	return 0xB
}

func (m *ApduDataUserMessage) InitializeParent(parent *ApduData) {
}

func NewApduDataUserMessage() *ApduData {
	child := &ApduDataUserMessage{
		Parent: NewApduData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataUserMessage(structType interface{}) *ApduDataUserMessage {
	castFunc := func(typ interface{}) *ApduDataUserMessage {
		if casted, ok := typ.(ApduDataUserMessage); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataUserMessage); ok {
			return casted
		}
		if casted, ok := typ.(ApduData); ok {
			return CastApduDataUserMessage(casted.Child)
		}
		if casted, ok := typ.(*ApduData); ok {
			return CastApduDataUserMessage(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataUserMessage) GetTypeName() string {
	return "ApduDataUserMessage"
}

func (m *ApduDataUserMessage) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataUserMessage) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataUserMessage) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataUserMessageParse(readBuffer utils.ReadBuffer) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataUserMessage"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataUserMessage"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataUserMessage{
		Parent: &ApduData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataUserMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataUserMessage"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataUserMessage"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataUserMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}