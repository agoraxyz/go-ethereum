package types

import (
	"math/big"

	"github.com/agoraxyz/go-ethereum/common"
)

// LegacyTx is the transaction data of the original Ethereum transactions.
type BaseSystemTx struct {
	Nonce    uint64          // nonce of sender account
	GasPrice *big.Int        // wei per gas
	Gas      uint64          // gas limit
	To       *common.Address `rlp:"nil"` // nil means contract creation
	Value    *big.Int        // wei amount
	Data     []byte          // contract invocation input data
	V, R, S  *big.Int        // signature values
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *BaseSystemTx) copy() TxData {
	cpy := &BaseSystemTx{
		Nonce: tx.Nonce,
		To:    copyAddressPtr(tx.To),
		Data:  common.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are initialized below.
		Value:    new(big.Int),
		GasPrice: new(big.Int),
		V:        new(big.Int),
		R:        new(big.Int),
		S:        new(big.Int),
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

// accessors for innerTx.
func (tx *BaseSystemTx) txType() byte              { return BaseSystemTxType }
func (tx *BaseSystemTx) chainID() *big.Int         { return deriveChainId(tx.V) }
func (tx *BaseSystemTx) accessList() AccessList    { return nil }
func (tx *BaseSystemTx) data() []byte              { return tx.Data }
func (tx *BaseSystemTx) gas() uint64               { return tx.Gas }
func (tx *BaseSystemTx) gasPrice() *big.Int        { return tx.GasPrice }
func (tx *BaseSystemTx) gasTipCap() *big.Int       { return tx.GasPrice }
func (tx *BaseSystemTx) gasFeeCap() *big.Int       { return tx.GasPrice }
func (tx *BaseSystemTx) value() *big.Int           { return tx.Value }
func (tx *BaseSystemTx) nonce() uint64             { return tx.Nonce }
func (tx *BaseSystemTx) to() *common.Address       { return tx.To }
func (tx *BaseSystemTx) blobGas() uint64           { return 0 }
func (tx *BaseSystemTx) blobGasFeeCap() *big.Int   { return nil }
func (tx *BaseSystemTx) blobHashes() []common.Hash { return nil }

func (tx *BaseSystemTx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	return dst.Set(tx.GasPrice)
}

func (tx *BaseSystemTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *BaseSystemTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.V, tx.R, tx.S = v, r, s
}
