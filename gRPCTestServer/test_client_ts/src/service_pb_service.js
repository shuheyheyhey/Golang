// package: 
// file: service.proto

var service_pb = require("./service_pb");
var grpc = require("grpc-web-client").grpc;

var Cat = (function () {
  function Cat() {}
  Cat.serviceName = "Cat";
  return Cat;
}());

Cat.GetMyCat = {
  methodName: "GetMyCat",
  service: Cat,
  requestStream: false,
  responseStream: false,
  requestType: service_pb.GetMyCatMessage,
  responseType: service_pb.MyCatResponse
};

exports.Cat = Cat;

function CatClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

CatClient.prototype.getMyCat = function getMyCat(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Cat.GetMyCat, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

exports.CatClient = CatClient;

