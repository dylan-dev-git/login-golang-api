import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';

import { DOCS_ROUTES } from './routes';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HttpClientJsonpModule } from '@angular/common/http';

import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from "@angular/material/form-field";
import { MatIconModule } from '@angular/material/icon';

import { HomeComponent } from './page/home/home.component';
import { UserLoginComponent } from './page/user-login/user-login.component';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    UserLoginComponent
  ],
  imports: [
    RouterModule.forRoot(DOCS_ROUTES),
    BrowserModule,
    BrowserAnimationsModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    HttpClientJsonpModule,
    MatFormFieldModule,
    MatInputModule,
    MatIconModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
