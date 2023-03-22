import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SharedFunctionsService {

  constructor() { }

  private subject = new Subject<any>();

  menuNavToggle() {
    this.subject.next({});
  }

  getMenuResponse(): Observable<any>{ 
    return this.subject.asObservable();
  }
}