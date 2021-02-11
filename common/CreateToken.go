package common
import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)
import "github.com/bwmarrin/snowflake"
var idGen,_ = snowflake.NewNode(1)
var secertKey = "Server"
func CreateToken(user string,password string)(string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"Id":idGen.Generate().String(),
		"NotBefore":time.Now().Unix(),
		"ExpiresAt":time.Now().Add(time.Hour).Unix(),
		"User":user+password,
	})
	return token.SignedString([]byte(secertKey))
}
func CheckToken(clientToken string)(jwt.MapClaims,error){
	token,err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secertKey),nil
	})
	if err!=nil{
		return nil,fmt.Errorf("pass token fail = %v",err)
	}
	claim,ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid{
		return nil,fmt.Errorf("error")
	}
	return claim,nil
}