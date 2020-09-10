package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func BuildVue() {
	dir := filepath.Dir("/Users/nguyentaidat/Documents/GOD/Project/builder/vue-test/")
	fmt.Println(dir)
	goExecutable, _ := exec.LookPath("yarn")
	cmdBuildVuePrj := &exec.Cmd{
		Path:   goExecutable,
		Dir:    dir,
		Args:   []string{dir, "build"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmdBuildVuePrj.Run(); err != nil {
		fmt.Println("Error", err)
	}
}
