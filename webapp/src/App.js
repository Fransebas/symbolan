import React, {useState} from 'react';
import './App.css';
import MathJax from 'react-mathjax2'


function FormulaList({formulaList, name}) {
    let functionsDOM = formulaList.map((functionString) => (
        <label>
            <MathJax.Node inline>{functionString} </MathJax.Node>
        </label>
    ));

    return (
        <div>
            {name} : {functionsDOM}
        </div>
    )
}

function App() {

    const [input, setInput] = useState("");

    const [original, setOriginal] = useState("");
    const [result, setResult] = useState("");

    const [pdfBase64, setPdfBase64] = useState(null);

    function changeInput(e) {
        setInput(e.target.value);
    }

    function onClick() {
        var xhr = new XMLHttpRequest();

        var response = {};
        // get a callback when the server responds
        xhr.addEventListener('load', () => {
            // update the state of the component with the result here
            response = JSON.parse(xhr.responseText);
            setOriginal(response.Original);
            setResult(response.Result);
            //setPdfBase64(xhr.responseText);
        });
        // open the request with the verb and the url
        xhr.open('POST', '/derivative');
        // send the request
        xhr.send(input)
    }

    let trigonometricsString = ["sin(x)", "cos(x)", "tan(x)"];
    let hiperbolicsString = ["sinh(x)", "cosh(x)", "tanh(x)"];
    let exponentAndLogsString = ["log(x)", "e^x"];
    let elementaryString = ["+", "-", "*", "/", "^"];
    let variablesString = ["x", "y", "z", "x_{56}"];
    let constantsString = ["a", "b", "c", "a_{88}"];

    return (
        <div className="App">
            <div className="column">
                <div>
                    <h1 style={{
                        margin: "0px",
                        bottomMargin: "0px"
                    }}> Symbolan </h1>
                    <h6 style={{
                        margin: "0px"
                    }}> (Very original name) </h6>
                </div>

                <div>
                    <h3> Source code here (WIP) </h3>
                    <a href="https://github.com/Fransebas/symbolan"> https://github.com/Fransebas/symbolan </a>
                </div>




                <h3> Derivate </h3>
                <div className="flexItem inputDiv">
                    <input value={input} onChange={changeInput}/>
                    <button id={"submitButton"} onClick={onClick}> Derivate </button>
                </div>

                <MathJax.Context input='tex'>
                    <div id="resultBox">
                        <h5 className="resultTitle"> Input: </h5>
                        <MathJax.Node>{original} </MathJax.Node>
                        <h5 className="resultTitle"> Differential: </h5>
                        <MathJax.Node>{result} </MathJax.Node>
                        <div className="bottomBox"/>
                    </div>
                </MathJax.Context>

                <div className="explanation">
                    <h5 className="resultTitle"> Valid functions </h5>
                    <MathJax.Context input='tex'>
                        <div>
                            <FormulaList formulaList={elementaryString} name={"Elementary Functions"}/>
                            <FormulaList formulaList={trigonometricsString} name={"Trigonometric"}/>
                            <FormulaList formulaList={hiperbolicsString} name={"Hiperbolics"}/>
                            <FormulaList formulaList={exponentAndLogsString} name={"Logs and exponentials"}/>
                            <FormulaList formulaList={variablesString} name={"Variables"}/>
                            <FormulaList formulaList={constantsString} name={"Constants"}/>
                            <p> Where <MathJax.Node inline> x </MathJax.Node> can be any function </p>
                            <h5> Example:</h5>
                            <p> sin(7x^2) + log(sinh(y)) + a_9 </p>
                        </div>
                    </MathJax.Context>
                    <div className="bottomBox"/>
                </div>

            </div>
        </div>
    );
}

export default App;
