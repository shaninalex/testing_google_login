import { Component } from '@angular/core';
import { AuthService, SignInlink, UserProfile } from '../auth.service';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html'
})
export class RegisterComponent {

    registerForm: FormGroup = new FormGroup({
        "name": new FormControl("test", [Validators.required]),
        "email": new FormControl("test@test.com", [Validators.required, Validators.email]),
        "password": new FormControl("test@test.com", [Validators.required]),
        "password_confirm": new FormControl("test@test.com", [Validators.required]) // TODO: validators equal
    })

    constructor(
        private authService: AuthService,
        private router: Router
    ) { }

    Submit(): void {
        if (this.registerForm.valid) {
            this.authService.regularRegister(this.registerForm.value).subscribe({
                next: result => {
                    if (result.status) {
                        this.router.navigate(['/auth/profile']);
                    }
                }
            })
        }
    }
}
