package stringUtils

import "strings"

// ToUpperFirstLetter 字符首字母大写
func ToUpperFirstLetter(str string) string {
	if str == "" {
		return str
	}
	return strings.ToUpper(str[0:1]) + str[1:len(str)]
} // ToUpperFirstLetter 字符首字母大写
func ToLowerFirstLetter(str string) string {
	if str == "" {
		return str
	}
	return strings.ToLower(str[0:1]) + str[1:len(str)]
}

// ToUnderScoreCase 将驼峰命名转下划线命名
func ToUnderScoreCase(str string) string {
	var retStr = ""
	for i, s := range str {
		if 64 < s && s < 91 && i != 0 {
			retStr += "_" + string(s+32)
		} else if i == 0 {
			retStr += string(s + 32)
		} else {
			retStr += string(s)
		}
	}
	return retStr
}

//ConvertToBigCamelCase 将下划线大写方式命名的字符串转换为驼峰式。如果转换前的下划线大写方式命名的字符串为空，则返回空字符串。 例如：HELLO_WORLD->HelloWorld
func ConvertToBigCamelCase(name string) string {
	if name == "" {
		return ""
	} else if !strings.Contains(name, "_") {
		// 不含下划线，仅将首字母大写
		return strings.ToUpper(name[0:1]) + name[1:len(name)]
	}
	var result string = ""
	camels := strings.Split(name, "_")
	for index := range camels {
		if camels[index] == "" {
			continue
		}
		camel := camels[index]
		result = result + strings.ToUpper(camel[0:1]) + strings.ToLower(camel[1:len(camel)])
	}
	return result
}

//ConvertToLittleCamelCase 将下划线大写方式命名的字符串转换为驼峰式,首字母小写。如果转换前的下划线大写方式命名的字符串为空，则返回空字符串。 例如：HELLO_WORLD->helloWorld
func ConvertToLittleCamelCase(name string) string {
	if name == "" {
		return ""
	} else if !strings.Contains(name, "_") {
		// 不含下划线，原值返回
		return name
	}
	var result string = ""
	camels := strings.Split(name, "_")
	for index := range camels {
		if camels[index] == "" {
			continue
		}
		camel := camels[index]
		if result == "" {
			result = strings.ToLower(camel)
		} else {
			result = result + strings.ToUpper(camel[0:1]) + strings.ToLower(camel[1:len(camel)])
		}
	}
	return result
}
