import { Component, OnInit } from '@angular/core';

import { GuardService } from 'src/app/services/guard/guard.Service';
import { MyinfoService } from 'src/app/services/myinfo/myinfo.Service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  userID: string = "";
  mydata: any;
  token: any;

  constructor(
    private guardService: GuardService,
    private myinfoService: MyinfoService
  ) { }

  ngOnInit(): void {
    this.userID = this.guardService.userID
    this.getMyInfo()
  }

  getMyInfo() {
    this.myinfoService.myinfo().subscribe(res => {
      if (res.status) {
        this.mydata = res.userData;
        this.token = sessionStorage.getItem('access_token')
      }
    }, err => {
      console.log(err);
    });
  }

}
