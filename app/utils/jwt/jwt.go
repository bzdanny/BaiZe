package jwt

import (
	"github.com/bzdanny/BaiZe/app/constant/constants"
	"github.com/bzdanny/BaiZe/app/setting"
	"github.com/bzdanny/BaiZe/app/system/systemModels"
	"github.com/bzdanny/BaiZe/app/utils/redisUtils"
	"github.com/dgrijalva/jwt-go"
)

var mySecret []byte
var issuer string

func Init(tokenConfig *setting.TokenConfig) {
	mySecret = []byte(tokenConfig.Secret)
	issuer = tokenConfig.Issuer
}

// GenToken 生成JWT
func GenToken(tokenId string) string {
	// 创建一个我们自己的声明的数据
	c := systemModels.JWT{
		tokenId,
		jwt.StandardClaims{
			Issuer: issuer, // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	signedString, err := token.SignedString(mySecret)
	if err != nil {
		panic(err)
	}
	return signedString
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (loginUser *systemModels.LoginUser, err error) {
	// 解析token
	var mc = new(systemModels.JWT)
	_, err = jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	loginUser, err = redisUtils.GetStruct(constants.LoginTokenKey+mc.TokenId, loginUser)

	//id := mc.TokenId
	return
}
