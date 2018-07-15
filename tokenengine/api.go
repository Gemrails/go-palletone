package tokenengine

import (
	"fmt"

	"github.com/palletone/go-palletone/tokenengine/btcd/txscript"
	"github.com/palletone/go-palletone/tokenengine/btcd/wire"
)

func GenerateLockScript(pubKeyHash []byte) []byte {
	//address to byte[20]
	lock, _ := txscript.NewScriptBuilder().AddOp(txscript.OP_DUP).AddOp(txscript.OP_HASH160).
		AddData(pubKeyHash).AddOp(txscript.OP_EQUALVERIFY).AddOp(txscript.OP_CHECKSIG).
		Script()
	return lock
}
func PickAddress(lockscript []byte) (string, error) {
	fmt.Println(lockscript)
	return "12gpXQVcCL2qhTNQgyLVdCFG2Qs2px98nV", nil
}
func GenerateP2PKHUnlockScript(sign []byte, pubKey []byte) []byte {
	unlock, _ := txscript.NewScriptBuilder().AddData(sign).AddData(pubKey).Script()
	return unlock
}
func ScriptValidate(utxoLockScript []byte, utxoAmount int64, tx *wire.MsgTx, inputIndex int) error {
	vm, err := txscript.NewEngine(utxoLockScript, tx, 0, txscript.StandardVerifyFlags, nil, nil, utxoAmount)
	if err != nil {
		fmt.Errorf("Failed to create script: %v", err)
		return err
	}
	return vm.Execute()
}
