/*
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

package cmd

import (
	"os"

	"github.com/urfave/cli/v2"

	kubectl_cli "k8s.io/component-base/cli"
	kubectl_cmd "k8s.io/kubectl/pkg/cmd"
	"k8s.io/kubectl/pkg/cmd/util"

	// Import to initialize client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func init() {
	Commands = append(Commands, KubectlCmd)
}

var KubectlCmd = &cli.Command{
	Name:        "kubectl",
	Aliases:     []string{},
	Usage:       "controls the Kubernetes cluster manager",
	Description: `controls the Kubernetes cluster manager`,
	ArgsUsage:   "",
	Action:      KubectlMain,
}

func KubectlMain(c *cli.Context) error {
	// mutate args so it's just kubectl
	os.Args = append([]string{"kubectl"}, c.Args().Slice()...)

	// run and handle kubectl
	if err := kubectl_cli.RunNoErrOutput(kubectl_cmd.NewDefaultKubectlCommand()); err != nil {
		// Pretty-print the error and exit with an error.
		util.CheckErr(err)
	}

	return nil
}