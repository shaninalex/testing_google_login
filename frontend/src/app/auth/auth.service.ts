import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, shareReplay } from 'rxjs';

export interface SignInlink {
    link: string
}

export interface UserProfile {
    Id: string
    Name: string
    Email: string
    Image: string
    Providers: Array<string>
}


@Injectable()
export class AuthService {

    constructor(private http: HttpClient) { }

    googleSignIn(): Observable<SignInlink> {
        return this.http.get<SignInlink>("/api/v1/auth/google/login").pipe(
            shareReplay()
        )
    }

    getProfile(): Observable<UserProfile> {
        return this.http.get<UserProfile>("/api/v1/user/profile").pipe(
            shareReplay()
        )
    } 
}
