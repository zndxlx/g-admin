package jwt

import (
    "github.com/dgrijalva/jwt-go"
    "errors"
    "g-admin/config"
    uuid "github.com/satori/go.uuid"
)

type JWT struct {
    SigningKey []byte
}

// Custom claims structure
type CustomClaims struct {
    UUID        uuid.UUID
    ID          int64
    Username    string
    NickName    string
    AuthorityId string
    BufferTime  int64
    jwt.StandardClaims
}

var (
    TokenExpired     = errors.New("Token is expired")
    TokenNotValidYet = errors.New("Token not active yet")
    TokenMalformed   = errors.New("That's not even a token")
    TokenInvalid     = errors.New("Couldn't handle this token:")
)

var _jwt *JWT

func Init()  {
    var _jwtConf = &config.Conf.JWT
    _jwt = &JWT{
        []byte(_jwtConf.SigningKey),
    }
}

// 创建一个token
func  CreateToken(claims CustomClaims) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(_jwt.SigningKey)
}

// 解析 token
func ParseToken(tokenString string) (*CustomClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
        return _jwt.SigningKey, nil
    })
    if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors&jwt.ValidationErrorMalformed != 0 {
                return nil, TokenMalformed
            } else if ve.Errors&jwt.ValidationErrorExpired != 0 {
                // Token is expired
                return nil, TokenExpired
            } else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
                return nil, TokenNotValidYet
            } else {
                return nil, TokenInvalid
            }
        }
    }
    if token != nil {
        if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
            return claims, nil
        }
        return nil, TokenInvalid
        
    } else {
        return nil, TokenInvalid
    }
}
