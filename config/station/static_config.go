package station

import (
	"fmt"
	"os"
	"path/filepath"
)

type StaticConfig struct {
	ProjectFolder string
	AppFolder     string
	DeviceFolder  string
	StaticFolder  string
}

func (u *StaticConfig) Check() {
	var err error
	if u.ProjectFolder == "" {
		u.ProjectFolder, err = os.Getwd()
		if err != nil {
			logger.Fatalf("get cwd failed %s", err)
		}
	}

	if u.AppFolder == "" {
		u.AppFolder = u.GetSubFolder("app")
	}

	if u.DeviceFolder == "" {
		u.DeviceFolder = u.GetSubFolder("device")
	}

	if u.StaticFolder == "" {
		u.DeviceFolder = u.GetSubFolder("static")
	}

}

func (u StaticConfig) String() string {
	return fmt.Sprintf("static:folder=%s", u.ProjectFolder)
}

func (u *StaticConfig) GetSubFolder(folder string) string {
	return filepath.Join(u.ProjectFolder, folder)
}
