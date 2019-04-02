/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.services.ApplicationCreateResponse', null, global);
goog.exportSymbol('proto.services.ApplicationDetails', null, global);
goog.exportSymbol('proto.services.GetApplicationsRequest', null, global);
goog.exportSymbol('proto.services.GetApplicationsResponse', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.services.ApplicationDetails = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.services.ApplicationDetails, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.services.ApplicationDetails.displayName = 'proto.services.ApplicationDetails';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.services.ApplicationDetails.prototype.toObject = function(opt_includeInstance) {
  return proto.services.ApplicationDetails.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.services.ApplicationDetails} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.services.ApplicationDetails.toObject = function(includeInstance, msg) {
  var f, obj = {
    applicationname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    orgname: jspb.Message.getFieldWithDefault(msg, 2, ""),
    cloud: jspb.Message.getFieldWithDefault(msg, 3, ""),
    awsaccesskey: jspb.Message.getFieldWithDefault(msg, 4, ""),
    awssecret: jspb.Message.getFieldWithDefault(msg, 5, ""),
    description: jspb.Message.getFieldWithDefault(msg, 6, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.services.ApplicationDetails}
 */
proto.services.ApplicationDetails.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.services.ApplicationDetails;
  return proto.services.ApplicationDetails.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.services.ApplicationDetails} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.services.ApplicationDetails}
 */
proto.services.ApplicationDetails.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setApplicationname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setOrgname(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setCloud(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setAwsaccesskey(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setAwssecret(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setDescription(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.services.ApplicationDetails.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.services.ApplicationDetails.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.services.ApplicationDetails} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.services.ApplicationDetails.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getApplicationname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getOrgname();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getCloud();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getAwsaccesskey();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getAwssecret();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getDescription();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
};


/**
 * optional string applicationName = 1;
 * @return {string}
 */
proto.services.ApplicationDetails.prototype.getApplicationname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.services.ApplicationDetails.prototype.setApplicationname = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string orgName = 2;
 * @return {string}
 */
proto.services.ApplicationDetails.prototype.getOrgname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.services.ApplicationDetails.prototype.setOrgname = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string cloud = 3;
 * @return {string}
 */
proto.services.ApplicationDetails.prototype.getCloud = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.services.ApplicationDetails.prototype.setCloud = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string awsAccessKey = 4;
 * @return {string}
 */
proto.services.ApplicationDetails.prototype.getAwsaccesskey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.services.ApplicationDetails.prototype.setAwsaccesskey = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string awsSecret = 5;
 * @return {string}
 */
proto.services.ApplicationDetails.prototype.getAwssecret = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/** @param {string} value */
proto.services.ApplicationDetails.prototype.setAwssecret = function(value) {
  jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string description = 6;
 * @return {string}
 */
proto.services.ApplicationDetails.prototype.getDescription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/** @param {string} value */
proto.services.ApplicationDetails.prototype.setDescription = function(value) {
  jspb.Message.setProto3StringField(this, 6, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.services.ApplicationCreateResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.services.ApplicationCreateResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.services.ApplicationCreateResponse.displayName = 'proto.services.ApplicationCreateResponse';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.services.ApplicationCreateResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.services.ApplicationCreateResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.services.ApplicationCreateResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.services.ApplicationCreateResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    applicationname: jspb.Message.getFieldWithDefault(msg, 1, ""),
    applicationid: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.services.ApplicationCreateResponse}
 */
proto.services.ApplicationCreateResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.services.ApplicationCreateResponse;
  return proto.services.ApplicationCreateResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.services.ApplicationCreateResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.services.ApplicationCreateResponse}
 */
proto.services.ApplicationCreateResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setApplicationname(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setApplicationid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.services.ApplicationCreateResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.services.ApplicationCreateResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.services.ApplicationCreateResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.services.ApplicationCreateResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getApplicationname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getApplicationid();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string applicationName = 1;
 * @return {string}
 */
proto.services.ApplicationCreateResponse.prototype.getApplicationname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.services.ApplicationCreateResponse.prototype.setApplicationname = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string applicationId = 2;
 * @return {string}
 */
proto.services.ApplicationCreateResponse.prototype.getApplicationid = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.services.ApplicationCreateResponse.prototype.setApplicationid = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.services.GetApplicationsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.services.GetApplicationsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.services.GetApplicationsRequest.displayName = 'proto.services.GetApplicationsRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.services.GetApplicationsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.services.GetApplicationsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.services.GetApplicationsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.services.GetApplicationsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    orgname: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.services.GetApplicationsRequest}
 */
proto.services.GetApplicationsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.services.GetApplicationsRequest;
  return proto.services.GetApplicationsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.services.GetApplicationsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.services.GetApplicationsRequest}
 */
proto.services.GetApplicationsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setOrgname(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.services.GetApplicationsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.services.GetApplicationsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.services.GetApplicationsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.services.GetApplicationsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getOrgname();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string orgName = 1;
 * @return {string}
 */
proto.services.GetApplicationsRequest.prototype.getOrgname = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.services.GetApplicationsRequest.prototype.setOrgname = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.services.GetApplicationsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.services.GetApplicationsResponse.repeatedFields_, null);
};
goog.inherits(proto.services.GetApplicationsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.services.GetApplicationsResponse.displayName = 'proto.services.GetApplicationsResponse';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.services.GetApplicationsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.services.GetApplicationsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.services.GetApplicationsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.services.GetApplicationsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.services.GetApplicationsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    applicationsList: jspb.Message.toObjectList(msg.getApplicationsList(),
    proto.services.ApplicationDetails.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.services.GetApplicationsResponse}
 */
proto.services.GetApplicationsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.services.GetApplicationsResponse;
  return proto.services.GetApplicationsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.services.GetApplicationsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.services.GetApplicationsResponse}
 */
proto.services.GetApplicationsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.services.ApplicationDetails;
      reader.readMessage(value,proto.services.ApplicationDetails.deserializeBinaryFromReader);
      msg.addApplications(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.services.GetApplicationsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.services.GetApplicationsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.services.GetApplicationsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.services.GetApplicationsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getApplicationsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.services.ApplicationDetails.serializeBinaryToWriter
    );
  }
};


/**
 * repeated ApplicationDetails applications = 1;
 * @return {!Array<!proto.services.ApplicationDetails>}
 */
proto.services.GetApplicationsResponse.prototype.getApplicationsList = function() {
  return /** @type{!Array<!proto.services.ApplicationDetails>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.services.ApplicationDetails, 1));
};


/** @param {!Array<!proto.services.ApplicationDetails>} value */
proto.services.GetApplicationsResponse.prototype.setApplicationsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.services.ApplicationDetails=} opt_value
 * @param {number=} opt_index
 * @return {!proto.services.ApplicationDetails}
 */
proto.services.GetApplicationsResponse.prototype.addApplications = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.services.ApplicationDetails, opt_index);
};


proto.services.GetApplicationsResponse.prototype.clearApplicationsList = function() {
  this.setApplicationsList([]);
};


goog.object.extend(exports, proto.services);
