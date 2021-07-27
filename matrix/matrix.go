package matrix
import "fmt"

/**1단계 가시화;행렬**/
func PrintMat(aa [][]float64) {
  y:=len(aa); x:=len(aa[0]) //크기 구하기
  fmt.Print("┌─"); for i:=0;i<x*4-2;i++ {fmt.Print(" ")}; fmt.Println("─┐")  //상단부 출력
  for i:=0;i<y;i++{  //행렬 원소 출력
    fmt.Print("│")
    for j:=0;j<x;j++ {
      fmt.Printf("%3.f ",aa[i][j])
    }
    fmt.Println("│")
  }
  fmt.Print("└─"); for i:=0;i<x*4-2;i++ {fmt.Print(" ")}; fmt.Println("─┘ ")  //하단부 출력
}

/**전치행렬=역행렬 구하기**/
func MatTrans(aa [][]float64) ([][]float64) {
	x:=len(aa[0]); y:=len(aa)  //크기 구하기
	newMat:=make([][]float64,x) //예외 처리 없이 행렬 생성
	for i:=range newMat {
		newMat[i]=make([]float64,y)
	}
	for i:=0; i<x; i++ {  //행렬 전치 연산
		for j:=0; j<y; j++ {
			newMat[i][j]=aa[j][i]
		}
	}
	PrintMat(newMat); fmt.Println(" >MatTrans success");
  return newMat
}

/**기본연산(덧셈,상수곱,곱셈)**/
func MatSum(aa [][]float64, bb [][]float64) ([][]float64) {
	ax:=len(aa[0]); ay:=len(aa) //크기 구하기
  bx:=len(bb[0]); by:=len(bb)
  var newMat [][] float64  //예외 처리와 함께 행렬 생성
  if ax!=bx || ay!=by {
    fmt.Println("Matrix size not match: SUM"); return newMat
  }
  x:=ax; y:=ay //새로운 행렬의 크기 구하기
  newMat=make([][]float64, y)
  for i:=0; i<y; i++ {
    newMat[i]=make([]float64, x)
    for j:=0; j<x; j++ {  //행렬의 합 연산
      newMat[i][j]=aa[i][j]+bb[i][j]
    }
  }
  PrintMat(newMat); fmt.Println(" >Matrix SUM success"); return newMat
}
func MatMulScalar(aa [][]float64, k float64) ([][]float64) {
  ax:=len(aa[0]); ay:=len(aa)
  x:=ax; y:=ay  //크기 구하기  &  새로운 행렬의 크기 구하기
  newMat:=make([][]float64, y)  //예외 처리 없이 행렬 생성
  for i:=0; i<y; i++ {  //행렬 상수곱 연산
    newMat[i]=make([]float64, x)
    for j:=0; j<x; j++ {
      newMat[i][j]=aa[i][j]*k
    }
  }
  return newMat;
}
func MatMul(aa [][]float64, bb [][]float64) ([][]float64) {
  ax:=len(aa[0]); ay:=len(aa)
  bx:=len(bb[0])  //크기 구하기
  var newMat [][]float64  //예외 처리와 함께 행렬 생성
  if ay!=bx {
    fmt.Println("Matrix size not match: MUL"); return newMat;
  }
  y:=ay; x:=bx; length:=ax  //새로운 행렬의 크기 구하기
  newMat=make([][]float64, ay)
  for i:=0; i<y; i++ {  //행렬 곱 연산
    newMat[i]=make([]float64, bx)
    for j:=0; j<x; j++ {
      newMat[i][j]=0
      for t:=0; t<length; t++ {
        newMat[i][j]+=aa[i][t]*bb[t][j]
      }
    }
  }
  PrintMat(newMat); fmt.Println(" >Matrix MUL success"); return newMat
}

/**텐서곱**/
func _Replace(newMat[][]float64, temp [][]float64, y int, x int, row int, col int) {
  for i:=0; i<row; i++ {
    for j:=0; j<col; j++ {
      newMat[y+i][x+j]=temp[i][j];
    }
  }
}
func TenMul(aa [][]float64, bb [][]float64) ([][]float64){
  ax:=len(aa[0]); ay:=len(aa)  //행렬의 크기 구하기
  bx:=len(bb[0]); by:=len(bb)
  nx:=ax*bx; ny:=ay*by //새로운 행렬의 크기 구하기
  newMat:=make([][]float64, ny)  //예외 처리 없이 행렬 생성
  for i:=0; i<ay*by; i++ {
    newMat[i]=make([]float64, nx)
  }
  for i:=0; i<ay; i++ {  //텐서곱 연산
    for j:=0; j<ax; j++ {
      tmpMat:=MatMulScalar(bb,aa[i][j])
      _Replace(newMat, tmpMat, i*by, j*bx, by, bx)  //새롭게 만든 배열 안에 상수배한 행렬 위치 재배열하기
    }
  }
  PrintMat(newMat); fmt.Println(" >Tensor MUL success");return newMat
}
