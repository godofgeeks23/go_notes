package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func run() {
	fmt.Printf("[run()] running: %v (PID:%v)\n", os.Args[2:], os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// this will create a new namespace for the process of its own so that it can't see what going on in the parent host
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("[run()] error:", err)
		panic(err)
	}
}

func child() {
	fmt.Printf("[child()] running %v (PID:%v)\n", os.Args[2:], os.Getpid())

	syscall.Sethostname([]byte("container"))
	syscall.Chroot("container-fs/filesystem")
	syscall.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("[child()] error:", err)
		panic(err)
	}

	defer syscall.Unmount("/proc", 0)
}

func main() {
	fmt.Println("container demo")
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("bad command")
	}
}
