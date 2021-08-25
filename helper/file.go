package helper

import "bytes"

/*
生成相对路径
*/
func GenRelativePath(fileName string) string {
	var buf bytes.Buffer
	buf.WriteString("./")
	buf.WriteString(fileName)
	return buf.String()
}
