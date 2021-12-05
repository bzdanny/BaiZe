package stringUtils

import "strings"

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	if str == "" {
		return str
	}
	var upperStr string
	b := []byte(str[:1])[0]
	if b >= 97 && b <= 122 {
		b -= 32 // string的码表相差32位
		upperStr += string(b)
	} else {
		return str
	}
	upperStr += str[1:]
	return upperStr
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

//将下划线大写方式命名的字符串转换为驼峰式。如果转换前的下划线大写方式命名的字符串为空，则返回空字符串。 例如：HELLO_WORLD->HelloWorld
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

//将下划线大写方式命名的字符串转换为驼峰式,首字母小写。如果转换前的下划线大写方式命名的字符串为空，则返回空字符串。 例如：HELLO_WORLD->helloWorld
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
