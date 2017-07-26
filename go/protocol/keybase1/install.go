// Auto-generated by avdl-compiler v1.3.17 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/install.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

// Install status describes state of install for a component or service.
type InstallStatus int

const (
	InstallStatus_UNKNOWN       InstallStatus = 0
	InstallStatus_ERROR         InstallStatus = 1
	InstallStatus_NOT_INSTALLED InstallStatus = 2
	InstallStatus_INSTALLED     InstallStatus = 4
)

func (o InstallStatus) DeepCopy() InstallStatus { return o }

var InstallStatusMap = map[string]InstallStatus{
	"UNKNOWN":       0,
	"ERROR":         1,
	"NOT_INSTALLED": 2,
	"INSTALLED":     4,
}

var InstallStatusRevMap = map[InstallStatus]string{
	0: "UNKNOWN",
	1: "ERROR",
	2: "NOT_INSTALLED",
	4: "INSTALLED",
}

type InstallAction int

const (
	InstallAction_UNKNOWN   InstallAction = 0
	InstallAction_NONE      InstallAction = 1
	InstallAction_UPGRADE   InstallAction = 2
	InstallAction_REINSTALL InstallAction = 3
	InstallAction_INSTALL   InstallAction = 4
)

func (o InstallAction) DeepCopy() InstallAction { return o }

var InstallActionMap = map[string]InstallAction{
	"UNKNOWN":   0,
	"NONE":      1,
	"UPGRADE":   2,
	"REINSTALL": 3,
	"INSTALL":   4,
}

var InstallActionRevMap = map[InstallAction]string{
	0: "UNKNOWN",
	1: "NONE",
	2: "UPGRADE",
	3: "REINSTALL",
	4: "INSTALL",
}

type ServiceStatus struct {
	Version        string        `codec:"version" json:"version"`
	Label          string        `codec:"label" json:"label"`
	Pid            string        `codec:"pid" json:"pid"`
	LastExitStatus string        `codec:"lastExitStatus" json:"lastExitStatus"`
	BundleVersion  string        `codec:"bundleVersion" json:"bundleVersion"`
	InstallStatus  InstallStatus `codec:"installStatus" json:"installStatus"`
	InstallAction  InstallAction `codec:"installAction" json:"installAction"`
	Status         Status        `codec:"status" json:"status"`
}

func (o ServiceStatus) DeepCopy() ServiceStatus {
	return ServiceStatus{
		Version:        o.Version,
		Label:          o.Label,
		Pid:            o.Pid,
		LastExitStatus: o.LastExitStatus,
		BundleVersion:  o.BundleVersion,
		InstallStatus:  o.InstallStatus.DeepCopy(),
		InstallAction:  o.InstallAction.DeepCopy(),
		Status:         o.Status.DeepCopy(),
	}
}

type ServicesStatus struct {
	Service []ServiceStatus `codec:"service" json:"service"`
	Kbfs    []ServiceStatus `codec:"kbfs" json:"kbfs"`
	Updater []ServiceStatus `codec:"updater" json:"updater"`
}

func (o ServicesStatus) DeepCopy() ServicesStatus {
	return ServicesStatus{
		Service: (func(x []ServiceStatus) []ServiceStatus {
			var ret []ServiceStatus
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Service),
		Kbfs: (func(x []ServiceStatus) []ServiceStatus {
			var ret []ServiceStatus
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Kbfs),
		Updater: (func(x []ServiceStatus) []ServiceStatus {
			var ret []ServiceStatus
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Updater),
	}
}

type FuseMountInfo struct {
	Path   string `codec:"path" json:"path"`
	Fstype string `codec:"fstype" json:"fstype"`
	Output string `codec:"output" json:"output"`
}

func (o FuseMountInfo) DeepCopy() FuseMountInfo {
	return FuseMountInfo{
		Path:   o.Path,
		Fstype: o.Fstype,
		Output: o.Output,
	}
}

type FuseStatus struct {
	Version       string          `codec:"version" json:"version"`
	BundleVersion string          `codec:"bundleVersion" json:"bundleVersion"`
	KextID        string          `codec:"kextID" json:"kextID"`
	Path          string          `codec:"path" json:"path"`
	KextStarted   bool            `codec:"kextStarted" json:"kextStarted"`
	InstallStatus InstallStatus   `codec:"installStatus" json:"installStatus"`
	InstallAction InstallAction   `codec:"installAction" json:"installAction"`
	MountInfos    []FuseMountInfo `codec:"mountInfos" json:"mountInfos"`
	Status        Status          `codec:"status" json:"status"`
}

func (o FuseStatus) DeepCopy() FuseStatus {
	return FuseStatus{
		Version:       o.Version,
		BundleVersion: o.BundleVersion,
		KextID:        o.KextID,
		Path:          o.Path,
		KextStarted:   o.KextStarted,
		InstallStatus: o.InstallStatus.DeepCopy(),
		InstallAction: o.InstallAction.DeepCopy(),
		MountInfos: (func(x []FuseMountInfo) []FuseMountInfo {
			var ret []FuseMountInfo
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.MountInfos),
		Status: o.Status.DeepCopy(),
	}
}

type ComponentResult struct {
	Name     string `codec:"name" json:"name"`
	Status   Status `codec:"status" json:"status"`
	ExitCode int    `codec:"exitCode" json:"exitCode"`
}

func (o ComponentResult) DeepCopy() ComponentResult {
	return ComponentResult{
		Name:     o.Name,
		Status:   o.Status.DeepCopy(),
		ExitCode: o.ExitCode,
	}
}

type InstallResult struct {
	ComponentResults []ComponentResult `codec:"componentResults" json:"componentResults"`
	Status           Status            `codec:"status" json:"status"`
	Fatal            bool              `codec:"fatal" json:"fatal"`
}

func (o InstallResult) DeepCopy() InstallResult {
	return InstallResult{
		ComponentResults: (func(x []ComponentResult) []ComponentResult {
			var ret []ComponentResult
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.ComponentResults),
		Status: o.Status.DeepCopy(),
		Fatal:  o.Fatal,
	}
}

type UninstallResult struct {
	ComponentResults []ComponentResult `codec:"componentResults" json:"componentResults"`
	Status           Status            `codec:"status" json:"status"`
}

func (o UninstallResult) DeepCopy() UninstallResult {
	return UninstallResult{
		ComponentResults: (func(x []ComponentResult) []ComponentResult {
			var ret []ComponentResult
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.ComponentResults),
		Status: o.Status.DeepCopy(),
	}
}

type FuseStatusArg struct {
	SessionID     int    `codec:"sessionID" json:"sessionID"`
	BundleVersion string `codec:"bundleVersion" json:"bundleVersion"`
}

func (o FuseStatusArg) DeepCopy() FuseStatusArg {
	return FuseStatusArg{
		SessionID:     o.SessionID,
		BundleVersion: o.BundleVersion,
	}
}

type InstallFuseArg struct {
}

func (o InstallFuseArg) DeepCopy() InstallFuseArg {
	return InstallFuseArg{}
}

type InstallKBFSArg struct {
}

func (o InstallKBFSArg) DeepCopy() InstallKBFSArg {
	return InstallKBFSArg{}
}

type UninstallKBFSArg struct {
}

func (o UninstallKBFSArg) DeepCopy() UninstallKBFSArg {
	return UninstallKBFSArg{}
}

type InstallCommandLinePrivilegedArg struct {
}

func (o InstallCommandLinePrivilegedArg) DeepCopy() InstallCommandLinePrivilegedArg {
	return InstallCommandLinePrivilegedArg{}
}

type InstallInterface interface {
	FuseStatus(context.Context, FuseStatusArg) (FuseStatus, error)
	InstallFuse(context.Context) (InstallResult, error)
	InstallKBFS(context.Context) (InstallResult, error)
	UninstallKBFS(context.Context) (UninstallResult, error)
	InstallCommandLinePrivileged(context.Context) (InstallResult, error)
}

func InstallProtocol(i InstallInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.install",
		Methods: map[string]rpc.ServeHandlerDescription{
			"fuseStatus": {
				MakeArg: func() interface{} {
					ret := make([]FuseStatusArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]FuseStatusArg)
					if !ok {
						err = rpc.NewTypeError((*[]FuseStatusArg)(nil), args)
						return
					}
					ret, err = i.FuseStatus(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"installFuse": {
				MakeArg: func() interface{} {
					ret := make([]InstallFuseArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.InstallFuse(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"installKBFS": {
				MakeArg: func() interface{} {
					ret := make([]InstallKBFSArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.InstallKBFS(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"uninstallKBFS": {
				MakeArg: func() interface{} {
					ret := make([]UninstallKBFSArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.UninstallKBFS(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"installCommandLinePrivileged": {
				MakeArg: func() interface{} {
					ret := make([]InstallCommandLinePrivilegedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.InstallCommandLinePrivileged(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type InstallClient struct {
	Cli rpc.GenericClient
}

func (c InstallClient) FuseStatus(ctx context.Context, __arg FuseStatusArg) (res FuseStatus, err error) {
	err = c.Cli.Call(ctx, "keybase.1.install.fuseStatus", []interface{}{__arg}, &res)
	return
}

func (c InstallClient) InstallFuse(ctx context.Context) (res InstallResult, err error) {
	err = c.Cli.Call(ctx, "keybase.1.install.installFuse", []interface{}{InstallFuseArg{}}, &res)
	return
}

func (c InstallClient) InstallKBFS(ctx context.Context) (res InstallResult, err error) {
	err = c.Cli.Call(ctx, "keybase.1.install.installKBFS", []interface{}{InstallKBFSArg{}}, &res)
	return
}

func (c InstallClient) UninstallKBFS(ctx context.Context) (res UninstallResult, err error) {
	err = c.Cli.Call(ctx, "keybase.1.install.uninstallKBFS", []interface{}{UninstallKBFSArg{}}, &res)
	return
}

func (c InstallClient) InstallCommandLinePrivileged(ctx context.Context) (res InstallResult, err error) {
	err = c.Cli.Call(ctx, "keybase.1.install.installCommandLinePrivileged", []interface{}{InstallCommandLinePrivilegedArg{}}, &res)
	return
}
