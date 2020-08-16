package server

import (
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"symbolan/Node"
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

func page(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requesting webpage")
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current dir" + dir)
	http.ServeFile(w, r, "./webapp/build/index.html")
}

func Run(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/derivative", derivative)
	mux.HandleFunc("/index", page)

	fs := http.FileServer(http.Dir("./webapp/build"))

	//mux.Handle("/webpage/", http.StripPrefix("/static/", fs))
	mux.Handle("/", fs)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	fmt.Println("Server running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
