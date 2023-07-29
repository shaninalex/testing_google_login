import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { AuthRoutingModule } from './auth-routing.module';
import { AuthComponent } from './auth.component';
import { LoginComponent } from './login/login.component';
import { AuthService } from './auth.service';
import { ProfileComponent } from './profile/profile.component';


@NgModule({
    declarations: [
        AuthComponent,
        LoginComponent,
        ProfileComponent,
    ],
    imports: [
        CommonModule,
        AuthRoutingModule,
        HttpClientModule
    ],
    providers: [
        AuthService
    ]
})
export class AuthModule { }
