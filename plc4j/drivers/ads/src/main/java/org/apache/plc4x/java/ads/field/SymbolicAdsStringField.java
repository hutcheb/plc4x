/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 */
package org.apache.plc4x.java.ads.field;

import org.apache.plc4x.java.ads.readwrite.types.AdsDataType;
import org.apache.plc4x.java.api.exceptions.PlcInvalidFieldException;
import org.w3c.dom.Document;
import org.w3c.dom.Element;

import java.util.Objects;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * ADS address witch is defined by symbolic name (e.g. {@code Main.items[0]}).
 */
public class SymbolicAdsStringField extends SymbolicAdsField implements AdsStringField {

    private static final Pattern SYMBOLIC_ADDRESS_STRING_PATTERN = Pattern.compile("^(?<symbolicAddress>.+):(?<adsDataType>'STRING'|'WSTRING')\\((?<stringLength>\\d{1,3})\\)(\\[(?<numberOfElements>\\d+)])?");

    private final int stringLength;

    private SymbolicAdsStringField(String symbolicAddress, AdsDataType adsDataType, int stringLength, Integer numberOfElements) {
        super(symbolicAddress, adsDataType, numberOfElements);
        this.stringLength = stringLength;
    }

    public static SymbolicAdsStringField of(String address) {
        Matcher matcher = SYMBOLIC_ADDRESS_STRING_PATTERN.matcher(address);
        if (!matcher.matches()) {
            throw new PlcInvalidFieldException(address, SYMBOLIC_ADDRESS_STRING_PATTERN, "{address}");
        }
        String symbolicAddress = matcher.group("symbolicAddress");

        String adsDataTypeString = matcher.group("adsDataType");
        AdsDataType adsDataType = AdsDataType.valueOf(adsDataTypeString);

        String stringLengthString = matcher.group("stringLength");
        Integer stringLength = stringLengthString != null ? Integer.valueOf(stringLengthString) : null;

        String numberOfElementsString = matcher.group("numberOfElements");
        Integer numberOfElements = numberOfElementsString != null ? Integer.valueOf(numberOfElementsString) : null;

        return new SymbolicAdsStringField(symbolicAddress, adsDataType, stringLength, numberOfElements);
    }

    public static boolean matches(String address) {
        return SYMBOLIC_ADDRESS_STRING_PATTERN.matcher(address).matches();
    }

    @Override
    public int getStringLength() {
        return stringLength;
    }

    @Override
    public String toString() {
        return "SymbolicAdsStringField{" +
            "symbolicAddress='" + getSymbolicAddress() + '\'' +
            ", stringLength=" + stringLength +
            '}';
    }

    @Override
    public void xmlSerialize(Element parent) {
        Document doc = parent.getOwnerDocument();
        Element messageElement = doc.createElement(getClass().getSimpleName());
        parent.appendChild(messageElement);

        Element symbolicAddressElement = doc.createElement("symbolicAddress");
        symbolicAddressElement.appendChild(doc.createTextNode(getSymbolicAddress()));
        messageElement.appendChild(symbolicAddressElement);

        Element numberOfElementsElement = doc.createElement("numberOfElements");
        numberOfElementsElement.appendChild(doc.createTextNode(Integer.toString(getNumberOfElements())));
        messageElement.appendChild(numberOfElementsElement);

        Element datatypeElement = doc.createElement("dataType");
        datatypeElement.appendChild(doc.createTextNode(getPlcDataType()));
        messageElement.appendChild(datatypeElement);
    }
}
