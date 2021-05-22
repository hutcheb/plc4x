/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.eip.readwrite.field;

import org.apache.plc4x.java.api.model.PlcField;
import org.apache.plc4x.java.eip.readwrite.types.CIPDataTypeCode;
import org.apache.plc4x.java.spi.utils.XmlSerializable;
import org.w3c.dom.Document;
import org.w3c.dom.Element;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class EipField implements PlcField, XmlSerializable {

    private static final Pattern ADDRESS_PATTERN =
        Pattern.compile("^%(?<tag>[a-zA-Z_.0-9]+\\[?[0-9]*\\]?):?(?<dataType>[A-Z]*):?(?<elementNb>[0-9]*)");

    private static final String TAG="tag";
    private static final String ELEMENTS="elementNb";
    private static final String TYPE="dataType";


    private final String tag;
    private CIPDataTypeCode type;
    private int  elementNb;

    public CIPDataTypeCode getType() {
        return type;
    }

    public void setType(CIPDataTypeCode type) {
        this.type = type;
    }

    public int getElementNb() {
        return elementNb;
    }

    public void setElementNb(int elementNb) {
        this.elementNb = elementNb;
    }

    public String getTag() {
        return tag;
    }

    public EipField(String tag) {
        this.tag = tag;
    }

    public EipField(String tag, int elementNb) {
        this.tag = tag;
        this.elementNb = elementNb;
    }

    public EipField(String tag, CIPDataTypeCode type, int elementNb) {
        this.tag = tag;
        this.type = type;
        this.elementNb = elementNb;
    }

    public EipField(String tag, CIPDataTypeCode type) {
        this.tag = tag;
        this.type = type;
    }

    public static boolean matches(String fieldQuery){
        return ADDRESS_PATTERN.matcher(fieldQuery).matches();
    }

    public static EipField of(String fieldString){
        Matcher matcher = ADDRESS_PATTERN.matcher(fieldString);
        if(matcher.matches()){
            String tag = matcher.group(TAG);
            int nb=0;
            CIPDataTypeCode type=null;
            if(!matcher.group(ELEMENTS).isEmpty()) {
                nb = Integer.parseInt(matcher.group(ELEMENTS));
            }
            if(!matcher.group(TYPE).isEmpty()) {
                type = CIPDataTypeCode.valueOf(matcher.group(TYPE));
            }
            if(nb!=0){
                if(type!=null){
                    return  new EipField(tag,type,nb);
                }
                return new EipField(tag, nb);
            }
            else{
                if(type!=null){
                    return  new EipField(tag,type);
                }
                return new EipField(tag);
            }
        }
        return null;
    }

    @Override
    public String getPlcDataType() {
        return type.toString();
    }

    @Override
    public Class<?> getDefaultJavaType() {
        switch (type){
            //ToDo differenciate Short, Integer and Long
            case INT:
            case DINT:
            case SINT:
            case LINT:
                return java.lang.Integer.class;
            case STRING:
            case STRING36:
                return java.lang.String.class;
            case REAL:
                return java.lang.Double.class;
            case BOOL:
                return java.lang.Boolean.class;
            default:
                return Object.class;
        }
    }

    @Override
    public void xmlSerialize(Element parent) {
        Document doc = parent.getOwnerDocument();
        Element messageElement = doc.createElement(getClass().getSimpleName());
        parent.appendChild(messageElement);

        Element blockNumberElement = doc.createElement("node");
        blockNumberElement.appendChild(doc.createTextNode(tag));
        messageElement.appendChild(blockNumberElement);

        Element indexElement = doc.createElement("type");
        indexElement.appendChild(doc.createTextNode(type == null ? "" : type.name()));
        messageElement.appendChild(indexElement);

        Element subIndexElement = doc.createElement("elementNb");
        subIndexElement.appendChild(doc.createTextNode(Integer.toString(elementNb)));
        messageElement.appendChild(subIndexElement);

        Element dataType = doc.createElement("defaultJavaType");
        dataType.appendChild(doc.createTextNode((type == null ? Object.class : getDefaultJavaType()).getName()));
        messageElement.appendChild(dataType);
    }
}
