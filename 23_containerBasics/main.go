package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

func run() {
	fmt.Printf("[run()] running: %v (PID:%v)\n", os.Args[2:], os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	if err := cmd.Run(); err != nil {
		fmt.Println("[run()] error:", err)
		panic(err)
	}
}

func child() {
	fmt.Printf("[child()] running %v (PID:%v)\n", os.Args[2:], os.Getpid())
	cg()

	syscall.Sethostname([]byte("container"))
	syscall.Chroot("container-fs")
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

func cg() {
	cgroupName := "mytestcgroup"
	cgroups := "/sys/fs/cgroup/"
	myCgroupPath := filepath.Join(cgroups, cgroupName)

	err := os.Mkdir(myCgroupPath, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	must(os.WriteFile(filepath.Join(myCgroupPath, "pids.max"), []byte("10"), 0700))

	must(os.WriteFile(filepath.Join(myCgroupPath, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))

}

func must(err error) {
	if err != nil {
		panic(err)
	}
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
