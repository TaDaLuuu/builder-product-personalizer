package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func StartVue() {
	dir := filepath.Dir("/Users/nguyentaidat/Documents/GOD/Project/builder/vue-test/")
	fmt.Println(dir)
	goExecutable, _ := exec.LookPath("yarn")
	cmdStartVuePrj := &exec.Cmd{
		Path:   goExecutable,
		Dir:    dir,
		Args:   []string{dir, "serve"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := cmdStartVuePrj.Run(); err != nil {
		fmt.Println("Error", err)
	}
}
