package main

import "fmt"


func main(){
	
var count,sum,num,flag int=0,0,1,0


fmt.Println(" first 50 prime numbers")
	

for count < 50 {
	   
 
num++
	   
	    

for i := 2; i <= num-1; {
	        

if num%i == 0 {
	           

 flag=1
	         

break 
	        
 
}
	     

 i++


}
	   
	  


if flag == 0  {
	      
 
 fmt.Println(num)

fmt.Println("sum of first 50 prime numbers is")
	      
  
sum+=num;
	       

 count++
 
}else {
    
 
flag=0


}

 
}

fmt.Println(sum)


}
