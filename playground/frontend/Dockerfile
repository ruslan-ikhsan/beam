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

ARG BUILDER=debian:11
FROM $BUILDER as build

COPY . /app
WORKDIR /app

RUN cd /app/playground_components &&\
    flutter pub get -v &&\
    flutter pub run build_runner build -v

RUN cd /app &&\
    flutter pub get -v &&\
    flutter pub run build_runner build -v &&\
    flutter build web -v

FROM nginx:1.21.3
COPY --from=build /app/nginx_default.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/build/web/ /usr/share/nginx/html
RUN cp -r /usr/share/nginx/html/assets/assets/* /usr/share/nginx/html/assets/
RUN gzip -kr /usr/share/nginx/html/assets/*
