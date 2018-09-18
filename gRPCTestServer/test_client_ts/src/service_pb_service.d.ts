// package: 
// file: service.proto

import * as service_pb from "./service_pb";
import {grpc} from "grpc-web-client";

type CatGetMyCat = {
  readonly methodName: string;
  readonly service: typeof Cat;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof service_pb.GetMyCatMessage;
  readonly responseType: typeof service_pb.MyCatResponse;
};

export class Cat {
  static readonly serviceName: string;
  static readonly GetMyCat: CatGetMyCat;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }
export type ServiceClientOptions = { transport: grpc.TransportConstructor; debug?: boolean }

interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}

export class CatClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  getMyCat(
    requestMessage: service_pb.GetMyCatMessage,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: service_pb.MyCatResponse|null) => void
  ): void;
  getMyCat(
    requestMessage: service_pb.GetMyCatMessage,
    callback: (error: ServiceError, responseMessage: service_pb.MyCatResponse|null) => void
  ): void;
}

