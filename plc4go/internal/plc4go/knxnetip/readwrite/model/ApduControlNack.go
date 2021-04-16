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
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ApduControlNack struct {
	Parent *ApduControl
}

// The corresponding interface
type IApduControlNack interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduControlNack) ControlType() uint8 {
	return 0x3
}

func (m *ApduControlNack) InitializeParent(parent *ApduControl) {
}

func NewApduControlNack() *ApduControl {
	child := &ApduControlNack{
		Parent: NewApduControl(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduControlNack(structType interface{}) *ApduControlNack {
	castFunc := func(typ interface{}) *ApduControlNack {
		if casted, ok := typ.(ApduControlNack); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduControlNack); ok {
			return casted
		}
		if casted, ok := typ.(ApduControl); ok {
			return CastApduControlNack(casted.Child)
		}
		if casted, ok := typ.(*ApduControl); ok {
			return CastApduControlNack(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduControlNack) GetTypeName() string {
	return "ApduControlNack"
}

func (m *ApduControlNack) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduControlNack) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduControlNack) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduControlNackParse(io *utils.ReadBuffer) (*ApduControl, error) {

	// Create a partially initialized instance
	_child := &ApduControlNack{
		Parent: &ApduControl{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduControlNack) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduControlNack) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *ApduControlNack) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func (m ApduControlNack) String() string {
	return string(m.Box("ApduControlNack", utils.DefaultWidth*2))
}

func (m ApduControlNack) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "ApduControlNack"
	}
	boxes := make([]utils.AsciiBox, 0)
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
