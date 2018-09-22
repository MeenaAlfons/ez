// +build linux

package utils

//
var Shell string = "/bin/sh"

//
var ShellFlag string = "-c"

//
func GetEvalCmd(command string) string {
	return "eval $(" + command + ")"
}
