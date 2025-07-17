package main

import (
	"fmt"
	"strconv"
	"strings"
)

//var environment *env.Environment
//var dbConn *database.Connection

func init() {
	//secrets := boost.Ignite("./../.env", pkg.GetDefaultValidator())
	//
	//environment = secrets
	//dbConn = boost.MakeDbConnection(environment)
}

func expandSimple(seed string) string {
	parts := strings.Split(seed, "<")
	repetitions := parts[0] // number | xy>

	rep, err := strconv.Atoi(repetitions)
	if err != nil {
		panic(fmt.Sprintf("init: Error converting repetitions: %s", err))
	}

	foo := strings.Split(parts[1], ">") // xy | >

	var output string
	for i := 0; i < rep; i++ {
		output += foo[0]
	}

	return output
}

func expand(seed string) string {
	if !strings.Contains(seed, "<") || !strings.Contains(seed, ">") {
		return seed
	}

	if strings.Count(seed, "<") == 1 && strings.Count(seed, ">") == 1 {
		return expandSimple(seed)
	}

	// look for patters like: 3<x>z, 2<x>x, end with (z or x)

	// expand("2<x>z") // "xxz"
	// expand("2<xy>3<z>")  // "xyxyzzz"
	// expand("3<x2<y>>")  // "xyyxyyxyy"

	//
	//fmt.Println("in: ", strings.Count(seed, "<"))
	//
	////parts := strings.Split(seed, "<")
	////cicles := parts[0]
	////
	////fmt.Println(cicles)

	return "nested: " + seed
}

func main() {
	fmt.Println("--> ", expand("xx"))
	fmt.Println("--> ", expand("4<xy>"))
	fmt.Println("--> ", expand("2<xy>3<z>")) //"2<x>z"

	//cli.ClearScreen()
	//
	//menu := panel.MakeMenu()
	//
	//for {
	//	err := menu.CaptureInput()
	//
	//	if err != nil {
	//		cli.Errorln(err.Error())
	//		continue
	//	}
	//
	//	switch menu.GetChoice() {
	//	case 1:
	//		if err = createBlogPost(menu); err != nil {
	//			cli.Errorln(err.Error())
	//			continue
	//		}
	//
	//		return
	//	case 2:
	//		if err = createNewApiAccount(menu); err != nil {
	//			cli.Errorln(err.Error())
	//			continue
	//		}
	//
	//		return
	//	case 3:
	//		if err = showApiAccount(menu); err != nil {
	//			cli.Errorln(err.Error())
	//			continue
	//		}
	//
	//		return
	//	case 4:
	//		if err = generateApiAccountsHTTPSignature(menu); err != nil {
	//			cli.Errorln(err.Error())
	//			continue
	//		}
	//
	//		return
	//	case 5:
	//		if err = generateAppEncryptionKey(); err != nil {
	//			cli.Errorln(err.Error())
	//			continue
	//		}
	//
	//		return
	//	case 0:
	//		cli.Successln("Goodbye!")
	//		return
	//	default:
	//		cli.Errorln("Unknown option. Try again.")
	//	}
	//
	//	cli.Blueln("Press Enter to continue...")
	//
	//	menu.PrintLine()
	//}
}

//
//func createBlogPost(menu panel.Menu) error {
//	input, err := menu.CapturePostURL()
//
//	if err != nil {
//		return err
//	}
//
//	httpClient := pkg.MakeDefaultClient(nil)
//	handler := posts.MakeHandler(input, httpClient, dbConn)
//
//	if _, err = handler.NotParsed(); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func createNewApiAccount(menu panel.Menu) error {
//	var err error
//	var account string
//	var handler *accounts.Handler
//
//	if account, err = menu.CaptureAccountName(); err != nil {
//		return err
//	}
//
//	if handler, err = accounts.MakeHandler(dbConn, environment); err != nil {
//		return err
//	}
//
//	if err = handler.CreateAccount(account); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func showApiAccount(menu panel.Menu) error {
//	var err error
//	var account string
//	var handler *accounts.Handler
//
//	if account, err = menu.CaptureAccountName(); err != nil {
//		return err
//	}
//
//	if handler, err = accounts.MakeHandler(dbConn, environment); err != nil {
//		return err
//	}
//
//	if err = handler.ReadAccount(account); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func generateApiAccountsHTTPSignature(menu panel.Menu) error {
//	var err error
//	var account string
//	var handler *accounts.Handler
//
//	if account, err = menu.CaptureAccountName(); err != nil {
//		return err
//	}
//
//	if handler, err = accounts.MakeHandler(dbConn, environment); err != nil {
//		return err
//	}
//
//	if err = handler.CreateSignature(account); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func generateAppEncryptionKey() error {
//	var err error
//	var key []byte
//
//	if key, err = auth.GenerateAESKey(); err != nil {
//		return err
//	}
//
//	decoded := fmt.Sprintf("%x", key)
//
//	cli.Successln("\n  The key was generated successfully.")
//	cli.Magentaln(fmt.Sprintf("  > Full key: %s", decoded))
//	cli.Cyanln(fmt.Sprintf("  > First half : %s", decoded[:32]))
//	cli.Cyanln(fmt.Sprintf("  > Second half: %s", decoded[32:]))
//	fmt.Println(" ")
//
//	return nil
//}
