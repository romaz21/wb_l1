package main

import "fmt"

type OldInterface interface {
	OldRequest() string
}

type OldImplementation struct{}

func (o *OldImplementation) OldRequest() string {
	return "Old Request"
}

type NewInterface interface {
	NewRequest() string
}

type NewImplementation struct{}

func (n *NewImplementation) NewRequest() string {
	return "New Request"
}

type Adapter struct {
	oldObject OldInterface
}

func (a *Adapter) NewRequest() string {
	return a.oldObject.OldRequest()
}

func ClientCode(newInterface NewInterface) {
	fmt.Println(newInterface.NewRequest())
}

func main() {
	oldObject := &OldImplementation{}
	adapter := &Adapter{oldObject}

	newObject := &NewImplementation{}

	ClientCode(newObject)
	ClientCode(adapter)
}
