package server

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"
	"symbolan/main/Node"
)

type DerivativeResponse struct {
	Original string
	Result   string
}

func derivative(w http.ResponseWriter, req *http.Request) {
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

	configNoSimplification := Node.SymbolanProcessorConfig{
		Simplify: Node.NONE,
	}

	symbolanProcessor := Node.NewSymbolanProcessor(&config)
	symbolanNoSimplyfierProcessor := Node.NewSymbolanProcessor(&configNoSimplification)
	node := symbolanProcessor.ProcessString(string(inputBytes))
	derivativeNode := *symbolanNoSimplyfierProcessor.ProcessString("D(x)")
	derivativeNode = derivativeNode.Eval(&map[string]*Node.Node{
		"x": node,
	})

	simplyfier := Node.NewSimplyfier()

	derivativeNode = simplyfier.Simplify(derivativeNode)

	latexStr := node.Latex()

	derivativeLatex := derivativeNode.Latex()

	res := DerivativeResponse{
		Original: latexStr,
		Result:   derivativeLatex,
	}

	jsonResponse, e := json.Marshal(res)

	if e == nil {
		_, _ = w.Write(jsonResponse)
	} else {
		fmt.Println(e)
		_, _ = w.Write([]byte(e.Error()))
	}
}

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/derivative", derivative)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":7777", handler))
}
