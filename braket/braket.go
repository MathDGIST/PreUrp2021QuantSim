package braket

import "fmt"

var p = fmt.Print
var pf = fmt.Printf
var pl = fmt.Println
var sf = fmt.Scanf
var s = fmt.Scan

func InputBraKet() {
  var a string
  var input []byte
  var arr []int
  var output []int
  zerone:=make([][]int,2)
  zerone[0]=[]int{1,0}
  zerone[1]=[]int{0,1}
  output=append(output,1)
  brk:=0
  s(&a)
  input=[]byte(a)

  if((input[0]!=byte('|')&&input[len(input)-1]!=byte('>')&&input[0]!=byte('<')&&input[len(input)-1]!=byte('|'))||(input[0]==byte('|')&&input[len(input)-1]==byte('|'))||(input[0]==byte('<')&&input[len(input)-1]==byte('>'))){
    pl("\n*******\n Error\n*******")
    brk+=1
  }else{
    for i:=1;i<len(input)-1;i++{
      if(input[i]-48!=0&&input[i]-48!=1){
        pl("\n*******\n Error\n*******")
        brk+=1
        break
      }
      arr=append(arr,int(input[i]-48))
    }
    cnt:=1
    for i:=0;i<len(arr);i++{

      cnt*=2;
      forcal:=make([]int,cnt)
      for j:=0;j<cnt;j++{
        forcal[j]=output[j/2]&zerone[arr[i]][j%2] 
      }
      // for j:=0;j<len(output);j++{
      //   output=output[1:]
      // }
      output=make([]int,cnt)
      copy(output,forcal)
    }
  }
  if(brk==0){
    var OUT [][]int
    if(input[0]==byte('<')){
      OUT=make([][]int,1)
      OUT[0]=make([]int,len(output))
      for i:=0;i<len(output);i++{
        OUT[0][i]=output[i]
      }
    }else{
      OUT=make([][]int,len(output))
      for i:=0;i<len(output);i++{
        OUT[i]=make([]int,1)
        OUT[i][0]=output[i]
      }
    }
    //Slice 완성. OUT 이차원 슬라이스가 bra 또는 ket의
    //행렬로 변환된 부분이다.

    //여기서 부터는 출력.
    pl("\n\nOUTPUT\n")
    for i:=0;i<len(OUT);i++{
      for j:=0;j<len(OUT[i]);j++{
        p(OUT[i][j]," ")
      }
      pl()
    }
  }
}