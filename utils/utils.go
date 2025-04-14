package utils

func ListofA2ZZ() (list []string) {
	for i := 'A'; i <= 'Z'; i++ {
		list = append(list, string(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			list = append(list, string(i)+string(j))
		}
	}
	// for i := 'A'; i <= 'Z'; i++ {
	// 	for j := 'A'; j <= 'Z'; j++ {
	// 		for k := 'A'; k <= 'Z'; k++ {
	// 			list = append(list, string(i)+string(j)+string(k))
	// 		}
	// 	}
	// }
	return
}
