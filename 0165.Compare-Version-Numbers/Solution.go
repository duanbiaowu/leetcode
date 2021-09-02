package leetcode

func compareVersion(version1 string, version2 string) int {
	for i, j := 0, 0; i < len(version1) || j < len(version2); {
		v1, v2 := 0, 0
		for ; i < len(version1) && version1[i] != '.'; i++ {
			v1 = v1*10 + int(version1[i]-'0')
		}
		for ; j < len(version2) && version2[j] != '.'; j++ {
			v2 = v2*10 + int(version2[j]-'0')
		}
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
		i++
		j++
	}
	return 0
}
