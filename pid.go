package pidfile

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

var (
	ErrPidExists = errors.New("pid file exists.")
	ErrLockFaild = errors.New("pid file lock faild")
	Debug        bool
)

func CheckExit(filename string) {
	pid, pid_err := GetValue(filename)
	if pid_err == nil {
		b, _ := IsActive(pid)
		if b {
			fmt.Printf("%s", "already running")
			os.Exit(1)
		}
	}
	pidValue, err := Create(filename)
	if err != nil {
		fmt.Printf("create pid:%s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("process pid:%d\n", pidValue)
}

func CheckLockFile(fd uintptr) error {
	err := syscall.Flock(int(fd), syscall.LOCK_EX|syscall.LOCK_NB)
	if err == syscall.EWOULDBLOCK {
		err = ErrLockFaild
	}
	return err
}

func Create(pidfile string) (int, error) {
	if _, err := os.Stat(pidfile); !os.IsNotExist(err) {
		// file exists
		if pid, _ := GetValue(pidfile); pid > 0 {
			if ok, _ := IsActive(pid); ok {
				return pid, ErrPidExists
			}
		}
	}

	if pf, err := os.OpenFile(pidfile, os.O_RDWR|os.O_CREATE, 0600); err != nil {
		return 0, errors.New(fmt.Sprintf("Error when create pid file: %s\n", err.Error()))
	} else {
		pid := os.Getpid()
		pf.Write([]byte(strconv.Itoa(pid)))
		if lock_err := CheckLockFile(pf.Fd()); lock_err != nil {
			return pid, lock_err
		}
		return pid, nil
	}
}

func IsActive(pid int) (bool, error) {
	if pid <= 0 {
		return false, errors.New("process id error.")
	}
	p, err := os.FindProcess(pid)
	if err != nil {
		if Debug {
			fmt.Printf("find process: %s\n", err.Error())
		}
		return false, err
	}

	if err := p.Signal(os.Signal(syscall.Signal(0))); err != nil {
		if Debug {
			fmt.Printf("send signal [0]: %s\n", err.Error())
		}
		return false, err
	}

	return true, nil
}

func GetValue(pidfile string) (int, error) {
	value, err := ioutil.ReadFile(pidfile)
	if err != nil {
		if Debug {
			fmt.Printf("read pid file: %s\n", err.Error())
		}
		return 0, err
	}
	pid, err := strconv.ParseInt(string(value), 10, 32)
	if err != nil {
		if Debug {
			fmt.Printf("trans pid to int: %s\n", err.Error())
		}
		return 0, err
	}
	return int(pid), nil
}

