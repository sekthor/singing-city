import { Component } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'songbird';

  constructor(translate: TranslateService){
    translate.addLangs(['en', 'english']);
    translate.addLangs(['de', 'deutsch']);
    translate.setDefaultLang("en")
  }

  
}
