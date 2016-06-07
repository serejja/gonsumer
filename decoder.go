/* Licensed to the Elodina Inc. under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

package gonsumer

// Decoder serves to decode given raw message bytes to something meaningful.
// Returns an error if fails to decoder given bytes.
type Decoder func(raw []byte) (interface{}, error)

// ByteDecoder is a default decoder implementation that does nothing and just returns the input untouched.
// Never returns an error.
var ByteDecoder = func(bytes []byte) (interface{}, error) {
	return bytes, nil
}

// StringDecoder converts the given bytes into a string. Never returns an error.
var StringDecoder = func(bytes []byte) (interface{}, error) {
	return string(bytes), nil
}
