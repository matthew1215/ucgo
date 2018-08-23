package demo

const CodeDemo = 10000

var MSG = map[int]string{
	CodeDemo: "demo",
}

func StatusText(code int) string {
	return MSG[code]
}
