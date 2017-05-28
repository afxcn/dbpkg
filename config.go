/**

Copyright (C) 2017  ZhiQiang Huang (email: afxcn@dbpkg.com)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

**/

package main

import (
	"os"
	"os/user"
	"path/filepath"
)

var (
	configDir string
)

func init() {
	configDir = os.Getenv("NGX_CONFIG")

	if configDir == "" {
		if u, err := user.Current(); err == nil {
			configDir = filepath.Join(u.HomeDir, ".config", "dbpkg")
		}
	}
}
