package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	JWT "github.com/dgrijalva/jwt-go"
	"os"
	"regexp"
	"time"
)

func test_jwt() {
	duration := time.Hour * 24 * 30 // 30 days
	expires := time.Now().Add(duration).Unix()

	// claims
	type MyCustomClaims struct {
		Foo      string `json:"foo"`
		ClientId string `json:"clientid"` // client/app uuid
		Arch     string `json:"arch"`     // arch (ia32/x64)
		Os       string `json:"os"`       // os type (win/osx/linux/...)
		OsVerion string `json:"osv"`      // os version
		HostId   string `json:"hostid"`   // os host UUID
		HDSerial string `json:"hds"`      // hard drive serial

		JWT.StandardClaims
	}

	claims := MyCustomClaims{
		"bar",
		"cid-343",
		"x64",
		"Win",
		"10.0.0.1563",
		"hid-343",
		"hds-343",
		JWT.StandardClaims{
			ExpiresAt: expires,
			Issuer:    "woodboard",
		},
	}

	// priv key
	priv_key, err := read_priv("/tmp/priv.pem")
	check(err)
	// pub key
	pub_key, err := read_pub("/tmp/pub.pem")
	check(err)

	// gen
	tkn, _ := gen_jwt(priv_key, claims)
	fmt.Printf("\n----- gen-ed token ----\n%v\n\n", tkn)

	// verify
	fmt.Println("\n---- verify ----: ")

	token, err := verify_jwt(pub_key, []byte(tkn))
	check(err)

	if !token.Valid {
		fmt.Printf("Error: Token is invalid")
	} else {
		fmt.Printf("  token verified")
	}
	printJSON(token.Claims)
	{
		claims2 := token.Claims.(JWT.MapClaims)
		expire2 := claims2["exp"].(float64)
		delta := time.Until(time.Unix(int64(expire2), 0))
		fmt.Printf("  %v to expire\n", delta)
	}
}

func gen_jwt(_key *ecdsa.PrivateKey, _claims JWT.Claims) (string, error) {
	// alg/signing-method
	alg_name := "ES384"
	method := JWT.GetSigningMethod(alg_name)
	if method == nil {
		return "", fmt.Errorf("can not find signing method: %v", alg_name)
	}

	// create a new token
	token := JWT.NewWithClaims(method, _claims)

	// write header (map string->interface{})
	token.Header["header-1"] = 2
	token.Header["header-2"] = "hi2"

	// sign
	out, err := token.SignedString(_key)
	check(err)
	return out, nil
}

// verify
func verify_jwt(_key *ecdsa.PublicKey, _token_data []byte) (*JWT.Token, error) {
	// trim whitespace
	token_str := regexp.MustCompile(`\s*$`).ReplaceAll(_token_data, []byte{})
	fmt.Fprintf(os.Stderr, "Token len: %v bytes\n", len(token_str))

	// parse
	token, err := JWT.Parse(string(token_str), func(t *JWT.Token) (interface{}, error) { return _key, nil })
	if token != nil {
		fmt.Fprintf(os.Stderr, "Header:\n  %v\n\n", token.Header)
		fmt.Fprintf(os.Stderr, "Claims:\n  %v\n\n", token.Claims)
	}
	check(err)
	return token, nil
}

// Print a json object in accordance with the prophecy (or the command line options)
func printJSON(j interface{}) error {
	var out []byte
	var err error

	out, err = json.MarshalIndent(j, "", "    ")
	if err == nil {
		fmt.Println(string(out))
	}
	return err
}

/*
// showToken pretty-prints the token on the command line.
func showToken() error {
  // get the token
  tokData, err := loadData(*flagShow)
  if err != nil {
    return fmt.Errorf("Couldn't read token: %v", err)
  }

  // trim possible whitespace from token
  tokData = regexp.MustCompile(`\s*$`).ReplaceAll(tokData, []byte{})
  if *flagDebug {
    fmt.Fprintf(os.Stderr, "Token len: %v bytes\n", len(tokData))
  }

  token, err := JWT.Parse(string(tokData), nil)
  if token == nil {
    return fmt.Errorf("malformed token: %v", err)
  }

  // Print the token details
  fmt.Println("Header:")
  if err := printJSON(token.Header); err != nil {
    return fmt.Errorf("Failed to output header: %v", err)
  }

  fmt.Println("Claims:")
  if err := printJSON(token.Claims); err != nil {
    return fmt.Errorf("Failed to output claims: %v", err)
  }

  return nil
}
*/
