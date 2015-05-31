class Circuit {
  constructor() {
    this.ws_ = new WebSocket("ws://circuit.xssrc.com:8080/__xssrc__/circuit/connection");
  }
}


const circuit = new Circuit;
circuit.run();

