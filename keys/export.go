package keys

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

type Client interface {
	Add(name, password string) (address string, mnemonic string, err sdk.Error)
	Recover(name, password, mnemonic string) (address string, err sdk.Error)
	RecoverWithHDPath(name, password, mnemonic, hdPath string) (address string, err sdk.Error)
	Import(name, password, privKeyArmor string) (address string, err sdk.Error)
	Export(name, password string) (privKeyArmor string, err sdk.Error)
	Delete(name, password string) sdk.Error
	Show(name, password string) (string, sdk.Error)
}
