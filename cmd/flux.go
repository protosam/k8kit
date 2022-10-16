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

	fluxcmd "github.com/fluxcd/flux2/cmd/flux"
)

func init() {
	Commands = append(Commands, FluxCmd)
}

var FluxCmd = &cli.Command{
	Name:        "flux",
	Aliases:     []string{},
	Usage:       "utility to manage Kubernetes CD pipelines with GitOps",
	Description: `utility to manage Kubernetes CD pipelines with GitOps`,
	ArgsUsage:   "",
	Action:      FluxMain,
}

func FluxMain(c *cli.Context) error {
	// mutate args so it's just kubectl
	os.Args = append([]string{"flux"}, c.Args().Slice()...)

	flux_cli := fluxcmd.NewDefaultFluxCommand()

	if err := flux_cli.Execute(); err != nil {
		ec := fluxcmd.GetExitStatus(err)
		os.Exit(ec)
	}
	return nil
}
