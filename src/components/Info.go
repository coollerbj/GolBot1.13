package components

import (
	"os/exec"
	"syscall"
)

func getUserName() string  {
	Info := exec.Command("cmd", "/C", "whoami")
	Info.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}//запускаем онсоль в скрытом режиме
	UserName, err := Info.Output()
	ChekErrors(err)

	return FormatString(string(UserName))
}


