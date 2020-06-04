package main
import "fmt"
import "net/http"
var point string
func main() {
  // point = 80;
  //  cal(point)
   handleRequest()
}
func handleRequest() {
  http.HandleFunc("/", cal)
  http.ListenAndServe(":8080", nil)
}
type test_struct struct {
    Test string
}
func cal(writer http.ResponseWriter, request *http.Request) {
   writer.Header().Set("Content-Type", "application/json")
   writer.WriteHeader(http.StatusCreated)
   
   fmt.Println(request.Body)
   writer.Write([]byte(point))
  // fmt.Println("You Point", point)
  //  fmt.Println("------------------------")
  // if point < 50 {
  //   fmt.Println("Your Grade F");
  // } else if point < 55 {
  //   fmt.Println("Your Grade D");
  // } else if point < 60 {
  //   fmt.Println("Your Grade D+");
  // } else if point < 65 {
  //   fmt.Println("Your Grade c");
  // } else if point < 70 {
  //   fmt.Println("Your Grade c+");
  // } else if point < 75 {
  //   fmt.Println("Your Grade B");
  // } else if point < 80 {
  //   fmt.Println("Your Grade B+");
  // } else if point >= 80 {
  //   fmt.Println("Your Grade A");
  // }
}
