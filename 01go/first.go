package main
import "fmt" 
//import "math"


func main() {
	////p1
	// var str string
	// var str2 string
	// var int1 int
	// fmt.Printf("Enter string")
	// fmt.Scan(&str,&str2,&int1)
	// fmt.Printf("Result:%s and %s and %d \n",str,str2,int1)
	

    // var a int=123
	// var b uint=0
	// b=a
    // fmt.Printf("a=%d,b=%d\n",a,b)


    // var a int=-12 // -128 - 128
	// var b uint8=0//64 bit range 0-255
	// b=uint8(a)
    // fmt.Printf("a=%d,b=%d\n",a,b)

	// var x int = 255
	// var r float64
	// r=math.Sqrt(float64(x))
	// fmt.Printf("x=%d and r=%f\n",x,r)

//type conversion
	var x int =42
	var y float64=float64(x)
	var z uint=uint(y)
	fmt.Printf("Value of x is %d and type is %T\n",x,x)
	fmt.Printf("Value of y is %f and type is %T\n",y,y)
	fmt.Printf("Value of z is %d and type is %T\n",z,z)
	


}
