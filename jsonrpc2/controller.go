package jsonrpc2

type Controller struct {
	Name 		string
	F 			Func
	Params 		interface{}
	Result 		interface{}
}