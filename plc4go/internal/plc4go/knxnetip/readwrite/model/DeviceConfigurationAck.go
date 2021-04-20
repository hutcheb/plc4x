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
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type DeviceConfigurationAck struct {
	DeviceConfigurationAckDataBlock *DeviceConfigurationAckDataBlock
	Parent                          *KnxNetIpMessage
}

// The corresponding interface
type IDeviceConfigurationAck interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *DeviceConfigurationAck) MsgType() uint16 {
	return 0x0311
}

func (m *DeviceConfigurationAck) InitializeParent(parent *KnxNetIpMessage) {
}

func NewDeviceConfigurationAck(deviceConfigurationAckDataBlock *DeviceConfigurationAckDataBlock) *KnxNetIpMessage {
	child := &DeviceConfigurationAck{
		DeviceConfigurationAckDataBlock: deviceConfigurationAckDataBlock,
		Parent:                          NewKnxNetIpMessage(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastDeviceConfigurationAck(structType interface{}) *DeviceConfigurationAck {
	castFunc := func(typ interface{}) *DeviceConfigurationAck {
		if casted, ok := typ.(DeviceConfigurationAck); ok {
			return &casted
		}
		if casted, ok := typ.(*DeviceConfigurationAck); ok {
			return casted
		}
		if casted, ok := typ.(KnxNetIpMessage); ok {
			return CastDeviceConfigurationAck(casted.Child)
		}
		if casted, ok := typ.(*KnxNetIpMessage); ok {
			return CastDeviceConfigurationAck(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *DeviceConfigurationAck) GetTypeName() string {
	return "DeviceConfigurationAck"
}

func (m *DeviceConfigurationAck) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *DeviceConfigurationAck) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (deviceConfigurationAckDataBlock)
	lengthInBits += m.DeviceConfigurationAckDataBlock.LengthInBits()

	return lengthInBits
}

func (m *DeviceConfigurationAck) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func DeviceConfigurationAckParse(io utils.ReadBuffer) (*KnxNetIpMessage, error) {
	io.PullContext("DeviceConfigurationAck")

	io.PullContext("deviceConfigurationAckDataBlock")

	// Simple Field (deviceConfigurationAckDataBlock)
	deviceConfigurationAckDataBlock, _deviceConfigurationAckDataBlockErr := DeviceConfigurationAckDataBlockParse(io)
	if _deviceConfigurationAckDataBlockErr != nil {
		return nil, errors.Wrap(_deviceConfigurationAckDataBlockErr, "Error parsing 'deviceConfigurationAckDataBlock' field")
	}
	io.CloseContext("deviceConfigurationAckDataBlock")

	io.CloseContext("DeviceConfigurationAck")

	// Create a partially initialized instance
	_child := &DeviceConfigurationAck{
		DeviceConfigurationAckDataBlock: deviceConfigurationAckDataBlock,
		Parent:                          &KnxNetIpMessage{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *DeviceConfigurationAck) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("DeviceConfigurationAck")

		// Simple Field (deviceConfigurationAckDataBlock)
		io.PushContext("deviceConfigurationAckDataBlock")
		_deviceConfigurationAckDataBlockErr := m.DeviceConfigurationAckDataBlock.Serialize(io)
		io.PopContext("deviceConfigurationAckDataBlock")
		if _deviceConfigurationAckDataBlockErr != nil {
			return errors.Wrap(_deviceConfigurationAckDataBlockErr, "Error serializing 'deviceConfigurationAckDataBlock' field")
		}

		io.PopContext("DeviceConfigurationAck")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *DeviceConfigurationAck) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "deviceConfigurationAckDataBlock":
				var data DeviceConfigurationAckDataBlock
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.DeviceConfigurationAckDataBlock = &data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *DeviceConfigurationAck) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.DeviceConfigurationAckDataBlock, xml.StartElement{Name: xml.Name{Local: "deviceConfigurationAckDataBlock"}}); err != nil {
		return err
	}
	return nil
}

func (m DeviceConfigurationAck) String() string {
	return string(m.Box("", 120))
}

func (m DeviceConfigurationAck) Box(name string, width int) utils.AsciiBox {
	boxName := "DeviceConfigurationAck"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Complex field (case complex)
		boxes = append(boxes, m.DeviceConfigurationAckDataBlock.Box("deviceConfigurationAckDataBlock", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
