###############################################################################
#  Licensed to the Apache Software Foundation (ASF) under one
#  or more contributor license agreements.  See the NOTICE file
#  distributed with this work for additional information
#  regarding copyright ownership.  The ASF licenses this file
#  to you under the Apache License, Version 2.0 (the
#  "License"); you may not use this file except in compliance
#  with the License.  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
# limitations under the License.
###############################################################################
#Dokerfile to set up the Beam Go SDK
ARG BASE_IMAGE=golang:1.18-bullseye
FROM $BASE_IMAGE

ARG BEAM_VERSION=2.40.0
ENV PREPARED_MOD_DIR=/opt/playground/prepared_folder/
ENV PIPELINES_FOLDER_NAME="executable_files"
RUN mkdir -p /opt/playground/ $PREPARED_MOD_DIR
WORKDIR $PREPARED_MOD_DIR

RUN go mod init $PIPELINES_FOLDER_NAME &&\
    go get github.com/apache/beam/sdks/v2/go/pkg/beam@v$BEAM_VERSION &&\
    go get github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/gcs@v$BEAM_VERSION &&\
    go mod download all

WORKDIR /server

COPY playground/backend .

RUN go build -o /server/app ./cmd

COPY configs /opt/playground/backend/configs/

COPY properties.yaml /opt/playground/backend/properties.yaml

ENTRYPOINT ["/server/app"]
