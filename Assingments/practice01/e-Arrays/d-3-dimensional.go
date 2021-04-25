package main

import "fmt"

func main() {
   /* an array with 5 rows and 3 columns*/
   var a = [5][3]int{ {0,0,0}, {1,2,3}, {2,4,5}, {3,6,9},{4,8,7}}
   var i, j int

   /* output each array element's value */
   for  i = 0; i < 5; i++ {
      for j = 0; j < 3; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
      }
   }
}