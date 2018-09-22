package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

//
func RunCommandWithStdout(command string, args ...string) bool {
	cmd := exec.Command(command, args...)
	return runCmdWithStdout(cmd)
}

//
func RunInShellWithStdout(command string) bool {
	cmd := exec.Command(Shell, ShellFlag, command)
	return runCmdWithStdout(cmd)
}

//
func RunInShellWithStdoutAndStdin(command, data string) bool {
	cmd := exec.Command(Shell, ShellFlag, command)
	cmd.Stdin = bytes.NewBuffer([]byte(data))
	return runCmdWithStdout(cmd)
}

//
func RunInShellWithStdoutAndCapture(command string) (bool, string) {
	cmd := exec.Command(Shell, ShellFlag, command)
	return runCmdWithStdoutAndCapture(cmd)
}

func forwardStream(w io.Writer, r io.Reader) error {
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			_, err := w.Write(buf[:n])
			if err != nil {
				return err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return err
		}
	}
	// never reached
	panic(true)
	return nil
}

//
func forwardAndCaptureStream(w io.Writer, r io.Reader) (string, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		d := buf[:n]
		out = append(out, d...)
		if n > 0 {
			_, err := w.Write(d)
			if err != nil {
				return string(out), err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return string(out), err
		}
	}
	// never reached
	panic(true)
	return "", nil
}

func runCmdWithStdout(cmd *exec.Cmd) bool {
	stdoutIn, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating Stdout pipe:", err)
		return false
	}

	var errStdout error
	go func() {
		errStdout = forwardStream(os.Stdout, stdoutIn)
	}()

	ok := run(cmd)

	if errStdout != nil {
		fmt.Println("Failure in stdout:", errStdout)
	}

	if !ok || errStdout != nil {
		return false
	}

	return true
}

func runCmdWithStdoutAndCapture(cmd *exec.Cmd) (bool, string) {
	stdoutIn, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating Stdout pipe:", err)
		return false, ""
	}

	var errStdout error
	var output string
	go func() {
		output, errStdout = forwardAndCaptureStream(os.Stdout, stdoutIn)
	}()

	ok := run(cmd)

	if errStdout != nil {
		fmt.Println("Failure in stdout:", errStdout)
	}

	if !ok || errStdout != nil {
		return false, output
	}

	return true, output
}

func run(cmd *exec.Cmd) bool {
	if err := cmd.Start(); err != nil {
		fmt.Println("Cannot Start Cmd:", err)
		return false
	}

	err := cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for command:", err)
	}

	if err != nil {
		return false
	}

	return true
}
