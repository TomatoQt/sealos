// Copyright © 2021 sealos.
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

package cmd

import (
	"fmt"
	"runtime"

	"github.com/labring/sealos/pkg/image/types"

	"github.com/spf13/cobra"

	"github.com/labring/sealos/pkg/image"
)

func NewPullCmd() *cobra.Command {
	var platform string
	var pullCmd = &cobra.Command{
		Use:     "pull",
		Short:   "pull cloud image",
		Example: `sealos pull labring/kubernetes:v1.24.0`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			registrySvc, err := image.NewRegistryService()
			if err != nil {
				return err
			}
			return registrySvc.Pull(types.ParsePlatform(platform), args[0])
		},
	}
	pullCmd.Flags().StringVar(&platform, "platform", fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH), "set the OS/ARCH/VARIANT of the image to the provided value instead of the current operating system and architecture of the host (for example linux/arm)")
	return pullCmd
}
