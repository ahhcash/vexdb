# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from . import ghastly_pb2 as ghastly__pb2

GRPC_GENERATED_VERSION = '1.69.0'
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
        + f' but the generated code in ghastly_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class GhastlyDBStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Put = channel.unary_unary(
                '/ghastlydb.GhastlyDB/Put',
                request_serializer=ghastly__pb2.PutRequest.SerializeToString,
                response_deserializer=ghastly__pb2.PutResponse.FromString,
                _registered_method=True)
        self.Get = channel.unary_unary(
                '/ghastlydb.GhastlyDB/Get',
                request_serializer=ghastly__pb2.GetRequest.SerializeToString,
                response_deserializer=ghastly__pb2.GetResponse.FromString,
                _registered_method=True)
        self.Delete = channel.unary_unary(
                '/ghastlydb.GhastlyDB/Delete',
                request_serializer=ghastly__pb2.DeleteRequest.SerializeToString,
                response_deserializer=ghastly__pb2.DeleteResponse.FromString,
                _registered_method=True)
        self.Search = channel.unary_unary(
                '/ghastlydb.GhastlyDB/Search',
                request_serializer=ghastly__pb2.SearchRequest.SerializeToString,
                response_deserializer=ghastly__pb2.SearchResponse.FromString,
                _registered_method=True)
        self.BulkPut = channel.stream_unary(
                '/ghastlydb.GhastlyDB/BulkPut',
                request_serializer=ghastly__pb2.PutRequest.SerializeToString,
                response_deserializer=ghastly__pb2.BulkPutResponse.FromString,
                _registered_method=True)
        self.BulkSearch = channel.unary_stream(
                '/ghastlydb.GhastlyDB/BulkSearch',
                request_serializer=ghastly__pb2.SearchRequest.SerializeToString,
                response_deserializer=ghastly__pb2.SearchResponse.FromString,
                _registered_method=True)
        self.HealthCheck = channel.unary_unary(
                '/ghastlydb.GhastlyDB/HealthCheck',
                request_serializer=ghastly__pb2.HealthCheckRequest.SerializeToString,
                response_deserializer=ghastly__pb2.HealthCheckResponse.FromString,
                _registered_method=True)


class GhastlyDBServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Put(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Delete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Search(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BulkPut(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def BulkSearch(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def HealthCheck(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_GhastlyDBServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Put': grpc.unary_unary_rpc_method_handler(
                    servicer.Put,
                    request_deserializer=ghastly__pb2.PutRequest.FromString,
                    response_serializer=ghastly__pb2.PutResponse.SerializeToString,
            ),
            'Get': grpc.unary_unary_rpc_method_handler(
                    servicer.Get,
                    request_deserializer=ghastly__pb2.GetRequest.FromString,
                    response_serializer=ghastly__pb2.GetResponse.SerializeToString,
            ),
            'Delete': grpc.unary_unary_rpc_method_handler(
                    servicer.Delete,
                    request_deserializer=ghastly__pb2.DeleteRequest.FromString,
                    response_serializer=ghastly__pb2.DeleteResponse.SerializeToString,
            ),
            'Search': grpc.unary_unary_rpc_method_handler(
                    servicer.Search,
                    request_deserializer=ghastly__pb2.SearchRequest.FromString,
                    response_serializer=ghastly__pb2.SearchResponse.SerializeToString,
            ),
            'BulkPut': grpc.stream_unary_rpc_method_handler(
                    servicer.BulkPut,
                    request_deserializer=ghastly__pb2.PutRequest.FromString,
                    response_serializer=ghastly__pb2.BulkPutResponse.SerializeToString,
            ),
            'BulkSearch': grpc.unary_stream_rpc_method_handler(
                    servicer.BulkSearch,
                    request_deserializer=ghastly__pb2.SearchRequest.FromString,
                    response_serializer=ghastly__pb2.SearchResponse.SerializeToString,
            ),
            'HealthCheck': grpc.unary_unary_rpc_method_handler(
                    servicer.HealthCheck,
                    request_deserializer=ghastly__pb2.HealthCheckRequest.FromString,
                    response_serializer=ghastly__pb2.HealthCheckResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'ghastlydb.GhastlyDB', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('ghastlydb.GhastlyDB', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class GhastlyDB(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Put(request,
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
            '/ghastlydb.GhastlyDB/Put',
            ghastly__pb2.PutRequest.SerializeToString,
            ghastly__pb2.PutResponse.FromString,
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
    def Get(request,
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
            '/ghastlydb.GhastlyDB/Get',
            ghastly__pb2.GetRequest.SerializeToString,
            ghastly__pb2.GetResponse.FromString,
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
    def Delete(request,
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
            '/ghastlydb.GhastlyDB/Delete',
            ghastly__pb2.DeleteRequest.SerializeToString,
            ghastly__pb2.DeleteResponse.FromString,
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
    def Search(request,
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
            '/ghastlydb.GhastlyDB/Search',
            ghastly__pb2.SearchRequest.SerializeToString,
            ghastly__pb2.SearchResponse.FromString,
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
    def BulkPut(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(
            request_iterator,
            target,
            '/ghastlydb.GhastlyDB/BulkPut',
            ghastly__pb2.PutRequest.SerializeToString,
            ghastly__pb2.BulkPutResponse.FromString,
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
    def BulkSearch(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(
            request,
            target,
            '/ghastlydb.GhastlyDB/BulkSearch',
            ghastly__pb2.SearchRequest.SerializeToString,
            ghastly__pb2.SearchResponse.FromString,
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
    def HealthCheck(request,
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
            '/ghastlydb.GhastlyDB/HealthCheck',
            ghastly__pb2.HealthCheckRequest.SerializeToString,
            ghastly__pb2.HealthCheckResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
