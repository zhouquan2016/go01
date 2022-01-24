package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"go01/api"
	"go01/security"
	"go01/util"
	"testing"
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
	sign, err := security.SignWithRsa(privateKey, src)
	if err != nil {
		panic(err)
	}
	fmt.Println(sign)

	b, err := security.VerifyWithRsa(sign, src, publicKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}

//func TestVerify(t *testing.T) {
//	publicKey := "-----BEGIN RSA PUBLIC KEY-----\nMIIBCgKCAQEAtbkotioMGMkmtxhnJqa0EnYESNc4nm6CSYhEugp1+vOKKxRt2kDH\n6Mf8bwfcv3jdPgVWYBV8JEl0BcaIpIh4yM+secShcXjAVdaT4AAK7X2lNxwH8jzy\npedos2b7A/ajXofCAk1J2LTEAwQ5nnnVuZkQK1AiF8EKQJ63pkNTeKUxO/r4XXFC\nj9t/Z0KJzXXTzZKg3s0Ildwz5uNsZwVQ5doBHXfe8z6iPy/kQ39PnH6vjz3bG7w8\n6QXcKlto7O0EGHBPc08wgfBcojYZuVZdZyi2LadqHN7MCBn7+yGJOldW9JFkA7QJ\nCuwUe4OhYilDIkEz/st3f+3EyiT/Y3SEjwIDAQAB\n-----END RSA PUBLIC KEY-----\n"
//	sign := "pm34wvzPORw3HnbSfj9ynLMEFIvYNxZrscktKnj+xGl6t44pnZPi8Z4F4Tem0XZ0/l3d2pWXh3UnwFObY1fs91jVrTotxNKyTc51GiwZfSOyn09kKZBaHNOZ/W3u04LyWUrX7+rrW3fpRqJEH1jVhpIKDNnb5K68lk7DTNsrqfqxS96jbduCOfgIZo2DGwvdAmXJUK4LwKNdSiwIZs1aYcXsFg4nKa0nOvZ55lED8eZzbzhmRV+Hr5M6wpTkR+AxYpiiWFXJjuikE5zGANkDH38qLeLkaLxqNdfLY5i/shs9LcsoKKR/Y/0KIXi/GYXrYHJPwhae/p0696R5aRD1tQ=="
//	//b, err := security.VerifyWithRsa(sign, publicKey)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(b)
//}

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

func TestHttpSign(t *testing.T) {

}
