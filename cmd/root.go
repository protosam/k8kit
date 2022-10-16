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
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

// Any cli command can be appended to this during init() so apps can just add
// all of the available commands as sub-commands.
var Commands = make([]*cli.Command, 0)

// store the original first element of os.Args
// this is primarily used when "/proc/self/exec" is not available
var realArgZero = os.Args[0]

func rexec(args ...string) error {
	// get file info
	realArgZeroFinfo, err := os.Stat(realArgZero)
	if err != nil {
		return err
	}

	// make sure file is a file instead of some other thing
	if !realArgZeroFinfo.Mode().IsRegular() {
		return fmt.Errorf("path %s is not a regular file", realArgZero)
	}

	// build absolute path
	realArgZero, err := filepath.Abs(realArgZero)
	if err != nil {
		return fmt.Errorf("error when trying to determine absolute path '%s': %s", realArgZero, err)
	}

	// run the command
	self := exec.Command(realArgZero, args...)

	// var stdout, stderr bytes.Buffer
	// self.Stdout = &stdout
	// self.Stderr = &stderr

	// pipe all standard pipes
	self.Stdin = os.Stdin
	self.Stderr = os.Stderr
	self.Stdout = os.Stdout

	// run the command
	return self.Run()
}
