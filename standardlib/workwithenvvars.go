package standardlib

import (
	"os"
	"strings"
)

type EnvVarMap map[string]string

func EnvVarToMap() EnvVarMap {
	envs := os.Environ()
	m := make(map[string]string)

	for _, v := range envs {
		env_name := strings.Split(v, "=")[0]
		env_val := strings.Split(v, "=")[1]
		m[env_name] = env_val
	}

	return m
}
