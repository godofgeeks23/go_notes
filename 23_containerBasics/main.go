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

	cg()

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

func cg() {
	cgroupName := "avi"
	cgroups := "/sys/fs/cgroup/"
	myCgroupPath := filepath.Join(cgroups, cgroupName)

	// create the directory - the kernel automatically detects this mkdir and creates the cgroup structure populated with files
	err := os.Mkdir(myCgroupPath, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	// 2. Set the limit - the file is named 'pids.max' and sits directly in your folder
	must(os.WriteFile(filepath.Join(myCgroupPath, "pids.max"), []byte("20"), 0700))

	// 3. Notify on release - removes the new cgroup in place after the container exits
	must(os.WriteFile(filepath.Join(myCgroupPath, "notify_on_release"), []byte("1"), 0700))

	// 4. Add the current process - add this process to be controlled by this cgroup
	must(os.WriteFile(filepath.Join(myCgroupPath, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))

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

func must(err error) {
	if err != nil {
		panic(err)
	}
}
