package util

//捕获异常 error
func Chk(err error) {
	if err != nil {
		panic(err)
	}
}
