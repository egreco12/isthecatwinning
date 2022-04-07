import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Player } from './players'
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ScoreService {

  constructor(private http: HttpClient) {}

  getJSON(): Observable<any> {
    return this.http.get("./assets/players.json");
  }
}
