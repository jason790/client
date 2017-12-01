// Auto-generated by avdl-compiler v1.3.22 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/revoke.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type RevokeKeyArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
	KeyID     KID `codec:"keyID" json:"keyID"`
}

type RevokeDeviceArg struct {
	SessionID int      `codec:"sessionID" json:"sessionID"`
	DeviceID  DeviceID `codec:"deviceID" json:"deviceID"`
	ForceSelf bool     `codec:"forceSelf" json:"forceSelf"`
	ForceLast bool     `codec:"forceLast" json:"forceLast"`
}

type RevokeSigsArg struct {
	SessionID    int      `codec:"sessionID" json:"sessionID"`
	SigIDQueries []string `codec:"sigIDQueries" json:"sigIDQueries"`
}

type RevokeInterface interface {
	RevokeKey(context.Context, RevokeKeyArg) error
	RevokeDevice(context.Context, RevokeDeviceArg) error
	RevokeSigs(context.Context, RevokeSigsArg) error
}

func RevokeProtocol(i RevokeInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.revoke",
		Methods: map[string]rpc.ServeHandlerDescription{
			"revokeKey": {
				MakeArg: func() interface{} {
					ret := make([]RevokeKeyArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]RevokeKeyArg)
					if !ok {
						err = rpc.NewTypeError((*[]RevokeKeyArg)(nil), args)
						return
					}
					err = i.RevokeKey(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"revokeDevice": {
				MakeArg: func() interface{} {
					ret := make([]RevokeDeviceArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]RevokeDeviceArg)
					if !ok {
						err = rpc.NewTypeError((*[]RevokeDeviceArg)(nil), args)
						return
					}
					err = i.RevokeDevice(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"revokeSigs": {
				MakeArg: func() interface{} {
					ret := make([]RevokeSigsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]RevokeSigsArg)
					if !ok {
						err = rpc.NewTypeError((*[]RevokeSigsArg)(nil), args)
						return
					}
					err = i.RevokeSigs(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type RevokeClient struct {
	Cli rpc.GenericClient
}

func (c RevokeClient) RevokeKey(ctx context.Context, __arg RevokeKeyArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.revoke.revokeKey", []interface{}{__arg}, nil)
	return
}

func (c RevokeClient) RevokeDevice(ctx context.Context, __arg RevokeDeviceArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.revoke.revokeDevice", []interface{}{__arg}, nil)
	return
}

func (c RevokeClient) RevokeSigs(ctx context.Context, __arg RevokeSigsArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.revoke.revokeSigs", []interface{}{__arg}, nil)
	return
}
