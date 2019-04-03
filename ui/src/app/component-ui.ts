import {ComponentDetail, Package} from 'autodevopsappservicemodel/clientjs/ComponentMgmt_pb'

export class ComponentUI {
    appComponent: ComponentDetail;

    constructor(private component: ComponentDetail) {
        this.appComponent = component;
    }

    c() {
        return this.appComponent;
    }

    getComponentNameString(): string {
        switch(this.appComponent.getComponenttype()) {
            case ComponentDetail.ComponentType.FRONTEND_WEB: return "Frontend Web";
            case ComponentDetail.ComponentType.BACKEND_API_RPC: return "Backend RPC Service";
            default: return "Package";
        }
    }

    getComponentAvatar(): string {
        switch(this.appComponent.getComponenttype()) {
            case ComponentDetail.ComponentType.FRONTEND_WEB: return "assets/angular-icon.png";
            case ComponentDetail.ComponentType.BACKEND_API_RPC: return "assets/grpc-icon.png";
            default: return "Package";
        }
    }

    getRepoIcon(): string {
        switch(this.appComponent.getPackage().getRepoprovider()) {
            case Package.RepoProvider.BITBUCKET: return "assets/bitbucket-icon.png";
            case Package.RepoProvider.GITHUB: return "assets/github-icon.png";
        }
    }

    getPipelineStateIcon(): string {
        switch(this.appComponent.getPipelinestate()) {
            case ComponentDetail.PipelineState.DEPLOYED: return "green-dot";
            case ComponentDetail.PipelineState.BROKEN: return "red-dot";
            case ComponentDetail.PipelineState.NOT_CONFIGURED: return "grey-dot";
        }
    }

    getPipelineState(): string {
        switch(this.appComponent.getPipelinestate()) {
            case ComponentDetail.PipelineState.DEPLOYED: return "Pipeline deployed";
            case ComponentDetail.PipelineState.BROKEN: return "Pipeline stalled";
            case ComponentDetail.PipelineState.NOT_CONFIGURED: return "Pipeline not configured";
        }
    }
}
