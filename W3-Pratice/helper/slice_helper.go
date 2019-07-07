package helper

func UniStrSlice(input *[]string) (uniqeArr []string) {
	checkArr := make(map[string]bool)
	for _, val := range *input {

		if _, ok := checkArr[val]; !ok {
			checkArr[val] = true;
			uniqeArr = append(uniqeArr, val)
		}
	}

	return
}