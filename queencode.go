package main

import "fmt"

func main(){

var robst,cobst,rqueen,cqueen int
var mat [5][5]int
var sum,n,b int=0,5,0
fmt.Println("Enter queen's position in row :")
fmt.Scanf("%d",&rqueen)
fmt.Println("Enter queen's position in column :")
fmt.Scanf("%d",&cqueen)
fmt.Println("Enter no.of obstacles :")
fmt.Scanf("%d",&b)
for i :=1;i <= b;i++ {
fmt.Println("Enter row of obstacle number",i)
fmt.Scanf("%d",&robst)
fmt.Println("Enter column of obstacle number",i)
fmt.Scanf("%d",&cobst)
mat[robst][cobst]=-1
}


mat[rqueen][cqueen]=1 
fmt.Println(mat)
//check all ways

for i :=rqueen+1; i <n; i = i+1 {
    if mat[i][cqueen] == 0 {
        sum+=1
    }else {
        break
    }
}
for i :=rqueen-1; i >=0; i = i-1 {
   if mat[i][cqueen] == 0 { 
        sum+=1
}else {
    break
}
}
//left and right

for i :=cqueen+1; i <n; i++ {
    if mat[rqueen][i] == 0 {
              sum+=1
    }else {
        break
    }
}
for i :=cqueen-1; i >=0; i-- {
   if mat[rqueen][i] == 0 { 
        sum+=1
}else {
    break
}
}
//diagonal right

 for row,col := rqueen-1,cqueen+1;row >= 0 && col < n; row, col = row-1, col+1 {
     if mat[row][col] ==0 {
         sum+=1
     }else {
         break
     }
     }
 for row,col := rqueen+1,cqueen-1;row < n && col >= 0; row, col = row+1, col-1 {
     if mat[row][col] ==0 {
         sum+=1
     }else {
         break
     }
     }
    
    //diagonal left 
    
     for row,col := rqueen-1,cqueen-1;row >= 0 && col >= 0; row, col = row-1, col-1 {
     if mat[row][col] ==0 {
         sum+=1
     }else {
         break
     }
     }
     
      for row,col := rqueen+1,cqueen+1;row < n && col < n; row, col = row+1, col+1 {
     if mat[row][col] ==0 {
         sum+=1
     }else {
         break
     }
     }
     fmt.Println("no.of movements possible is:")
     fmt.Println(sum)
}

