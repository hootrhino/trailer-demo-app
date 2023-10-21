# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: trailer.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rtrailer.proto\x12\x07trailer\"\x14\n\x06\x43onfig\x12\n\n\x02kv\x18\x01 \x01(\x0c\"\t\n\x07Request\")\n\x08Response\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0f\n\x07message\x18\x02 \x01(\t\"+\n\x0eServiceRequest\x12\x0b\n\x03\x63md\x18\x01 \x01(\x0c\x12\x0c\n\x04\x61rgs\x18\x02 \x01(\x0c\"-\n\x0fServiceResponse\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\x0c\"\x7f\n\x0eStatusResponse\x12.\n\x06status\x18\x01 \x01(\x0e\x32\x1e.trailer.StatusResponse.Status\x12\x0f\n\x07message\x18\x02 \x01(\t\",\n\x06Status\x12\x0b\n\x07RUNNING\x10\x00\x12\t\n\x05\x45RROR\x10\x01\x12\n\n\x06STOPED\x10\x02\"\x0f\n\rSchemaRequest\"/\n\x07\x44\x61taRow\x12$\n\x06\x63olumn\x18\x01 \x03(\x0b\x32\x14.trailer.ColumnValue\"M\n\x06\x43olumn\x12\x0c\n\x04name\x18\x01 \x01(\t\x12 \n\x04type\x18\x02 \x01(\x0e\x32\x12.trailer.ValueType\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\"X\n\x0b\x43olumnValue\x12\x0c\n\x04name\x18\x01 \x01(\t\x12 \n\x04type\x18\x02 \x01(\x0e\x32\x12.trailer.ValueType\x12\n\n\x02id\x18\x03 \x01(\t\x12\r\n\x05value\x18\x04 \x01(\x0c\"Q\n\x0eSchemaResponse\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x05\x12\x0f\n\x07message\x18\x02 \x01(\t\x12 \n\x07\x63olumns\x18\x03 \x03(\x0b\x32\x0f.trailer.Column\" \n\x0f\x44\x61taRowsRequest\x12\r\n\x05query\x18\x01 \x01(\x0c\"1\n\x10\x44\x61taRowsResponse\x12\x1d\n\x03row\x18\x01 \x03(\x0b\x32\x10.trailer.DataRow*-\n\tValueType\x12\n\n\x06STRING\x10\x00\x12\n\n\x06NUMBER\x10\x01\x12\x08\n\x04\x42OOL\x10\x02\x32\x8a\x03\n\x07Trailer\x12,\n\x04Init\x12\x0f.trailer.Config\x1a\x11.trailer.Response\"\x00\x12.\n\x05Start\x12\x10.trailer.Request\x1a\x11.trailer.Response\"\x00\x12\x35\n\x06Status\x12\x10.trailer.Request\x1a\x17.trailer.StatusResponse\"\x00\x12>\n\x07Service\x12\x17.trailer.ServiceRequest\x1a\x18.trailer.ServiceResponse\"\x00\x12>\n\x05Query\x12\x18.trailer.DataRowsRequest\x1a\x19.trailer.DataRowsResponse\"\x00\x12;\n\x06Schema\x12\x16.trailer.SchemaRequest\x1a\x17.trailer.SchemaResponse\"\x00\x12-\n\x04Stop\x12\x10.trailer.Request\x1a\x11.trailer.Response\"\x00\x42 \n\x07trailerB\x07TrailerP\x00Z\n./;trailerb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'trailer_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\007trailerB\007TrailerP\000Z\n./;trailer'
  _globals['_VALUETYPE']._serialized_start=726
  _globals['_VALUETYPE']._serialized_end=771
  _globals['_CONFIG']._serialized_start=26
  _globals['_CONFIG']._serialized_end=46
  _globals['_REQUEST']._serialized_start=48
  _globals['_REQUEST']._serialized_end=57
  _globals['_RESPONSE']._serialized_start=59
  _globals['_RESPONSE']._serialized_end=100
  _globals['_SERVICEREQUEST']._serialized_start=102
  _globals['_SERVICEREQUEST']._serialized_end=145
  _globals['_SERVICERESPONSE']._serialized_start=147
  _globals['_SERVICERESPONSE']._serialized_end=192
  _globals['_STATUSRESPONSE']._serialized_start=194
  _globals['_STATUSRESPONSE']._serialized_end=321
  _globals['_STATUSRESPONSE_STATUS']._serialized_start=277
  _globals['_STATUSRESPONSE_STATUS']._serialized_end=321
  _globals['_SCHEMAREQUEST']._serialized_start=323
  _globals['_SCHEMAREQUEST']._serialized_end=338
  _globals['_DATAROW']._serialized_start=340
  _globals['_DATAROW']._serialized_end=387
  _globals['_COLUMN']._serialized_start=389
  _globals['_COLUMN']._serialized_end=466
  _globals['_COLUMNVALUE']._serialized_start=468
  _globals['_COLUMNVALUE']._serialized_end=556
  _globals['_SCHEMARESPONSE']._serialized_start=558
  _globals['_SCHEMARESPONSE']._serialized_end=639
  _globals['_DATAROWSREQUEST']._serialized_start=641
  _globals['_DATAROWSREQUEST']._serialized_end=673
  _globals['_DATAROWSRESPONSE']._serialized_start=675
  _globals['_DATAROWSRESPONSE']._serialized_end=724
  _globals['_TRAILER']._serialized_start=774
  _globals['_TRAILER']._serialized_end=1168
# @@protoc_insertion_point(module_scope)
