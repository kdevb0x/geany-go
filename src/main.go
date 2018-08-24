/* Copyright 2018  Kris Davidson (KiDoDa)
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main // import "github.com/kidoda/geany-go"

import (
	"log"
)

func main() {
	err := PluginSetInfo(pluginInfo)
	if err != nil {
		log.Fatalf("failed to set plugin info: %s", err)
	}
}
