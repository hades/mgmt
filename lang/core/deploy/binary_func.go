// Mgmt
// Copyright (C) 2013-2024+ James Shubin and the project contributors
// Written by James Shubin <james@shubin.ca> and the project contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//
// Additional permission under GNU GPL version 3 section 7
//
// If you modify this program, or any covered work, by linking or combining it
// with embedded mcl code and modules (and that the embedded mcl code and
// modules which link with this program, contain a copy of their source code in
// the authoritative form) containing parts covered by the terms of any other
// license, the licensors of this program grant you additional permission to
// convey the resulting work. Furthermore, the licensors of this program grant
// the original author, James Shubin, additional permission to update this
// additional permission if he deems it necessary to achieve the goals of this
// additional permission.

package coredeploy

import (
	"context"
	"os"
	"path/filepath"

	"github.com/purpleidea/mgmt/lang/funcs/simple"
	"github.com/purpleidea/mgmt/lang/types"
)

func init() {
	simple.ModuleRegister(ModuleName, "binary_path", &simple.Scaffold{
		T: types.NewType("func() str"),
		F: BinaryPath,
	})
}

// BinaryPath returns the path to the binary of this program. This is useful for
// bootstrapping new machines when we want to get the path to copy it over from.
func BinaryPath(ctx context.Context, input []types.Value) (types.Value, error) {
	p1, err := os.Executable()
	if err != nil {
		return nil, err
	}

	p2, err := filepath.EvalSymlinks(p1)
	if err != nil {
		return nil, err
	}

	p3, err := filepath.Abs(p2)
	if err != nil {
		return nil, err
	}

	return &types.StrValue{
		V: p3,
	}, nil
}
