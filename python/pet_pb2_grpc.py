# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import pet_pb2 as proto_dot_pet__pb2

GRPC_GENERATED_VERSION = '1.66.2'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in proto/pet_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class PetServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.RegisterNewPet = channel.unary_unary(
                '/pet.PetService/RegisterNewPet',
                request_serializer=proto_dot_pet__pb2.RegisterNewPetRequest.SerializeToString,
                response_deserializer=proto_dot_pet__pb2.RegisterNewPetReply.FromString,
                _registered_method=True)
        self.SearchPet = channel.unary_unary(
                '/pet.PetService/SearchPet',
                request_serializer=proto_dot_pet__pb2.SearchPetRequest.SerializeToString,
                response_deserializer=proto_dot_pet__pb2.SearchPetReply.FromString,
                _registered_method=True)


class PetServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def RegisterNewPet(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SearchPet(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_PetServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'RegisterNewPet': grpc.unary_unary_rpc_method_handler(
                    servicer.RegisterNewPet,
                    request_deserializer=proto_dot_pet__pb2.RegisterNewPetRequest.FromString,
                    response_serializer=proto_dot_pet__pb2.RegisterNewPetReply.SerializeToString,
            ),
            'SearchPet': grpc.unary_unary_rpc_method_handler(
                    servicer.SearchPet,
                    request_deserializer=proto_dot_pet__pb2.SearchPetRequest.FromString,
                    response_serializer=proto_dot_pet__pb2.SearchPetReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'pet.PetService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('pet.PetService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class PetService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def RegisterNewPet(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/pet.PetService/RegisterNewPet',
            proto_dot_pet__pb2.RegisterNewPetRequest.SerializeToString,
            proto_dot_pet__pb2.RegisterNewPetReply.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def SearchPet(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/pet.PetService/SearchPet',
            proto_dot_pet__pb2.SearchPetRequest.SerializeToString,
            proto_dot_pet__pb2.SearchPetReply.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)