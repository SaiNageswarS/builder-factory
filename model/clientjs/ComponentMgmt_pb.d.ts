export class ComponentCreateResponse {
  constructor ();
  getComponentname(): string;
  setComponentname(a: string): void;
  getComponentid(): number;
  setComponentid(a: number): void;
  toObject(): ComponentCreateResponse.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ComponentCreateResponse;
}

export namespace ComponentCreateResponse {
  export type AsObject = {
    Componentname: string;
    Componentid: number;
  }
}

export class ComponentDetail {
  constructor ();
  getComponentname(): string;
  setComponentname(a: string): void;
  getApplicationname(): string;
  setApplicationname(a: string): void;
  getComponenttype(): ComponentDetail.ComponentType;
  setComponenttype(a: ComponentDetail.ComponentType): void;
  getOrgname(): string;
  setOrgname(a: string): void;
  getPackage(): Package;
  setPackage(a: Package): void;
  getPipelinestate(): ComponentDetail.PipelineState;
  setPipelinestate(a: ComponentDetail.PipelineState): void;
  toObject(): ComponentDetail.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ComponentDetail;
}

export namespace ComponentDetail {
  export type AsObject = {
    Componentname: string;
    Applicationname: string;
    Componenttype: ComponentDetail.ComponentType;
    Orgname: string;
    Package: Package;
    Pipelinestate: ComponentDetail.PipelineState;
  }

  export enum ComponentType { 
    FRONTEND_WEB = 0,
    BACKEND_API_REST = 1,
    BACKEND_API_RPC = 2,
    SCHEDULED_JOB = 3,
    WORKFLOW = 4,
    MAPREDUCE_CLUSTER = 5,
  }

  export enum PipelineState { 
    DEPLOYED = 0,
    NOT_CONFIGURED = 1,
    BROKEN = 2,
  }
}

export class GetComponentRequest {
  constructor ();
  getOrgname(): string;
  setOrgname(a: string): void;
  getAppname(): string;
  setAppname(a: string): void;
  toObject(): GetComponentRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => GetComponentRequest;
}

export namespace GetComponentRequest {
  export type AsObject = {
    Orgname: string;
    Appname: string;
  }
}

export class GetComponentResponse {
  constructor ();
  getComponentsList(): ComponentDetail[];
  setComponentsList(a: ComponentDetail[]): void;
  toObject(): GetComponentResponse.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => GetComponentResponse;
}

export namespace GetComponentResponse {
  export type AsObject = {
    ComponentsList: ComponentDetail[];
  }
}

export class Package {
  constructor ();
  getName(): string;
  setName(a: string): void;
  getRepourl(): string;
  setRepourl(a: string): void;
  getLanguage(): string;
  setLanguage(a: string): void;
  getRepoprovider(): Package.RepoProvider;
  setRepoprovider(a: Package.RepoProvider): void;
  getRepoweburl(): string;
  setRepoweburl(a: string): void;
  toObject(): Package.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => Package;
}

export namespace Package {
  export type AsObject = {
    Name: string;
    Repourl: string;
    Language: string;
    Repoprovider: Package.RepoProvider;
    Repoweburl: string;
  }

  export enum RepoProvider { 
    BITBUCKET = 0,
    GITHUB = 1,
  }
}

