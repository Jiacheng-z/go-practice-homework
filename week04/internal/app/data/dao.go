package data

import "fmt"

type dao struct {

}

func (d *dao) Get(id int64) string {
	return fmt.Sprintf("Test Data(id:%v)", id)
}
