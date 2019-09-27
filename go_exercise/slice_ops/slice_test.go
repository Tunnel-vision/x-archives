package slice_ops

import (
	"fmt"
	"testing"
)

func printSlice(s []int)  {
	fmt.Printf("s=%v, len(s)=%d,cap=%d\n",s,len(s),cap(s))
}


func TestSliceAutoLonger(t *testing.T)  {
	var s [] int
	t.Log(s)
	for i :=0; i<10;i++  {
		s = append(s,i)
		printSlice(s)
	}
	t.Log(s)
	for i:=0;i<10;i++{
		s = s[1:]
		printSlice(s)
	}
	t.Log(s)
}

func TestMakeSlice(t *testing.T)  {
	s1 := make([]int,10)
	t.Logf("s=%v, len(s)=%d,cap=%d\n",s1,len(s1),cap(s1))

	s2 :=make([]int,10,32)
	t.Logf("s=%v, len(s)=%d,cap=%d\n",s2,len(s2),cap(s2))
}

func TestCopySlice(t *testing.T)  {
	var s1 = []int{1,3,4,5,7,9}
	var s2 = make([]int,2,4)
	copy(s2,s1)
	t.Logf("s=%v, len(s)=%d,cap=%d\n",s2,len(s2),cap(s2))
	var s3 = []int{11}
	copy(s3,s1)
	t.Logf("s=%v, len(s)=%d,cap=%d\n",s3,len(s3),cap(s3))
	s2 = append(s2,100)
	t.Logf("s=%v, len(s)=%d,cap=%d\n",s2,len(s2),cap(s2))

}

func TestSliceOutOfBound(t *testing.T)  {
	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[2:6]
	t.Logf("s=%v, len(s)=%d,cap=%d\n",s1,len(s1),cap(s1))

	s2 := s1[3:5]
	t.Logf("s=%v, len(s)=%d,cap=%d\n",s2,len(s2),cap(s2))
}


//[] 只能访问 len(arr) 范围内的元素, [:] 只能访问 cap(arr) 范围内的元素,一般而言 cap>=len 所以某些情况看起来越界,其实并不没有越界,只是二者的标准不同!
//我们知道切片 slice 的内部数据结构是基于动态数组,存在三个重要的变量,分别是指针 ptr,个数 len 和容量 cap ,理解了这三个变量如何实现动态数组就不会掉进切片的坑了!
//
//个数 len 是通过下标访问时的有效范围,超过 len 后会报越界错误,而容量 cap 是往后能看到的最大范围,动态数组的本质也是控制这两个变量实现有效数组的访问.












