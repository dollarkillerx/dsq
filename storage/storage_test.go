/**
 * @Author: DollarKillerX
 * @Description: storage_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:50 2019/12/18
 */
package storage

import (
	"log"
	"testing"
)

func TestCollection_Flush(t *testing.T) {
	kvDb, e := New("test")
	if e != nil {
		panic(e)
	}

	lists, e := kvDb.List()
	if e != nil {
		panic(e)
	}
	log.Println(lists)

	e = kvDb.Put("hello", "hello kv db")
	if e != nil {
		panic(e)
	}

	lists, e = kvDb.List()
	if e != nil {
		panic(e)
	}
	log.Println(lists)

	//kvDb.Get()
	e = kvDb.Put("hrc", "我的天啊")
	if e != nil {
		panic(e)
	}
}
