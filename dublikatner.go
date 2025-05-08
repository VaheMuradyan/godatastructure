package main

func AreDisjoint(arr1, arr2 []int) bool {
	set1 := make(map[int]struct{})
	for _, num := range arr1 {
		set1[num] = struct{}{} // Populate the map using struct{} to save space
	}

	for _, num := range arr2 {
		if _, exists := set1[num]; exists {
			return false // If found, the sets are not disjoint.
		}
	}
	return true
}

func RemoveDuplicates(arr []int) []int {
	nums := make(map[int]struct{})
	for _, num := range arr {
		nums[num] = struct{}{} // Add the number if it's not already present, ignoring duplicates
	}

	result := make([]int, 0, len(nums))
	for num := range nums {
		result = append(result, num) // Each unique element is added to the result slice
	}
	return result
}
