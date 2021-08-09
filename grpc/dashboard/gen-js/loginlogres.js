// source: dashboard.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

goog.provide('proto.dashboardgrpc.LoginLogRes');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.dashboardgrpc.LoginLog');
goog.require('proto.dashboardgrpc.Result');

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
proto.dashboardgrpc.LoginLogRes = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.dashboardgrpc.LoginLogRes.repeatedFields_, null);
};
goog.inherits(proto.dashboardgrpc.LoginLogRes, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dashboardgrpc.LoginLogRes.displayName = 'proto.dashboardgrpc.LoginLogRes';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.dashboardgrpc.LoginLogRes.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.dashboardgrpc.LoginLogRes.prototype.toObject = function(opt_includeInstance) {
  return proto.dashboardgrpc.LoginLogRes.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dashboardgrpc.LoginLogRes} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dashboardgrpc.LoginLogRes.toObject = function(includeInstance, msg) {
  var f, obj = {
    result: (f = msg.getResult()) && proto.dashboardgrpc.Result.toObject(includeInstance, f),
    logsList: jspb.Message.toObjectList(msg.getLogsList(),
    proto.dashboardgrpc.LoginLog.toObject, includeInstance)
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
 * @return {!proto.dashboardgrpc.LoginLogRes}
 */
proto.dashboardgrpc.LoginLogRes.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dashboardgrpc.LoginLogRes;
  return proto.dashboardgrpc.LoginLogRes.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dashboardgrpc.LoginLogRes} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dashboardgrpc.LoginLogRes}
 */
proto.dashboardgrpc.LoginLogRes.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.dashboardgrpc.Result;
      reader.readMessage(value,proto.dashboardgrpc.Result.deserializeBinaryFromReader);
      msg.setResult(value);
      break;
    case 2:
      var value = new proto.dashboardgrpc.LoginLog;
      reader.readMessage(value,proto.dashboardgrpc.LoginLog.deserializeBinaryFromReader);
      msg.addLogs(value);
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
proto.dashboardgrpc.LoginLogRes.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dashboardgrpc.LoginLogRes.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dashboardgrpc.LoginLogRes} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dashboardgrpc.LoginLogRes.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResult();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.dashboardgrpc.Result.serializeBinaryToWriter
    );
  }
  f = message.getLogsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      2,
      f,
      proto.dashboardgrpc.LoginLog.serializeBinaryToWriter
    );
  }
};


/**
 * optional Result result = 1;
 * @return {?proto.dashboardgrpc.Result}
 */
proto.dashboardgrpc.LoginLogRes.prototype.getResult = function() {
  return /** @type{?proto.dashboardgrpc.Result} */ (
    jspb.Message.getWrapperField(this, proto.dashboardgrpc.Result, 1));
};


/**
 * @param {?proto.dashboardgrpc.Result|undefined} value
 * @return {!proto.dashboardgrpc.LoginLogRes} returns this
*/
proto.dashboardgrpc.LoginLogRes.prototype.setResult = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.dashboardgrpc.LoginLogRes} returns this
 */
proto.dashboardgrpc.LoginLogRes.prototype.clearResult = function() {
  return this.setResult(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.dashboardgrpc.LoginLogRes.prototype.hasResult = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated LoginLog logs = 2;
 * @return {!Array<!proto.dashboardgrpc.LoginLog>}
 */
proto.dashboardgrpc.LoginLogRes.prototype.getLogsList = function() {
  return /** @type{!Array<!proto.dashboardgrpc.LoginLog>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.dashboardgrpc.LoginLog, 2));
};


/**
 * @param {!Array<!proto.dashboardgrpc.LoginLog>} value
 * @return {!proto.dashboardgrpc.LoginLogRes} returns this
*/
proto.dashboardgrpc.LoginLogRes.prototype.setLogsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 2, value);
};


/**
 * @param {!proto.dashboardgrpc.LoginLog=} opt_value
 * @param {number=} opt_index
 * @return {!proto.dashboardgrpc.LoginLog}
 */
proto.dashboardgrpc.LoginLogRes.prototype.addLogs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 2, opt_value, proto.dashboardgrpc.LoginLog, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.dashboardgrpc.LoginLogRes} returns this
 */
proto.dashboardgrpc.LoginLogRes.prototype.clearLogsList = function() {
  return this.setLogsList([]);
};


