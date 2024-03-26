package baseapp

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//
// Side channel
//

// BeginSideBlock implements the ABCI application interface.
func (app *BaseApp) BeginSideBlock(req abci.RequestBeginSideBlock) (res abci.ResponseBeginSideBlock) {
	if app.beginSideBlocker != nil {
		res = app.beginSideBlocker(app.deliverState.ctx, req)
	}

	return
}

// DeliverSideTx implements the ABCI application interface.
func (app *BaseApp) DeliverSideTx(req abci.RequestDeliverSideTx) (res abci.ResponseDeliverSideTx) {
	tx, err := app.txDecoder(req.Tx)
	if err != nil {
		// r := err.Result()
		res = abci.ResponseDeliverSideTx{
			Result: abci.SideTxResultType_Skip,
			// Code:      uint32(r.Code),
			// Codespace: string(r.Codespace),
		}
	} else {
		res = app.runSideTx(req.Tx, tx, req)
	}

	return
}

// runSideTx processes a side transaction. App can make an external call here.
func (app *BaseApp) runSideTx(txBytes []byte, tx sdk.Tx, req abci.RequestDeliverSideTx) (res abci.ResponseDeliverSideTx) {
	defer func() {
		if r := recover(); r != nil {
			res = abci.ResponseDeliverSideTx{
				Result: abci.SideTxResultType_Skip, // skip proposal
				// Code:      uint32(sdk.CodeInternal),
				// Codespace: string(sdk.CodespaceRoot),
			}
		}
	}()

	var msgs = tx.GetMsgs()
	if err := validateBasicTxMsgs(msgs); err != nil {
		// r := err.Result()
		res = abci.ResponseDeliverSideTx{
			Result: abci.SideTxResultType_Skip, // skip proposal
			// Code:      uint32(r.Code),
			// Codespace: string(r.Codespace),
		}
		return
	}

	if app.deliverSideTxHandler != nil {
		// get deliver-tx context
		ctx := app.getContextForTx(runTxModeDeliver, txBytes)

		res = app.deliverSideTxHandler(ctx, tx, req)
	} else {
		res = abci.ResponseDeliverSideTx{
			Result: abci.SideTxResultType_Skip, // skip proposal
		}
	}

	return
}
