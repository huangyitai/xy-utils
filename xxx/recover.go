package xxx

import "fmt"

// RecoverToErr 将recover方法返回结果转换为error
func RecoverToErr(r interface{}) error {
	if r == nil {
		return nil
	}
	if err, ok := r.(error); ok {
		return err
	}
	return fmt.Errorf("%v", r)
}
