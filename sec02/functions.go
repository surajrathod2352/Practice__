package main
import "fmt"
func main(){
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
   fof(sum,qa); 

}
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
 }// apply all the funtion uses to his too  

 func fof(sum func(v ...int),qa func(s string, o ...int)){
//fmt.Println("fof was caled")
//pqp:= sum+qa;
//fmt.Println(pqp)
 }

 