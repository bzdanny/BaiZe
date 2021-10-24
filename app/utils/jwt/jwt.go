package jwt

import (
	redis "baize/app/common/redis"
	"baize/app/constant/constants"
	setting "baize/app/setting"
	loginModels "baize/app/system/models/loginModels"
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
	c := loginModels.JWT{
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
func ParseToken(tokenString string) (*loginModels.LoginUser, error) {
	// 解析token
	var mc = new(loginModels.JWT)
	_, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	loginUser, err := redis.GetCacheLoginUser(constants.LoginTokenKey + mc.TokenId)

	//id := mc.TokenId
	return loginUser, err
}
