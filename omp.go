package omp

import (
	"fmt"
	"strings"
)

// OMP .
type OMP interface {
	// Target related
	CreateTarget(t *Target) (string, error)
	ModifyTarget(t *Target) error
	GetTargets() ([]OMPtarget, error)
	DeleteTarget(id string) error

	// configs
	CreateConfig(id string, name string) (string, error)
	GetConfigs() ([]Config, error)

	// tasks
	CreateTask(t *Task) (string, error)
	GetTasks(id string) ([]Task, error)
	DeleteTask(id string) error
	ModifyTask(t *Task) error
	StartTask(id string) (string, error)
	StopTask(t *Task) error

	// scanners
	CreateScanner(s *Scanner, credID string) (string, error)
	GetScanners() ([]Scanner, error)

	// credentials
	CreateCredential(cred Credential) (string, error)
	GetCredentials( /*cred Credential*/ ) ([]Credential, error)

	// result
	GetResults(r *Result, taskID string, filter ...string) ([]Result, error)
}

// Version 版本类型
type Version int

// Version .
const (
	Version4 Version = 4
	Version7         = 7
)

// New .
func New(addr, username, password string) (OMP, error) {
	c, err := newConnector(addr)
	if err != nil {
		return nil, err
	}
	err = c.Auth(username, password)
	if err != nil {
		return nil, err
	}
	v, err := c.GetVersion()
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(v, "7.0") {
		return &OMPv7{c}, nil
	}
	return nil, fmt.Errorf("server version is %s, not supported yet", v)
}