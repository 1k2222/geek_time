package main

import (
	"fmt"
	"geek_time/beauty_of_design_pattern/067_iterator/snapshot_array/snapshot_array"
)

// 编号为0~19的20个人站成1排，从0开始报数，报到偶数的出列，剩余的继续按规则报数，重复三轮
func main() {
	arr := snapshot_array.NewSnapshotArrayList()

	for i := 0; i < 20; i++ {
		arr.Add(i)
	}

	for round := 0; round < 3; round++ {
		survivor := make([]int, 0)
		for i, iter := 0, arr.Iterator(); iter.HasNext(); i++ {
			t, _ := iter.Next().(int)
			survivor = append(survivor, t)
			if i%2 == 0 {
				arr.Remove(t)
			}
		}
		fmt.Printf("Round %d:%v\n", round, survivor)
		fmt.Printf("\t the 2th value of next round is %d\n", arr.Get(1))
	}

	survivor := make([]int, 0)
	for iter := arr.Iterator(); iter.HasNext(); {
		t, _ := iter.Next().(int)
		survivor = append(survivor, t)
	}
	fmt.Printf("Round %d:%v\n", 3, survivor)
}
