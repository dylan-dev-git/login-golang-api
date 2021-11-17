import { Injectable } from '@angular/core';
import { Router, CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';

@Injectable({
  providedIn: 'root'
})

export class GuardService implements CanActivate {

  userID: string = "";

  constructor(
    public router: Router
  ) { }

  canActivate(): boolean {

    if (this.userID == "") {
      this.router.navigate(['login']);
      return false;
    }
    return true;
  }

}