/*
Copyright 2021 Flant CJSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hooks

import (
	"fmt"
	"io/ioutil"

	"github.com/flant/addon-operator/pkg/module_manager/go_hook"
	"github.com/flant/addon-operator/sdk"
)

var _ = sdk.RegisterFunc(&go_hook.HookConfig{
	OnStartup: &go_hook.OrderedConfig{Order: 5},
}, discoverApiserverCA)

func discoverApiserverCA(input *go_hook.HookInput) error {
	caPath := "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"

	content, err := ioutil.ReadFile(caPath)
	if err != nil {
		return fmt.Errorf("cannot find kubernetes ca: %v, (not in pod?)", err)
	}

	input.Values.Set("global.discovery.kubernetesCA", string(content))
	return nil
}
