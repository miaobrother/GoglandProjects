package lib

//import "encoding/xml"

type DomainXml struct{
	XmlNmae	string `xml:"name"`
	RDevices SChildrenGroup `xml:"devices"`

}
type SChildrenGroup struct {

	Graphics VncPort `xml:"graphics"`
	Interface SInterface `xml:"interface"`
}
type VncPort struct {
	Port string `xml:"port,attr"`
}
type SInterface struct {
	Target DevNetwokr `xml:"target"`
}


type DevNetwokr struct {
	Dev string `xml:"dev,attr"`
}
