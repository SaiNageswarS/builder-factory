import * as grpcWeb from 'grpc-web';
import {
  ComponentCreateResponse,
  ComponentDetail,
  GetComponentRequest,
  GetComponentResponse,
  Package} from './ComponentMgmt_pb';

export class ComponentMgmtClient {
  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; });

  create(
    request: ComponentDetail,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: ComponentCreateResponse) => void
  ): grpcWeb.ClientReadableStream<ComponentCreateResponse>;

  getComponents(
    request: GetComponentRequest,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: GetComponentResponse) => void
  ): grpcWeb.ClientReadableStream<GetComponentResponse>;

}

export class ComponentMgmtPromiseClient {
  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; });

  create(
    request: ComponentDetail,
    metadata: grpcWeb.Metadata
  ): Promise<ComponentCreateResponse>;

  getComponents(
    request: GetComponentRequest,
    metadata: grpcWeb.Metadata
  ): Promise<GetComponentResponse>;

}

