# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: group.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
from ntypes import ntypes_pb2 as ntypes_dot_ntypes__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='group.proto',
  package='charonrpc',
  syntax='proto3',
  serialized_pb=_b('\n\x0bgroup.proto\x12\tcharonrpc\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x13ntypes/ntypes.proto\"\xdc\x01\n\x05Group\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12.\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12!\n\ncreated_by\x18\x05 \x01(\x0b\x32\r.ntypes.Int64\x12.\n\nupdated_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12!\n\nupdated_by\x18\x07 \x01(\x0b\x32\r.ntypes.Int64\"G\n\x12\x43reateGroupRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12#\n\x0b\x64\x65scription\x18\x02 \x01(\x0b\x32\x0e.ntypes.String\"6\n\x13\x43reateGroupResponse\x12\x1f\n\x05group\x18\x01 \x01(\x0b\x32\x10.charonrpc.Group\"\x1d\n\x0fGetGroupRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"3\n\x10GetGroupResponse\x12\x1f\n\x05group\x18\x01 \x01(\x0b\x32\x10.charonrpc.Group\"V\n\x11ListGroupsRequest\x12\x1d\n\x06offset\x18\x64 \x01(\x0b\x32\r.ntypes.Int64\x12\x1c\n\x05limit\x18\x65 \x01(\x0b\x32\r.ntypes.Int64J\x04\x08\x01\x10\x64\"6\n\x12ListGroupsResponse\x12 \n\x06groups\x18\x01 \x03(\x0b\x32\x10.charonrpc.Group\" \n\x12\x44\x65leteGroupRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"c\n\x12ModifyGroupRequest\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x1c\n\x04name\x18\x02 \x01(\x0b\x32\x0e.ntypes.String\x12#\n\x0b\x64\x65scription\x18\x03 \x01(\x0b\x32\x0e.ntypes.String\"6\n\x13ModifyGroupResponse\x12\x1f\n\x05group\x18\x01 \x01(\x0b\x32\x10.charonrpc.Group\"R\n\x1aSetGroupPermissionsRequest\x12\x10\n\x08group_id\x18\x01 \x01(\x03\x12\x13\n\x0bpermissions\x18\x02 \x03(\t\x12\r\n\x05\x66orce\x18\x03 \x01(\x08\"R\n\x1bSetGroupPermissionsResponse\x12\x0f\n\x07\x63reated\x18\x01 \x01(\x03\x12\x0f\n\x07removed\x18\x02 \x01(\x03\x12\x11\n\tuntouched\x18\x03 \x01(\x03\")\n\x1bListGroupPermissionsRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"3\n\x1cListGroupPermissionsResponse\x12\x13\n\x0bpermissions\x18\x01 \x03(\t2\xbd\x04\n\x0cGroupManager\x12I\n\x06\x43reate\x12\x1d.charonrpc.CreateGroupRequest\x1a\x1e.charonrpc.CreateGroupResponse\"\x00\x12I\n\x06Modify\x12\x1d.charonrpc.ModifyGroupRequest\x1a\x1e.charonrpc.ModifyGroupResponse\"\x00\x12@\n\x03Get\x12\x1a.charonrpc.GetGroupRequest\x1a\x1b.charonrpc.GetGroupResponse\"\x00\x12\x45\n\x04List\x12\x1c.charonrpc.ListGroupsRequest\x1a\x1d.charonrpc.ListGroupsResponse\"\x00\x12\x45\n\x06\x44\x65lete\x12\x1d.charonrpc.DeleteGroupRequest\x1a\x1a.google.protobuf.BoolValue\"\x00\x12\x64\n\x0fListPermissions\x12&.charonrpc.ListGroupPermissionsRequest\x1a\'.charonrpc.ListGroupPermissionsResponse\"\x00\x12\x61\n\x0eSetPermissions\x12%.charonrpc.SetGroupPermissionsRequest\x1a&.charonrpc.SetGroupPermissionsResponse\"\x00\x42,Z*github.com/piotrkowalczuk/charon/charonrpcb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,google_dot_protobuf_dot_wrappers__pb2.DESCRIPTOR,ntypes_dot_ntypes__pb2.DESCRIPTOR,])




_GROUP = _descriptor.Descriptor(
  name='Group',
  full_name='charonrpc.Group',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='charonrpc.Group.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='charonrpc.Group.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='charonrpc.Group.description', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='charonrpc.Group.created_at', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_by', full_name='charonrpc.Group.created_by', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='charonrpc.Group.updated_at', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='updated_by', full_name='charonrpc.Group.updated_by', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=113,
  serialized_end=333,
)


_CREATEGROUPREQUEST = _descriptor.Descriptor(
  name='CreateGroupRequest',
  full_name='charonrpc.CreateGroupRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='charonrpc.CreateGroupRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='charonrpc.CreateGroupRequest.description', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=335,
  serialized_end=406,
)


_CREATEGROUPRESPONSE = _descriptor.Descriptor(
  name='CreateGroupResponse',
  full_name='charonrpc.CreateGroupResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='group', full_name='charonrpc.CreateGroupResponse.group', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=408,
  serialized_end=462,
)


_GETGROUPREQUEST = _descriptor.Descriptor(
  name='GetGroupRequest',
  full_name='charonrpc.GetGroupRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='charonrpc.GetGroupRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=464,
  serialized_end=493,
)


_GETGROUPRESPONSE = _descriptor.Descriptor(
  name='GetGroupResponse',
  full_name='charonrpc.GetGroupResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='group', full_name='charonrpc.GetGroupResponse.group', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=495,
  serialized_end=546,
)


_LISTGROUPSREQUEST = _descriptor.Descriptor(
  name='ListGroupsRequest',
  full_name='charonrpc.ListGroupsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='offset', full_name='charonrpc.ListGroupsRequest.offset', index=0,
      number=100, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limit', full_name='charonrpc.ListGroupsRequest.limit', index=1,
      number=101, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=548,
  serialized_end=634,
)


_LISTGROUPSRESPONSE = _descriptor.Descriptor(
  name='ListGroupsResponse',
  full_name='charonrpc.ListGroupsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='groups', full_name='charonrpc.ListGroupsResponse.groups', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=636,
  serialized_end=690,
)


_DELETEGROUPREQUEST = _descriptor.Descriptor(
  name='DeleteGroupRequest',
  full_name='charonrpc.DeleteGroupRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='charonrpc.DeleteGroupRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=692,
  serialized_end=724,
)


_MODIFYGROUPREQUEST = _descriptor.Descriptor(
  name='ModifyGroupRequest',
  full_name='charonrpc.ModifyGroupRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='charonrpc.ModifyGroupRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='charonrpc.ModifyGroupRequest.name', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='charonrpc.ModifyGroupRequest.description', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=726,
  serialized_end=825,
)


_MODIFYGROUPRESPONSE = _descriptor.Descriptor(
  name='ModifyGroupResponse',
  full_name='charonrpc.ModifyGroupResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='group', full_name='charonrpc.ModifyGroupResponse.group', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=827,
  serialized_end=881,
)


_SETGROUPPERMISSIONSREQUEST = _descriptor.Descriptor(
  name='SetGroupPermissionsRequest',
  full_name='charonrpc.SetGroupPermissionsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='group_id', full_name='charonrpc.SetGroupPermissionsRequest.group_id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='permissions', full_name='charonrpc.SetGroupPermissionsRequest.permissions', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='force', full_name='charonrpc.SetGroupPermissionsRequest.force', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=883,
  serialized_end=965,
)


_SETGROUPPERMISSIONSRESPONSE = _descriptor.Descriptor(
  name='SetGroupPermissionsResponse',
  full_name='charonrpc.SetGroupPermissionsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='created', full_name='charonrpc.SetGroupPermissionsResponse.created', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='removed', full_name='charonrpc.SetGroupPermissionsResponse.removed', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='untouched', full_name='charonrpc.SetGroupPermissionsResponse.untouched', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=967,
  serialized_end=1049,
)


_LISTGROUPPERMISSIONSREQUEST = _descriptor.Descriptor(
  name='ListGroupPermissionsRequest',
  full_name='charonrpc.ListGroupPermissionsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='charonrpc.ListGroupPermissionsRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1051,
  serialized_end=1092,
)


_LISTGROUPPERMISSIONSRESPONSE = _descriptor.Descriptor(
  name='ListGroupPermissionsResponse',
  full_name='charonrpc.ListGroupPermissionsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='permissions', full_name='charonrpc.ListGroupPermissionsResponse.permissions', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1094,
  serialized_end=1145,
)

_GROUP.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_GROUP.fields_by_name['created_by'].message_type = ntypes_dot_ntypes__pb2._INT64
_GROUP.fields_by_name['updated_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_GROUP.fields_by_name['updated_by'].message_type = ntypes_dot_ntypes__pb2._INT64
_CREATEGROUPREQUEST.fields_by_name['description'].message_type = ntypes_dot_ntypes__pb2._STRING
_CREATEGROUPRESPONSE.fields_by_name['group'].message_type = _GROUP
_GETGROUPRESPONSE.fields_by_name['group'].message_type = _GROUP
_LISTGROUPSREQUEST.fields_by_name['offset'].message_type = ntypes_dot_ntypes__pb2._INT64
_LISTGROUPSREQUEST.fields_by_name['limit'].message_type = ntypes_dot_ntypes__pb2._INT64
_LISTGROUPSRESPONSE.fields_by_name['groups'].message_type = _GROUP
_MODIFYGROUPREQUEST.fields_by_name['name'].message_type = ntypes_dot_ntypes__pb2._STRING
_MODIFYGROUPREQUEST.fields_by_name['description'].message_type = ntypes_dot_ntypes__pb2._STRING
_MODIFYGROUPRESPONSE.fields_by_name['group'].message_type = _GROUP
DESCRIPTOR.message_types_by_name['Group'] = _GROUP
DESCRIPTOR.message_types_by_name['CreateGroupRequest'] = _CREATEGROUPREQUEST
DESCRIPTOR.message_types_by_name['CreateGroupResponse'] = _CREATEGROUPRESPONSE
DESCRIPTOR.message_types_by_name['GetGroupRequest'] = _GETGROUPREQUEST
DESCRIPTOR.message_types_by_name['GetGroupResponse'] = _GETGROUPRESPONSE
DESCRIPTOR.message_types_by_name['ListGroupsRequest'] = _LISTGROUPSREQUEST
DESCRIPTOR.message_types_by_name['ListGroupsResponse'] = _LISTGROUPSRESPONSE
DESCRIPTOR.message_types_by_name['DeleteGroupRequest'] = _DELETEGROUPREQUEST
DESCRIPTOR.message_types_by_name['ModifyGroupRequest'] = _MODIFYGROUPREQUEST
DESCRIPTOR.message_types_by_name['ModifyGroupResponse'] = _MODIFYGROUPRESPONSE
DESCRIPTOR.message_types_by_name['SetGroupPermissionsRequest'] = _SETGROUPPERMISSIONSREQUEST
DESCRIPTOR.message_types_by_name['SetGroupPermissionsResponse'] = _SETGROUPPERMISSIONSRESPONSE
DESCRIPTOR.message_types_by_name['ListGroupPermissionsRequest'] = _LISTGROUPPERMISSIONSREQUEST
DESCRIPTOR.message_types_by_name['ListGroupPermissionsResponse'] = _LISTGROUPPERMISSIONSRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Group = _reflection.GeneratedProtocolMessageType('Group', (_message.Message,), dict(
  DESCRIPTOR = _GROUP,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.Group)
  ))
_sym_db.RegisterMessage(Group)

CreateGroupRequest = _reflection.GeneratedProtocolMessageType('CreateGroupRequest', (_message.Message,), dict(
  DESCRIPTOR = _CREATEGROUPREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.CreateGroupRequest)
  ))
_sym_db.RegisterMessage(CreateGroupRequest)

CreateGroupResponse = _reflection.GeneratedProtocolMessageType('CreateGroupResponse', (_message.Message,), dict(
  DESCRIPTOR = _CREATEGROUPRESPONSE,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.CreateGroupResponse)
  ))
_sym_db.RegisterMessage(CreateGroupResponse)

GetGroupRequest = _reflection.GeneratedProtocolMessageType('GetGroupRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETGROUPREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.GetGroupRequest)
  ))
_sym_db.RegisterMessage(GetGroupRequest)

GetGroupResponse = _reflection.GeneratedProtocolMessageType('GetGroupResponse', (_message.Message,), dict(
  DESCRIPTOR = _GETGROUPRESPONSE,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.GetGroupResponse)
  ))
_sym_db.RegisterMessage(GetGroupResponse)

ListGroupsRequest = _reflection.GeneratedProtocolMessageType('ListGroupsRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTGROUPSREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.ListGroupsRequest)
  ))
_sym_db.RegisterMessage(ListGroupsRequest)

ListGroupsResponse = _reflection.GeneratedProtocolMessageType('ListGroupsResponse', (_message.Message,), dict(
  DESCRIPTOR = _LISTGROUPSRESPONSE,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.ListGroupsResponse)
  ))
_sym_db.RegisterMessage(ListGroupsResponse)

DeleteGroupRequest = _reflection.GeneratedProtocolMessageType('DeleteGroupRequest', (_message.Message,), dict(
  DESCRIPTOR = _DELETEGROUPREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.DeleteGroupRequest)
  ))
_sym_db.RegisterMessage(DeleteGroupRequest)

ModifyGroupRequest = _reflection.GeneratedProtocolMessageType('ModifyGroupRequest', (_message.Message,), dict(
  DESCRIPTOR = _MODIFYGROUPREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.ModifyGroupRequest)
  ))
_sym_db.RegisterMessage(ModifyGroupRequest)

ModifyGroupResponse = _reflection.GeneratedProtocolMessageType('ModifyGroupResponse', (_message.Message,), dict(
  DESCRIPTOR = _MODIFYGROUPRESPONSE,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.ModifyGroupResponse)
  ))
_sym_db.RegisterMessage(ModifyGroupResponse)

SetGroupPermissionsRequest = _reflection.GeneratedProtocolMessageType('SetGroupPermissionsRequest', (_message.Message,), dict(
  DESCRIPTOR = _SETGROUPPERMISSIONSREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.SetGroupPermissionsRequest)
  ))
_sym_db.RegisterMessage(SetGroupPermissionsRequest)

SetGroupPermissionsResponse = _reflection.GeneratedProtocolMessageType('SetGroupPermissionsResponse', (_message.Message,), dict(
  DESCRIPTOR = _SETGROUPPERMISSIONSRESPONSE,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.SetGroupPermissionsResponse)
  ))
_sym_db.RegisterMessage(SetGroupPermissionsResponse)

ListGroupPermissionsRequest = _reflection.GeneratedProtocolMessageType('ListGroupPermissionsRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTGROUPPERMISSIONSREQUEST,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.ListGroupPermissionsRequest)
  ))
_sym_db.RegisterMessage(ListGroupPermissionsRequest)

ListGroupPermissionsResponse = _reflection.GeneratedProtocolMessageType('ListGroupPermissionsResponse', (_message.Message,), dict(
  DESCRIPTOR = _LISTGROUPPERMISSIONSRESPONSE,
  __module__ = 'group_pb2'
  # @@protoc_insertion_point(class_scope:charonrpc.ListGroupPermissionsResponse)
  ))
_sym_db.RegisterMessage(ListGroupPermissionsResponse)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z*github.com/piotrkowalczuk/charon/charonrpc'))

_GROUPMANAGER = _descriptor.ServiceDescriptor(
  name='GroupManager',
  full_name='charonrpc.GroupManager',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=1148,
  serialized_end=1721,
  methods=[
  _descriptor.MethodDescriptor(
    name='Create',
    full_name='charonrpc.GroupManager.Create',
    index=0,
    containing_service=None,
    input_type=_CREATEGROUPREQUEST,
    output_type=_CREATEGROUPRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Modify',
    full_name='charonrpc.GroupManager.Modify',
    index=1,
    containing_service=None,
    input_type=_MODIFYGROUPREQUEST,
    output_type=_MODIFYGROUPRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Get',
    full_name='charonrpc.GroupManager.Get',
    index=2,
    containing_service=None,
    input_type=_GETGROUPREQUEST,
    output_type=_GETGROUPRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='List',
    full_name='charonrpc.GroupManager.List',
    index=3,
    containing_service=None,
    input_type=_LISTGROUPSREQUEST,
    output_type=_LISTGROUPSRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Delete',
    full_name='charonrpc.GroupManager.Delete',
    index=4,
    containing_service=None,
    input_type=_DELETEGROUPREQUEST,
    output_type=google_dot_protobuf_dot_wrappers__pb2._BOOLVALUE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ListPermissions',
    full_name='charonrpc.GroupManager.ListPermissions',
    index=5,
    containing_service=None,
    input_type=_LISTGROUPPERMISSIONSREQUEST,
    output_type=_LISTGROUPPERMISSIONSRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SetPermissions',
    full_name='charonrpc.GroupManager.SetPermissions',
    index=6,
    containing_service=None,
    input_type=_SETGROUPPERMISSIONSREQUEST,
    output_type=_SETGROUPPERMISSIONSRESPONSE,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_GROUPMANAGER)

DESCRIPTOR.services_by_name['GroupManager'] = _GROUPMANAGER

# @@protoc_insertion_point(module_scope)
