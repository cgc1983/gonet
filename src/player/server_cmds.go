package player

import "strings"
import "math/rand"

func (ud *UserData) exec_srv(msg string) string {
	params:= strings.SplitN(msg, " ", 2)
	switch params[0] {
	case "MESG": return ud.S_mesg(params[1]);
	case "ATTACKED": return ud.S_attacked(params[1]);
	}

	return ""
}

func (ud *UserData) S_mesg(p string) string {
	msg := []string{"mesg",p}
	return strings.Join(msg, " ")
}

func (ud *UserData) S_attacked(p string) string {
	msg := []string{"attacked",p}

	//
	nums := make([]int, 100)
	for i:=0;i<len(nums);i++ {
		nums[i] = rand.Int()
	}


	// insertion sort
	for i:=1; i<len(nums);i++ {
		cur := nums[i]
		j:=i-1

		for j>=1 && nums[j] > cur {
			nums[j+1] = nums[j]
			j=j-1
		}

		nums[j+1] = cur
	}

	/*
	for i:=1; i<len(nums);i++ {
		println(nums[i])
	}
	*/

	return strings.Join(msg, " ")
}