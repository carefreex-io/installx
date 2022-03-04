package main

import (
	"bytes"
	"install/common"
	"log"
	"os"
	"os/exec"
	"runtime"
)

var sysCmd string

func init() {
	if runtime.GOOS == "windows" {
		sysCmd = "sh"
	} else {
		sysCmd = "/bin/sh"
	}
}

func Install(frameworkInfo common.FrameworkInfo, options common.Options) {
	checkList(options.Path)

	cloneFramework(frameworkInfo.Addr, options.Path)

	cleanGit(options.Path)

	replaceCarefreeTarget(frameworkInfo.ReplaceTarget, options.Name, options.Path)
}

func checkList(path string) {
	exists, err := pathExists(path)
	if err != nil {
		log.Fatalf("%v check path exists failed: path=%v, err=%v", common.ErrorStr, path, err)
	}
	if exists {
		log.Fatalf("%v path=%v exists", common.ErrorStr, path)
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func cloneFramework(addr string, path string) {
	arg := "git clone " + addr + " " + path

	execCmd(arg)
}

func cleanGit(path string) {
	arg := "rm -rf " + path + "/.git"

	execCmd(arg)
}

func replaceCarefreeTarget(replaceTarget string, name string, path string) {
	arg := "find " + path + " -name '*.go' -print | xargs perl -pi -e 's|\"" + replaceTarget + "/|\"" + name + "/|g'"

	arg += " && find " + path + " -name 'go.mod' -print | xargs perl -pi -e 's|module " + replaceTarget + "|module " + name + "|g'"

	arg += " && find " + path + " -name 'Makefile' -print | xargs perl -pi -e 's|" + replaceTarget + "|" + name + "|g'"

	execCmd(arg)
}

func execCmd(arg string) {
	cmd := exec.Command(sysCmd, "-c", arg)
	var stdout, stderr bytes.Buffer
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	_, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if err != nil {
		log.Fatalf("%v %v\n", common.ErrorStr, errStr)
	}
}
