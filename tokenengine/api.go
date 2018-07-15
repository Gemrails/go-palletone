package tokenengine

import (
	"github.com/palletone/go-palletone/common"
	"github.com/palletone/go-palletone/common/log"
	"github.com/palletone/go-palletone/tokenengine/btcd/txscript"
	"github.com/palletone/go-palletone/tokenengine/btcd/wire"
)

//Generate a P2PKH lock script, just only need input 20bytes public key hash.
//You can use Address.Bytes() to get address hash.
func GenerateP2PKHLockScript(pubKeyHash []byte) []byte {

	lock, _ := txscript.NewScriptBuilder().AddOp(txscript.OP_DUP).AddOp(txscript.OP_HASH160).
		AddData(pubKeyHash).AddOp(txscript.OP_EQUALVERIFY).AddOp(txscript.OP_CHECKSIG).
		Script()
	return lock
}
func GenerateP2SHLockScript(scriptHash []byte) []byte {
	//Mock
	lock, _ := txscript.NewScriptBuilder().AddOp(txscript.OP_DUP).AddOp(txscript.OP_HASH160).
		AddData(scriptHash).AddOp(txscript.OP_EQUALVERIFY).AddOp(txscript.OP_CHECKSIG).
		Script()
	return lock
}
func GenerateLockScript(address common.Address) []byte {
	//Mock TODO
	t, _ := address.Validate()
	if t == common.PublicKeyHash {
		return GenerateP2PKHLockScript(address.Bytes())
	} else {
		return GenerateP2SHLockScript(address.Bytes())
	}
}

//Give a lock script, and parse it then pick the address string out.
func PickAddress(lockscript []byte) (string, error) {
	log.Debug(string(lockscript))
	//Mock
	return "12gpXQVcCL2qhTNQgyLVdCFG2Qs2px98nV", nil
}

//根据签名和公钥信息生成解锁脚本
//Use signature and public key to generate a P2PKH unlock script
func GenerateP2PKHUnlockScript(sign []byte, pubKey []byte) []byte {
	unlock, _ := txscript.NewScriptBuilder().AddData(sign).AddData(pubKey).Script()
	return unlock
}

//根据收集到的签名和脚本生成解锁脚本
//Use collection signatures and redeem script to unlock
func GenerateP2SHUnlockScript(signs [][]byte, redeemScript []byte) []byte {
	builder := txscript.NewScriptBuilder()
	for _, sign := range signs {
		builder.AddData(sign)
	}
	unlock, _ := builder.AddData(redeemScript).Script()
	return unlock
}

//validate this transaction and input index script can unlock the utxo.
func ScriptValidate(utxoLockScript []byte, utxoAmount int64, tx *wire.MsgTx, inputIndex int) error {
	vm, err := txscript.NewEngine(utxoLockScript, tx, 0, txscript.StandardVerifyFlags, nil, nil, utxoAmount)
	if err != nil {
		log.Error("Failed to create script: ", err)
		return err
	}
	return vm.Execute()
}
