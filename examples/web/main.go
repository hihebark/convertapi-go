package main

import (
	"fmt"
	"github.com/hihebark/convertapi-go"
	"github.com/hihebark/convertapi-go/config"
	"github.com/hihebark/convertapi-go/param"
	"os"
)

func main() {
	config.Default.Secret = os.Getenv("CONVERTAPI_SECRET") // Get your secret at https://www.convertapi.com/a

	fmt.Println("Converting remote PPTX to PDF")
	pptxRes := convertapi.ConvDef("web", "pdf",
		param.NewString("url", "https://en.wikipedia.org/wiki/Data_conversion"),
		param.NewString("filename", "web-example"),
	)

	if files, errs := pptxRes.ToPath("/tmp"); errs == nil {
		fmt.Println("PDF file saved to: ", files[0].Name())
	} else {
		fmt.Println(errs)
	}
}
