package app

import (
	"github.com/dgrijalva/jwt-go"
	"tinnydouban/global"
	"tinnydouban/pkg/util"
	"time"
)

//这块主要涉及 JWT 的一些基本属性，
//第一个是 GetJWTSecret 方法，用于获取该项目的 JWT Secret，目前我们是直接使用配置所配置的 Secret，
//第二个是 Claims 结构体，分为两大块，第一块是我们所嵌入的 AppKey 和 AppSecret，用于我们自定义的认证信息，
//第二块是 jwt.StandardClaims 结构体，它是 jwt-go 库中预定义的
type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

//在 GenerateToken 方法中，它承担了整个流程中比较重要的职责，
//也就是生成 JWT Token 的行为，
//主体的函数流程逻辑是根据客户端传入的 AppKey 和 AppSecret
//以及在项目配置中所设置的签发者（Issuer）和过期时间（ExpiresAt），
//根据指定的算法生成签名后的 Token。这其中涉及两个的内部方法，如下：
//
//jwt.NewWithClaims：
//根据 Claims 结构体创建 Token 实例，它一共包含两个形参，
//第一个参数是 SigningMethod，其包含 SigningMethodHS256、SigningMethodHS384、SigningMethodHS512
//三种 crypto.Hash 加密算法的方案。
//第二个参数是 Claims，主要是用于传递用户所预定义的一些权利要求，便于后续的加密、校验等行为。
//tokenClaims.SignedString：生成签名字符串，根据所传入 Secret 不同，进行签名并返回标准的 Token。
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

///在 ParseToken 方法中，它主要的功能是解析和校验 Token，
//承担着与 GenerateToken 相对的功能，其函数流程主要是解析传入的 Token，
//然后根据 Claims 的相关属性要求进行校验。这其中涉及两个的内部方法，如下：
//
//ParseWithClaims：用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回 *Token。
//Valid：验证基于时间的声明，例如：过期时间（ExpiresAt）、签发者（Issuer）、生效时间（Not Before），需要注意的是，如果没有任何声明在令牌中，仍然会被认为是有效的。
//至此我们就完成了 JWT 令牌的生成、解析、校验的方法编写，
//我们会在后续的应用中间件中对其进行调用，使其能够在应用程序中将一整套的动作给串联起来。
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
