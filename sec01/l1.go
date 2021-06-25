//Constant and variablle
package main
import "fmt"
 const pi=3.14
 const zz="hi this is a constant string"
 const ww= 22

 func main(){
	 fmt.Printf("pq %v, %T \n",pi,pi)
	 fmt.Printf("pq %v, %T\n",zz,zz)
	 fmt.Printf("pq %v, %T\n",ww,ww)
	 var a bool
	 fmt.Println(a)
	 a=true
	 fmt.Println(a)
	  var ax,aa,as= 1.2, 22,"hey yo"
	  fmt.Println(ax,aa,as)
	 ax =11 +ax
	 fmt.Println(ax,aa,as)
	 //________________________________________________
	 // If statment
	 x:=2;
	 if x>10{fmt.Println("greater than x")
	} else {fmt.Println("less than 10")}

	/////_________________________________________________
	//switch statment
	switch x=1;{
	case x==1:fmt.Println("equal to 1")
	case x<1: fmt.Println("less than 1")
	case x>1: fmt.Println("greater than 1")}

	//for stATMENT________________
	for x=1;x<20;x++{
		fmt.Printf("value of x is = %v \n" ,x);
		switch {
			case x==10:fmt.Println("***_____equal to 10")
			case x<10: fmt.Println("less than 10")
			case x>10: fmt.Println("greater than 10")}
		
	}
// function________________
 prtmsg(10)

 // just testing
 fmt.Printf("testing :%v, %T",'+','+')

 //------------------------------------------
 //function with multiple parameters
p, q:= 3, 10;fmt.Printf("p=%v;q=%v \n",p,q)
p, q=swap(p,q);fmt.Printf("p=%v;q=%v \n",p,q)

//varadic functon----------------------
sum();
sum(1);
sum(1,2,3,4,5);
qa("a1",1);
qa("a1",1,2);
qa("a1",1,2,3,4);
foo();
 }
 //----------------------------------------------------
 //----------------------------------------------------
 //----------------------------------------------------
 //----------------------------------------------------
 func prtmsg(max int){

	for x:=1;x<max;x++{
		fmt.Printf("value of x is = %v \n" ,x);
		switch {
			case x==10:fmt.Println("***------equal to 10")
			case x<10: fmt.Println("--------less than 10")
			case x>10: fmt.Println("-----greater than 10")}}
 }
 //----------------------------------------------------
 //----------------------------------------------------
 //function for multiple parameter
 func swap(p int, q int)(int,int){
	 var temp int
	 temp=p;
	 p=q;
	 q=temp;
	 return p,q ;
 }
 //--------------------------------------------------------
 //--------------------------------------------------------
 //Varadic Function
 func sum(v ...int){
	 fmt.Printf("value is %v and type is %T\n",v,v);
	 fmt.Println(len(v));
 }
 //hybred varadic func
 func qa(s string, o ...int){
	 fmt.Println(s,o)
 }
 //anaynomus function
 var foo=func(){
	 fmt.Println("foo is called")
 }