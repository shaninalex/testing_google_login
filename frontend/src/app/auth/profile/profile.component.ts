import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html'
})
export class ProfileComponent implements OnInit {

    constructor(private authService: AuthService) { }

    ngOnInit(): void {
        this.authService.getProfile().subscribe({
            next: result => {
                console.log(result);
            }
        })
    }

}
