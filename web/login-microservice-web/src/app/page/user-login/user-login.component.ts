import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';

import { FormBuilder, FormControl, FormGroup, ValidationErrors, Validators } from '@angular/forms';
import { LoginService } from 'src/app/services/login/login.Service';
import { GuardService } from 'src/app/services/guard/guard.Service';

@Component({
  selector: 'app-user-login',
  templateUrl: './user-login.component.html',
  styleUrls: ['./user-login.component.scss']
})
export class UserLoginComponent implements OnInit {

  loginForm!: FormGroup;
  hide = true;
  loginFailed = "";

  constructor(
    private loginService: LoginService,
    private router: Router,
    private fb: FormBuilder,
    private guardService: GuardService
  ) {
    this.formInit();
  }

  formInit(): void {
    this.loginForm = this.fb.group({
      email : ['', [Validators.required]],
      password : ['', [Validators.required]]
    });
  }

  ngOnInit() {
  }

  signin() {
    this.loginService.login(this.loginForm.value).subscribe(res => {
      if (res.status) {
        sessionStorage.setItem('access_token', res.token);
        this.guardService.userID = res.userid
        this.router.navigate(['home'])
      }
    }, err => {
      console.log(err);
      this.loginFailed = "Wrong ID/PW"
    });
  }



}
