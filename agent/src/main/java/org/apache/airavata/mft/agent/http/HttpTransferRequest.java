/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.apache.airavata.mft.agent.http;

import org.apache.airavata.mft.common.AuthToken;
import org.apache.airavata.mft.core.api.Connector;
import org.apache.airavata.mft.core.api.MetadataCollector;

public class HttpTransferRequest {
    private Connector otherConnector;
    private MetadataCollector otherMetadataCollector;
    private ConnectorParams connectorParams;
    private String resourceId;
    private String childResourcePath;
    private String credentialToken;
    private long createdTime = System.currentTimeMillis();
    private AuthToken authToken;

    public Connector getOtherConnector() {
        return otherConnector;
    }

    public HttpTransferRequest setOtherConnector(Connector otherConnector) {
        this.otherConnector = otherConnector;
        return this;
    }

    public MetadataCollector getOtherMetadataCollector() {
        return otherMetadataCollector;
    }

    public HttpTransferRequest setOtherMetadataCollector(MetadataCollector otherMetadataCollector) {
        this.otherMetadataCollector = otherMetadataCollector;
        return this;
    }

    public String getResourceId() {
        return resourceId;
    }

    public HttpTransferRequest setResourceId(String resourceId) {
        this.resourceId = resourceId;
        return this;
    }

    public String getChildResourcePath() {
        return childResourcePath;
    }

    public HttpTransferRequest setChildResourcePath(String childResourcePath) {
        this.childResourcePath = childResourcePath;
        return this;
    }

    public String getCredentialToken() {
        return credentialToken;
    }

    public HttpTransferRequest setCredentialToken(String credentialToken) {
        this.credentialToken = credentialToken;
        return this;
    }

    public ConnectorParams getConnectorParams() {
        return connectorParams;
    }

    public HttpTransferRequest setConnectorParams(ConnectorParams connectorParams) {
        this.connectorParams = connectorParams;
        return this;
    }

    public long getCreatedTime() {
        return createdTime;
    }

    public HttpTransferRequest setCreatedTime(long createdTime) {
        this.createdTime = createdTime;
        return this;
    }

    public AuthToken getAuthToken() {
        return authToken;
    }

    public HttpTransferRequest setAuthToken(AuthToken authToken) {
        this.authToken = authToken;
        return this;
    }
}
