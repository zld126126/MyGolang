package util

//捕获异常 error
func Catch(err error) {
	if err != nil {
		panic(err)
	}
}
