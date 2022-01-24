package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"go01/api"
	"go01/security"
	"go01/util"
	"reflect"
	"testing"
	"time"
)

func TestMd5(t *testing.T) {
	plain := "hello"
	encodeString := security.Md5(plain)
	fmt.Println(plain, ",md5:", encodeString, ",len:", len(encodeString))
	t.FailNow()
}

type Person struct {
	B      bool
	Int1   int
	Int2   int8
	Int3   int16
	Int4   int32
	Int5   int64
	Uint1  uint
	Uint2  uint8
	Uint3  uint16
	Uint4  uint32
	Uint5  uint64
	Float1 float32
	Float2 float64
	//Float3 complex64
	//Float4 complex128
	Array2  [16]int
	Array1  [16]int
	Map2    map[string]string
	Ptr1    *Person
	Slice1  []Person
	Slice2  []int
	String1 string
	Struct1 *Person
	Struct2 Tag
}

type Tag struct {
	Id   int
	Name string
}

func TestGenString(t *testing.T) {
	a := Person{Map2: map[string]string{"Array2": "xx", "Array1": "xx", "B": "xx"}}
	bs, _ := json.Marshal(a)
	fmt.Println(string(bs))

}

func TestFloat(t *testing.T) {
	a := 1.123456789455666555885245524524522
	fmt.Println(a)
}

func TestRsa(t *testing.T) {
	privateKey, publicKey, err := security.GenRsaKeys(2048)
	if err != nil {
		panic(err)
	}
	fmt.Println("private key:", privateKey, "public key:", publicKey)
}

func TestError(t *testing.T) {
	//var pathError *fs.PathError
	//var err error = &fs.PathError{Path: "/"}
	//if errors.As(err, &pathError) {
	//	fmt.Println("Failed at path:", pathError.Path)
	//} else {
	//	fmt.Println(err)
	//}
	var err error = &util.ServiceError{Message: "xxx"}
	var target *util.ServiceError
	if errors.As(err, &target) {
		fmt.Println("yes")
	} else {
		fmt.Println(err)
	}
}

func TestSign(t *testing.T) {
	privateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAtbkotioMGMkmtxhnJqa0EnYESNc4nm6CSYhEugp1+vOKKxRt\n2kDH6Mf8bwfcv3jdPgVWYBV8JEl0BcaIpIh4yM+secShcXjAVdaT4AAK7X2lNxwH\n8jzypedos2b7A/ajXofCAk1J2LTEAwQ5nnnVuZkQK1AiF8EKQJ63pkNTeKUxO/r4\nXXFCj9t/Z0KJzXXTzZKg3s0Ildwz5uNsZwVQ5doBHXfe8z6iPy/kQ39PnH6vjz3b\nG7w86QXcKlto7O0EGHBPc08wgfBcojYZuVZdZyi2LadqHN7MCBn7+yGJOldW9JFk\nA7QJCuwUe4OhYilDIkEz/st3f+3EyiT/Y3SEjwIDAQABAoIBAQCVCjUfokClm2RL\nSpJfGt4ZPf2pmApxtgTZNg/X3XiBx3gMeQbXt8W9XzfkADjGOOSLV2lOCJD2Wd/0\nVK4A84TMfbbhb6lRHU3xmLlTP/W2bhQwrKa/v8jq1G4EpXM9/uxiPUmbBbSJLUiN\niMuQ9C+btrOSMnijrskv0nyguhsdmEqDyl9GK6Om+8s+InbTYF2UEkfWmt3byMij\no/ZHeHOrNMAdtWgaO8tuP4lVgXSoSF1jGZ7iOpFJfqbDQFF6hzAH6qV9GMwMYp33\nuvn/JjJ/qv5hJyV8iuke/09CfZbElCGmUQQRxBH8V1oBaHFnelFXREaJAkRSd29g\n67l8MN8BAoGBAMHLrKBCbvpy275NmIX4xU9WcHp/fmry0oLbajW8VM/ki9VIGgJL\noSc7pXSifXsJ/QCsujv92h3AGVuFIe9bXOq3wVfzG3kssekzpOAY3HnsQem8Eg7P\nkzE1V/sbcUH7TtrgTlcWlWRyTp6azjnXkSAPmzRfy0ohiPt2tvZYPbLBAoGBAPAN\nfjITqQwA/xiQgbUo6La9QusXI09twhrVeskgEdmVycJkgoRrxqRki3AvnGHUEPfw\nZcwKQ2FVclrR3aKhHW/iPNf46xjg+yancKYfZOB4fAaxIX5z62+k7Z6s7qfEpOmm\nV9LuR/nDLbRmasSFXbZ0Yfovdjsew5ODoi+akBtPAoGAP+JIUbwUoXLjhWRG90L+\nqByyj28f2Vmak5CI/pXKz41jmzdde4w635gF/uDhxIGSaXqHGeeg01XUBhtpCGJx\nyt4wWjHFyLg3HczseQM5CugbAlYBDejXeM1A54IwX7PcVsLCkGrdbHNR/27AtMif\nCpaabzV06kcLxPlpbuO/0wECgYAh7OY+0YR8i8+Bezq5jZSF7u18KAL3gL2D6VO3\nBO9A3uhqrqW7bTffl84VBsWFFeFoCPN6CKFJKjhFGY/HIhn06/ZJV4ZyN6mG5vcG\noz0wdBajI3lmU5+cYaSrXilEUIg19SpIRyCo7aqR6j+AkpCR7pTCNN7ysABX4qyT\nKbbgBQKBgQCWSNbR4X1SSaFrTypxuHy3Gba30FckDEftE7UApqZBeM8zGPhvqb1P\nBDv9M9FCk54uLj4wvYTgyfb6p9LglFrs32lvEcejUHUSC4o9hksr+Exd9LylF11y\nu8ByLUudj4uwt+DkC/M4hiW4KGUM7PKx6jWVxJMbkfBMUnfIM7KwnA==\n-----END RSA PRIVATE KEY-----\n"
	publicKey := "-----BEGIN RSA PUBLIC KEY-----\nMIIBCgKCAQEAtbkotioMGMkmtxhnJqa0EnYESNc4nm6CSYhEugp1+vOKKxRt2kDH\n6Mf8bwfcv3jdPgVWYBV8JEl0BcaIpIh4yM+secShcXjAVdaT4AAK7X2lNxwH8jzy\npedos2b7A/ajXofCAk1J2LTEAwQ5nnnVuZkQK1AiF8EKQJ63pkNTeKUxO/r4XXFC\nj9t/Z0KJzXXTzZKg3s0Ildwz5uNsZwVQ5doBHXfe8z6iPy/kQ39PnH6vjz3bG7w8\n6QXcKlto7O0EGHBPc08wgfBcojYZuVZdZyi2LadqHN7MCBn7+yGJOldW9JFkA7QJ\nCuwUe4OhYilDIkEz/st3f+3EyiT/Y3SEjwIDAQAB\n-----END RSA PUBLIC KEY-----\n"
	src := "xwww"
	sign, err := security.SignWithPrivateKey(privateKey, []byte(src))
	if err != nil {
		panic(err)
	}
	fmt.Println(sign)

	b, err := security.VerifyWithPublicKey(sign, []byte(src), publicKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}

func TestDate(t *testing.T) {
	ds1 := "\"2020-10-12\""
	ds2 := "\"2020-10-12 14:15:12\""
	var date api.JsonDate
	var datetime api.JsonDateTime
	err := json.Unmarshal([]byte(ds1), &date)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(ds2), &datetime)
	if err != nil {
		panic(err)
	}
	bytes, err := json.Marshal(&date)
	if err != nil {
		panic(err)
	}
	fmt.Println("date:", string(bytes))
	bytes, err = json.Marshal(&datetime)
	if err != nil {
		panic(err)
	}
	fmt.Println("datetime:", string(bytes))

}

func TestCreateMerchant(t *testing.T) {
	query := &api.MerchantCreateQuery{
		BusinessNo:  "xxxxxqqaaaqzq",
		CompanyName: "小明的公司",
		UserName:    "张三",
		PhoneNo:     "15652087252",
		IdentifyNo:  "36253291992000999",
	}
	vo := new(util.ApiResult)
	util.HttpPost("https://localhost:8080/merchant/create", query, vo)
	fmt.Println("resp result:", vo)
}

func TestHttpSign(t *testing.T) {
	myPrivate := "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAzUXp2skVoOWa54Piza/cVoWx1O5lA3Uw6oxoiq2bpVgWQa11\nU9XjVZxcDGRT2ti/68/xpUW0SOrAGdMFU1W4IdufHBafGtIJK2risHM5Zbs5yWXH\npanbzDDkwKye4mtVjw9mvoDn2SX1AtfjYD94sX80caje2q6N3Z3Nqjabnekf4Ugm\ndOTgFW4i3t+uTX/kjMXxz2z6ntv8smkMU0+BC1edh0oshfPSmR2C0O25JajXrOy1\nhQefKV+1rXOblt0I5RnlfNN4FWY3CA0OBfY28UnChMDL8Fqde3Q0HZuWL3CQZ5Cm\njtltO/9vPnBzcDMI7ZxDQIJefAD/LzhKmiafiQIDAQABAoIBAGdr0Q3dABuncusw\nBIYzE2U4SMeKMLSUR8f0Y/tyypd8kQIYHaPWgIoylCHeLm11hQSPKSVmwBV4yL56\nFhXNe077sudY8mJ17g7E9F+IPiPmN5AjynnJ4aa3/5ifoBFdmKbSUKaslaNPP2lm\nUfW9JNbwywB034r4JRvCFMusZa0hz2JGZzZaxVU7PSKwJB99rq5wEjM6ugGkm5dm\n+oJVz7N5eP63FWHI1EbuEH4KSjAKa3XIs0LZF6URLCeDk1fHiKrlmZLzEmHxoHBY\n//h9pLnHNVEu+ZlLJqEfSSczz2AoxpD4TVrza15n0FpsZtrWova0ygJ4jyKOA9v0\nkecY1rECgYEA5iFbc82sGt4iW8Vc6EBMq8WrJ1mxVKRPsvr+nK7vX/kQBphF7NYh\nV+O4vrkS4uctTcVY+YnNHfsFntKztwKUSnp6mG5zs0RlRddNWb3w9KnZcGemt5hk\n9j9kTkzhL8LgUAyRsbtcPApn5kFdx/moipS1FmDxH5f3IgRAjhRghS0CgYEA5Fk4\nhP88NxX/7RshDBCk25mUjmji/F5nP+CwXLLlBGnnapWbaw8JWP39uU7Licm/PCeh\nzkw2/RcOI1jhGb9QnBkOtaYszhbKFAKFedOqMmm1QMBdWN8LQwQJ7Pm2G68vTwDI\nH7Lphh9m1vWznU3TWd+9Q6Bt6vRneMfhLwfldU0CgYBv3jNInehVvznQhqZdODSN\nI6/Jly4+yQj9Jkny/I7choQl0IvdT1IyGT7LZ8goiNxa+93+n+AX/NiElkRKaHSR\np+xEgRy5PjxEsR6my1CAfrh2jcxbXzdlBPfLoE+vGfpUr2d7kGqLvTs4kPn3pgMq\nKpPE5ycPhp3Z6VtANeOdmQKBgQC3lSeEfXMa0nwUUzhJa+QujRXLPHYo8XjO+USw\n8j5/QumDXk46TzdzrZfb8IZg+6rcoPxMUCfxaCS8tEsdGRuks9yIm9XvxpCmb7SB\n0nNF3oiTS1SBC6kRMuEL/WK0dS5ytK0wFvX6F6rlbVn/QL+HxQJzqTpVqDK7/u6C\nkjWfDQKBgQCtnGU6ULLLcxygA2DUalMktO1QiCT6aFL5NLfIE8HYxsOD1xWjGhEr\nyhFdrkqMWEbBJ5okwS0jDpa8QuX0ZqqSiSOVg26f+IPsBnpaABI+AOV8YJf4wx0B\nxZ5anZBDxmgc7mcREJiXaog1xDrTMzgYcq5Zj1IKZmiiZ7VYfP3v3w==\n\n-----END RSA PRIVATE KEY-----"
	//privateKey1 := "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAtbkotioMGMkmtxhnJqa0EnYESNc4nm6CSYhEugp1+vOKKxRt\n2kDH6Mf8bwfcv3jdPgVWYBV8JEl0BcaIpIh4yM+secShcXjAVdaT4AAK7X2lNxwH\n8jzypedos2b7A/ajXofCAk1J2LTEAwQ5nnnVuZkQK1AiF8EKQJ63pkNTeKUxO/r4\nXXFCj9t/Z0KJzXXTzZKg3s0Ildwz5uNsZwVQ5doBHXfe8z6iPy/kQ39PnH6vjz3b\nG7w86QXcKlto7O0EGHBPc08wgfBcojYZuVZdZyi2LadqHN7MCBn7+yGJOldW9JFk\nA7QJCuwUe4OhYilDIkEz/st3f+3EyiT/Y3SEjwIDAQABAoIBAQCVCjUfokClm2RL\nSpJfGt4ZPf2pmApxtgTZNg/X3XiBx3gMeQbXt8W9XzfkADjGOOSLV2lOCJD2Wd/0\nVK4A84TMfbbhb6lRHU3xmLlTP/W2bhQwrKa/v8jq1G4EpXM9/uxiPUmbBbSJLUiN\niMuQ9C+btrOSMnijrskv0nyguhsdmEqDyl9GK6Om+8s+InbTYF2UEkfWmt3byMij\no/ZHeHOrNMAdtWgaO8tuP4lVgXSoSF1jGZ7iOpFJfqbDQFF6hzAH6qV9GMwMYp33\nuvn/JjJ/qv5hJyV8iuke/09CfZbElCGmUQQRxBH8V1oBaHFnelFXREaJAkRSd29g\n67l8MN8BAoGBAMHLrKBCbvpy275NmIX4xU9WcHp/fmry0oLbajW8VM/ki9VIGgJL\noSc7pXSifXsJ/QCsujv92h3AGVuFIe9bXOq3wVfzG3kssekzpOAY3HnsQem8Eg7P\nkzE1V/sbcUH7TtrgTlcWlWRyTp6azjnXkSAPmzRfy0ohiPt2tvZYPbLBAoGBAPAN\nfjITqQwA/xiQgbUo6La9QusXI09twhrVeskgEdmVycJkgoRrxqRki3AvnGHUEPfw\nZcwKQ2FVclrR3aKhHW/iPNf46xjg+yancKYfZOB4fAaxIX5z62+k7Z6s7qfEpOmm\nV9LuR/nDLbRmasSFXbZ0Yfovdjsew5ODoi+akBtPAoGAP+JIUbwUoXLjhWRG90L+\nqByyj28f2Vmak5CI/pXKz41jmzdde4w635gF/uDhxIGSaXqHGeeg01XUBhtpCGJx\nyt4wWjHFyLg3HczseQM5CugbAlYBDejXeM1A54IwX7PcVsLCkGrdbHNR/27AtMif\nCpaabzV06kcLxPlpbuO/0wECgYAh7OY+0YR8i8+Bezq5jZSF7u18KAL3gL2D6VO3\nBO9A3uhqrqW7bTffl84VBsWFFeFoCPN6CKFJKjhFGY/HIhn06/ZJV4ZyN6mG5vcG\noz0wdBajI3lmU5+cYaSrXilEUIg19SpIRyCo7aqR6j+AkpCR7pTCNN7ysABX4qyT\nKbbgBQKBgQCWSNbR4X1SSaFrTypxuHy3Gba30FckDEftE7UApqZBeM8zGPhvqb1P\nBDv9M9FCk54uLj4wvYTgyfb6p9LglFrs32lvEcejUHUSC4o9hksr+Exd9LylF11y\nu8ByLUudj4uwt+DkC/M4hiW4KGUM7PKx6jWVxJMbkfBMUnfIM7KwnA==\n-----END RSA PRIVATE KEY-----\n"

	publicKey := "-----BEGIN RSA PUBLIC KEY-----\nMIIBCgKCAQEA1OcATmzTFlQJdfEE9In6MDrIDVjhi0Y7r/g9LkCoyOC0VakkszR9\nomFMVG4SVVoV92JftvstJMiAbOuMEM8Rgv8j+mVmVu2qk42DTnmeNzmyTpis3FfD\n8qY5ecH5Pa1xzTTj1eVv6X22WX7nVDIQ8gx6K5oEgAYThNyK2Sp9XFgUZTTn5MgI\naNwVMlb/PQyV20d04Ds3MyHuuIvvGbRiQF1dlzhSSWsRoN5qFBa089o77m4yG8Ou\n/+MoXkhIbCU/zu4NGhgplnOdNWwuUb8iU3qo89L2qNjauBTlp1YIFg87pueIkIys\neL0h6bZIFHNukJOgw1u+bm+07au92FxM2QIDAQAB\n-----END RSA PUBLIC KEY-----"
	planQuery := &api.PayQuery{
		MerchantQuery: api.MerchantQuery{
			MerchantId: 2,
			SecretKey:  "ddac0389790578d9372beebe7f0ac2d0",
		},
		TradeId:   "xsdc",
		PayMoney:  100,
		TradeDate: api.JsonDate(time.Now()),
	}
	bs, err := json.Marshal(planQuery)
	util.AssertError(err, "json转换异常")
	signData, err := security.SignWithPrivateKey(myPrivate, bs)
	util.AssertError(err, "加签失败")

	encrptData, err := security.EncryptWithPublicKey(bs, publicKey)
	util.AssertError(err, "publicKey encode fail!")

	signQuery := api.SignQuery{
		SignData:    signData,
		EncryptData: encrptData,
	}
	vo := new(util.ApiResult)
	util.HttpPost("https://localhost:8080/pay", signQuery, vo)
}

func TestF(t *testing.T) {
	a := &api.PayQuery{
		MerchantQuery: api.MerchantQuery{
			MerchantId: 10,
			SecretKey:  "xxxx",
		},
		TradeId:   "",
		PayMoney:  0,
		TradeDate: api.JsonDate{},
	}
	m := reflect.ValueOf(a).FieldByName("MerchantQuery").Interface().(api.MerchantQuery)
	fmt.Println(m)
}
