package gossip

import (
    "github.com/iotadevelopment/go/packages/transaction"
    "reflect"
)

type neighborEvents struct {
    IncomingConnection     *callbackEvent
    ReceiveData            *dataEvent
    ReceiveTransactionData *dataEvent
    ReceiveTransaction     *transactionEvent
    Disconnect             *callbackEvent
    Error                  *errorEvent
}

//region callbackEvent /////////////////////////////////////////////////////////////////////////////////////////////////////

type callbackEvent struct {
    callbacks map[uintptr]Callback
}

func (this *callbackEvent) Attach(callback Callback) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *callbackEvent) Detach(callback Callback) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *callbackEvent) Trigger() {
    for _, callback := range this.callbacks {
        callback()
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region dataEvent /////////////////////////////////////////////////////////////////////////////////////////////////////

type dataEvent struct {
    callbacks map[uintptr]DataConsumer
}

func (this *dataEvent) Attach(callback DataConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *dataEvent) Detach(callback DataConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *dataEvent) Trigger(data []byte) {
    for _, callback := range this.callbacks {
        callback(data)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region transactionEvent //////////////////////////////////////////////////////////////////////////////////////////////

type transactionEvent struct {
    callbacks map[uintptr]TransactionConsumer
}

func (this *transactionEvent) Attach(callback TransactionConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *transactionEvent) Detach(callback TransactionConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *transactionEvent) Trigger(transaction *transaction.Transaction) {
    for _, callback := range this.callbacks {
        callback(transaction)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region errorEvent ////////////////////////////////////////////////////////////////////////////////////////////////////

type errorEvent struct {
    callbacks map[uintptr]ErrorConsumer
}

func (this *errorEvent) Attach(callback ErrorConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *errorEvent) Detach(callback ErrorConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *errorEvent) Trigger(err error) {
    for _, callback := range this.callbacks {
        callback(err)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////