package main

import (
  "main/matrix"
  "main/braket"
)
func main() {
  a:=[][]float64{{16,4,28},{4,40,12},{28,12,8}}
  b:=[][]float64{{0.5,0.5,0.75},{0.5, 1.25, 0.25}, {0.75, 0.25, 1}}
  matrix.PrintMat(a)
  matrix.MatTrans(a)
  matrix.MatSum(a,b)
  matrix.MatMulScalar(a,2)
  matrix.MatMul(a,b)
  matrix.TenMul(a,b)
  braket.InputBraKet()

}