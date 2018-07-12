/*
 * Code generated by protoc-gen-pbrpc. DO NOT EDIT.
 * source: services.proto
 */

package sample

import (
    "context"
    "reflect"

    pbrpc "github.com/let-z-go/pbrpc"
)

// ClientService ///////////////////////////////////////////////////////////////

type ClientServiceClient struct {
    Channel pbrpc.Channel
    Context context.Context
}

func (self ClientServiceClient) GetNickname(autoRetry bool) (*GetNicknameResponse, error) {
    responseType := methodTableOfClientService[0].ResponseType
    response, e := self.Channel.CallMethod(self.Context, clientServiceName, 0, &pbrpc.Void{}, responseType, autoRetry)

    if e != nil {
        return nil, e
    }

    return response.(*GetNicknameResponse), e
}

type ClientServiceHandlerBase struct {}

func (ClientServiceHandlerBase) GetName() string { return clientServiceName }
func (ClientServiceHandlerBase) GetMethodTable() pbrpc.MethodTable { return methodTableOfClientService }

var _ = pbrpc.ServiceHandler(ClientServiceHandlerBase{})

// `ClientServiceHandler` template:
//
// type ClientServiceHandler struct {
//     ClientServiceHandlerBase
// }
//
// func (ClientServiceHandler) GetNickname(context_ context.Context, channel pbrpc.Channel) (*GetNicknameResponse, error)

const clientServiceName = "CLIENT"

var methodTableOfClientService = pbrpc.MethodTable{
    /* [0]: ClientService.GetNickname */ {
        reflect.TypeOf(pbrpc.Void{}),
        reflect.TypeOf(GetNicknameResponse{}),

        func(serviceHandler pbrpc.ServiceHandler, context_ context.Context, channel pbrpc.Channel, _ pbrpc.IncomingMessage) (pbrpc.OutgoingMessage, pbrpc.ErrorCode) {
            methodHandler, ok := serviceHandler.(interface { GetNickname(context.Context, pbrpc.Channel) (*GetNicknameResponse, error) })

            if !ok {
                return nil, pbrpc.ErrorNotImplemented
            }

            response, e := methodHandler.GetNickname(context_, channel)

            if e != nil {
                if e, ok := e.(pbrpc.Error); ok && e.IsInitiative() {
                    return nil, e.GetCode()
                } else {
                    return nil, pbrpc.ErrorInternalServer
                }
            }

            return response, 0
        },
    },
}

// ServerService ///////////////////////////////////////////////////////////////

type ServerServiceClient struct {
    Channel pbrpc.Channel
    Context context.Context
}

func (self ServerServiceClient) SayHello(request *SayHelloRequest, autoRetry bool) (*SayHelloResponse, error) {
    responseType := methodTableOfServerService[0].ResponseType
    response, e := self.Channel.CallMethod(self.Context, serverServiceName, 0, request, responseType, autoRetry)

    if e != nil {
        return nil, e
    }

    return response.(*SayHelloResponse), e
}

type ServerServiceHandlerBase struct {}

func (ServerServiceHandlerBase) GetName() string { return serverServiceName }
func (ServerServiceHandlerBase) GetMethodTable() pbrpc.MethodTable { return methodTableOfServerService }

var _ = pbrpc.ServiceHandler(ServerServiceHandlerBase{})

// `ServerServiceHandler` template:
//
// type ServerServiceHandler struct {
//     ServerServiceHandlerBase
// }
//
// func (ServerServiceHandler) SayHello(context_ context.Context, channel pbrpc.Channel, request *SayHelloRequest) (*SayHelloResponse, error)

const serverServiceName = "SERVER"

var methodTableOfServerService = pbrpc.MethodTable{
    /* [0]: ServerService.SayHello */ {
        reflect.TypeOf(SayHelloRequest{}),
        reflect.TypeOf(SayHelloResponse{}),

        func(serviceHandler pbrpc.ServiceHandler, context_ context.Context, channel pbrpc.Channel, request pbrpc.IncomingMessage) (pbrpc.OutgoingMessage, pbrpc.ErrorCode) {
            methodHandler, ok := serviceHandler.(interface { SayHello(context.Context, pbrpc.Channel, *SayHelloRequest) (*SayHelloResponse, error) })

            if !ok {
                return nil, pbrpc.ErrorNotImplemented
            }

            response, e := methodHandler.SayHello(context_, channel, request.(*SayHelloRequest))

            if e != nil {
                if e, ok := e.(pbrpc.Error); ok && e.IsInitiative() {
                    return nil, e.GetCode()
                } else {
                    return nil, pbrpc.ErrorInternalServer
                }
            }

            return response, 0
        },
    },
}