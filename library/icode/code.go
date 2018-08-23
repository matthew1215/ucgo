package icode

const Ok = 200
const ParamsError = 1000
const NotLogin = 1001

var MSG = map[int]string{
	Ok:          "ok",
	ParamsError: "ParamsError",
	NotLogin:    "Unauthorized",
}

func StatusText(code int) string {
	return MSG[code]
}
