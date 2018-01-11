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
package org.apache.plc4x.java.api.messages.specific;

import org.apache.plc4x.java.api.messages.PlcWriteResponse;
import org.apache.plc4x.java.api.messages.items.WriteResponseItem;

import java.util.List;
import java.util.Optional;

public class TypeSafePlcWriteResponse<T> extends PlcWriteResponse {

    public TypeSafePlcWriteResponse(PlcWriteResponse plcWriteResponse) {
        super(plcWriteResponse.getRequest(), plcWriteResponse.getResponseItems());
    }

    public TypeSafePlcWriteResponse(TypeSafePlcWriteRequest<T> request, WriteResponseItem<T> responseItem) {
        super(request, responseItem);
    }

    @SuppressWarnings("unchecked")
    public TypeSafePlcWriteResponse(TypeSafePlcWriteRequest<T> request, List<WriteResponseItem<T>> responseItems) {
        super(request, responseItems);
    }

    public TypeSafePlcWriteRequest<T> getRequest() {
        return (TypeSafePlcWriteRequest<T>) super.getRequest();
    }

    @SuppressWarnings("unchecked")
    public List<? extends WriteResponseItem<T>> getResponseItems() {
        return (List<WriteResponseItem<T>>) super.getResponseItems();
    }

    @SuppressWarnings("unchecked")
    @Override
    public Optional<WriteResponseItem<T>> getResponseItem() {
        return (Optional<WriteResponseItem<T>>) super.getResponseItem();
    }
}
