package main

import (
	"fmt"

	"git.arvaninternal.ir/argo/lang"
)

func main(){
    lang.SetTranslationFilesPath("translate")
    lang.SetDefaultLocale("fa")
    lang.SetFallbackLocales("en")
    lang.Init()

	fmt.Println(lang.TrBy("fa", "msg.wallet_created"))
	fmt.Println(lang.TrBy("en", "msg.wallet_created"))
	fmt.Println(lang.TrBy("fr", "msg.wallet_created"))

}