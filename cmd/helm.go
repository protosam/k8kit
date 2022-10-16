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
	"log"
	"os"

	"github.com/urfave/cli/v2"

	helmcmd "helm.sh/helm/v3/cmd/helm"
)

func init() {
	Commands = append(Commands, HelmCmd)
}

var HelmCmd = &cli.Command{
	Name:        "helm",
	Aliases:     []string{},
	Usage:       "Kubernetes package manager",
	Description: `Kubernetes package manager`,
	ArgsUsage:   "",
	Action:      HelmMain,
}

func HelmMain(c *cli.Context) error {
	// mutate args so it's just kubectl
	os.Args = append([]string{"helm"}, c.Args().Slice()...)

	helm_cli, err := helmcmd.NewDefaultHelmCommand()
	if err != nil {
		log.Fatal(err)
	}

	if err := helm_cli.Execute(); err != nil {
		helmcmd.Debug("%+v", err)
		ec := helmcmd.GetExitCode(err)
		os.Exit(ec)
	}
	return nil
}
