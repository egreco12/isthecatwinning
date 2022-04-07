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
  firstPlace: Player = new Player();
  otherPlayer: Player = new Player();
  differential: number = 0;

  constructor(private scoreService: ScoreService) {
      }

  getScore(): void {
  }

  ngOnInit(): void {
    this.scoreService.getJSON().subscribe(
	res => {
	 console.log('HTTP response', res);
         res.forEach((player: Player)=> {
           this.players.push(player);
	 });
         this.firstPlace = this.players[0];
         this.otherPlayer = this.players[1];
         this.differential = -1*(this.firstPlace.totalScore - this.otherPlayer.totalScore);
        },
	err => console.log('HTTP Error', err)
    );
  }
}
