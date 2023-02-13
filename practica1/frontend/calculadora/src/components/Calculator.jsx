import React, {useState} from 'react';
import './Calculator.css';
import Container from '@mui/material/Container';
import { Box } from '@mui/system';

export default function Calculator() {
    const [num, setNum] = useState(0);
    const [oldNum, setOldNum] = useState(0);
    const [operator, setOperator] = useState();
    
    function inputNum (e) {
        let input = e.target.value
        if(num === 0){
            setNum(input);
        } else {
            setNum(num+input);
        }
    }

    function clear() {
        setNum(0);
    }

    function operatorHandler(e) {
        let operatorInput = e.target.value;
        setOperator(operatorInput);
        setOldNum(num);
        setNum(0);
    }

    function calculate() {
        if (operator === "/") {
            //setNum(parseFloat(oldNum) / parseFloat(num));
            (async () => {
                const rawResponse = await fetch('http://localhost:8000/dividir', {
                  method: 'POST',
                  headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                  },
                  body: JSON.stringify({num1: Number(oldNum), num2: Number(num)})
                });
                const content = await rawResponse.json();
              
                console.log(content[0].resultado);
                setNum(content[0].resultado)
              })();

        } else if (operator === "X") {
            //setNum(parseFloat(oldNum) * parseFloat(num));
            (async () => {
                const rawResponse = await fetch('http://localhost:8000/multiplicar', {
                  method: 'POST',
                  headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                  },
                  body: JSON.stringify({num1: Number(oldNum), num2: Number(num)})
                });
                const content = await rawResponse.json();
              
                console.log(content[0].resultado);
                setNum(content[0].resultado)
              })();

        } else if (operator === "-") {
            //setNum(parseFloat(oldNum) - parseFloat(num));
            (async () => {
                const rawResponse = await fetch('http://localhost:8000/restar', {
                  method: 'POST',
                  headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                  },
                  body: JSON.stringify({num1: Number(oldNum), num2: Number(num)})
                });
                const content = await rawResponse.json();
              
                console.log(content[0].resultado);
                setNum(content[0].resultado)
              })();
        }else if (operator === "+") {
            //setNum(parseFloat(oldNum) + parseFloat(num));
            (async () => {
                const rawResponse = await fetch('http://localhost:8000/sumar', {
                  method: 'POST',
                  headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                  },
                  body: JSON.stringify({num1: Number(oldNum), num2: Number(num)})
                });
                const content = await rawResponse.json();
              
                console.log(content[0].resultado);
                setNum(content[0].resultado)
              })();
        }
    }
    
    return (
        <div>
            <Box m={5}/>
            <Container maxWidth="xs">
                <div className='wrapper'>
                <Box m={12}/>
                    <h1 className='resultado'>{num}</h1>
                    <button onClick={clear}>AC</button>
                    <button style={{visibility: "hidden"}}>k</button>
                    {/* */}
                    <button style={{visibility: "hidden"}}>k</button>
                    <button className='orange' onClick={operatorHandler} value={'/'}>/</button>
                    <button className='grey' onClick={inputNum} value={7}>7</button>
                    <button className='grey' onClick={inputNum} value={8}>8</button>
                    <button className='grey' onClick={inputNum} value={9}>9</button>
                    <button className='orange' onClick={operatorHandler} value={'X'}>X</button>
                    <button className='grey' onClick={inputNum} value={4}>4</button>
                    <button className='grey' onClick={inputNum} value={5}>5</button>
                    <button className='grey' onClick={inputNum} value={6}>6</button>
                    <button className='orange' onClick={operatorHandler} value={'-'}>-</button>
                    <button className='grey' onClick={inputNum} value={1}>1</button>
                    <button className='grey' onClick={inputNum} value={2}>2</button>
                    <button className='grey' onClick={inputNum} value={3}>3</button>
                    <button className='orange' onClick={operatorHandler} value={'+'}>+</button>
                    <button className='grey' onClick={inputNum} value={0}>0</button>
                    <button style={{visibility: "hidden"}}>k</button> {/* */}
                    <button style={{visibility: "hidden"}}>k</button>
                    <button className='orange' onClick={calculate}>=</button>
                </div>
            </Container>
        </div>
    )
}