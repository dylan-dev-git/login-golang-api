import { Routes } from '@angular/router';

import { GuardService as Guard } from './services/guard/guard.Service';
import { HomeComponent } from './page/home/home.component';
import { UserLoginComponent } from './page/user-login/user-login.component';

export const DOCS_ROUTES: Routes = [
    {
        path: 'login',
        component: UserLoginComponent,
    },
    {
        path: 'home',
        canActivate: [Guard],
        component: HomeComponent
    }
]