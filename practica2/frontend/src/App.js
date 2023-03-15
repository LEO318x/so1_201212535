import React from "react";
import "./styles.css";
import Nav from "./pages/Nav";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Home from "./pages/Home"
import Ram from "./pages/Ram";
import Cpu from "./pages/Cpu";
import Cpuram from "./pages/Cpuram";

export default function App() {
  return (
    <Router>
      <div className="App">
        <Nav />
        <Switch>
          <Route path="/" exact component={Home} />
          <Route path="/cpuram" component={Cpuram} />
          <Route path="/ram" component={Ram} />
          <Route path="/cpu" component={Cpu} />
          
        </Switch>
      </div>
    </Router>
  );
}
