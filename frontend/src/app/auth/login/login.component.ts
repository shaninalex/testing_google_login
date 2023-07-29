import { Component } from '@angular/core';
import { AuthService, SignInlink } from '../auth.service';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html'
})
export class LoginComponent {

    constructor(private authService: AuthService) { }

    signInWithGoogle() {
        this.authService.googleSignIn().subscribe({
            next: (result: SignInlink) => {
                window.location.href = decodeURIComponent(result.link);
            }
        })
    }
}
