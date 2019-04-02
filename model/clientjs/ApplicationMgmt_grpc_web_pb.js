/**
 * @fileoverview gRPC-Web generated client stub for services
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.services = require('./ApplicationMgmt_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.services.ApplicationMgmtClient =
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
proto.services.ApplicationMgmtPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!proto.services.ApplicationMgmtClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.services.ApplicationMgmtClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.services.ApplicationDetails,
 *   !proto.services.ApplicationCreateResponse>}
 */
const methodInfo_ApplicationMgmt_Create = new grpc.web.AbstractClientBase.MethodInfo(
  proto.services.ApplicationCreateResponse,
  /** @param {!proto.services.ApplicationDetails} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.services.ApplicationCreateResponse.deserializeBinary
);


/**
 * @param {!proto.services.ApplicationDetails} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.services.ApplicationCreateResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.services.ApplicationCreateResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.services.ApplicationMgmtClient.prototype.create =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/services.ApplicationMgmt/Create',
      request,
      metadata,
      methodInfo_ApplicationMgmt_Create,
      callback);
};


/**
 * @param {!proto.services.ApplicationDetails} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.services.ApplicationCreateResponse>}
 *     The XHR Node Readable Stream
 */
proto.services.ApplicationMgmtPromiseClient.prototype.create =
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
 *   !proto.services.GetApplicationsRequest,
 *   !proto.services.GetApplicationsResponse>}
 */
const methodInfo_ApplicationMgmt_GetAllApplications = new grpc.web.AbstractClientBase.MethodInfo(
  proto.services.GetApplicationsResponse,
  /** @param {!proto.services.GetApplicationsRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.services.GetApplicationsResponse.deserializeBinary
);


/**
 * @param {!proto.services.GetApplicationsRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.services.GetApplicationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.services.GetApplicationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.services.ApplicationMgmtClient.prototype.getAllApplications =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/services.ApplicationMgmt/GetAllApplications',
      request,
      metadata,
      methodInfo_ApplicationMgmt_GetAllApplications,
      callback);
};


/**
 * @param {!proto.services.GetApplicationsRequest} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.services.GetApplicationsResponse>}
 *     The XHR Node Readable Stream
 */
proto.services.ApplicationMgmtPromiseClient.prototype.getAllApplications =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.getAllApplications(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.services;

