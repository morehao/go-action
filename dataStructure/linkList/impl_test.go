package linkList

import "testing"

func TestFunc(t *testing.T) {
	linkList := CreateLinkList()
	linkList.Add(1)
	linkList.Add(2)
	linkList.Add(3)
	linkList.Add(4)
	linkList.Scan()
	// linkList.Insert(2, 4)
	// linkList.Scan()
	linkList.Delete(4)
	linkList.Scan()
}
