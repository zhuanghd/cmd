package cmd

import (
	"log"
	"os/exec"
	"strings"
	"syscall"
)

// BashExec 执行 bash 命令返回执行结果
func BashExec(cmdStr string) string {
	cmd := exec.Command("bash", "-c", cmdStr)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// HideWindowExec 隐藏窗口执行
func HideWindowExec(cmdStr string, arg ...string) string {
	cmd := exec.Command(cmdStr, arg...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return ""
	}
	return strings.TrimSpace(string(out))
}
