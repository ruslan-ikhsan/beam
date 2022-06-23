/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import 'package:highlight/highlight.dart';
import 'package:highlight/languages/go.dart';
import 'package:highlight/languages/java.dart';
import 'package:highlight/languages/python.dart';
import 'package:highlight/languages/scala.dart';
import 'package:playground/config.g.dart';
import 'package:playground/constants/params.dart';

enum SDK {
  java,
  go,
  python,
  scio,
}

SDK getDefaultSdk() {
  return SDK.values.byName(
    Uri.base.queryParameters[kSdkParam] ?? SDK.java.name,
  );
}

extension SDKToString on SDK {
  String get displayName {
    switch (this) {
      case SDK.go:
        return 'Go';
      case SDK.java:
        return 'Java';
      case SDK.python:
        return 'Python';
      case SDK.scio:
        return 'SCIO';
    }
  }
}

extension SdkToRoute on SDK {
  String get getRoute {
    switch (this) {
      case SDK.java:
        return kApiJavaClientURL;
      case SDK.go:
        return kApiGoClientURL;
      case SDK.python:
        return kApiPythonClientURL;
      case SDK.scio:
        return kApiScioClientURL;
      default:
        return '';
    }
  }
}

extension SdkToHighlightMode on SDK {
  Mode get highlight {
    switch (this) {
      case SDK.java:
        return java;
      case SDK.go:
        return go;
      case SDK.python:
        return python;
      case SDK.scio:
        return scala;
    }
  }
}
