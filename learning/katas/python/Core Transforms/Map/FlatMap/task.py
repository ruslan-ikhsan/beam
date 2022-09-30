#   Licensed to the Apache Software Foundation (ASF) under one
#   or more contributor license agreements.  See the NOTICE file
#   distributed with this work for additional information
#   regarding copyright ownership.  The ASF licenses this file
#   to you under the Apache License, Version 2.0 (the
#   "License"); you may not use this file except in compliance
#   with the License.  You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

# beam-playground:
#   name: FlatMap
#   description: Task from katas to implement a function that maps each input sentence
#     into words split by whitespace (" ").
#   multifile: false
#   context_line: 30
#   categories:
#     - Core Transforms
#   complexity: BASIC

import apache_beam as beam

from log_elements import LogElements

with beam.Pipeline() as p:

  (p | beam.Create(['Apache Beam', 'Unified Batch and Streaming'])
     | beam.FlatMap(lambda sentence: sentence.split())
     | LogElements())
