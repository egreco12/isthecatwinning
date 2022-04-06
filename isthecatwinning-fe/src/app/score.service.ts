import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Player, MOCK_PLAYERS } from './players';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ScoreService {

  constructor(private http: HttpClient) {
    this.getJSON().subscribe(data => {
      console.log(data);
    });
  }

  getPlayers(): Player[] {
    return MOCK_PLAYERS;
  }

  getJSON(): Observable<any> {
    return this.http.get("./assets/players.json");
  }
}
