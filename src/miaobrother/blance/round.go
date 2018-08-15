package main

type Roundblance struct {
	currentIndex int
}

func (r *Roundblance) DoBlance(AddrList []string) string {
	l := len(AddrList)
	r.currentIndex = r.currentIndex % l
	addr := AddrList[r.currentIndex]
	r.currentIndex++
	return addr

}
