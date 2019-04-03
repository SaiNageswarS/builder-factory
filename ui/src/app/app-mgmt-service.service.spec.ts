import { TestBed } from '@angular/core/testing';

import { AppMgmtServiceService } from './app-mgmt-service.service';

describe('AppMgmtServiceService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: AppMgmtServiceService = TestBed.get(AppMgmtServiceService);
    expect(service).toBeTruthy();
  });
});
