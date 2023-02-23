import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CatchComponent } from '../catch.component';
import { HomeComponent } from '../home.component';
import { LoginEditorComponent } from '../login-editor.component';
import { registerFormComponent } from '../registerform.component';

export const routes: Routes = [
  { path: 'login', component: LoginEditorComponent },
  { path: 'register', component: registerFormComponent },
  { path: 'home', component: HomeComponent },
  { path: 'test', component: CatchComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}
