package utils

import "fmt"

// BitwiseOrToUint64BySlice 对切片中的所有元素进行按位或运算
func BitwiseOrToUint64BySlice(slice []uint64) uint64 {
	result := uint64(0)
	for _, v := range slice {
		if v == 1 {
			result |= v
			continue
		}
		menuType := (v - 1) << 1
		fmt.Println("menuType:", v, menuType)
		result |= menuType
	}
	// 转为二进制
	return result
}

// BitwiseOrToSliceByUint64 将一个 uint64 类型的值按位或运算，返回包含相应位的切片
func BitwiseOrToSliceByUint64(n uint64) []uint64 {
	var result []uint64
	for i := uint64(1); i <= n; i <<= 1 {
		fmt.Println("i:", i)
		if i == 0 {
			result = append(result, i)
			continue
		}
		result = append(result, i-1)
	}
	return result
}

// func convertVisibilityPolicyToArray(visibilityPolicy int) []int {
// 	visibilityPolicyList := []int{}
// 	if visibilityPolicy&VisibilityNormalUser != 0 {
// 		visibilityPolicyList = append(visibilityPolicyList, VisibilityNormalUser)
// 	}
// 	if visibilityPolicy&VisibilityPlatformAdmin != 0 {
// 		visibilityPolicyList = append(visibilityPolicyList, VisibilityPlatformAdmin)
// 	}
// 	if visibilityPolicy&VisibilityCompanyAdmin != 0 {
// 		visibilityPolicyList = append(visibilityPolicyList, VisibilityCompanyAdmin)
// 	}
// 	return nil
// }
