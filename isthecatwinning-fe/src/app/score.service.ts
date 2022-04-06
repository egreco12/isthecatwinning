import { Injectable } from '@angular/core';
import { Player, MOCK_PLAYERS } from './players';

@Injectable({
  providedIn: 'root'
})
export class ScoreService {

  constructor() { }

  getPlayers(): Player[] {
    return MOCK_PLAYERS;
}
}
