import * as grpcWeb from 'grpc-web';
import {
  ApplicationCreateResponse,
  ApplicationDetails,
  GetApplicationsRequest,
  GetApplicationsResponse} from './ApplicationMgmt_pb';

export class ApplicationMgmtClient {
  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; });

  create(
    request: ApplicationDetails,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: ApplicationCreateResponse) => void
  ): grpcWeb.ClientReadableStream<ApplicationCreateResponse>;

  getAllApplications(
    request: GetApplicationsRequest,
    metadata: grpcWeb.Metadata,
    callback: (err: grpcWeb.Error,
               response: GetApplicationsResponse) => void
  ): grpcWeb.ClientReadableStream<GetApplicationsResponse>;

}

export class ApplicationMgmtPromiseClient {
  constructor (hostname: string,
               credentials: null | { [index: string]: string; },
               options: null | { [index: string]: string; });

  create(
    request: ApplicationDetails,
    metadata: grpcWeb.Metadata
  ): Promise<ApplicationCreateResponse>;

  getAllApplications(
    request: GetApplicationsRequest,
    metadata: grpcWeb.Metadata
  ): Promise<GetApplicationsResponse>;

}

