import { Injectable } from '@angular/core';
import { Observable, of, throwError } from 'rxjs';
import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { catchError, tap, map } from 'rxjs/operators';
import * as ENV from 'src/app/env';

var headers = new HttpHeaders();
headers = headers.set(ENV.APIKEY, ENV.APIVALUE);

@Injectable({
  providedIn: 'root'
})

export class LoginService {

  constructor(private http: HttpClient) { }

  login(loginForm: any): Observable<any> {
    return this.http.post<any>(ENV.APIURL + 'api/auth/login', { loginForm: loginForm }, { headers: headers }).pipe(
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
