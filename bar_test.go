// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package i3ipc

import (
	"testing"
)

func TestGetBarConfig(t *testing.T) {
	ipc, _ := GetIPCSocket()

	go startTestIPCSocket(testMessages["bar"])

	ids, err := ipc.GetBarIds()
	if err != nil {
		t.Errorf("Getting bar IDs failed: %v", err)
	}

	id := ids[0]
	//bar, err := GetBarConfig(ipc)
	_, err = ipc.GetBarConfig(id)
	if err != nil {
		t.Errorf("Getting bar config failed: %v", err)
	}
}
