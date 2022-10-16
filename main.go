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

package main

import (
	"os"

	"github.com/protosam/k8kit/cmd"
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:    "k8kit",
	Version: cmd.VERSION,
	Authors: []*cli.Author{
		{
			Name:  "protosam",
			Email: "github.com/protosam",
		},
	},
	Copyright: "(c) github.com/protosam/k8kit",
	HelpName:  "k8kit",
	Usage:     "A vendored selection of cluster tools.",
	Commands:  cmd.Commands,
}

func main() {
	App.Run(os.Args)
}
