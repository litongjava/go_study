package main
import(
	"fmt"
	"net/http"
)

//创建处理器函数
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"hello,world",r.URL.Path)
}

func main(){
	http.HandleFunc("/",handler)
	fmt.Printf("服务已经启动\n")
	//创建路由,使用默认的多路器
	http.ListenAndServe(":8080",nil)
}