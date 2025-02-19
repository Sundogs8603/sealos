/*
Copyright 2022 cuisongliu@qq.com.

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

package kubernetes

import "path"

const copyKubeAdminConfigCommand = `rm -rf $HOME/.kube/config && mkdir -p $HOME/.kube && cp /etc/kubernetes/admin.conf $HOME/.kube/config`

func (k *KubeadmRuntime) copyKubeConfigFileToNodes(hosts []string) error {
	src := k.getContentData().AdminFile()
	dst := path.Join(".kube", "config")
	return k.sendFileToHosts(hosts, src, dst)
}

func (k *KubeadmRuntime) copyMasterKubeConfig(host string) error {
	return k.sshCmdAsync(host, copyKubeAdminConfigCommand)
}
