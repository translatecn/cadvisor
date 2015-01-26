// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"

	"github.com/google/cadvisor/utils/cpuload/netlink"
)

func main() {
	n, err := netlink.New()
	if err != nil {
		log.Printf("Failed to create cpu load util: %s", err)
		return
	}
	defer n.Stop()

	paths := []string{"/sys/fs/cgroup/cpu", "/sys/fs/cgroup/cpu/docker"}
	names := []string{"/", "/docker"}
	for i, path := range paths {
		stats, err := n.GetCpuLoad(path, names[i])
		if err != nil {
			log.Printf("Error getting cpu load for %q: %s", path, err)
		}
		log.Printf("Task load for %s: %+v", path, stats)
	}
}