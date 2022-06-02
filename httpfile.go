package main
/* 在局域网中将当前文件夹中所有文件http服务
go run main.go
go build main.go
./main.exe
*/
import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func getIp() string {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
func main() {
	h := http.FileServer(http.Dir("."))
	getIp()
	fmt.Println("http://" + getIp() + ":80")
	http.ListenAndServe(":80", h)
}
