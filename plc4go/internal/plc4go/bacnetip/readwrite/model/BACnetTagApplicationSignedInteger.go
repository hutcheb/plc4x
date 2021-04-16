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
	"encoding/hex"
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BACnetTagApplicationSignedInteger struct {
	Data   []int8
	Parent *BACnetTag
}

// The corresponding interface
type IBACnetTagApplicationSignedInteger interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetTagApplicationSignedInteger) ContextSpecificTag() uint8 {
	return 0
}

func (m *BACnetTagApplicationSignedInteger) InitializeParent(parent *BACnetTag, typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) {
	m.Parent.TypeOrTagNumber = typeOrTagNumber
	m.Parent.LengthValueType = lengthValueType
	m.Parent.ExtTagNumber = extTagNumber
	m.Parent.ExtLength = extLength
}

func NewBACnetTagApplicationSignedInteger(data []int8, typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) *BACnetTag {
	child := &BACnetTagApplicationSignedInteger{
		Data:   data,
		Parent: NewBACnetTag(typeOrTagNumber, lengthValueType, extTagNumber, extLength),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetTagApplicationSignedInteger(structType interface{}) *BACnetTagApplicationSignedInteger {
	castFunc := func(typ interface{}) *BACnetTagApplicationSignedInteger {
		if casted, ok := typ.(BACnetTagApplicationSignedInteger); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTagApplicationSignedInteger); ok {
			return casted
		}
		if casted, ok := typ.(BACnetTag); ok {
			return CastBACnetTagApplicationSignedInteger(casted.Child)
		}
		if casted, ok := typ.(*BACnetTag); ok {
			return CastBACnetTagApplicationSignedInteger(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTagApplicationSignedInteger) GetTypeName() string {
	return "BACnetTagApplicationSignedInteger"
}

func (m *BACnetTagApplicationSignedInteger) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetTagApplicationSignedInteger) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *BACnetTagApplicationSignedInteger) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetTagApplicationSignedIntegerParse(io *utils.ReadBuffer, lengthValueType uint8, extLength uint8) (*BACnetTag, error) {

	// Array field (data)
	// Length array
	data := make([]int8, 0)
	_dataLength := utils.InlineIf(bool(bool((lengthValueType) == (5))), func() uint16 { return uint16(extLength) }, func() uint16 { return uint16(lengthValueType) })
	_dataEndPos := io.GetPos() + uint16(_dataLength)
	for io.GetPos() < _dataEndPos {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data = append(data, _item)
	}

	// Create a partially initialized instance
	_child := &BACnetTagApplicationSignedInteger{
		Data:   data,
		Parent: &BACnetTag{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetTagApplicationSignedInteger) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Array Field (data)
		if m.Data != nil {
			for _, _element := range m.Data {
				_elementErr := io.WriteInt8(8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetTagApplicationSignedInteger) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "data":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.Data = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *BACnetTagApplicationSignedInteger) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	_encodedData := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Data))
	_encodedData = strings.ToUpper(_encodedData)
	if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	return nil
}

func (m BACnetTagApplicationSignedInteger) String() string {
	return string(m.Box("BACnetTagApplicationSignedInteger", utils.DefaultWidth*2))
}

func (m BACnetTagApplicationSignedInteger) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "BACnetTagApplicationSignedInteger"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("Data", m.Data, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
