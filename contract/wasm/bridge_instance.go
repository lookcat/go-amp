package wasm

import (
	"errors"

	"github.com/ampchain/go-amp/contract"
	"github.com/ampchain/go-amp/contract/bridge"
	"github.com/ampchain/go-amp/contract/wasm/vm"
	"github.com/ampchain/go-amp/pb"
)

type bridgeInstance struct {
	ctx        *bridge.Context
	vmInstance vm.Instance
	codeDesc   *pb.WasmCodeDesc
}

func (v *bridgeInstance) guessEntry() (string, error) {
	switch v.codeDesc.GetRuntime() {
	case "go":
		return "run", nil
	case "c":
		return "_" + v.ctx.Method, nil
	default:
		return "", errors.New("bad runtime")
	}
}

func (v *bridgeInstance) getEntry() (string, error) {
	return v.guessEntry()
}

func (v *bridgeInstance) Exec() error {
	entry, err := v.getEntry()
	if err != nil {
		return err
	}
	return v.vmInstance.Exec(entry)
}

func (v *bridgeInstance) ResourceUsed() contract.Limits {
	return v.vmInstance.ResourceUsed()
}

func (v *bridgeInstance) Release() {
	v.vmInstance.Release()
}
