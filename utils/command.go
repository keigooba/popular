package utils

import (
	"fmt"
	"os/exec"
)

// コマンドの実行
func Command() error {
	cmd := exec.Command("sh", "start.sh")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	fmt.Println("コマンドが正常に実行されました。")
	return nil
}
