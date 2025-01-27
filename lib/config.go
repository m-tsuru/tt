package lib

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"gopkg.in/ini.v1"
)

type Config struct {
	Core *Core
	User *User
}

func (c *Config) ToMap() {
	panic("unimplemented")
}

type Core struct {
	Mode  *string
	Refer *string
	TZ    *string
}

type User struct {
	Name  *string
	Email *string
	TZ    *string
}

func GetGlobalConfigPath(userArg *string) (error, string) {
	var homeDir string
	if userArg != nil {
		homeDir = *userArg
	} else {
		switch runtime.GOOS {
		case "windows":
			homeDir = os.Getenv("HOMEPATH")
		case "linux", "darwin":
			homeDir = os.Getenv("HOME")
		default:
			err := fmt.Errorf("Unsupported Platform: Please define your home directory.")
			return err, ""
		}
	}

	settingConfPath := homeDir + "/.tt/config"

	_, err := os.Stat(settingConfPath)
	if os.IsNotExist(err) {
		println("info: Global Config is not found. Create New Configuration File.")
		err = nil
		_, err := os.Stat(homeDir + "/.tt")
		if os.IsNotExist(err) {
			os.Mkdir(homeDir+"/.tt", 0755)
		}
		f, err := os.Create(settingConfPath)
		if err != nil {
			return err, settingConfPath
		}
		defer f.Close()
	}

	return nil, settingConfPath
}

func GetGlobalConfig(path *string) (error, *Config) {
	var config Config
	err, globalConfPath := GetGlobalConfigPath(path)
	if err != nil {
		return err, nil
	}
	raw, err := ini.Load(globalConfPath)
	if err == os.ErrNotExist {
		f, err := os.Create(globalConfPath)
		if err != nil {
			return err, nil
		}
		f.Close()
		return nil, &config
	} else if err != nil {
		return err, nil
	}
	config = Config{
		Core: &Core{
			Mode:  ptr(raw.Section("core").Key("mode").String()),
			Refer: ptr(raw.Section("core").Key("refer").String()),
			TZ:    ptr(raw.Section("core").Key("tz").String()),
		},
		User: &User{
			Name:  ptr(raw.Section("user").Key("name").String()),
			Email: ptr(raw.Section("user").Key("email").String()),
			TZ:    ptr(raw.Section("user").Key("tz").String()),
		},
	}
	return nil, &config
}

func GetConfig(key string, path string) (error, *string) {
	parts := strings.Split(key, ".")
	if len(parts) == 1 {
		err := fmt.Errorf("key does not contain a section: %s", parts[0])
		return err, nil
	} else if len(parts) > 3 || len(parts) < 0 {
		err := fmt.Errorf("key is not valid.")
		return err, nil
	}
	raw, err := ini.Load(path)
	if err != nil {
		return err, nil
	}
	res := ptr(raw.Section(parts[0]).Key(parts[1]).String())
	return nil, res
}

func SetConfig(key string, value *string, path string) (error, *string) {
	parts := strings.Split(key, ".")
	if len(parts) == 1 {
		err := fmt.Errorf("key does not contain a section: %s", parts[0])
		return err, nil
	} else if len(parts) > 3 || len(parts) < 0 {
		err := fmt.Errorf("key is not valid.")
		return err, nil
	}
	raw, err := ini.Load(path)
	raw.Section(parts[0]).Key(parts[1]).SetValue(*value)
	err = raw.SaveTo(path)
	if err != nil {
		return err, nil
	}
	return nil, value
}

func EditGlobalConfig(path string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return fmt.Errorf("Environment Variable EDITOR is not set.\n")
	}
	cmd := exec.Command(editor, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func ptr(s string) *string {
	return &s
}

func init() {}
