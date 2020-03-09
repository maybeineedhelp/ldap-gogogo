package main

import (
	"fmt"

	ldapv3 "gopkg.in/ldap.v3"
)

func main() {
	accountName := "jiandan"
	password := "123456"
	conn, err := ldapv3.Dial("tcp", fmt.Sprintf("localhost:12396"))
	if err != nil {
		fmt.Println("connerr------", err)
		return
	}
	defer conn.Close()

	// if !m.insecure {
	// 	if err := conn.StartTLS(&tls.Config{InsecureSkipVerify: true}); err != nil {
	// 		return false, nil, err
	// 	}
	// }

	if err := conn.Bind("cn=admin,dc=applysquare,dc=org", "test2admin"); err != nil {
		fmt.Println("adminBinderr------", err)
		return
	}

	baseDN := "ou=People,dc=applysquare,dc=org"
	req := ldapv3.NewSearchRequest(
		baseDN,
		ldapv3.ScopeWholeSubtree, ldapv3.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(uid=%s)", accountName),
		[]string{"dn"},
		nil,
	)

	sr, err := conn.Search(req)
	if err != nil {
		fmt.Println("Searcherr------", err)
		return
	}
	if len(sr.Entries) != 1 {
		fmt.Println("sr.Entries>1------")
	}

	dn := sr.Entries[0].DN
	fmt.Println("Searchresult------", dn)
	err = conn.Bind(dn, password)
	fmt.Println("userBinderr------", err)
	if ldapv3.IsErrorWithCode(err, ldapv3.ErrorNetwork) {
		fmt.Println("userBinderrWithCode------", err)
	}
	fmt.Println("login ok")
	return
}
