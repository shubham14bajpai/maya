// Copyright © 2019 The OpenEBS Authors
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

package command

import (
	"github.com/openebs/maya/cmd/maya-agent/storage/block"
	"github.com/spf13/cobra"
)

//NewSubCmdUnMount unmounts specified mounted disk
func NewSubCmdUnMount() *cobra.Command {
	var disk string
	getCmd := &cobra.Command{
		Use:   "unmount",
		Short: "unmount mounted disk",
		Long:  `specified mounted disk on the storage area network is unmounted`,
		Run: func(cmd *cobra.Command, args []string) {

			block.UnMount(disk)

		},
	}
	getCmd.Flags().StringVar(&disk, "disk", "sdc",
		"disk name")
	return getCmd
}
