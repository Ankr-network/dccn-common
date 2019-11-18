package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/Ankr-network/dccn-common/access"
)

func main() {
	fmt.Println("test rbac authorize service")

	e := access.NewAuthorizationService("127.0.0.1:6801")

	fmt.Println(strings.Repeat("==", 20))
	tbl := []struct {
		In  string
		Exp bool
	}{
		{
			In:  "Owner,data/auth,read",
			Exp: true,
		},
		{
			In:  "Admin,ankr:team:hello:team/name,update",
			Exp: false,
		},
		{
			In:  "Admin,ankr:team:world:ankr:team/name,update",
			Exp: false,
		},
		{
			In:  "Admin,ankr:team:world:ankr:team,delete",
			Exp: false,
		},
		{
			In:  "Admin,ankr:payr:hello,delete",
			Exp: false,
		},
		{
			In:  "Admin,ankr:fee:hello,read",
			Exp: false,
		},
		{
			In:  "Admin,ankr:fee:hello,list",
			Exp: false,
		},
		{
			In:  "Admin,ankr:fee:hello,delete",
			Exp: false,
		},
		{
			In:  "Admin,ankr:payr:hello,read",
			Exp: false,
		},
		{
			In:  "Admin,ankr:payr:hello,list",
			Exp: false,
		},
		{
			In:  "Admin,ankr:hub:hello,read",
			Exp: true,
		},
		{
			In:  "Admin,ankr:hub:hello,list",
			Exp: true,
		},
		// member
		{
			In:  "Member,ankr:payr:hello,read",
			Exp: false,
		},
		{
			In:  "Member,ankr:payr:hello,list",
			Exp: false,
		},
		{
			In:  "Member,ankr:hub:hello,read",
			Exp: true,
		},
		{
			In:  "Member,ankr:hub:hello,list",
			Exp: true,
		},
	}

	for _, v := range tbl {
		s := strings.Split(v.In, ",")
		ok, err := e.Authorize(context.Background(), s[0], s[1], s[2])
		fmt.Printf("input: %s output: %v, expect: %v error: %v \n", v.In, ok, v.Exp, err)
	}

	// output:
	//		input: Owner,data/auth,read output: true, expect: true error: <nil>
	//		input: Admin,ankr:team:hello:team/name,update output: false, expect: false error: <nil>
	//		input: Admin,ankr:team:world:ankr:team/name,update output: false, expect: false error: <nil>
	//		input: Admin,ankr:team:world:ankr:team,delete output: false, expect: false error: <nil>
	//		input: Admin,ankr:payr:hello,delete output: false, expect: false error: <nil>
	//		input: Admin,ankr:fee:hello,read output: false, expect: false error: <nil>
	//		input: Admin,ankr:fee:hello,list output: false, expect: false error: <nil>
	//		input: Admin,ankr:fee:hello,delete output: false, expect: false error: <nil>
	//		input: Admin,ankr:payr:hello,read output: false, expect: false error: <nil>
	//		input: Admin,ankr:payr:hello,list output: false, expect: false error: <nil>
	//		input: Admin,ankr:hub:hello,read output: true, expect: true error: <nil>
	//		input: Admin,ankr:hub:hello,list output: true, expect: true error: <nil>
	//		input: Member,ankr:payr:hello,read output: false, expect: false error: <nil>
	//		input: Member,ankr:payr:hello,list output: false, expect: false error: <nil>
	//		input: Member,ankr:hub:hello,read output: true, expect: true error: <nil>
	//		input: Member,ankr:hub:hello,list output: true, expect: true error: <nil>

}
