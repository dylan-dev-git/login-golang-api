import { Injectable } from '@angular/core';
import { Observable, of, throwError } from 'rxjs';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { catchError, tap, map } from 'rxjs/operators';
import * as ENV from 'src/app/env';

var headers = new HttpHeaders();
headers = headers.set(ENV.APIKEY, ENV.APIVALUE);
headers = headers.set('Authorization', 'Bearer ' + sessionStorage.getItem('access_token'));

@Injectable({
  providedIn: 'root'
})

export class MyinfoService {

  constructor(private http: HttpClient) { }

  myinfo(): Observable<any> {
    return this.http.get<any>(ENV.APIURL + 'api/user/userinfo', { headers: headers }).pipe(
      map(this.extractData),
      catchError(this.handleError)
    );
  }



  private extractData(res: Response) {
    const body = res;
    return body || {};
  }

  private handleError(error: HttpErrorResponse) {
    if (error.error instanceof ErrorEvent) {
      console.error('An error occurred: ', error.error.message);

      return throwError(error.error.message);
    } else {
      console.error(
        `Backend returned code ${error.status}, ` + `body was: ${error.error}`
      );

      return throwError(error.error.message);
    }
  }
}
