import { Component } from '@angular/core';
import { ScoreService } from './score.service'
import { Player } from './players';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'isthecatwinning';
  players: Player[] = [];

  constructor(private scoreService: ScoreService) {}

  getScore(): void {
    this.players = this.scoreService.getPlayers();	
  }

  ngOnInit(): void {
    this.getScore();
  }
}
