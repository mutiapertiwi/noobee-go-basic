package main

import "log"

func main() {
	sorted := InsertionSort(6, 5, 3, 1, 8, 7, 2, 4)
	log.Println(sorted)
}

//fungsi sebelumnya
func BubbleSort(nums ...int) (sorted []int) {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			log.Println(nums[i], ">", nums[j])
			if nums[i] > nums[j] {
				log.Println("change", nums[i], "in position", i, "to", nums[j], "in position", j)
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
				log.Println(nums)
			}
		}

		log.Println("hasil sort batch", i+1, ":", nums)
	}
	return
}

func InsertionSort(nums ...int) (sorted []int) {
	for i := 1; i < len(nums); i++ {
		for k := i; k > 0; k-- {
			log.Println(nums[k-1], ">", nums[k])
			if nums[k-1] > nums[k] {
				log.Println("change", nums[k-1], "in position", k-1, "to", nums[k], "in position", k)

				temp := nums[k]
				nums[k] = nums[k-1]
				nums[k-1] = temp

				log.Println(nums)
			}
		}
		log.Println("hasil sort batch", i+1, ":", nums)
	}

	sorted = nums
	return
}
