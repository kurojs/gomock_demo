//go:generate mockgen -source=structure.go -destination=./mock/structure.go -package=mock
package main

type Stacker interface {
	Append(interface{})
	Pop() interface{}
}

type Queuer interface {
	Enque(interface{})
	Dequeue() interface{}
}
