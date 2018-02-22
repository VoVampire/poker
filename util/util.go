package util

// 三项表达式
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// 统计二进制中1的个数（有效位数为16位）
func CountOne(a uint64) uint64 {
	a = ((a & 0xAAAA) >> 1) + (a & 0x5555)
	a = ((a & 0xCCCC) >> 2) + (a & 0x3333)
	a = ((a & 0xF0F0) >> 4) + (a & 0x0F0F)
	a = ((a & 0xFF00) >> 8) + (a & 0x00FF)
	return a
}