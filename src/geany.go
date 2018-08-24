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

/*
#cgo LDFLAGS: -lgeany
#include <stdlib.h>
#include <geany/geanyplugin.h>
#include <geany/plugindata.h>
*/
import "C"

import (
	"unsafe"
)

type PluginInfo struct {
	Pname        string
	Pdescription string
	Pversion     string
	Pauthor      string
}

const pluginInfo = &PluginInfo{
	Pname:        "gocode",
	Pdescription: "Geany integration of gocode daemon for Golang autocompletion.",
	Pversion:     "0.0.1",
	Pauthor:      "Kristofer D. Davidson (KiDoDa)",
}

// cgo call to void plugin_set_info()
func PluginSetInfo(info *PluginInfo) error {
	cname := C.CString(info.Pname)
	cdesc := C.CString(info.Pdescription)
	cver := C.CString(info.Pversion)
	cauth := C.CString(info.Pauthor)
	defer func() {
		C.Free(unsafe.Pointer(cname))
		C.Free(unsafe.Pointer(cdesc))
		C.Free(unsafe.Pointer(cver))
		C.Free(unsafe.Pointer(cauth))
	}()
	setInfo := C.plugin_set_info(cname, cdesc, cver, cauth)
	return nil
}
