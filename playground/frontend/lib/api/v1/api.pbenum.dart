///
//  Generated code. Do not modify.
//  source: api/v1/api.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Sdk extends $pb.ProtobufEnum {
  static const Sdk SDK_UNSPECIFIED = Sdk._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'SDK_UNSPECIFIED');
  static const Sdk SDK_JAVA = Sdk._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'SDK_JAVA');
  static const Sdk SDK_GO = Sdk._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'SDK_GO');
  static const Sdk SDK_PYTHON = Sdk._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'SDK_PYTHON');
  static const Sdk SDK_SCIO = Sdk._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'SDK_SCIO');

  static const $core.List<Sdk> values = <Sdk> [
    SDK_UNSPECIFIED,
    SDK_JAVA,
    SDK_GO,
    SDK_PYTHON,
    SDK_SCIO,
  ];

  static final $core.Map<$core.int, Sdk> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Sdk? valueOf($core.int value) => _byValue[value];

  const Sdk._($core.int v, $core.String n) : super(v, n);
}

class Status extends $pb.ProtobufEnum {
  static const Status STATUS_UNSPECIFIED = Status._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_UNSPECIFIED');
  static const Status STATUS_VALIDATING = Status._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_VALIDATING');
  static const Status STATUS_VALIDATION_ERROR = Status._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_VALIDATION_ERROR');
  static const Status STATUS_PREPARING = Status._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_PREPARING');
  static const Status STATUS_PREPARATION_ERROR = Status._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_PREPARATION_ERROR');
  static const Status STATUS_COMPILING = Status._(5, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_COMPILING');
  static const Status STATUS_COMPILE_ERROR = Status._(6, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_COMPILE_ERROR');
  static const Status STATUS_EXECUTING = Status._(7, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_EXECUTING');
  static const Status STATUS_FINISHED = Status._(8, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_FINISHED');
  static const Status STATUS_RUN_ERROR = Status._(9, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_RUN_ERROR');
  static const Status STATUS_ERROR = Status._(10, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_ERROR');
  static const Status STATUS_RUN_TIMEOUT = Status._(11, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_RUN_TIMEOUT');
  static const Status STATUS_CANCELED = Status._(12, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_CANCELED');

  static const $core.List<Status> values = <Status> [
    STATUS_UNSPECIFIED,
    STATUS_VALIDATING,
    STATUS_VALIDATION_ERROR,
    STATUS_PREPARING,
    STATUS_PREPARATION_ERROR,
    STATUS_COMPILING,
    STATUS_COMPILE_ERROR,
    STATUS_EXECUTING,
    STATUS_FINISHED,
    STATUS_RUN_ERROR,
    STATUS_ERROR,
    STATUS_RUN_TIMEOUT,
    STATUS_CANCELED,
  ];

  static final $core.Map<$core.int, Status> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Status? valueOf($core.int value) => _byValue[value];

  const Status._($core.int v, $core.String n) : super(v, n);
}

class PrecompiledObjectType extends $pb.ProtobufEnum {
  static const PrecompiledObjectType EXAMPLE = PrecompiledObjectType._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'EXAMPLE');
  static const PrecompiledObjectType KATA = PrecompiledObjectType._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'KATA');
  static const PrecompiledObjectType UNIT_TEST = PrecompiledObjectType._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'UNIT_TEST');

  static const $core.List<PrecompiledObjectType> values = <PrecompiledObjectType> [
    EXAMPLE,
    KATA,
    UNIT_TEST,
  ];

  static final $core.Map<$core.int, PrecompiledObjectType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static PrecompiledObjectType? valueOf($core.int value) => _byValue[value];

  const PrecompiledObjectType._($core.int v, $core.String n) : super(v, n);
}

