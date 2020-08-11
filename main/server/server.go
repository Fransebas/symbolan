package server

import (
	"encoding/base64"
	"fmt"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"
	"symbolan/main/Latex"
	"symbolan/main/Node"
)

type DerivativeResponse struct {
	original string
	result   string
}

func Run() {
	derivate := func(w http.ResponseWriter, req *http.Request) {
		if (*req).Method == "OPTIONS" {
			w.Write([]byte("ok"))
			return
		}
		//w.Header().Set("Content-type", "application/pdf")
		inputBytes, _ := ioutil.ReadAll(req.Body)

		fmt.Println("input " + string(inputBytes))

		config := Node.SymbolanProcessorConfig{
			Simplify: Node.NUMERIC_FUNCTIONS_2,
		}

		symbolanProcessor := Node.NewSymbolanProcessor(&config)
		node := symbolanProcessor.ProcessString(string(inputBytes))
		derivativeNode := symbolanProcessor.ProcessString("D(x)")
		derivativeNode.Eval(&map[string]*Node.Node{
			"x": node,
		})

		firstMessage := "The input equation is "
		secondMessage := "The derivative is "

		latexStr := node.Latex()

		derivativeLatex := derivativeNode.Latex()

		latexDoc := fmt.Sprintf("\\documentclass[preview]{standalone}\n\\begin{document}\n%v\n%v\n%v\n%v\n\\end{document}", firstMessage, latexStr, secondMessage, derivativeLatex)
		fmt.Println(latexDoc)
		pdf, e := Latex.ProcessLatex(latexDoc)

		if e == nil {
			_, _ = w.Write([]byte(base64.StdEncoding.EncodeToString(pdf)))
		} else {
			fmt.Println(e)
			_, _ = w.Write([]byte(e.Error()))
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/derivate", derivate)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":7777", handler))
}
