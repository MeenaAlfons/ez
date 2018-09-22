// +build windows

package utils

//
var Shell string = "cmd"

//
var ShellFlag string = "/c"

//
func GetEvalCmd(command string) string {
	// return "@FOR /f \"tokens=* delims=^L\" %i IN ('" + command + "') DO %i"
	return "@FOR /f \"tokens=*\" %i IN ('" + command + "') DO @%i"
}
