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
    let lang = localStorage.getItem("lang")
    if (lang === null) {
      lang = navigator.language.slice(0,2);
      lang = translate.getLangs().find(l => l === lang) || "en"
      localStorage.setItem("lang", lang)
    }
    translate.setDefaultLang(lang)
  }

  
}
