package Latex

import (
	"github.com/rwestlund/gotex"
	"log"
)

func ProcessLatex(latexStr string) ([]byte, error) {
	var pdf, err = gotex.Render(latexStr, gotex.Options{
		Command: "/Library/TeX/texbin/pdflatex",
	})

	if err != nil {
		log.Println("render failed ", err)
		return nil, err
	} else {
		// Do something with the PDF file, like send it to an HTTP client
		// or write it to a file.
		return pdf, nil
		//_ = ioutil.WriteFile("image.pdf", pdf, 0644)
	}
}
