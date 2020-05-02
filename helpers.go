package svcw

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	modTmp       string
	moduleRegexp = regexp.MustCompile("module (.+)")
)

func getGoModDir() (modPath string) {
	mod := getGoModFilePath()
	modPath, _ = filepath.Split(mod)
	return
}

func getGoModFilePath() (modPath string) {
	if len(modTmp) > 0 {
		return modTmp
	}
	cmd := exec.Command("go", "env", "GOMOD")
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	_ = cmd.Run()
	mod := stdout.String()
	mod = strings.Trim(mod, "\n")
	modTmp = mod
	return mod
}

func getImportPath(importPath string) (servicePkgPath string) {
	mod := getGoModFilePath()
	m, _ := ioutil.ReadFile(mod)
	f := moduleRegexp.FindStringSubmatch(strings.Split(string(m), "\n")[0])
	if len(f) == 2 {
		modP := f[1]
		servicePkgPath = strings.Replace(importPath, modP, filepath.Dir(filepath.ToSlash(mod)), -1)
	}
	return
}
