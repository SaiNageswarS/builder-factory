export class ApplicationCreateResponse {
  constructor ();
  getApplicationname(): string;
  setApplicationname(a: string): void;
  getApplicationid(): string;
  setApplicationid(a: string): void;
  toObject(): ApplicationCreateResponse.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ApplicationCreateResponse;
}

export namespace ApplicationCreateResponse {
  export type AsObject = {
    Applicationname: string;
    Applicationid: string;
  }
}

export class ApplicationDetails {
  constructor ();
  getApplicationname(): string;
  setApplicationname(a: string): void;
  getOrgname(): string;
  setOrgname(a: string): void;
  getCloud(): string;
  setCloud(a: string): void;
  getAwsaccesskey(): string;
  setAwsaccesskey(a: string): void;
  getAwssecret(): string;
  setAwssecret(a: string): void;
  getDescription(): string;
  setDescription(a: string): void;
  toObject(): ApplicationDetails.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ApplicationDetails;
}

export namespace ApplicationDetails {
  export type AsObject = {
    Applicationname: string;
    Orgname: string;
    Cloud: string;
    Awsaccesskey: string;
    Awssecret: string;
    Description: string;
  }
}

export class GetApplicationsRequest {
  constructor ();
  getOrgname(): string;
  setOrgname(a: string): void;
  toObject(): GetApplicationsRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => GetApplicationsRequest;
}

export namespace GetApplicationsRequest {
  export type AsObject = {
    Orgname: string;
  }
}

export class GetApplicationsResponse {
  constructor ();
  getApplicationsList(): ApplicationDetails[];
  setApplicationsList(a: ApplicationDetails[]): void;
  toObject(): GetApplicationsResponse.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => GetApplicationsResponse;
}

export namespace GetApplicationsResponse {
  export type AsObject = {
    ApplicationsList: ApplicationDetails[];
  }
}

