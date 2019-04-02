/**
 * @fileoverview gRPC-Web generated client stub for services
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.services = require('./ComponentMgmt_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.services.ComponentMgmtClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.services.ComponentMgmtPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!proto.services.ComponentMgmtClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.services.ComponentMgmtClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.services.ComponentDetail,
 *   !proto.services.ComponentCreateResponse>}
 */
const methodInfo_ComponentMgmt_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.services.ComponentCreateResponse,
  /** @param {!proto.services.ComponentDetail} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.services.ComponentCreateResponse.deserializeBinary
);


/**
 * @param {!proto.services.ComponentDetail} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.services.ComponentCreateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.services.ComponentCreateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.services.ComponentMgmtClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/services.ComponentMgmt/Create',
      request,
      metadata,
      methodInfo_ComponentMgmt_Create,
      callback);
};


/**
 * @param {!proto.services.ComponentDetail} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.services.ComponentCreateResponse>}
 *     The XHR Node Readable Stream
 */
proto.services.ComponentMgmtPromiseClient.prototype.create =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.create(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.services.GetComponentRequest,
 *   !proto.services.GetComponentResponse>}
 */
const methodInfo_ComponentMgmt_GetComponents = new grpc.web.AbstractClientBase.MethodInfo(
  proto.services.GetComponentResponse,
  /** @param {!proto.services.GetComponentRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.services.GetComponentResponse.deserializeBinary
);


/**
 * @param {!proto.services.GetComponentRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.services.GetComponentResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.services.GetComponentResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.services.ComponentMgmtClient.prototype.getComponents =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/services.ComponentMgmt/GetComponents',
      request,
      metadata,
      methodInfo_ComponentMgmt_GetComponents,
      callback);
};


/**
 * @param {!proto.services.GetComponentRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.services.GetComponentResponse>}
 *     The XHR Node Readable Stream
 */
proto.services.ComponentMgmtPromiseClient.prototype.getComponents =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getComponents(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.services;

