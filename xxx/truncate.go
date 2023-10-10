package xxx

// TruncateToMultiLine 按前后缀长度截断字符串并换行
func TruncateToMultiLine(str string, prefixLen, suffixLen int) string {
	if len(str) <= prefixLen+suffixLen {
		return str
	}
	if prefixLen == 0 {
		return "......\n" + str[len(str)-suffixLen:]
	}
	if suffixLen == 0 {
		return str[:prefixLen] + "\n......"
	}
	return str[:prefixLen] + "\n......\n" + str[len(str)-suffixLen:]
}

// TruncateToSingleLine 按前后缀长度截断字符串
func TruncateToSingleLine(str string, prefixLen, suffixLen int) string {
	if len(str) <= prefixLen+suffixLen {
		return str
	}
	if prefixLen == 0 {
		return "......" + str[len(str)-suffixLen:]
	}
	if suffixLen == 0 {
		return str[:prefixLen] + "......"
	}
	return str[:prefixLen] + "......" + str[len(str)-suffixLen:]
}

// TruncateBytes 按前后缀长度截断字节数组
func TruncateBytes(bs []byte, prefixLen, suffixLen int) []byte {
	if len(bs) <= prefixLen+suffixLen {
		return bs
	}

	res := make([]byte, 0, prefixLen+suffixLen+5)

	if prefixLen == 0 {
		res = append(res, "......\n"...)
		return append(res, bs[len(bs)-suffixLen:]...)
	}
	if suffixLen == 0 {
		res = append(res, bs[:prefixLen]...)
		return append(res, "\n......"...)
	}

	res = append(res, bs[:prefixLen]...)
	res = append(res, "\n......\n"...)
	return append(res, bs[len(bs)-suffixLen:]...)
}

// JSONOrProtojsonTruncateToJSONStr 区分普通对象或者proto消息序列化成JSON串，并且按照前后缀长度截断
func JSONOrProtojsonTruncateToJSONStr(v interface{}, prefixLen, suffixLen int) string {
	return TruncateToMultiLine(JSONOrProtojsonToJSONStr(v), prefixLen, suffixLen)
}
