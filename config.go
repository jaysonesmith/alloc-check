package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/spf13/viper"
)

func ViperDoubleDip() error {
	projectName := viper.GetString("PROJECT_NAME")
	instanceName := viper.GetString("INSTANCE_NAME")
	sqlName := viper.GetString("SQL_NAME")

	var missingVars []string
	for _, k := range []string{"PROJECT_NAME", "INSTANCE_NAME", "SQL_NAME"} {
		if viper.GetString(k) == "" {
			missingVars = append(missingVars, k)
		}
	}

	if len(missingVars) != 0 {
		return fmt.Errorf("%v required", missingVars)
	}

	_ = projectName
	_ = instanceName
	_ = sqlName

	return nil
}

func OSDoubleDip() error {
	projectName := os.Getenv("PROJECT_NAME")
	instanceName := os.Getenv("INSTANCE_NAME")
	sqlName := os.Getenv("SQL_NAME")

	var missingVars []string
	for _, k := range []string{"PROJECT_NAME", "INSTANCE_NAME", "SQL_NAME"} {
		if os.Getenv(k) == "" {
			missingVars = append(missingVars, k)
		}
	}

	if len(missingVars) != 0 {
		return fmt.Errorf("%v required", missingVars)
	}

	_ = projectName
	_ = instanceName
	_ = sqlName

	return nil
}

func ViperSingleDip() error {
	var projectName, instanceName, sqlName string
	var envBindings = map[string]func(in string){
		"PROJECT_NAME":  func(in string) { projectName = in },
		"INSTANCE_NAME": func(in string) { instanceName = in },
		"SQL_NAME":      func(in string) { sqlName = in },
	}

	missingVars := make([]string, 0, len(envBindings))
	for k, v := range envBindings {
		found := viper.GetString(k)
		if found != "" {
			v(found)
			continue
		}
		missingVars = append(missingVars, k)
	}

	if len(missingVars) != 0 {
		return fmt.Errorf("%v required", missingVars)
	}

	_ = projectName
	_ = instanceName
	_ = sqlName

	return nil
}

func OSSingleDip() error {
	var projectName, instanceName, sqlName string
	var envBindings = map[string]func(in string){
		"PROJECT_NAME":  func(in string) { projectName = in },
		"INSTANCE_NAME": func(in string) { instanceName = in },
		"SQL_NAME":      func(in string) { sqlName = in },
	}

	missingVars := make([]string, 0, len(envBindings))
	for k, v := range envBindings {
		found := os.Getenv(k)
		if found != "" {
			v(found)
			continue
		}
		missingVars = append(missingVars, k)
	}

	if len(missingVars) != 0 {
		return fmt.Errorf("%v required", missingVars)
	}

	_ = projectName
	_ = instanceName
	_ = sqlName

	return nil
}

func ViperSingleDipStruct() error {
	type gcpConfig struct {
		projectName  string
		instanceName string
		sqlName      string
	}

	cfg := &gcpConfig{
		projectName:  viper.GetString("PROJECT_NAME"),
		instanceName: viper.GetString("INSTANCE_NAME"),
		sqlName:      viper.GetString("SQL_NAME"),
	}

	fmt.Printf("cfg: %+v\n", cfg)
	fmt.Printf("cfg size: %+v\n", reflect.TypeOf(cfg).Size())

	return nil
}

func ViperGetBool() bool {
	set := viper.IsSet("TEST_BOOL")
	if !set {
		return false
	}
	return viper.GetBool("TEST_BOOL")
}

func OSGetBool() bool {
	env, ok := os.LookupEnv("TEST_BOOL")
	if !ok {
		return false
	}

	b, err := strconv.ParseBool(env)
	if err != nil {
		return false
	}

	return b
}
