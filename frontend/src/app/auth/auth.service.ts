import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, shareReplay } from 'rxjs';

export interface SignInlink {
    link: string
}

export interface UserProfile {
    id: string
    name: string
    email: string
    image: string
    providers: Array<string>
}

export interface RegisterPayload {
    name: string
    email: string
    password: string
    password_confirm: string
}

export interface RegularRegisterResponse {
    status: boolean
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

    regularRegister(payload: RegisterPayload): Observable<RegularRegisterResponse> {
        return this.http.post<RegularRegisterResponse>("/api/v1/auth/register", payload).pipe(
            shareReplay()
        )
    }

    regularLogin(payload: RegisterPayload): Observable<RegularRegisterResponse> {
        return this.http.post<RegularRegisterResponse>("/api/v1/auth/login", payload).pipe(
            shareReplay()
        )
    }
    
}
