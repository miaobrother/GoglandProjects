package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	Width,Height int
}

type Rect struct {

}
//面积
func (r Rect) Area(p Params,ret *int) error  {
	*ret = p.Width * p.Height
	return nil
}

//  周长
func (r *Rect) Perimeter(p Params,ret *int) error  {
	*ret = (p.Width + p.Height) *2
	return nil
}

func main() {
	rect := new(Rect)

	rpc.Register(rect)

	rpc.HandleHTTP()
	err := http.ListenAndServe(":4321",nil)
	if err != nil{
		log.Fatal(err)
	}
}