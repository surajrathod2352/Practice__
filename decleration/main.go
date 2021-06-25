package main
import "fmt"
/* you can define types like give a diff name to int, flot to asn or anything as you wanst
for example 
type name typename
// type num int [and you can do it for almost anyhing
 */
func main(){
	a:=12.22;
	b:=13.23
	pas,as:=avg(a,b)
	fmt.Printf("valuse of sum is $%.2f and value of avj is %.2f ",pas,as)
}
func avg(a float64,b float64)(s float64,av float64){
	s=a+b
	av=s/2
	return
}