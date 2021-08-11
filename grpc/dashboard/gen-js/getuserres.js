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

goog.provide('proto.dashboardgrpc.GetUserRes');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.dashboardgrpc.Result');
goog.require('proto.dashboardgrpc.User');

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
proto.dashboardgrpc.GetUserRes = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.dashboardgrpc.GetUserRes, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dashboardgrpc.GetUserRes.displayName = 'proto.dashboardgrpc.GetUserRes';
}



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
proto.dashboardgrpc.GetUserRes.prototype.toObject = function(opt_includeInstance) {
  return proto.dashboardgrpc.GetUserRes.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dashboardgrpc.GetUserRes} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dashboardgrpc.GetUserRes.toObject = function(includeInstance, msg) {
  var f, obj = {
    result: (f = msg.getResult()) && proto.dashboardgrpc.Result.toObject(includeInstance, f),
    user: (f = msg.getUser()) && proto.dashboardgrpc.User.toObject(includeInstance, f)
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
 * @return {!proto.dashboardgrpc.GetUserRes}
 */
proto.dashboardgrpc.GetUserRes.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dashboardgrpc.GetUserRes;
  return proto.dashboardgrpc.GetUserRes.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dashboardgrpc.GetUserRes} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dashboardgrpc.GetUserRes}
 */
proto.dashboardgrpc.GetUserRes.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = new proto.dashboardgrpc.User;
      reader.readMessage(value,proto.dashboardgrpc.User.deserializeBinaryFromReader);
      msg.setUser(value);
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
proto.dashboardgrpc.GetUserRes.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dashboardgrpc.GetUserRes.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dashboardgrpc.GetUserRes} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dashboardgrpc.GetUserRes.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResult();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.dashboardgrpc.Result.serializeBinaryToWriter
    );
  }
  f = message.getUser();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.dashboardgrpc.User.serializeBinaryToWriter
    );
  }
};


/**
 * optional Result Result = 1;
 * @return {?proto.dashboardgrpc.Result}
 */
proto.dashboardgrpc.GetUserRes.prototype.getResult = function() {
  return /** @type{?proto.dashboardgrpc.Result} */ (
    jspb.Message.getWrapperField(this, proto.dashboardgrpc.Result, 1));
};


/**
 * @param {?proto.dashboardgrpc.Result|undefined} value
 * @return {!proto.dashboardgrpc.GetUserRes} returns this
*/
proto.dashboardgrpc.GetUserRes.prototype.setResult = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.dashboardgrpc.GetUserRes} returns this
 */
proto.dashboardgrpc.GetUserRes.prototype.clearResult = function() {
  return this.setResult(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.dashboardgrpc.GetUserRes.prototype.hasResult = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional User User = 2;
 * @return {?proto.dashboardgrpc.User}
 */
proto.dashboardgrpc.GetUserRes.prototype.getUser = function() {
  return /** @type{?proto.dashboardgrpc.User} */ (
    jspb.Message.getWrapperField(this, proto.dashboardgrpc.User, 2));
};


/**
 * @param {?proto.dashboardgrpc.User|undefined} value
 * @return {!proto.dashboardgrpc.GetUserRes} returns this
*/
proto.dashboardgrpc.GetUserRes.prototype.setUser = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.dashboardgrpc.GetUserRes} returns this
 */
proto.dashboardgrpc.GetUserRes.prototype.clearUser = function() {
  return this.setUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.dashboardgrpc.GetUserRes.prototype.hasUser = function() {
  return jspb.Message.getField(this, 2) != null;
};

