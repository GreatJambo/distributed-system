# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: pet.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'pet.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\tpet.proto\x12\x03pet\"P\n\x03Pet\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06gender\x18\x02 \x01(\t\x12\x0b\n\x03\x61ge\x18\x03 \x01(\x05\x12\r\n\x05\x62reed\x18\x04 \x01(\t\x12\x0f\n\x07picture\x18\x05 \x01(\t\"b\n\x15RegisterNewPetRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06gender\x18\x02 \x01(\t\x12\x0b\n\x03\x61ge\x18\x03 \x01(\x05\x12\r\n\x05\x62reed\x18\x04 \x01(\t\x12\x0f\n\x07picture\x18\x05 \x01(\t\"0\n\x13RegisterNewPetReply\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0b\n\x03msg\x18\x02 \x01(\t\"^\n\x10SearchPetRequest\x12\x0e\n\x04name\x18\x01 \x01(\tH\x00\x12\x10\n\x06gender\x18\x02 \x01(\tH\x00\x12\r\n\x03\x61ge\x18\x03 \x01(\x05H\x00\x12\x0f\n\x05\x62reed\x18\x04 \x01(\tH\x00\x42\x08\n\x06\x64\x65tail\"(\n\x0eSearchPetReply\x12\x16\n\x04pets\x18\x01 \x03(\x0b\x32\x08.pet.Pet2\x8d\x01\n\nPetService\x12\x46\n\x0eRegisterNewPet\x12\x1a.pet.RegisterNewPetRequest\x1a\x18.pet.RegisterNewPetReply\x12\x37\n\tSearchPet\x12\x15.pet.SearchPetRequest\x1a\x13.pet.SearchPetReplyB\x06Z\x04/petb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'pet_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\004/pet'
  _globals['_PET']._serialized_start=18
  _globals['_PET']._serialized_end=98
  _globals['_REGISTERNEWPETREQUEST']._serialized_start=100
  _globals['_REGISTERNEWPETREQUEST']._serialized_end=198
  _globals['_REGISTERNEWPETREPLY']._serialized_start=200
  _globals['_REGISTERNEWPETREPLY']._serialized_end=248
  _globals['_SEARCHPETREQUEST']._serialized_start=250
  _globals['_SEARCHPETREQUEST']._serialized_end=344
  _globals['_SEARCHPETREPLY']._serialized_start=346
  _globals['_SEARCHPETREPLY']._serialized_end=386
  _globals['_PETSERVICE']._serialized_start=389
  _globals['_PETSERVICE']._serialized_end=530
# @@protoc_insertion_point(module_scope)
